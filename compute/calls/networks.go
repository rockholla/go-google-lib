package calls

import (
	v1 "google.golang.org/api/compute/v1"
	googleapi "google.golang.org/api/googleapi"
)

// NetworkGetCallInterface is an interface to get a network
type NetworkGetCallInterface interface {
	Do(call *v1.NetworksGetCall, opts ...googleapi.CallOption) (*v1.Network, error)
}

// NetworkDeleteCallInterface is an interface to delete a network
type NetworkDeleteCallInterface interface {
	Do(call *v1.NetworksDeleteCall, opts ...googleapi.CallOption) (*v1.Operation, error)
}

// NetworkGetCall is the default implementation for NetworkGetCallInterface
type NetworkGetCall struct{}

// NetworkDeleteCall is the default implementation for NetworkDeleteCallInterface
type NetworkDeleteCall struct{}

// Do performs the call, the default implementation of the interface
func (c *NetworkGetCall) Do(call *v1.NetworksGetCall, opts ...googleapi.CallOption) (*v1.Network, error) {
	return call.Do(opts...)
}

// Do performs the call, the default implementation of the interface
func (c *NetworkDeleteCall) Do(call *v1.NetworksDeleteCall, opts ...googleapi.CallOption) (*v1.Operation, error) {
	return call.Do(opts...)
}
