package cloudresourcemanager

import (
	"context"
	"fmt"
	"regexp"

	v1 "google.golang.org/api/cloudresourcemanager/v1"
)

// EnsureOrganizationRoles makes sure that a particular member has the supplied roles on the organization
func (crm *CloudResourceManager) EnsureOrganizationRoles(organization string, member string, roles []string) error {
	ctx := context.Background()
	if matched, _ := regexp.Match("^organizations\\/", []byte(organization)); !matched {
		organization = fmt.Sprintf("organizations/%s", organization)
	}
	crm.log.Info("Ensuring member %s has roles in %s:", member, organization)
	organizationsService := v1.NewOrganizationsService(crm.V1)
	organizationPolicyGetCall := organizationsService.GetIamPolicy(organization, &v1.GetIamPolicyRequest{}).Context(ctx)
	policy, err := crm.Calls.OrganizationsGetIAMPolicy.Do(organizationPolicyGetCall)
	if err != nil {
		return err
	}
	for _, role := range roles {
		crm.log.ListItem(role)
		binding := getExistingOrganizationRoleBinding(policy.Bindings, role)
		if binding == nil {
			policy.Bindings = append(policy.Bindings, &v1.Binding{
				Members: []string{member},
				Role:    role,
			})
		} else if !memberExistsInOrganizationRoleBinding(binding, member) {
			binding.Members = append(binding.Members, member)
		}
	}
	setIAMPolicyRequest := &v1.SetIamPolicyRequest{
		Policy: policy,
	}
	organizationSetPolicyCall := organizationsService.SetIamPolicy(organization, setIAMPolicyRequest).Context(ctx)
	_, err = crm.Calls.OrganizationsSetIAMPolicy.Do(organizationSetPolicyCall)
	if err != nil {
		return err
	}
	return nil
}

func getExistingOrganizationRoleBinding(bindings []*v1.Binding, role string) *v1.Binding {
	for _, binding := range bindings {
		if binding.Role == role {
			return binding
		}
	}
	return nil
}

func memberExistsInOrganizationRoleBinding(binding *v1.Binding, member string) bool {
	for _, bindingMember := range binding.Members {
		if bindingMember == member {
			return true
		}
	}
	return false
}
