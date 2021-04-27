package iam

import (
	"context"
	"errors"
	"fmt"
	"testing"

	gax "github.com/googleapis/gax-go/v2"
	loggermock "github.com/rockholla/go-lib/mocks/custom-mocks/logger"
	adminpb "google.golang.org/genproto/googleapis/iam/admin/v1"
)

const (
	testServiceAccountName           = "test-sa"
	testProjectID                    = "000000000"
	testServiceAccountEmail          = "test-sa@test"
	testServiceAccountFullName       = "projects/000000000/serviceAccounts/test-sa@test"
	testServiceAccountKeyID          = "key-id"
	testServiceAccountKeyPrivateData = "private-key"
	testCredentials                  = `{
  "client_id": "xxxxxxx.apps.googleusercontent.com",
  "client_secret": "xxxxxxxxxxxxxxx",
  "refresh_token": "xxxxxxxxx",
  "type": "authorized_user"
}`
)

var (
	triggerNotFound = false
)

type adminV1Mock struct{}

func (mock *adminV1Mock) GetServiceAccount(ctx context.Context, req *adminpb.GetServiceAccountRequest, opts ...gax.CallOption) (*adminpb.ServiceAccount, error) {
	if triggerNotFound {
		triggerNotFound = false
		return &adminpb.ServiceAccount{}, errors.New("notfound")
	}
	return &adminpb.ServiceAccount{
		Name:        testServiceAccountFullName,
		ProjectId:   testProjectID,
		DisplayName: testServiceAccountName,
		Email:       testServiceAccountEmail,
	}, nil
}

func (mock *adminV1Mock) CreateServiceAccount(ctx context.Context, req *adminpb.CreateServiceAccountRequest, opts ...gax.CallOption) (*adminpb.ServiceAccount, error) {
	return &adminpb.ServiceAccount{
		Name:        testServiceAccountFullName,
		ProjectId:   testProjectID,
		DisplayName: testServiceAccountName,
		Email:       testServiceAccountEmail,
	}, nil
}

func (mock *adminV1Mock) CreateServiceAccountKey(ctx context.Context, req *adminpb.CreateServiceAccountKeyRequest, opts ...gax.CallOption) (*adminpb.ServiceAccountKey, error) {
	return &adminpb.ServiceAccountKey{
		PrivateKeyData: []byte(testServiceAccountKeyPrivateData),
		Name:           fmt.Sprintf("%s/keys/%s", testServiceAccountFullName, testServiceAccountKeyID),
	}, nil
}

func (mock *adminV1Mock) DeleteServiceAccount(ctx context.Context, req *adminpb.DeleteServiceAccountRequest, opts ...gax.CallOption) error {
	if triggerNotFound {
		triggerNotFound = false
		return errors.New("notfound")
	}
	return nil
}

func setMocks(iam *IAM) {
	iam.AdminV1 = &adminV1Mock{}
}

func TestInitialize(t *testing.T) {
	iam := &IAM{}
	err := iam.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during iam.Initialize() with blank credentials: %s", err)
	}
	err = iam.Initialize(testCredentials, loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during iam.Initialize() with explicit credentials: %s", err)
	}
}

func TestEnsureServiceAccount(t *testing.T) {
	iam := &IAM{}
	err := iam.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during iam.Initialize() with blank credentials: %s", err)
	}
	setMocks(iam)
	serviceAccount := &ServiceAccount{
		Name: testServiceAccountName,
	}
	err = iam.EnsureServiceAccount(testProjectID, serviceAccount, false)
	if err != nil {
		t.Errorf("Got unexpected error during iam.EnsureServiceAccount() for sa that exists, not force creating a key: %s", err)
	}
	if serviceAccount.Name != testServiceAccountName {
		t.Errorf("Expecting result sa display name from iam.EnsureServiceAccount() for sa that exists, not force creating a key to be \"%s\", but got: %s",
			testServiceAccountName, serviceAccount.Name)
	}
	if serviceAccount.Key != "" {
		t.Errorf("Expecting result key from iam.EnsureServiceAccount() for sa that exists, not force creating a key to be blank, but got: %s",
			serviceAccount.Key)
	}
}

func TestEnsureServiceAccountNotFound(t *testing.T) {
	iam := &IAM{}
	err := iam.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during iam.Initialize() with blank credentials: %s", err)
	}
	setMocks(iam)
	triggerNotFound = true
	serviceAccount := &ServiceAccount{
		Name: testServiceAccountName,
	}
	err = iam.EnsureServiceAccount(testProjectID, serviceAccount, false)
	if err != nil {
		t.Errorf("Got unexpected error during iam.EnsureServiceAccount() for sa that doesn't exist: %s", err)
	}
	if serviceAccount.Name != testServiceAccountName {
		t.Errorf("Expecting result sa display name from iam.EnsureServiceAccount() for sa that doesn't exist to be \"%s\", but got: %s",
			testServiceAccountName, serviceAccount.Name)
	}
	if serviceAccount.Key != "" {
		t.Errorf("Expecting empty result key from iam.EnsureServiceAccount() for new service account when createNewKey is false, but got: %s",
			serviceAccount.Key)
	}
}

func TestEnsureServiceAccountNotFoundNewKey(t *testing.T) {
	iam := &IAM{}
	err := iam.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during iam.Initialize() with blank credentials: %s", err)
	}
	setMocks(iam)
	triggerNotFound = true
	serviceAccount := &ServiceAccount{
		Name: testServiceAccountName,
	}
	err = iam.EnsureServiceAccount(testProjectID, serviceAccount, true)
	if err != nil {
		t.Errorf("Got unexpected error during iam.EnsureServiceAccount() for sa that doesn't exist: %s", err)
	}
	if serviceAccount.Name != testServiceAccountName {
		t.Errorf("Expecting result sa display name from iam.EnsureServiceAccount() for sa that doesn't exist to be \"%s\", but got: %s",
			testServiceAccountName, serviceAccount.Name)
	}
	if serviceAccount.Key != testServiceAccountKeyPrivateData {
		t.Errorf("Expecting result key from iam.EnsureServiceAccount() for sa that doesn't exist and createNewKey == true to be \"%s\", but got: %s",
			testServiceAccountKeyPrivateData,
			serviceAccount.Key)
	}
}

func TestDeleteServiceAccount(t *testing.T) {
	iam := &IAM{}
	err := iam.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during iam.Initialize() with blank credentials: %s", err)
	}
	setMocks(iam)
	err = iam.DeleteServiceAccount("project", "service-account")
	if err != nil {
		t.Errorf("Got unexpected error from DeleteServiceAccount(): %s", err)
	}
}

func TestDeleteServiceAccountNotFound(t *testing.T) {
	iam := &IAM{}
	err := iam.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during iam.Initialize() with blank credentials: %s", err)
	}
	setMocks(iam)
	triggerNotFound = true
	err = iam.DeleteServiceAccount("project", "service-account")
	if err != nil {
		t.Errorf("Got unexpected error from DeleteServiceAccount(): %s", err)
	}
}
