// Package cloudidentity is the library for google cloud identity operations
package cloudidentity

import (
	"context"
	"fmt"
	"strings"

	"github.com/rockholla/go-google-lib/cloudidentity/calls"
	"github.com/rockholla/go-lib/logger"
	v1 "google.golang.org/api/cloudidentity/v1beta1"
	v1beta1 "google.golang.org/api/cloudidentity/v1beta1"
	"google.golang.org/api/option"
)

// Interface represents functionality for CloudBilling
type Interface interface {
	Initialize(impersonateServiceAccountEmail string, log logger.Interface) error
	EnsureGroup(name string, domain string, customerID string) error
}

// CloudIdentity wraps google-provided apis for interacting with google.golang.org/api/cloudbilling/*
type CloudIdentity struct {
	log     logger.Interface
	V1Beta1 *v1beta1.Service
	Calls   *Calls
}

// Calls are interfaces for making the actual calls to various underlying apis
type Calls struct {
	GroupGet    calls.GroupGetCallInterface
	GroupCreate calls.GroupCreateCallInterface
}

// Initialize sets up necessary google-provided sdks and other local data
func (ci *CloudIdentity) Initialize(impersonateServiceAccountEmail string, log logger.Interface) error {
	var err error
	ctx := context.Background()
	ci.log = log
	ci.Calls = &Calls{
		GroupGet:    &calls.GroupGetCall{},
		GroupCreate: &calls.GroupCreateCall{},
	}
	if impersonateServiceAccountEmail != "" {
		if ci.V1Beta1, err = v1beta1.NewService(ctx, option.ImpersonateCredentials(impersonateServiceAccountEmail)); err != nil {
			return err
		}
	} else {
		if ci.V1Beta1, err = v1.NewService(ctx); err != nil {
			return err
		}
	}
	return nil
}

// EnsureGroup will make sure that a cloud identity group exists
func (ci *CloudIdentity) EnsureGroup(name string, domain string, customerID string) error {
	ctx := context.Background()
	groupsService := v1beta1.NewGroupsService(ci.V1Beta1)
	groupKeyID := fmt.Sprintf("%s@%s", name, domain)
	fullCustomerID := fmt.Sprintf("customers/%s", customerID)
	ci.log.Info("Ensuring cloud identity %s in %s:", groupKeyID, fullCustomerID)
	groupGetCall := groupsService.Get(name).Context(ctx)
	_, err := ci.Calls.GroupGet.Do(groupGetCall)
	if err != nil {
		if !strings.Contains(err.Error(), "was not found") {
			return err
		}
	}
	groupCreateCall := groupsService.Create(&v1beta1.Group{
		DisplayName: name,
		GroupKey: &v1beta1.EntityKey{
			Id: groupKeyID,
		},
		Parent: fullCustomerID,
	}).Context(ctx)
	if _, err = ci.Calls.GroupCreate.Do(groupCreateCall); err != nil {
		return err
	}
	return nil
}
