// Package cloudresourcemanager is the library for google cloud resource manager operations
package cloudresourcemanager

import (
	"context"

	"github.com/rockholla/go-google-lib/cloudresourcemanager/calls"
	"github.com/rockholla/go-lib/logger"
	v1 "google.golang.org/api/cloudresourcemanager/v1"
	v2beta1 "google.golang.org/api/cloudresourcemanager/v2beta1"
	"google.golang.org/api/option"
	suv1 "google.golang.org/api/serviceusage/v1"
)

// Interface represents functionality for CloudResourceManager
type Interface interface {
	Initialize(credentials string, log logger.Interface) error
	GetFolder(displayName string, parent string) (string, error)
	EnsureFolder(displayName string, parent string) (string, error)
	EnsureFolderRoles(folder string, member string, roles []string) error
	SetFolderOrgPolicy(folder string, policy *v1.OrgPolicy) error
	GetProject(name string, parent string) (*v1.Project, error)
	DeleteProject(id string) error
	GetProjectByID(id string) (*v1.Project, error)
	EnsureProject(name string, parent string) (string, int64, error)
	EnableProjectServices(projectID string, services []string) error
	EnsureProjectRoles(project string, member string, roles []string) error
	EnsureOrganizationRoles(organization string, member string, roles []string) error
}

// CloudResourceManager wraps google-provided apis for interacting with google.golang.org/api/cloudresourcemanager/*
type CloudResourceManager struct {
	log     logger.Interface
	V1      *v1.Service
	V2Beta1 *v2beta1.Service
	SUV1    *suv1.Service
	Calls   *Calls
}

// Calls are interfaces for making the actual calls to various underlying apis
type Calls struct {
	FoldersSearch             calls.FoldersSearchCallInterface
	FoldersCreate             calls.FoldersCreateCallInterface
	FoldersGetIAMPolicy       calls.FoldersGetIAMPolicyCallInterface
	FoldersSetIAMPolicy       calls.FoldersSetIAMPolicyCallInterface
	FoldersSetOrgPolicy       calls.FoldersSetOrgPolicyCallInterface
	ProjectsList              calls.ProjectsListCallInterface
	ProjectsGet               calls.ProjectsGetCallInterface
	ProjectsCreate            calls.ProjectsCreateCallInterface
	ProjectsDelete            calls.ProjectsDeleteCallInterface
	ProjectsGetIAMPolicy      calls.ProjectsGetIAMPolicyCallInterface
	ProjectsSetIAMPolicy      calls.ProjectsSetIAMPolicyCallInterface
	ServiceEnable             calls.ServiceEnableCallInterface
	OrganizationsGetIAMPolicy calls.OrganizationsGetIAMPolicyCallInterface
	OrganizationsSetIAMPolicy calls.OrganizationsSetIAMPolicyCallInterface
}

// Initialize sets up necessary google-provided sdks and other local data
func (crm *CloudResourceManager) Initialize(credentials string, log logger.Interface) error {
	var err error
	ctx := context.Background()
	crm.log = log
	crm.Calls = &Calls{
		FoldersSearch:             &calls.FoldersSearchCall{},
		FoldersCreate:             &calls.FoldersCreateCall{},
		FoldersGetIAMPolicy:       &calls.FoldersGetIAMPolicyCall{},
		FoldersSetIAMPolicy:       &calls.FoldersSetIAMPolicyCall{},
		FoldersSetOrgPolicy:       &calls.FoldersSetOrgPolicyCall{},
		ProjectsList:              &calls.ProjectsListCall{},
		ProjectsGet:               &calls.ProjectsGetCall{},
		ProjectsCreate:            &calls.ProjectsCreateCall{},
		ProjectsDelete:            &calls.ProjectsDeleteCall{},
		ProjectsGetIAMPolicy:      &calls.ProjectsGetIAMPolicyCall{},
		ProjectsSetIAMPolicy:      &calls.ProjectsSetIAMPolicyCall{},
		ServiceEnable:             &calls.ServiceEnableCall{},
		OrganizationsGetIAMPolicy: &calls.OrganizationsGetIAMPolicyCall{},
		OrganizationsSetIAMPolicy: &calls.OrganizationsSetIAMPolicyCall{},
	}
	if credentials != "" {
		if crm.V1, err = v1.NewService(ctx, option.WithCredentialsJSON([]byte(credentials))); err != nil {
			return err
		}
		if crm.V2Beta1, err = v2beta1.NewService(ctx, option.WithCredentialsJSON([]byte(credentials))); err != nil {
			return err
		}
		if crm.SUV1, err = suv1.NewService(ctx, option.WithCredentialsJSON([]byte(credentials))); err != nil {
			return err
		}
	} else {
		if crm.V1, err = v1.NewService(ctx); err != nil {
			return err
		}
		if crm.V2Beta1, err = v2beta1.NewService(ctx); err != nil {
			return err
		}
		if crm.SUV1, err = suv1.NewService(ctx); err != nil {
			return err
		}
	}
	return nil
}
