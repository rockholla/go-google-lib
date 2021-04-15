package cloudresourcemanager

import (
	"testing"

	loggermock "github.com/rockholla/go-lib/mocks/custom-mocks/logger/logger"
)

const (
	testRole        = "role1"
	testMember      = "test@test"
	testCredentials = `{
  "client_id": "xxxxxxx.apps.googleusercontent.com",
  "client_secret": "xxxxxxxxxxxxxxx",
  "refresh_token": "xxxxxxxxx",
  "type": "authorized_user"
}`
)

func TestInitialize(t *testing.T) {
	crm := &CloudResourceManager{}
	err := crm.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during cloudresourcemanager.TestInitialize() with blank credentials: %s", err)
	}
	err = crm.Initialize(testCredentials, loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during cloudresourcemanager.Initialize() with explicit credentials: %s", err)
	}
}
