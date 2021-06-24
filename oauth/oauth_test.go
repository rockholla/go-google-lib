package oauth

import (
	"testing"

	loggermock "github.com/rockholla/go-lib/mocks/custom-mocks/logger"
)

var (
	testCredentials = `{
  "client_id": "xxxxxxx.apps.googleusercontent.com",
  "client_secret": "xxxxxxxxxxxxxxx",
  "refresh_token": "xxxxxxxxx",
  "type": "authorized_user"
}`
)

func TestInitialize(t *testing.T) {
	o := &OAuth{}
	err := o.Initialize("", loggermock.GetLogMock(), []string{})
	if err != nil {
		t.Errorf("Got unexpected error during compute.Initialize() with blank credentials: %s", err)
	}
	err = o.Initialize(testCredentials, loggermock.GetLogMock(), []string{})
	if err != nil {
		t.Errorf("Got unexpected error during compute.Initialize() with explicit credentials: %s", err)
	}
}
