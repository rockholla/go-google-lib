package cloudresourcemanager

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"time"

	v2beta1 "google.golang.org/api/cloudresourcemanager/v2beta1"
)

// GetFolder returns an existing folder name, blank if none found
func (crm *CloudResourceManager) GetFolder(displayName string, parent string) (string, error) {
	ctx := context.Background()
	foldersService := v2beta1.NewFoldersService(crm.V2Beta1)
	query := fmt.Sprintf("displayName=%s AND lifecycleState=ACTIVE", displayName)
	if parent != "" {
		query = fmt.Sprintf("%s AND parent=%s", query, parent)
	}
	folderSearchRequest := &v2beta1.SearchFoldersRequest{
		PageSize: 1,
		Query:    query,
	}
	folderSearchCall := foldersService.Search(folderSearchRequest).Context(ctx)
	folderSearchResponse, err := crm.Calls.FoldersSearch.Do(folderSearchCall)
	if err != nil {
		return "", err
	}
	if len(folderSearchResponse.Folders) == 0 {
		return "", nil
	}
	return folderSearchResponse.Folders[0].Name, nil
}

// EnsureFolder will make sure that a folder exists, creates it if it doesn't already exist, nothing if it does,
// returns either new or existing folder name
func (crm *CloudResourceManager) EnsureFolder(displayName string, parent string) (string, error) {
	crm.log.InfoPart("Ensuring that folder %s exists", displayName)
	if parent != "" {
		crm.log.InfoPart(" in %s...", parent)
	}
	ctx := context.Background()
	foldersService := v2beta1.NewFoldersService(crm.V2Beta1)
	name, err := crm.GetFolder(displayName, parent)
	if err != nil {
		crm.log.InfoPart("\n")
		return "", err
	}
	if name != "" {
		crm.log.InfoPart("already exists\n")
		return name, nil
	}
	crm.log.InfoPart("creating\n")
	folder := &v2beta1.Folder{
		DisplayName: displayName,
	}
	folderCreateCall := foldersService.Create(folder).Context(ctx).Parent(parent)
	folderCreateOperation, err := crm.Calls.FoldersCreate.Do(folderCreateCall)
	if err != nil {
		return "", err
	}
	if folderCreateOperation.Error != nil {
		return "", errors.New(folderCreateOperation.Error.Message)
	}
	for name == "" {
		name, err = crm.GetFolder(displayName, parent)
		if err != nil {
			return "", err
		}
		if name == "" {
			time.Sleep(3 * time.Second)
		}
	}
	return name, nil
}

// EnsureFolderRoles makes sure that a particular member has the supplied roles on the folder
func (crm *CloudResourceManager) EnsureFolderRoles(folder string, member string, roles []string) error {
	ctx := context.Background()
	if matched, _ := regexp.Match("^folders\\/", []byte(folder)); !matched {
		folder = fmt.Sprintf("folders/%s", folder)
	}
	crm.log.Info("Ensuring member %s has roles in %s:", member, folder)
	foldersService := v2beta1.NewFoldersService(crm.V2Beta1)
	folderGetPolicyCall := foldersService.GetIamPolicy(folder, &v2beta1.GetIamPolicyRequest{}).Context(ctx)
	policy, err := crm.Calls.FoldersGetIAMPolicy.Do(folderGetPolicyCall)
	if err != nil {
		return err
	}
	for _, role := range roles {
		crm.log.ListItem(role)
		binding := getExistingFolderRoleBinding(policy.Bindings, role)
		if binding == nil {
			policy.Bindings = append(policy.Bindings, &v2beta1.Binding{
				Members: []string{member},
				Role:    role,
			})
		} else if !memberExistsInFolderRoleBinding(binding, member) {
			binding.Members = append(binding.Members, member)
		}
	}
	setIAMPolicyRequest := &v2beta1.SetIamPolicyRequest{
		Policy: policy,
	}
	folderSetPolicyCall := foldersService.SetIamPolicy(folder, setIAMPolicyRequest).Context(ctx)
	_, err = crm.Calls.FoldersSetIAMPolicy.Do(folderSetPolicyCall)
	if err != nil {
		return err
	}
	return nil
}

func getExistingFolderRoleBinding(bindings []*v2beta1.Binding, role string) *v2beta1.Binding {
	for _, binding := range bindings {
		if binding.Role == role {
			return binding
		}
	}
	return nil
}

func memberExistsInFolderRoleBinding(binding *v2beta1.Binding, member string) bool {
	for _, bindingMember := range binding.Members {
		if bindingMember == member {
			return true
		}
	}
	return false
}
