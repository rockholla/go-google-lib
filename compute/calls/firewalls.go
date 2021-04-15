package calls

import (
	v1 "google.golang.org/api/compute/v1"
	googleapi "google.golang.org/api/googleapi"
)

// FirewallsListCallInterface is an interface to list all disks
type FirewallsListCallInterface interface {
	Do(call *v1.FirewallsListCall, opts ...googleapi.CallOption) (*v1.FirewallList, error)
}

// FirewallDeleteCallInterface is an interface to delete a disk
type FirewallDeleteCallInterface interface {
	Do(call *v1.FirewallsDeleteCall, opts ...googleapi.CallOption) (*v1.Operation, error)
}

// FirewallsListCall is the default implementation for FirewallsListCallInterface
type FirewallsListCall struct{}

// FirewallDeleteCall is the default implementation for FirewallDeleteCallInterface
type FirewallDeleteCall struct{}

// Do performs the call, the default implementation of the interface
func (c *FirewallsListCall) Do(call *v1.FirewallsListCall, opts ...googleapi.CallOption) (*v1.FirewallList, error) {
	return call.Do(opts...)
}

// Do performs the call, the default implementation of the interface
func (c *FirewallDeleteCall) Do(call *v1.FirewallsDeleteCall, opts ...googleapi.CallOption) (*v1.Operation, error) {
	return call.Do(opts...)
}
