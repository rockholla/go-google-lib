// Package cloudidentity is the library for google cloud identity operations
package cloudidentity

import (
	"context"
	"fmt"
	"strings"

	"github.com/rockholla/go-google-lib/cloudidentity/calls"
	"github.com/rockholla/go-lib/logger"
	v1beta1 "google.golang.org/api/cloudidentity/v1beta1"
	"google.golang.org/api/option"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Interface represents functionality for CloudBilling
type Interface interface {
	Initialize(impersonateServiceAccountEmail string, log logger.Interface) error
	EnsureGroup(name string, domain string, customerID string) (*v1beta1.Group, error)
}

// CloudIdentity wraps google-provided apis for interacting with google.golang.org/api/cloudbilling/*
type CloudIdentity struct {
	log     logger.Interface
	V1Beta1 *v1beta1.Service
	Calls   *Calls
}

// Calls are interfaces for making the actual calls to various underlying apis
type Calls struct {
	GroupCreate calls.GroupCreateCallInterface
	GroupLookup calls.GroupLookupCallInterface
}

// Initialize sets up necessary google-provided sdks and other local data
func (ci *CloudIdentity) Initialize(impersonateServiceAccountEmail string, log logger.Interface) error {
	var err error
	ctx := context.Background()
	ci.log = log
	ci.Calls = &Calls{
		GroupCreate: &calls.GroupCreateCall{},
		GroupLookup: &calls.GroupLookupCall{},
	}
	if impersonateServiceAccountEmail != "" {
		if ci.V1Beta1, err = v1beta1.NewService(ctx, option.ImpersonateCredentials(impersonateServiceAccountEmail)); err != nil {
			return err
		}
	} else {
		if ci.V1Beta1, err = v1beta1.NewService(ctx); err != nil {
			return err
		}
	}
	return nil
}

// EnsureGroup will make sure that a cloud identity group exists
func (ci *CloudIdentity) EnsureGroup(name string, domain string, customerID string) (*v1beta1.Group, error) {
	ctx := context.Background()
	groupsService := v1beta1.NewGroupsService(ci.V1Beta1)
	groupKeyID := fmt.Sprintf("%s@%s", name, domain)
	fullCustomerID := fmt.Sprintf("customers/%s", customerID)
	ci.log.InfoPart("Ensuring cloud identity group %s exists in %s...", groupKeyID, fullCustomerID)
	group := &v1beta1.Group{
		DisplayName: name,
		GroupKey: &v1beta1.EntityKey{
			Id: groupKeyID,
		},
		Parent: fullCustomerID,
		Labels: map[string]string{
			"cloudidentity.googleapis.com/groups.discussion_forum": "",
		},
	}
	groupCreateCall := groupsService.Create(group).Context(ctx).InitialGroupConfig("WITH_INITIAL_OWNER")
	if _, err := ci.Calls.GroupCreate.Do(groupCreateCall); err != nil {
		if s, ok := status.FromError(err); ok {
			if s.Code() != codes.AlreadyExists {
				return nil, err
			}
		} else if !strings.Contains(err.Error(), "alreadyExists") {
			return nil, err
		}
		ci.log.InfoPart("already exists\n")
	} else {
		ci.log.InfoPart("created\n")
	}
	groupLookupCall := groupsService.Lookup().Context(ctx).GroupKeyId(groupKeyID)
	lookupResponse, err := ci.Calls.GroupLookup.Do(groupLookupCall)
	if err != nil {
		return nil, err
	}
	group.Name = lookupResponse.Name
	return group, nil
}
