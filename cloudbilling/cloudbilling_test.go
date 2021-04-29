package cloudbilling

import (
	"strings"
	"testing"

	loggermock "github.com/rockholla/go-lib/mocks/custom-mocks/logger"
	v1 "google.golang.org/api/cloudbilling/v1"
	googleapi "google.golang.org/api/googleapi"
)

const (
	testRole               = "role1"
	testBillingAccountName = "billingAccounts/tests"
	testMember             = "test@test"
	testCredentials        = `{
  "client_id": "xxxxxxx.apps.googleusercontent.com",
  "client_secret": "xxxxxxxxxxxxxxx",
  "refresh_token": "xxxxxxxxx",
  "type": "authorized_user"
}`
)

type projectsUpdateBillingInfoMock struct{}
type billingAccountsGetIAMPolicyMock struct{}
type billingAccountsGetIAMPolicyExistingMemberMock struct{}
type billingAccountsGetIAMPolicyExistingRoleMock struct{}
type billingAccountsSetIAMPolicyMock struct{}

// Do is the mock for default projectsUpdateBillingInfoMock
func (c *projectsUpdateBillingInfoMock) Do(call *v1.ProjectsUpdateBillingInfoCall, opts ...googleapi.CallOption) (*v1.ProjectBillingInfo, error) {
	return &v1.ProjectBillingInfo{
		BillingAccountName: testBillingAccountName,
	}, nil
}

// Do is the mock for default billingAccountsGetIAMPolicy
func (c *billingAccountsGetIAMPolicyMock) Do(call *v1.BillingAccountsGetIamPolicyCall, opts ...googleapi.CallOption) (*v1.Policy, error) {
	return &v1.Policy{}, nil
}

// Do is the mock for billingAccountsGetIamPolicy that includes an existing member in role
func (c *billingAccountsGetIAMPolicyExistingMemberMock) Do(call *v1.BillingAccountsGetIamPolicyCall, opts ...googleapi.CallOption) (*v1.Policy, error) {
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

// Do is the mock for billingAccountsGetIAMPolicy that includes an existing member in role
func (c *billingAccountsGetIAMPolicyExistingRoleMock) Do(call *v1.BillingAccountsGetIamPolicyCall, opts ...googleapi.CallOption) (*v1.Policy, error) {
	var bindings []*v1.Binding
	bindings = append(bindings, &v1.Binding{
		Role:    testRole,
		Members: []string{},
	})
	return &v1.Policy{
		Bindings: bindings,
	}, nil
}

// Do is the mock for default billingAccountsSetIAMPolicy
func (c *billingAccountsSetIAMPolicyMock) Do(call *v1.BillingAccountsSetIamPolicyCall, opts ...googleapi.CallOption) (*v1.Policy, error) {
	return &v1.Policy{}, nil
}

func setCallMockDefaults(cb *CloudBilling) {
	cb.Calls = &Calls{
		ProjectsUpdateBillingInfo:   &projectsUpdateBillingInfoMock{},
		BillingAccountsGetIAMPolicy: &billingAccountsGetIAMPolicyMock{},
		BillingAccountsSetIAMPolicy: &billingAccountsSetIAMPolicyMock{},
	}
}

func TestInitialize(t *testing.T) {
	cb := &CloudBilling{}
	err := cb.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during cloudbilling.Initialize() with blank credentials: %s", err)
	}
	err = cb.Initialize(testCredentials, loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during cloudbilling.Initialize() with explicit credentials: %s", err)
	}
}

func TestSetProjectBillingAccount(t *testing.T) {
	cb := &CloudBilling{}
	err := cb.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error for cloudbilling.Initialize() with blank credentials: %s", err)
	}
	setCallMockDefaults(cb)
	billingAccountName, err := cb.SetProjectBillingAccount("000000", "0000000")
	if err != nil {
		t.Errorf("Got unexpected error during cloudbilling.SetProjectBillingAccount(): %s", err)
	}
	if billingAccountName != testBillingAccountName {
		t.Errorf("Got unexpected result/billing account name from cloudbilling.SetProjectBillingAccount(): %s", billingAccountName)
	}
}

func TestEnsureRoles(t *testing.T) {
	cb := &CloudBilling{}
	err := cb.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error for cloudbilling.Initialize() with blank credentials: %s", err)
	}
	setCallMockDefaults(cb)
	err = cb.EnsureRoles(testBillingAccountName, testMember, []string{testRole, "role2"})
	if err != nil {
		t.Errorf("Got unexpected error for cloudbilling.EnsureRoles(): %s", err)
	}
}

func TestEnsureRolesAlternativeBillingAccount(t *testing.T) {
	cb := &CloudBilling{}
	err := cb.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error for cloudbilling.Initialize() with blank credentials: %s", err)
	}
	setCallMockDefaults(cb)
	err = cb.EnsureRoles(strings.Replace(testBillingAccountName, "billingAccounts/", "", 1), testMember, []string{testRole, "role2"})
	if err != nil {
		t.Errorf("Got unexpected error for cloudbilling.EnsureRoles() with alternative billing account format: %s", err)
	}
}

func TestEnsureRolesExistingMember(t *testing.T) {
	cb := &CloudBilling{}
	err := cb.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error for cloudbilling.Initialize() with blank credentials: %s", err)
	}
	setCallMockDefaults(cb)
	cb.Calls.BillingAccountsGetIAMPolicy = &billingAccountsGetIAMPolicyExistingMemberMock{}
	err = cb.EnsureRoles(testBillingAccountName, testMember, []string{testRole, "role2"})
	if err != nil {
		t.Errorf("Got unexpected error for cloudbilling.EnsureRoles() with existing member: %s", err)
	}
}

func TestEnsureRolesExistingRoleWithoutMember(t *testing.T) {
	cb := &CloudBilling{}
	err := cb.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error for cloudbilling.Initialize() with blank credentials: %s", err)
	}
	setCallMockDefaults(cb)
	cb.Calls.BillingAccountsGetIAMPolicy = &billingAccountsGetIAMPolicyExistingRoleMock{}
	err = cb.EnsureRoles(testBillingAccountName, testMember, []string{testRole, "role2"})
	if err != nil {
		t.Errorf("Got unexpected error for cloudbilling.EnsureRoles() with existing role, but member not within: %s", err)
	}
}

func TestRemoveRolesExistingMemberRole(t *testing.T) {
	cb := &CloudBilling{}
	err := cb.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error for cloudbilling.Initialize() with blank credentials: %s", err)
	}
	setCallMockDefaults(cb)
	cb.Calls.BillingAccountsGetIAMPolicy = &billingAccountsGetIAMPolicyExistingMemberMock{}
	err = cb.RemoveRoles(testBillingAccountName, testMember, []string{testRole, "role2"})
	if err != nil {
		t.Errorf("Got unexpected error for cloudbilling.TestRemoveRolesExistingMemberRole(): %s", err)
	}
}

func TestRemoveRolesExistingRoleNoMember(t *testing.T) {
	cb := &CloudBilling{}
	err := cb.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error for cloudbilling.Initialize() with blank credentials: %s", err)
	}
	setCallMockDefaults(cb)
	cb.Calls.BillingAccountsGetIAMPolicy = &billingAccountsGetIAMPolicyExistingRoleMock{}
	err = cb.RemoveRoles(testBillingAccountName, testMember, []string{testRole, "role2"})
	if err != nil {
		t.Errorf("Got unexpected error for cloudbilling.TestRemoveRolesExistingRoleNoMember(): %s", err)
	}
}
