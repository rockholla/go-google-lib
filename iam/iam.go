// Package iam is the library for google cloud iam operations
package iam

import (
	"context"
	"fmt"
	"strings"

	adminv1 "cloud.google.com/go/iam/admin/apiv1"
	"github.com/rockholla/go-google-lib/logger"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	adminpb "google.golang.org/genproto/googleapis/iam/admin/v1"
)

// Interface represents functionality for IAM
type Interface interface {
	Initialize(credentials string, log logger.Interface) error
	EnsureServiceAccount(projectID string, serviceAccount *ServiceAccount, createNewKey bool) error
	DeleteServiceAccount(projectID string, serviceAccountName string) error
}

// AdminV1 is an interface for the underlying IAM sdk/library for api interaction
type AdminV1 interface {
	GetServiceAccount(ctx context.Context, req *adminpb.GetServiceAccountRequest, opts ...gax.CallOption) (*adminpb.ServiceAccount, error)
	CreateServiceAccount(ctx context.Context, req *adminpb.CreateServiceAccountRequest, opts ...gax.CallOption) (*adminpb.ServiceAccount, error)
	CreateServiceAccountKey(ctx context.Context, req *adminpb.CreateServiceAccountKeyRequest, opts ...gax.CallOption) (*adminpb.ServiceAccountKey, error)
	DeleteServiceAccount(ctx context.Context, req *adminpb.DeleteServiceAccountRequest, opts ...gax.CallOption) error
}

// IAM wraps google-provided apis for interacting with cloud.google.com/go/iam/*
type IAM struct {
	log     logger.Interface
	AdminV1 AdminV1
}

// ServiceAccount is an object representing a service account
type ServiceAccount struct {
	Name  string
	Email string
	Key   string
}

// Initialize sets up necessary google-provided sdks and other local data
func (iam *IAM) Initialize(credentials string, log logger.Interface) error {
	var err error
	ctx := context.Background()
	iam.log = log
	if credentials != "" {
		if iam.AdminV1, err = adminv1.NewIamClient(ctx, option.WithCredentialsJSON([]byte(credentials))); err != nil {
			return err
		}
	} else {
		if iam.AdminV1, err = adminv1.NewIamClient(ctx); err != nil {
			return err
		}
	}
	return nil
}

// EnsureServiceAccount will make sure that a service account and key exists for a particular service account name
// in the specified project ID. You can also instruct to force create a new/additional key if one already exists
func (iam *IAM) EnsureServiceAccount(projectID string, serviceAccount *ServiceAccount, createNewKey bool) error {
	serviceAccount.setEmail(projectID)
	ctx := context.Background()
	createServiceAccount := false
	iam.log.Info(`Ensuring that service account %s exists in project %s`, serviceAccount.Name, projectID)
	getServiceAccountRequest := &adminpb.GetServiceAccountRequest{
		Name: fmt.Sprintf("projects/%s/serviceAccounts/%s", projectID, serviceAccount.Email),
	}
	_, err := iam.AdminV1.GetServiceAccount(ctx, getServiceAccountRequest)
	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "notfound") {
			createServiceAccount = true
		} else {
			return err
		}
	}
	if createServiceAccount {
		createServiceAccountRequest := &adminpb.CreateServiceAccountRequest{
			Name:      fmt.Sprintf("projects/%s", projectID),
			AccountId: serviceAccount.Name,
			ServiceAccount: &adminpb.ServiceAccount{
				DisplayName: serviceAccount.Name,
			},
		}
		_, err := iam.AdminV1.CreateServiceAccount(ctx, createServiceAccountRequest)
		if err != nil {
			return err
		}
	}
	if createServiceAccount || createNewKey {
		createServiceAccountKeyRequest := &adminpb.CreateServiceAccountKeyRequest{
			Name: fmt.Sprintf("projects/%s/serviceAccounts/%s", projectID, serviceAccount.Email),
		}
		serviceAccountKey, err := iam.AdminV1.CreateServiceAccountKey(ctx, createServiceAccountKeyRequest)
		if err != nil {
			return err
		}
		serviceAccount.Key = string(serviceAccountKey.PrivateKeyData)
		return nil
	}
	return nil
}

// DeleteServiceAccount will remove a service account from a project
func (iam *IAM) DeleteServiceAccount(projectID string, serviceAccountName string) error {
	serviceAccount := &ServiceAccount{
		Name: serviceAccountName,
	}
	serviceAccount.setEmail(projectID)
	ctx := context.Background()
	iam.log.Info(`Deleting service account %s in project %s`, serviceAccountName, projectID)
	deleteServiceAccountRequest := &adminpb.DeleteServiceAccountRequest{
		Name: fmt.Sprintf("projects/%s/serviceAccounts/%s", projectID, serviceAccount.Email),
	}
	err := iam.AdminV1.DeleteServiceAccount(ctx, deleteServiceAccountRequest)
	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "notfound") {
			return nil
		}
		return err
	}
	return nil
}

func (serviceAccount *ServiceAccount) setEmail(projectID string) {
	serviceAccount.Email = fmt.Sprintf("%s@%s.iam.gserviceaccount.com", serviceAccount.Name, projectID)
}
