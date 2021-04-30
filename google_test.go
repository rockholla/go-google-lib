package google

import (
	"testing"

	loggermock "github.com/rockholla/go-lib/mocks/custom-mocks/logger"
)

const (
	testCredentials = `{
  "client_id": "xxxxxxx.apps.googleusercontent.com",
  "client_secret": "xxxxxxxxxxxxxxx",
  "refresh_token": "xxxxxxxxx",
  "type": "service_account"
}`
)

func TestSetCredentials(t *testing.T) {
	g := &Google{}
	g.Initialize("", loggermock.GetLogMock())
	g.Initialize("/some/path", loggermock.GetLogMock())
}

func TestGetCloudResourceManager(t *testing.T) {
	var err error
	g := &Google{}
	_, err = g.GetCloudResourceManager()
	if err != nil {
		t.Errorf("Got unexpected error from google.GetCloudResourceManager(): %s", err)
	}
	_, err = g.GetCloudResourceManager()
	if err != nil {
		t.Errorf("Got unexpected error from google.GetCloudResourceManager() second run: %s", err)
	}
}

func TestGetCloudBilling(t *testing.T) {
	var err error
	g := &Google{}
	_, err = g.GetCloudBilling()
	if err != nil {
		t.Errorf("Got unexpected error from google.GetCloudBilling(): %s", err)
	}
	_, err = g.GetCloudBilling()
	if err != nil {
		t.Errorf("Got unexpected error from google.GetCloudBilling() second run: %s", err)
	}
}

func TestGetIAM(t *testing.T) {
	var err error
	g := &Google{}
	_, err = g.GetIAM()
	if err != nil {
		t.Errorf("Got unexpected error from google.GetIAM(): %s", err)
	}
	_, err = g.GetIAM()
	if err != nil {
		t.Errorf("Got unexpected error from google.GetIAM() second run: %s", err)
	}
}

func TestGetDeploymentManager(t *testing.T) {
	var err error
	g := &Google{}
	_, err = g.GetDeploymentManager()
	if err != nil {
		t.Errorf("Got unexpected error from google.GetDeploymentManager(): %s", err)
	}
	_, err = g.GetDeploymentManager()
	if err != nil {
		t.Errorf("Got unexpected error from google.GetDeploymentManager() second run: %s", err)
	}
}

func TestGetStorage(t *testing.T) {
	var err error
	g := &Google{}
	_, err = g.GetStorage()
	if err != nil {
		t.Errorf("Got unexpected error from google.GetStorage(): %s", err)
	}
	_, err = g.GetStorage()
	if err != nil {
		t.Errorf("Got unexpected error from google.GetStorage() second run: %s", err)
	}
}

func TestGetCompute(t *testing.T) {
	var err error
	g := &Google{}
	_, err = g.GetCompute()
	if err != nil {
		t.Errorf("Got unexpected error from google.GetCompute(): %s", err)
	}
	_, err = g.GetCompute()
	if err != nil {
		t.Errorf("Got unexpected error from google.GetCompute() second run: %s", err)
	}
}

func TestGetDNS(t *testing.T) {
	var err error
	g := &Google{}
	_, err = g.GetDNS()
	if err != nil {
		t.Errorf("Got unexpected error from google.GetDNS(): %s", err)
	}
	_, err = g.GetDNS()
	if err != nil {
		t.Errorf("Got unexpected error from google.GetDNS() second run: %s", err)
	}
}

func TestGetCloudIdentity(t *testing.T) {
	var err error
	g := &Google{}
	_, err = g.GetCloudIdentity("")
	if err != nil {
		t.Errorf("Got unexpected error from google.GetCloudIdentity(): %s", err)
	}
	_, err = g.GetCloudIdentity("")
	if err != nil {
		t.Errorf("Got unexpected error from google.GetCloudIdentity() second run: %s", err)
	}
}

func TestGetCloudIdentityImpersonate(t *testing.T) {
	var err error
	g := &Google{}
	_, err = g.GetCloudIdentity("impersonate@sa")
	if err != nil {
		t.Errorf("Got unexpected error from google.TestGetCloudIdentityCustomCreds(): %s", err)
	}
	_, err = g.GetCloudIdentity("impersonate@sa")
	if err != nil {
		t.Errorf("Got unexpected error from google.TestGetCloudIdentityCustomCreds() second run: %s", err)
	}
}

func TestGetAdmin(t *testing.T) {
	var err error
	g := &Google{}
	g.Initialize("/other/credentials", loggermock.GetLogMock())
	_, err = g.GetAdmin(testCredentials, "go-google-lib.tests", "admin")
	if err != nil {
		t.Errorf("Got unexpected error from google.GetAdmin(): %s", err)
	}
	_, err = g.GetAdmin(testCredentials, "go-google-lib.tests", "admin")
	if err != nil {
		t.Errorf("Got unexpected error from google.GetAdmin() second run: %s", err)
	}
}
