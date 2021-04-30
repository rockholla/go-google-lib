package calls

import (
	v1beta1 "google.golang.org/api/cloudidentity/v1beta1"
	googleapi "google.golang.org/api/googleapi"
)

// GroupGetCallInterface is an interface to a call to create a cloud identity group
type GroupGetCallInterface interface {
	Do(call *v1beta1.GroupsGetCall, opts ...googleapi.CallOption) (*v1beta1.Group, error)
}

// GroupCreateCallInterface is an interface to a call to create a cloud identity group
type GroupCreateCallInterface interface {
	Do(call *v1beta1.GroupsCreateCall, opts ...googleapi.CallOption) (*v1beta1.Operation, error)
}

// GroupGetCall is the default implementation for GroupGetCallInterface
type GroupGetCall struct{}

// GroupCreateCall is the default implementation for GroupCreateCallInterface
type GroupCreateCall struct{}

// Do performs the call, the default implementation of the interface
func (c *GroupGetCall) Do(call *v1beta1.GroupsGetCall, opts ...googleapi.CallOption) (*v1beta1.Group, error) {
	return call.Do(opts...)
}

// Do performs the call, the default implementation of the interface
func (c *GroupCreateCall) Do(call *v1beta1.GroupsCreateCall, opts ...googleapi.CallOption) (*v1beta1.Operation, error) {
	return call.Do(opts...)
}
