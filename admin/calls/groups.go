// Package calls are mockable remote calls for operations
package calls

import (
	dirv1 "google.golang.org/api/admin/directory/v1"
	googleapi "google.golang.org/api/googleapi"
)

// GroupsInsertCallInterface is an interface to a call to insert a group into Google admin
type GroupsInsertCallInterface interface {
	Do(call *dirv1.GroupsInsertCall, opts ...googleapi.CallOption) (*dirv1.Group, error)
}

// GroupsUpdateCallInterface is an interface to a call to update a group into Google admin
type GroupsUpdateCallInterface interface {
	Do(call *dirv1.GroupsUpdateCall, opts ...googleapi.CallOption) (*dirv1.Group, error)
}

// GroupsGetCallInterface is an interface to a call to get a single group in Google admin
type GroupsGetCallInterface interface {
	Do(call *dirv1.GroupsGetCall, opts ...googleapi.CallOption) (*dirv1.Group, error)
}

// GroupsDeleteCallInterface is an interface to a call to delete a single group in Google admin
type GroupsDeleteCallInterface interface {
	Do(call *dirv1.GroupsDeleteCall, opts ...googleapi.CallOption) error
}

// GroupsInsertCall is the default implementation for GroupsInsertCallInterface
type GroupsInsertCall struct{}

// GroupsUpdateCall is the default implementation for GroupsUpdateCallInterface
type GroupsUpdateCall struct{}

// GroupsGetCall is the default implementation for GroupsGetCallInterface
type GroupsGetCall struct{}

// GroupsDeleteCall is the default implementation for GroupsDeleteCallInterface
type GroupsDeleteCall struct{}

// Do performs the call, the default implementation of the interface
func (c *GroupsInsertCall) Do(call *dirv1.GroupsInsertCall, opts ...googleapi.CallOption) (*dirv1.Group, error) {
	return call.Do(opts...)
}

// Do performs the call, the default implementation of the interface
func (c *GroupsUpdateCall) Do(call *dirv1.GroupsUpdateCall, opts ...googleapi.CallOption) (*dirv1.Group, error) {
	return call.Do(opts...)
}

// Do performs the call, the default implementation of the interface
func (c *GroupsGetCall) Do(call *dirv1.GroupsGetCall, opts ...googleapi.CallOption) (*dirv1.Group, error) {
	return call.Do(opts...)
}

// Do performs the call, the default implementation of the interface
func (c *GroupsDeleteCall) Do(call *dirv1.GroupsDeleteCall, opts ...googleapi.CallOption) error {
	return call.Do(opts...)
}
