package cloudresourcemanager

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"

	v1 "google.golang.org/api/cloudresourcemanager/v1"
	suv1 "google.golang.org/api/serviceusage/v1"
)

// MakeProjectID will construct a custom project ID based on the name of the project and it's parent
func MakeProjectID(name string, parent string) (string, error) {
	parentParts := strings.Split(parent, "/")
	if len(parentParts) != 2 {
		return "", fmt.Errorf("Expecting the parent argument to be like [type]/[ID], e.g. folders/92737276394872, but got: %s", parent)
	}
	timestamp := fmt.Sprintf("%d", time.Now().Unix())
	base := fmt.Sprintf("%s-%s", name, parentParts[1])
	if len(base) > 25 {
		base = base[:25]
	}
	return fmt.Sprintf("%s%s", base, timestamp[len(timestamp)-5:]), nil
}

// GetProject returns an existing project object, nil if none found
func (crm *CloudResourceManager) GetProject(name string, parent string) (*v1.Project, error) {
	ctx := context.Background()
	parentParts := strings.Split(parent, "/")
	projectsService := v1.NewProjectsService(crm.V1)
	projectsListCall := projectsService.List().Context(ctx)
	filter := fmt.Sprintf("name:%s parent.type:folder parent.id:%s lifecycleState:ACTIVE", name, parentParts[1])
	projectsListCall = projectsListCall.Filter(filter)
	listProjectsResponse, err := crm.Calls.ProjectsList.Do(projectsListCall)
	if err != nil {
		return nil, err
	}

	if len(listProjectsResponse.Projects) == 0 {
		return nil, nil
	}
	return listProjectsResponse.Projects[0], nil
}

// GetProjectByID gets an existing project object, found by its ID
func (crm *CloudResourceManager) GetProjectByID(id string) (*v1.Project, error) {
	ctx := context.Background()
	projectsService := v1.NewProjectsService(crm.V1)
	projectsGetCall := projectsService.Get(id).Context(ctx)
	project, err := crm.Calls.ProjectsGet.Do(projectsGetCall)
	if err != nil {
		return nil, err
	}
	if project.LifecycleState != "ACTIVE" {
		return nil, nil
	}
	return project, nil
}

// EnsureProject will make sure that a project exists, creates it if it doesn't already exist, nothing if it does,
// returns either new or existing project ID and project number
func (crm *CloudResourceManager) EnsureProject(name string, parent string) (string, int64, error) {
	crm.log.InfoPart("Ensuring that project %s exists", name)
	if parent != "" {
		crm.log.InfoPart(" in %s...", parent)
	}
	ctx := context.Background()
	projectsService := v1.NewProjectsService(crm.V1)
	existingProject, err := crm.GetProject(name, parent)
	if err != nil {
		crm.log.InfoPart("\n")
		return "", 0, err
	}
	if existingProject != nil {
		crm.log.InfoPart("already exists\n")
		return existingProject.ProjectId, existingProject.ProjectNumber, nil
	}
	crm.log.InfoPart("creating\n")
	projectID, err := MakeProjectID(name, parent)
	if err != nil {
		return "", 0, err
	}
	parentParts := strings.Split(parent, "/")
	parentResource := &v1.ResourceId{
		Type: strings.TrimRight(parentParts[0], "s"),
		Id:   parentParts[1],
	}
	project := &v1.Project{
		Name:      name,
		Parent:    parentResource,
		ProjectId: projectID,
	}
	projectCreateCall := projectsService.Create(project).Context(ctx)
	projectCreateOperation, err := crm.Calls.ProjectsCreate.Do(projectCreateCall)
	if err != nil {
		return "", 0, err
	}
	if projectCreateOperation.Error != nil {
		return "", 0, errors.New(projectCreateOperation.Error.Message)
	}
	for existingProject == nil {
		existingProject, err = crm.GetProject(name, parent)
		if err != nil {
			return "", 0, err
		}
		if existingProject == nil {
			time.Sleep(3 * time.Second)
		}
	}
	return existingProject.ProjectId, existingProject.ProjectNumber, nil
}

// EnableProjectServices will enable 1 or many services in a project
func (crm *CloudResourceManager) EnableProjectServices(projectID string, services []string) error {
	ctx := context.Background()
	servicesService := suv1.NewServicesService(crm.SUV1)
	crm.log.Info("Ensuring service APIs are enabled in project %s:", projectID)
	project, err := crm.GetProjectByID(projectID)
	if err != nil {
		return err
	}
	for _, service := range services {
		crm.log.ListItem(service)
	}
	enableServiceRequest := &suv1.BatchEnableServicesRequest{
		ServiceIds: services,
	}
	servicesEnableCall := servicesService.BatchEnable(fmt.Sprintf("projects/%d", project.ProjectNumber), enableServiceRequest).Context(ctx)
	servicesEnableOperation, err := crm.Calls.ServicesEnable.Do(servicesEnableCall)
	if err != nil {
		return err
	}
	if servicesEnableOperation.Error != nil {
		return errors.New(servicesEnableOperation.Error.Message)
	}
	// TODO: wait for services to actually be enabled, proven difficult in current state of affairs
	// with Google SDK/API here
	return nil
}

// EnsureProjectRoles makes sure that a particular member has the supplied roles on the project
func (crm *CloudResourceManager) EnsureProjectRoles(project string, member string, roles []string) error {
	ctx := context.Background()
	if matched, _ := regexp.Match("^projects\\/", []byte(project)); matched {
		project = strings.Replace(project, "projects/", "", 1)
	}
	crm.log.Info("Ensuring member %s has roles in projects/%s:", member, project)
	projectsService := v1.NewProjectsService(crm.V1)
	projectPolicyGetCall := projectsService.GetIamPolicy(project, &v1.GetIamPolicyRequest{}).Context(ctx)
	policy, err := crm.Calls.ProjectsGetIAMPolicy.Do(projectPolicyGetCall)
	if err != nil {
		return err
	}
	for _, role := range roles {
		crm.log.ListItem(role)
		binding := getExistingProjectRoleBinding(policy.Bindings, role)
		if binding == nil {
			policy.Bindings = append(policy.Bindings, &v1.Binding{
				Members: []string{member},
				Role:    role,
			})
		} else if !memberExistsInProjectRoleBinding(binding, member) {
			binding.Members = append(binding.Members, member)
		}
	}
	setIAMPolicyRequest := &v1.SetIamPolicyRequest{
		Policy: policy,
	}
	projectSetPolicyCall := projectsService.SetIamPolicy(project, setIAMPolicyRequest).Context(ctx)
	_, err = crm.Calls.ProjectsSetIAMPolicy.Do(projectSetPolicyCall)
	if err != nil {
		return err
	}
	return nil
}

// DeleteProject will delete a Google Cloud project ID
func (crm *CloudResourceManager) DeleteProject(id string) error {
	crm.log.InfoPart("Deleting project %s...", id)
	existingProject, err := crm.GetProjectByID(id)
	if err != nil {
		crm.log.InfoPart("error\n")
		return fmt.Errorf("error determining if project to delete exists: %s", err)
	}
	if existingProject == nil {
		return nil
	}
	ctx := context.Background()
	projectsService := v1.NewProjectsService(crm.V1)
	projectDeleteCall := projectsService.Delete(id).Context(ctx)
	projectDeleteEmpty, err := crm.Calls.ProjectsDelete.Do(projectDeleteCall)
	if err != nil {
		crm.log.InfoPart("error\n")
		return err
	}
	if projectDeleteEmpty.ServerResponse.HTTPStatusCode >= 300 || projectDeleteEmpty.ServerResponse.HTTPStatusCode < 200 {
		crm.log.InfoPart("error\n")
		return fmt.Errorf("error deleting project: status code %d", projectDeleteEmpty.ServerResponse.HTTPStatusCode)
	}
	crm.log.InfoPart("done\n")
	return nil
}

func getExistingProjectRoleBinding(bindings []*v1.Binding, role string) *v1.Binding {
	for _, binding := range bindings {
		if binding.Role == role {
			return binding
		}
	}
	return nil
}

func memberExistsInProjectRoleBinding(binding *v1.Binding, member string) bool {
	for _, bindingMember := range binding.Members {
		if bindingMember == member {
			return true
		}
	}
	return false
}
