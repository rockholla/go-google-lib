// Package cloudbilling is the library for google cloud billing operations
package cloudbilling

import (
	"context"
	"fmt"
	"regexp"

	"github.com/rockholla/go-google-lib/google/cloudbilling/calls"
	"github.com/rockholla/go-google-lib/logger"
	v1 "google.golang.org/api/cloudbilling/v1"
	"google.golang.org/api/option"
)

// Interface represents functionality for CloudBilling
type Interface interface {
	Initialize(credentials string, log logger.Interface) error
	SetProjectBillingAccount(projectID string, billingAccountID string) (string, error)
	EnsureRoles(billingAccount string, member string, roles []string) error
}

// CloudBilling wraps google-provided apis for interacting with google.golang.org/api/cloudbilling/*
type CloudBilling struct {
	log   logger.Interface
	V1    *v1.APIService
	Calls *Calls
}

// Calls are interfaces for making the actual calls to various underlying apis
type Calls struct {
	ProjectsUpdateBillingInfo   calls.ProjectsUpdateBillingInfoCallInterface
	BillingAccountsGetIAMPolicy calls.BillingAccountsGetIAMPolicyCallInterface
	BillingAccountsSetIAMPolicy calls.BillingAccountsSetIAMPolicyCallInterface
}

// Initialize sets up necessary google-provided sdks and other local data
func (cb *CloudBilling) Initialize(credentials string, log logger.Interface) error {
	var err error
	ctx := context.Background()
	cb.log = log
	cb.Calls = &Calls{
		ProjectsUpdateBillingInfo:   &calls.ProjectsUpdateBillingInfoCall{},
		BillingAccountsGetIAMPolicy: &calls.BillingAccountsGetIAMPolicyCall{},
		BillingAccountsSetIAMPolicy: &calls.BillingAccountsSetIAMPolicyCall{},
	}
	if credentials != "" {
		if cb.V1, err = v1.NewService(ctx, option.WithCredentialsJSON([]byte(credentials))); err != nil {
			return err
		}
	} else {
		if cb.V1, err = v1.NewService(ctx); err != nil {
			return err
		}
	}
	return nil
}

// SetProjectBillingAccount will update the billing account attached to a project, returns billing account name
func (cb *CloudBilling) SetProjectBillingAccount(projectID string, billingAccountID string) (string, error) {
	ctx := context.Background()
	cb.log.Info("Assigning billing account ID %s to project %s", billingAccountID, projectID)
	projectsService := v1.NewProjectsService(cb.V1)
	updateBillingInfoCall := projectsService.UpdateBillingInfo(fmt.Sprintf("projects/%s", projectID), &v1.ProjectBillingInfo{
		Name:               fmt.Sprintf("projects/%s/billlingInfo", projectID),
		BillingAccountName: fmt.Sprintf("billingAccounts/%s", billingAccountID),
		BillingEnabled:     true,
	}).Context(ctx)
	result, err := cb.Calls.ProjectsUpdateBillingInfo.Do(updateBillingInfoCall)
	if err != nil {
		return "", err
	}
	return result.BillingAccountName, nil
}

// EnsureRoles makes sure that a particular member has the supplied roles on the billing account
func (cb *CloudBilling) EnsureRoles(billingAccount string, member string, roles []string) error {
	ctx := context.Background()
	if matched, _ := regexp.Match("^billingAccounts\\/", []byte(billingAccount)); !matched {
		billingAccount = fmt.Sprintf("billingAccounts/%s", billingAccount)
	}
	cb.log.Info("Ensuring member %s has roles in %s:", member, billingAccount)
	billingAccountsService := v1.NewBillingAccountsService(cb.V1)
	billingAccountGetPolicyCall := billingAccountsService.GetIamPolicy(billingAccount).Context(ctx)
	policy, err := cb.Calls.BillingAccountsGetIAMPolicy.Do(billingAccountGetPolicyCall)
	if err != nil {
		return err
	}
	for _, role := range roles {
		cb.log.ListItem(role)
		binding := getExistingRoleBinding(policy.Bindings, role)
		if binding == nil {
			policy.Bindings = append(policy.Bindings, &v1.Binding{
				Members: []string{member},
				Role:    role,
			})
		} else if !memberExistsInRoleBinding(binding, member) {
			binding.Members = append(binding.Members, member)
		}
	}
	setIAMPolicyRequest := &v1.SetIamPolicyRequest{
		Policy: policy,
	}
	billingAccountSetPolicyCall := billingAccountsService.SetIamPolicy(billingAccount, setIAMPolicyRequest).Context(ctx)
	_, err = cb.Calls.BillingAccountsSetIAMPolicy.Do(billingAccountSetPolicyCall)
	if err != nil {
		return err
	}
	return nil
}

func getExistingRoleBinding(bindings []*v1.Binding, role string) *v1.Binding {
	for _, binding := range bindings {
		if binding.Role == role {
			return binding
		}
	}
	return nil
}

func memberExistsInRoleBinding(binding *v1.Binding, member string) bool {
	for _, bindingMember := range binding.Members {
		if bindingMember == member {
			return true
		}
	}
	return false
}
