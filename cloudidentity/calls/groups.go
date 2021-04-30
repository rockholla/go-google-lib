package calls

import (
	v1beta1 "google.golang.org/api/cloudidentity/v1beta1"
	googleapi "google.golang.org/api/googleapi"
)

// GroupCreateCallInterface is an interface to a call to create a cloud identity group
type GroupCreateCallInterface interface {
	Do(call *v1beta1.GroupsCreateCall, opts ...googleapi.CallOption) (*v1beta1.Operation, error)
}

// GroupLookupCallInterface is an interface to a call to create a cloud identity group
type GroupLookupCallInterface interface {
	Do(call *v1beta1.GroupsLookupCall, opts ...googleapi.CallOption) (*v1beta1.LookupGroupNameResponse, error)
}

// GroupCreateCall is the default implementation for GroupCreateCallInterface
type GroupCreateCall struct{}

// GroupLookupCall is the default implementation for GroupLookupCallInterface
type GroupLookupCall struct{}

// Do performs the call, the default implementation of the interface
func (c *GroupCreateCall) Do(call *v1beta1.GroupsCreateCall, opts ...googleapi.CallOption) (*v1beta1.Operation, error) {
	return call.Do(opts...)
}

// Do performs the call, the default implementation of the interface
func (c *GroupLookupCall) Do(call *v1beta1.GroupsLookupCall, opts ...googleapi.CallOption) (*v1beta1.LookupGroupNameResponse, error) {
	return call.Do(opts...)
}
