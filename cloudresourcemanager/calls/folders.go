// Package calls are mockable remote calls for operations
package calls

import (
	v1 "google.golang.org/api/cloudresourcemanager/v1"
	v2beta1 "google.golang.org/api/cloudresourcemanager/v2beta1"
	googleapi "google.golang.org/api/googleapi"
)

// FoldersSearchCallInterface is an interface to a call to search for a folder
type FoldersSearchCallInterface interface {
	Do(call *v2beta1.FoldersSearchCall, opts ...googleapi.CallOption) (*v2beta1.SearchFoldersResponse, error)
}

// FoldersCreateCallInterface is an interface to a call to create a folder
type FoldersCreateCallInterface interface {
	Do(call *v2beta1.FoldersCreateCall, opts ...googleapi.CallOption) (*v2beta1.Operation, error)
}

// FoldersGetIAMPolicyCallInterface is an interface to a call to get the iam policy for a folder
type FoldersGetIAMPolicyCallInterface interface {
	Do(call *v2beta1.FoldersGetIamPolicyCall, opts ...googleapi.CallOption) (*v2beta1.Policy, error)
}

// FoldersSetIAMPolicyCallInterface is an interface to a call to set the iam policy for a folder
type FoldersSetIAMPolicyCallInterface interface {
	Do(call *v2beta1.FoldersSetIamPolicyCall, opts ...googleapi.CallOption) (*v2beta1.Policy, error)
}

// FoldersSetOrgPolicyCallInterface is an interface to a call to set an org policy constraint on a folder
type FoldersSetOrgPolicyCallInterface interface {
	Do(call *v1.FoldersSetOrgPolicyCall, opts ...googleapi.CallOption) (*v1.OrgPolicy, error)
}

// FoldersSearchCall is the default implementation for FoldersSearchCallInterface
type FoldersSearchCall struct{}

// FoldersCreateCall is the default implementation for FoldersCreateCallInterface
type FoldersCreateCall struct{}

// FoldersGetIAMPolicyCall is the default implementation for FoldersGetIAMPolicyCallInterface
type FoldersGetIAMPolicyCall struct{}

// FoldersSetIAMPolicyCall is the default implementation for FoldersSetIAMPolicyCallInterface
type FoldersSetIAMPolicyCall struct{}

// FoldersSetOrgPolicyCall is the default implementation for FoldersSetOrgPolicyCallInterface
type FoldersSetOrgPolicyCall struct{}

// Do performs the call, the default implementation of the interface
func (c *FoldersSearchCall) Do(call *v2beta1.FoldersSearchCall, opts ...googleapi.CallOption) (*v2beta1.SearchFoldersResponse, error) {
	return call.Do(opts...)
}

// Do performs the call, the default implementation of the interface
func (c *FoldersCreateCall) Do(call *v2beta1.FoldersCreateCall, opts ...googleapi.CallOption) (*v2beta1.Operation, error) {
	return call.Do(opts...)
}

// Do performs the call, the default implementation of the interface
func (c *FoldersGetIAMPolicyCall) Do(call *v2beta1.FoldersGetIamPolicyCall, opts ...googleapi.CallOption) (*v2beta1.Policy, error) {
	return call.Do(opts...)
}

// Do performs the call, the default implementation of the interface
func (c *FoldersSetIAMPolicyCall) Do(call *v2beta1.FoldersSetIamPolicyCall, opts ...googleapi.CallOption) (*v2beta1.Policy, error) {
	return call.Do(opts...)
}

// Do performs the call, the default implementation of the interface
func (c *FoldersSetOrgPolicyCall) Do(call *v1.FoldersSetOrgPolicyCall, opts ...googleapi.CallOption) (*v1.OrgPolicy, error) {
	return call.Do(opts...)
}
