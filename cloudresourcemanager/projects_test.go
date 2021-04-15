package cloudresourcemanager

import (
	"regexp"
	"strings"
	"testing"

	loggermock "github.com/rockholla/go-lib/mocks/custom-mocks/logger/logger"
	v1 "google.golang.org/api/cloudresourcemanager/v1"
	googleapi "google.golang.org/api/googleapi"
	smv1 "google.golang.org/api/servicemanagement/v1"
)

const (
	testProjectName            = "projects/tests"
	testProjectNumber          = 10928177192
	testProjectID              = "tests-019281921"
	testProjectNameDoesntExist = "projects/new"
	testProjectParentFolder    = "folders/1111111111"
)

var (
	listCount = 0
)

type projectsListMock struct{}
type projectsListMockNoResults struct{}
type projectsListNoResultThenResult struct{}
type projectsGetMock struct{}
type projectsCreateMock struct{}
type projectsDeleteMock struct{}
type projectsGetIAMPolicyMock struct{}
type projectsGetIAMPolicyExistingMemberMock struct{}
type projectsGetIAMPolicyExistingRoleMock struct{}
type projectsSetIAMPolicyMock struct{}
type servicesEnableMock struct{}

// Do is the mock for default projectsListMock
func (c *projectsListMock) Do(call *v1.ProjectsListCall, opts ...googleapi.CallOption) (*v1.ListProjectsResponse, error) {
	var projects []*v1.Project
	projects = append(projects, &v1.Project{
		Name:          testProjectName,
		ProjectId:     testProjectID,
		ProjectNumber: testProjectNumber,
		Parent: &v1.ResourceId{
			Type: "folder",
			Id:   strings.Replace(testProjectParentFolder, "folders/", "", 1),
		},
		LifecycleState: "ACTIVE",
	})
	return &v1.ListProjectsResponse{
		Projects: projects,
	}, nil
}

// Do is the mock for projectsList that doesn't include any results
func (c *projectsListMockNoResults) Do(call *v1.ProjectsListCall, opts ...googleapi.CallOption) (*v1.ListProjectsResponse, error) {
	var projects []*v1.Project
	return &v1.ListProjectsResponse{
		Projects: projects,
	}, nil
}

// Do is the mock for projectsList that will return nothing on the first time, will return a result on the second
func (c *projectsListNoResultThenResult) Do(call *v1.ProjectsListCall, opts ...googleapi.CallOption) (*v1.ListProjectsResponse, error) {
	var projects []*v1.Project
	if listCount == 0 {
		listCount++
	} else {
		projects = append(projects, &v1.Project{
			Name:          testProjectName,
			ProjectId:     testProjectID,
			ProjectNumber: testProjectNumber,
			Parent: &v1.ResourceId{
				Type: "folder",
				Id:   strings.Replace(testProjectParentFolder, "folders/", "", 1),
			},
			LifecycleState: "ACTIVE",
		})
		listCount = 0
	}
	return &v1.ListProjectsResponse{
		Projects: projects,
	}, nil
}

// Do is the mock for default projectsGetMock
func (c *projectsGetMock) Do(call *v1.ProjectsGetCall, opts ...googleapi.CallOption) (*v1.Project, error) {
	return &v1.Project{
		Name:           testProjectName,
		ProjectNumber:  testProjectNumber,
		LifecycleState: "ACTIVE",
	}, nil
}

// Do is the mock for default projectsCreateMock
func (c *projectsCreateMock) Do(call *v1.ProjectsCreateCall, opts ...googleapi.CallOption) (*v1.Operation, error) {
	return &v1.Operation{
		Error: nil,
	}, nil
}

// Do is the mock for default projectsDeleteMock
func (c *projectsDeleteMock) Do(call *v1.ProjectsDeleteCall, opts ...googleapi.CallOption) (*v1.Empty, error) {
	return &v1.Empty{
		ServerResponse: googleapi.ServerResponse{
			HTTPStatusCode: 200,
		},
	}, nil
}

// Do is the mock for default projectsGetIAMPolicy
func (c *projectsGetIAMPolicyMock) Do(call *v1.ProjectsGetIamPolicyCall, opts ...googleapi.CallOption) (*v1.Policy, error) {
	return &v1.Policy{}, nil
}

// Do is the mock for projectsGetIamPolicy that includes an existing member in role
func (c *projectsGetIAMPolicyExistingMemberMock) Do(call *v1.ProjectsGetIamPolicyCall, opts ...googleapi.CallOption) (*v1.Policy, error) {
	var bindings []*v1.Binding
	bindings = append(bindings, &v1.Binding{
		Role: testRole,
		Members: []string{
			testMember,
		},
	})
	return &v1.Policy{
		Bindings: bindings,
	}, nil
}

// Do is the mock for projectsGetIAMPolicy that includes an existing member in role
func (c *projectsGetIAMPolicyExistingRoleMock) Do(call *v1.ProjectsGetIamPolicyCall, opts ...googleapi.CallOption) (*v1.Policy, error) {
	var bindings []*v1.Binding
	bindings = append(bindings, &v1.Binding{
		Role:    testRole,
		Members: []string{},
	})
	return &v1.Policy{
		Bindings: bindings,
	}, nil
}

// Do is the mock for default projectsSetIAMPolicy
func (c *projectsSetIAMPolicyMock) Do(call *v1.ProjectsSetIamPolicyCall, opts ...googleapi.CallOption) (*v1.Policy, error) {
	return &v1.Policy{}, nil
}

// Do is the mock for default servicesEnable
func (c *servicesEnableMock) Do(call *smv1.ServicesEnableCall, opts ...googleapi.CallOption) (*smv1.Operation, error) {
	return &smv1.Operation{}, nil
}

func setProjectsCallMockDefaults(crm *CloudResourceManager) {
	crm.Calls = &Calls{
		ProjectsList:         &projectsListMock{},
		ProjectsGet:          &projectsGetMock{},
		ProjectsCreate:       &projectsCreateMock{},
		ProjectsDelete:       &projectsDeleteMock{},
		ProjectsGetIAMPolicy: &projectsGetIAMPolicyMock{},
		ProjectsSetIAMPolicy: &projectsSetIAMPolicyMock{},
		ServicesEnable:       &servicesEnableMock{},
	}
}

func TestMakeProjectID(t *testing.T) {
	_, err := MakeProjectID("project", "invalid-parent")
	if err == nil {
		t.Error("cloudresourcemanager.MakeProjectID() didn't fail as expected with invalid parent")
	}
	projectID, err := MakeProjectID("project", "folders/test")
	if err != nil {
		t.Errorf("Got unexpected error during cloudresourcemanager.MakeProjectID(): %s", err)
	}
	matched, _ := regexp.MatchString("^project-test[0-9]+", projectID)
	if !matched {
		t.Errorf("Expected cloudresourcemanager.MakeProjectID() result to match \"^project-test[0-9]+\", instead got: %s", projectID)
	}
}

func TestMakeProjectIDTruncate(t *testing.T) {
	projectID, err := MakeProjectID("project", "folders/123456789012345678901234567890")
	if err != nil {
		t.Errorf("Got unexpected error during cloudresourcemanager.MakeProjectID() that should be truncated: %s", err)
	}
	if len(projectID) != 30 {
		t.Errorf("Expected cloudresourcemanager.MakeProjectID() trucated result to return length of 30, instead got: %d", len(projectID))
	}
}

func TestGetProject(t *testing.T) {
	crm := &CloudResourceManager{}
	err := crm.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during cloudresourcemanager.Initialize() with blank credentials: %s", err)
	}
	setProjectsCallMockDefaults(crm)
	project, err := crm.GetProject(testProjectName, testProjectParentFolder)
	if err != nil {
		t.Errorf("Got unexpected error during cloudresourcemanager.GetProject(): %s", err)
	}
	if project.Name != testProjectName {
		t.Errorf("Didn't get expected testProjectName \"%s\" from cloudresourcemanager.GetProject(), instead: %s", testProjectName, project.Name)
	}
}

func TestGetProjectByID(t *testing.T) {
	crm := &CloudResourceManager{}
	err := crm.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during cloudresourcemanager.Initialize() with blank credentials: %s", err)
	}
	setProjectsCallMockDefaults(crm)
	project, err := crm.GetProjectByID("test-project-id")
	if err != nil {
		t.Errorf("Got unexpected error during cloudresourcemanager.GetProjectByID(): %s", err)
	}
	if project.Name != testProjectName {
		t.Errorf("Didn't get expected testProjectName \"%s\" from cloudresourcemanager.GetProjectByID(), instead: %s", testProjectName, project.Name)
	}
}

func TestGetProjectNoResults(t *testing.T) {
	crm := &CloudResourceManager{}
	err := crm.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during cloudresourcemanager.Initialize() with blank credentials: %s", err)
	}
	setProjectsCallMockDefaults(crm)
	crm.Calls.ProjectsList = &projectsListMockNoResults{}
	project, err := crm.GetProject(testProjectName, testProjectParentFolder)
	if err != nil {
		t.Errorf("Got unexpected error during cloudresourcemanager.GetProject() with no results: %s", err)
	}
	if project != nil {
		t.Errorf("Didn't get empty project result from cloudresourcemanager.GetProject() with no results: %s", err)
	}
}

func TestEnsureProjectAlreadyExists(t *testing.T) {
	crm := &CloudResourceManager{}
	err := crm.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during cloudresourcemanager.Initialize() with blank credentials: %s", err)
	}
	setProjectsCallMockDefaults(crm)
	ID, number, err := crm.EnsureProject(testProjectName, testProjectParentFolder)
	if err != nil {
		t.Errorf("Got unexpected error during cloudresourcemanager.EnsureProject() for project that already exists: %s", err)
	}
	if ID == "" {
		t.Error("Got unexpected blank project ID from cloudresourcemanager.EnsureProject() for project that already exists")
	}
	if number != testProjectNumber {
		t.Errorf("Got unexpected project number from cloudresourcemanager.EnsureProject() for project that already exists: %d", number)
	}
}

func TestEnsureProjectDoesntExist(t *testing.T) {
	crm := &CloudResourceManager{}
	err := crm.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during cloudresourcemanager.Initialize() with blank credentials: %s", err)
	}
	setProjectsCallMockDefaults(crm)
	crm.Calls.ProjectsList = &projectsListNoResultThenResult{}
	ID, number, err := crm.EnsureProject(testProjectName, testProjectParentFolder)
	if err != nil {
		t.Errorf("Got unexpected error during cloudresourcemanager.EnsureProject() for project that doesn't exist: %s", err)
	}
	if ID == "" {
		t.Error("Got unexpected blank project ID from cloudresourcemanager.EnsureProject() for project that already exists")
	}
	if number != testProjectNumber {
		t.Errorf("Got unexpected project number from cloudresourcemanager.EnsureProject() for project that doesn't exist: %d", number)
	}
}

func TestEnsureProjectRoles(t *testing.T) {
	crm := &CloudResourceManager{}
	err := crm.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error for cloudresourcemanager.Initialize() with blank credentials: %s", err)
	}
	setProjectsCallMockDefaults(crm)
	err = crm.EnsureProjectRoles(testProjectName, testMember, []string{testRole, "role2"})
	if err != nil {
		t.Errorf("Got unexpected error for cloudresourcemanager.EnsureProjectRoles(): %s", err)
	}
}

func TestEnsureProjectRolesAlternativeProject(t *testing.T) {
	crm := &CloudResourceManager{}
	err := crm.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error for cloudresourcemanager.Initialize() with blank credentials: %s", err)
	}
	setProjectsCallMockDefaults(crm)
	err = crm.EnsureProjectRoles(strings.Replace(testProjectName, "projects/", "", 1), testMember, []string{testRole, "role2"})
	if err != nil {
		t.Errorf("Got unexpected error for cloudresourcemanager.EnsureProjectRoles() with alternative project format: %s", err)
	}
}

func TestEnsureProjectRolesExistingMember(t *testing.T) {
	crm := &CloudResourceManager{}
	err := crm.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error for cloudresourcemanager.Initialize() with blank credentials: %s", err)
	}
	setProjectsCallMockDefaults(crm)
	crm.Calls.ProjectsGetIAMPolicy = &projectsGetIAMPolicyExistingMemberMock{}
	err = crm.EnsureProjectRoles(testProjectName, testMember, []string{testRole, "role2"})
	if err != nil {
		t.Errorf("Got unexpected error for cloudresourcemanager.EnsureProjectRoles() with existing member: %s", err)
	}
}

func TestEnsureProjectRolesExistingRoleWithoutMember(t *testing.T) {
	crm := &CloudResourceManager{}
	err := crm.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error for cloudresourcemanager.Initialize() with blank credentials: %s", err)
	}
	setProjectsCallMockDefaults(crm)
	crm.Calls.ProjectsGetIAMPolicy = &projectsGetIAMPolicyExistingRoleMock{}
	err = crm.EnsureProjectRoles(testProjectName, testMember, []string{testRole, "role2"})
	if err != nil {
		t.Errorf("Got unexpected error for cloudresourcemanager.EnsureProjectRoles() with existing role, but member not within: %s", err)
	}
}

func TestEnableProjectServices(t *testing.T) {
	crm := &CloudResourceManager{}
	err := crm.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error for cloudresourcemanager.Initialize() with blank credentials: %s", err)
	}
	setProjectsCallMockDefaults(crm)
	err = crm.EnableProjectServices(testProjectName, []string{"one", "two", "three"})
	if err != nil {
		t.Errorf("Got unexpected error for cloudresourcemanager.EnableProjectServices(): %s", err)
	}
}

func TestDeleteProject(t *testing.T) {
	crm := &CloudResourceManager{}
	err := crm.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during cloudresourcemanager.Initialize() with blank credentials: %s", err)
	}
	setProjectsCallMockDefaults(crm)
	crm.Calls.ProjectsDelete = &projectsDeleteMock{}
	err = crm.DeleteProject(testProjectID)
	if err != nil {
		t.Errorf("Got unexpected error during cloudresourcemanager.DeleteProject(): %s", err)
	}
}
