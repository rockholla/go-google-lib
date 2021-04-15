package cloudresourcemanager

import (
	"strings"
	"testing"

	loggermock "github.com/rockholla/go-lib/mocks/custom-mocks/logger"
	v1 "google.golang.org/api/cloudresourcemanager/v1"
	googleapi "google.golang.org/api/googleapi"
)

const (
	testOrganizationName = "organizations/0000000000000"
)

type organizationsGetIAMPolicyMock struct{}
type organizationsGetIAMPolicyExistingMemberMock struct{}
type organizationsGetIAMPolicyExistingRoleMock struct{}
type organizationsSetIAMPolicyMock struct{}

// Do is the mock for default organizationsGetIAMPolicyMock
func (c *organizationsGetIAMPolicyMock) Do(call *v1.OrganizationsGetIamPolicyCall, opts ...googleapi.CallOption) (*v1.Policy, error) {
	return &v1.Policy{}, nil
}

// Do is the mock for organizationsGetIAMPolicyMock that includes an existing member in role
func (c *organizationsGetIAMPolicyExistingMemberMock) Do(call *v1.OrganizationsGetIamPolicyCall, opts ...googleapi.CallOption) (*v1.Policy, error) {
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

// Do is the mock for organizationsGetIAMPolicyMock that includes an existing member in role
func (c *organizationsGetIAMPolicyExistingRoleMock) Do(call *v1.OrganizationsGetIamPolicyCall, opts ...googleapi.CallOption) (*v1.Policy, error) {
	var bindings []*v1.Binding
	bindings = append(bindings, &v1.Binding{
		Role:    testRole,
		Members: []string{},
	})
	return &v1.Policy{
		Bindings: bindings,
	}, nil
}

// Do is the mock for default organizationsSetIAMPolicyMock
func (c *organizationsSetIAMPolicyMock) Do(call *v1.OrganizationsSetIamPolicyCall, opts ...googleapi.CallOption) (*v1.Policy, error) {
	return &v1.Policy{}, nil
}

func setOrganizationsCallMockDefaults(crm *CloudResourceManager) {
	crm.Calls = &Calls{
		OrganizationsGetIAMPolicy: &organizationsGetIAMPolicyMock{},
		OrganizationsSetIAMPolicy: &organizationsSetIAMPolicyMock{},
	}
}

func TestEnsureOrganizationRoles(t *testing.T) {
	crm := &CloudResourceManager{}
	err := crm.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error for cloudresourcemanager.Initialize() with blank credentials: %s", err)
	}
	setOrganizationsCallMockDefaults(crm)
	err = crm.EnsureOrganizationRoles(testOrganizationName, testMember, []string{testRole, "role2"})
	if err != nil {
		t.Errorf("Got unexpected error for cloudresourcemanager.EnsureOrganizationRoles(): %s", err)
	}
}

func TestEnsureOrganizationRolesAlternativeOrganization(t *testing.T) {
	crm := &CloudResourceManager{}
	err := crm.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error for cloudresourcemanager.Initialize() with blank credentials: %s", err)
	}
	setOrganizationsCallMockDefaults(crm)
	err = crm.EnsureOrganizationRoles(strings.Replace(testOrganizationName, "organizations/", "", 1), testMember, []string{testRole, "role2"})
	if err != nil {
		t.Errorf("Got unexpected error for cloudresourcemanager.EnsureOrganizationRoles() with alternative organization format: %s", err)
	}
}

func TestEnsureOrganizationRolesExistingMember(t *testing.T) {
	crm := &CloudResourceManager{}
	err := crm.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error for cloudresourcemanager.Initialize() with blank credentials: %s", err)
	}
	setOrganizationsCallMockDefaults(crm)
	crm.Calls.OrganizationsGetIAMPolicy = &organizationsGetIAMPolicyExistingMemberMock{}
	err = crm.EnsureOrganizationRoles(testOrganizationName, testMember, []string{testRole, "role2"})
	if err != nil {
		t.Errorf("Got unexpected error for cloudresourcemanager.EnsureOrganizationRoles() with existing member: %s", err)
	}
}

func TestEnsureOrganizationRolesExistingRoleWithoutMember(t *testing.T) {
	crm := &CloudResourceManager{}
	err := crm.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error for cloudresourcemanager.Initialize() with blank credentials: %s", err)
	}
	setOrganizationsCallMockDefaults(crm)
	crm.Calls.OrganizationsGetIAMPolicy = &organizationsGetIAMPolicyExistingRoleMock{}
	err = crm.EnsureOrganizationRoles(testOrganizationName, testMember, []string{testRole, "role2"})
	if err != nil {
		t.Errorf("Got unexpected error for cloudresourcemanager.EnsureOrganizationRoles() with existing role, but member not within: %s", err)
	}
}
