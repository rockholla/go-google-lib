package admin

import (
	"testing"

	loggermock "github.com/rockholla/go-lib/mocks/custom-mocks/logger/logger"
	dirv1 "google.golang.org/api/admin/directory/v1"
	"google.golang.org/api/googleapi"
)

var (
	triggerGroupNotFound  = false
	triggerMemberNotFound = false
	testDomain            = "go-google-lib.tests"
	testAdminUsername     = "admin"
	testCredentials       = `{
  "client_id": "xxxxxxx.apps.googleusercontent.com",
  "client_secret": "xxxxxxxxxxxxxxx",
  "refresh_token": "xxxxxxxxx",
  "type": "service_account"
}`
)

var (
	testAPIGroup  = &dirv1.Group{}
	testAPIMember = &dirv1.Member{}
)

type groupsInsertMock struct{}
type groupsUpdateMock struct{}
type groupsGetMock struct{}
type groupsDeleteMock struct{}
type membersGetMock struct{}
type membersInsertMock struct{}
type membersDeleteMock struct{}

func (c *groupsInsertMock) Do(call *dirv1.GroupsInsertCall, opts ...googleapi.CallOption) (*dirv1.Group, error) {
	return testAPIGroup, nil
}

func (c *groupsUpdateMock) Do(call *dirv1.GroupsUpdateCall, opts ...googleapi.CallOption) (*dirv1.Group, error) {
	return testAPIGroup, nil
}

func (c *groupsGetMock) Do(call *dirv1.GroupsGetCall, opts ...googleapi.CallOption) (*dirv1.Group, error) {
	if triggerGroupNotFound {
		triggerGroupNotFound = false
		return nil, nil
	}
	return testAPIGroup, nil
}

func (c *groupsDeleteMock) Do(call *dirv1.GroupsDeleteCall, opts ...googleapi.CallOption) error {
	return nil
}

func (c *membersGetMock) Do(call *dirv1.MembersGetCall, opts ...googleapi.CallOption) (*dirv1.Member, error) {
	if triggerMemberNotFound {
		triggerMemberNotFound = false
		return nil, nil
	}
	return testAPIMember, nil
}

func (c *membersInsertMock) Do(call *dirv1.MembersInsertCall, opts ...googleapi.CallOption) (*dirv1.Member, error) {
	return testAPIMember, nil
}

func (c *membersDeleteMock) Do(call *dirv1.MembersDeleteCall, opts ...googleapi.CallOption) error {
	return nil
}

func setCallMockDefaults(a *Admin) {
	a.Calls = &Calls{
		GroupsInsert:  &groupsInsertMock{},
		GroupsUpdate:  &groupsUpdateMock{},
		GroupsGet:     &groupsGetMock{},
		GroupsDelete:  &groupsDeleteMock{},
		MembersGet:    &membersGetMock{},
		MembersInsert: &membersInsertMock{},
		MembersDelete: &membersDeleteMock{},
	}
}

func TestInitialize(t *testing.T) {
	a := &Admin{}
	err := a.Initialize("", testDomain, testAdminUsername, loggermock.GetLogMock())
	if err == nil {
		t.Errorf("Didn't get expected error during admin.Initialize() with blank credentials")
	}
	err = a.Initialize(testCredentials, testDomain, testAdminUsername, loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during admin.Initialize() with explicit credentials: %s", err)
	}
}

func TestEnsureGroupNew(t *testing.T) {
	a := &Admin{}
	err := a.Initialize(testCredentials, testDomain, testAdminUsername, loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error for admin.Initialize() with blank credentials: %s", err)
	}
	triggerGroupNotFound = true
	setCallMockDefaults(a)
	_, err = a.EnsureGroup("test-group", "test group")
	if err != nil {
		t.Errorf("Got unexpected error during admin.EnsureGroup(): %s", err)
	}
}

func TestEnsureGroupExists(t *testing.T) {
	a := &Admin{}
	err := a.Initialize(testCredentials, testDomain, testAdminUsername, loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error for admin.Initialize() with blank credentials: %s", err)
	}
	setCallMockDefaults(a)
	_, err = a.EnsureGroup("test-group", "test group")
	if err != nil {
		t.Errorf("Got unexpected error during admin.EnsureGroup(): %s", err)
	}
}

func TestDeleteGroup(t *testing.T) {
	a := &Admin{}
	err := a.Initialize(testCredentials, testDomain, testAdminUsername, loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error for admin.Initialize() with blank credentials: %s", err)
	}
	setCallMockDefaults(a)
	err = a.DeleteGroup("test-group")
	if err != nil {
		t.Errorf("Got unexpected error during admin.DeleteGroup(): %s", err)
	}
}

func TestEnsureMembershipNew(t *testing.T) {
	a := &Admin{}
	err := a.Initialize(testCredentials, testDomain, testAdminUsername, loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error for admin.Initialize() with blank credentials: %s", err)
	}
	triggerMemberNotFound = true
	setCallMockDefaults(a)
	_, err = a.EnsureMembership("test-group", "test-member")
	if err != nil {
		t.Errorf("Got unexpected error during admin.EnsureMembership(): %s", err)
	}
}

func TestEnsureMembershipExists(t *testing.T) {
	a := &Admin{}
	err := a.Initialize(testCredentials, testDomain, testAdminUsername, loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error for admin.Initialize() with blank credentials: %s", err)
	}
	setCallMockDefaults(a)
	_, err = a.EnsureMembership("test-group", "test-member")
	if err != nil {
		t.Errorf("Got unexpected error during admin.EnsureMembership(): %s", err)
	}
}

func TestDeleteMembership(t *testing.T) {
	a := &Admin{}
	err := a.Initialize(testCredentials, testDomain, testAdminUsername, loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error for admin.Initialize() with blank credentials: %s", err)
	}
	setCallMockDefaults(a)
	err = a.DeleteMembership("test-group", "test-member")
	if err != nil {
		t.Errorf("Got unexpected error during admin.DeleteMembership(): %s", err)
	}
}
