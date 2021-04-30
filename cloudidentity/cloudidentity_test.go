package cloudidentity

import (
	"errors"
	"testing"

	loggermock "github.com/rockholla/go-lib/mocks/custom-mocks/logger"
	v1beta1 "google.golang.org/api/cloudidentity/v1beta1"
	googleapi "google.golang.org/api/googleapi"
)

var triggerGroupNotFound = false

type groupGetMock struct{}
type groupCreateMock struct{}

func (c *groupGetMock) Do(call *v1beta1.GroupsGetCall, opts ...googleapi.CallOption) (*v1beta1.Group, error) {
	if triggerGroupNotFound {
		triggerGroupNotFound = false
		return nil, errors.New("was not found")
	}
	return &v1beta1.Group{}, nil
}

func (c *groupCreateMock) Do(call *v1beta1.GroupsCreateCall, opts ...googleapi.CallOption) (*v1beta1.Operation, error) {
	return &v1beta1.Operation{}, nil
}

func setCallMockDefaults(ci *CloudIdentity) {
	ci.Calls = &Calls{
		GroupGet:    &groupGetMock{},
		GroupCreate: &groupCreateMock{},
	}
}

func TestInitialize(t *testing.T) {
	ci := &CloudIdentity{}
	err := ci.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during cloudidentity.Initialize() with blank credentials: %s", err)
	}
	err = ci.Initialize("impersonate@sa", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during cloudidentity.Initialize() with explicit credentials: %s", err)
	}
}

func TestEnsureGroupNotExists(t *testing.T) {
	ci := &CloudIdentity{}
	err := ci.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error for cloudidentity.Initialize() with blank credentials: %s", err)
	}
	setCallMockDefaults(ci)
	triggerGroupNotFound = true
	err = ci.EnsureGroup("name", "domain", "customer-id")
	if err != nil {
		t.Errorf("Got unexpected error during cloudidentity.EnsureGroup(): %s", err)
	}
}
