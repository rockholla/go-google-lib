package calls

import (
	dirv1 "google.golang.org/api/admin/directory/v1"
	googleapi "google.golang.org/api/googleapi"
)

// MembersGetCallInterface is an interface to a call to get a member of a group in Google admin
type MembersGetCallInterface interface {
	Do(call *dirv1.MembersGetCall, opts ...googleapi.CallOption) (*dirv1.Member, error)
}

// MembersInsertCallInterface is an interface to a call to insert a member into a group in Google admin
type MembersInsertCallInterface interface {
	Do(call *dirv1.MembersInsertCall, opts ...googleapi.CallOption) (*dirv1.Member, error)
}

// MembersDeleteCallInterface is an interface to a call to delete a member in Google admin
type MembersDeleteCallInterface interface {
	Do(call *dirv1.MembersDeleteCall, opts ...googleapi.CallOption) error
}

// MembersGetCall is the default implementation for MembersGetCallInterface
type MembersGetCall struct{}

// MembersInsertCall is the default implementation for MembersInsertCallInterface
type MembersInsertCall struct{}

// MembersDeleteCall is the default implementation for MembersDeleteCallInterface
type MembersDeleteCall struct{}

// Do performs the call, the default implementation of the interface
func (c *MembersGetCall) Do(call *dirv1.MembersGetCall, opts ...googleapi.CallOption) (*dirv1.Member, error) {
	return call.Do(opts...)
}

// Do performs the call, the default implementation of the interface
func (c *MembersInsertCall) Do(call *dirv1.MembersInsertCall, opts ...googleapi.CallOption) (*dirv1.Member, error) {
	return call.Do(opts...)
}

// Do performs the call, the default implementation of the interface
func (c *MembersDeleteCall) Do(call *dirv1.MembersDeleteCall, opts ...googleapi.CallOption) error {
	return call.Do(opts...)
}
