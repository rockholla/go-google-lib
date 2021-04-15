// Package calls are mockable remote calls for operations
package calls

import (
	v1 "google.golang.org/api/cloudbilling/v1"
	googleapi "google.golang.org/api/googleapi"
)

// BillingAccountsGetIAMPolicyCallInterface is an interface to a call to get the IAM policy for a billing account
type BillingAccountsGetIAMPolicyCallInterface interface {
	Do(call *v1.BillingAccountsGetIamPolicyCall, opts ...googleapi.CallOption) (*v1.Policy, error)
}

// BillingAccountsSetIAMPolicyCallInterface is an interface to a call to set the IAM policy for a billing account
type BillingAccountsSetIAMPolicyCallInterface interface {
	Do(call *v1.BillingAccountsSetIamPolicyCall, opts ...googleapi.CallOption) (*v1.Policy, error)
}

// BillingAccountsGetIAMPolicyCall is the default implementation for BillingAccountsGetIAMPolicyCallInterface
type BillingAccountsGetIAMPolicyCall struct{}

// BillingAccountsSetIAMPolicyCall is the default implementation for BillingAccountsSetIAMPolicyCallInterface
type BillingAccountsSetIAMPolicyCall struct{}

// Do performs the call, the default implementation of the interface
func (c *BillingAccountsGetIAMPolicyCall) Do(call *v1.BillingAccountsGetIamPolicyCall, opts ...googleapi.CallOption) (*v1.Policy, error) {
	return call.Do(opts...)
}

// Do performs the call, the default implementation of the interface
func (c *BillingAccountsSetIAMPolicyCall) Do(call *v1.BillingAccountsSetIamPolicyCall, opts ...googleapi.CallOption) (*v1.Policy, error) {
	return call.Do(opts...)
}
