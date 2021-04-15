package cloudresourcemanager

import (
	"strings"
	"testing"

	loggermock "github.com/rockholla/go-lib/mocks/custom-mocks/logger"
	v2beta1 "google.golang.org/api/cloudresourcemanager/v2beta1"
	googleapi "google.golang.org/api/googleapi"
)

const (
	testFolderName            = "folders/tests"
	testFolderNameDoesntExist = "folders/new"
)

var (
	searchCount = 0
)

type foldersSearchMock struct{}
type foldersSearchMockNoResults struct{}
type foldersSearchNoResultThenResult struct{}
type foldersCreateMock struct{}
type foldersGetIAMPolicyMock struct{}
type foldersGetIAMPolicyExistingMemberMock struct{}
type foldersGetIAMPolicyExistingRoleMock struct{}
type foldersSetIAMPolicyMock struct{}

// Do is the mock for default foldersSearchMock
func (c *foldersSearchMock) Do(call *v2beta1.FoldersSearchCall, opts ...googleapi.CallOption) (*v2beta1.SearchFoldersResponse, error) {
	var folders []*v2beta1.Folder
	folders = append(folders, &v2beta1.Folder{
		Name: testFolderName,
	})
	return &v2beta1.SearchFoldersResponse{
		Folders: folders,
	}, nil
}

// Do is the mock for foldersSearch that doesn't include any results
func (c *foldersSearchMockNoResults) Do(call *v2beta1.FoldersSearchCall, opts ...googleapi.CallOption) (*v2beta1.SearchFoldersResponse, error) {
	var folders []*v2beta1.Folder
	return &v2beta1.SearchFoldersResponse{
		Folders: folders,
	}, nil
}

// Do is the mock for foldersSearch that will return nothing on the first time, will return a result on the second
func (c *foldersSearchNoResultThenResult) Do(call *v2beta1.FoldersSearchCall, opts ...googleapi.CallOption) (*v2beta1.SearchFoldersResponse, error) {
	var folders []*v2beta1.Folder
	if searchCount == 0 {
		searchCount++
	} else {
		folders = append(folders, &v2beta1.Folder{
			Name: testFolderName,
		})
		searchCount = 0
	}
	return &v2beta1.SearchFoldersResponse{
		Folders: folders,
	}, nil
}

// Do is the mock for default foldersCreateMock
func (c *foldersCreateMock) Do(call *v2beta1.FoldersCreateCall, opts ...googleapi.CallOption) (*v2beta1.Operation, error) {
	return &v2beta1.Operation{
		Error: nil,
	}, nil
}

// Do is the mock for default foldersGetIAMPolicy
func (c *foldersGetIAMPolicyMock) Do(call *v2beta1.FoldersGetIamPolicyCall, opts ...googleapi.CallOption) (*v2beta1.Policy, error) {
	return &v2beta1.Policy{}, nil
}

// Do is the mock for foldersGetIamPolicy that includes an existing member in role
func (c *foldersGetIAMPolicyExistingMemberMock) Do(call *v2beta1.FoldersGetIamPolicyCall, opts ...googleapi.CallOption) (*v2beta1.Policy, error) {
	var bindings []*v2beta1.Binding
	bindings = append(bindings, &v2beta1.Binding{
		Role: testRole,
		Members: []string{
			testMember,
		},
	})
	return &v2beta1.Policy{
		Bindings: bindings,
	}, nil
}

// Do is the mock for foldersGetIAMPolicy that includes an existing member in role
func (c *foldersGetIAMPolicyExistingRoleMock) Do(call *v2beta1.FoldersGetIamPolicyCall, opts ...googleapi.CallOption) (*v2beta1.Policy, error) {
	var bindings []*v2beta1.Binding
	bindings = append(bindings, &v2beta1.Binding{
		Role:    testRole,
		Members: []string{},
	})
	return &v2beta1.Policy{
		Bindings: bindings,
	}, nil
}

// Do is the mock for default foldersSetIAMPolicy
func (c *foldersSetIAMPolicyMock) Do(call *v2beta1.FoldersSetIamPolicyCall, opts ...googleapi.CallOption) (*v2beta1.Policy, error) {
	return &v2beta1.Policy{}, nil
}

func setFoldersCallMockDefaults(crm *CloudResourceManager) {
	crm.Calls = &Calls{
		FoldersSearch:       &foldersSearchMock{},
		FoldersCreate:       &foldersCreateMock{},
		FoldersGetIAMPolicy: &foldersGetIAMPolicyMock{},
		FoldersSetIAMPolicy: &foldersSetIAMPolicyMock{},
	}
}

func TestGetFolder(t *testing.T) {
	crm := &CloudResourceManager{}
	err := crm.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during cloudresourcemanager.Initialize() with blank credentials: %s", err)
	}
	setFoldersCallMockDefaults(crm)
	_, err = crm.GetFolder("000000", "folders/111111111")
	if err != nil {
		t.Errorf("Got unexpected error during cloudresourcemanager.GetFolder(): %s", err)
	}
}

func TestGetFolderNoResults(t *testing.T) {
	crm := &CloudResourceManager{}
	err := crm.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during cloudresourcemanager.Initialize() with blank credentials: %s", err)
	}
	setFoldersCallMockDefaults(crm)
	crm.Calls.FoldersSearch = &foldersSearchMockNoResults{}
	result, err := crm.GetFolder("000000", "folders/111111111")
	if err != nil {
		t.Errorf("Got unexpected error during cloudresourcemanager.GetFolder() with no results: %s", err)
	}
	if result != "" {
		t.Errorf("Didn't get expected blank string from result of cloudresourcemanager.GetFolder() with no results: %s", err)
	}
}

func TestEnsureFolderAlreadyExists(t *testing.T) {
	crm := &CloudResourceManager{}
	err := crm.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during cloudresourcemanager.Initialize() with blank credentials: %s", err)
	}
	setFoldersCallMockDefaults(crm)
	name, err := crm.EnsureFolder(testFolderName, "folders/000000000")
	if err != nil {
		t.Errorf("Got unexpected error during cloudresourcemanager.EnsureFolder() for folder that already exists: %s", err)
	}
	if name != testFolderName {
		t.Errorf("Got unexpected result from cloudresourcemanager.EnsureFolder() for folder that already exists: %s", err)
	}
}

func TestEnsureFolderDoesntExist(t *testing.T) {
	crm := &CloudResourceManager{}
	err := crm.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during cloudresourcemanager.Initialize() with blank credentials: %s", err)
	}
	setFoldersCallMockDefaults(crm)
	crm.Calls.FoldersSearch = &foldersSearchNoResultThenResult{}
	name, err := crm.EnsureFolder(testFolderName, "folders/11111111")
	if err != nil {
		t.Errorf("Got unexpected error during cloudresourcemanager.EnsureFolder() for folder that doesn't exist: %s", err)
	}
	if name != testFolderName {
		t.Errorf("Got unexpected result from cloudresourcemanager.EnsureFolder() for folder that doesn't exist: %s", err)
	}
}

func TestEnsureFolderRoles(t *testing.T) {
	crm := &CloudResourceManager{}
	err := crm.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error for cloudresourcemanager.Initialize() with blank credentials: %s", err)
	}
	setFoldersCallMockDefaults(crm)
	err = crm.EnsureFolderRoles(testFolderName, testMember, []string{testRole, "role2"})
	if err != nil {
		t.Errorf("Got unexpected error for cloudresourcemanager.EnsureFolderRoles(): %s", err)
	}
}

func TestEnsureFolderRolesAlternativeFolder(t *testing.T) {
	crm := &CloudResourceManager{}
	err := crm.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error for cloudresourcemanager.Initialize() with blank credentials: %s", err)
	}
	setFoldersCallMockDefaults(crm)
	err = crm.EnsureFolderRoles(strings.Replace(testFolderName, "folders/", "", 1), testMember, []string{testRole, "role2"})
	if err != nil {
		t.Errorf("Got unexpected error for cloudresourcemanager.EnsureFolderRoles() with alternative folder format: %s", err)
	}
}

func TestEnsureFolderRolesExistingMember(t *testing.T) {
	crm := &CloudResourceManager{}
	err := crm.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error for cloudresourcemanager.Initialize() with blank credentials: %s", err)
	}
	setFoldersCallMockDefaults(crm)
	crm.Calls.FoldersGetIAMPolicy = &foldersGetIAMPolicyExistingMemberMock{}
	err = crm.EnsureFolderRoles(testFolderName, testMember, []string{testRole, "role2"})
	if err != nil {
		t.Errorf("Got unexpected error for cloudresourcemanager.EnsureFolderRoles() with existing member: %s", err)
	}
}

func TestEnsureFolderRolesExistingRoleWithoutMember(t *testing.T) {
	crm := &CloudResourceManager{}
	err := crm.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error for cloudresourcemanager.Initialize() with blank credentials: %s", err)
	}
	setFoldersCallMockDefaults(crm)
	crm.Calls.FoldersGetIAMPolicy = &foldersGetIAMPolicyExistingRoleMock{}
	err = crm.EnsureFolderRoles(testFolderName, testMember, []string{testRole, "role2"})
	if err != nil {
		t.Errorf("Got unexpected error for cloudresourcemanager.EnsureFolderRoles() with existing role, but member not within: %s", err)
	}
}
