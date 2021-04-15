package calls

import (
	v1 "google.golang.org/api/cloudresourcemanager/v1"
	googleapi "google.golang.org/api/googleapi"
)

// OrganizationsGetIAMPolicyCallInterface is an interface to a call to get the iam policy for a organization
type OrganizationsGetIAMPolicyCallInterface interface {
	Do(call *v1.OrganizationsGetIamPolicyCall, opts ...googleapi.CallOption) (*v1.Policy, error)
}

// OrganizationsSetIAMPolicyCallInterface is an interface to a call to set the iam policy for a organization
type OrganizationsSetIAMPolicyCallInterface interface {
	Do(call *v1.OrganizationsSetIamPolicyCall, opts ...googleapi.CallOption) (*v1.Policy, error)
}

// OrganizationsGetIAMPolicyCall is the default implementation for OrganizationsGetIAMPolicyCallInterface
type OrganizationsGetIAMPolicyCall struct{}

// OrganizationsSetIAMPolicyCall is the default implementation for OrganizationsSetIAMPolicyCallInterface
type OrganizationsSetIAMPolicyCall struct{}

// Do performs the call, the default implementation of the interface
func (c *OrganizationsGetIAMPolicyCall) Do(call *v1.OrganizationsGetIamPolicyCall, opts ...googleapi.CallOption) (*v1.Policy, error) {
	return call.Do(opts...)
}

// Do performs the call, the default implementation of the interface
func (c *OrganizationsSetIAMPolicyCall) Do(call *v1.OrganizationsSetIamPolicyCall, opts ...googleapi.CallOption) (*v1.Policy, error) {
	return call.Do(opts...)
}
