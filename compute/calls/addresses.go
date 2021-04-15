package calls

import (
	v1 "google.golang.org/api/compute/v1"
	googleapi "google.golang.org/api/googleapi"
)

// AddressesListCallInterface is an interface to list all addresses
type AddressesListCallInterface interface {
	Do(call *v1.AddressesAggregatedListCall, opts ...googleapi.CallOption) (*v1.AddressAggregatedList, error)
}

// AddressDeleteCallInterface is an interface to delete an address
type AddressDeleteCallInterface interface {
	Do(call *v1.AddressesDeleteCall, opts ...googleapi.CallOption) (*v1.Operation, error)
}

// AddressesListCall is the default implementation for AddressesListCallInterface
type AddressesListCall struct{}

// AddressDeleteCall is the default implementation for AddressDeleteCallInterface
type AddressDeleteCall struct{}

// Do performs the call, the default implementation of the interface
func (c *AddressesListCall) Do(call *v1.AddressesAggregatedListCall, opts ...googleapi.CallOption) (*v1.AddressAggregatedList, error) {
	return call.Do(opts...)
}

// Do performs the call, the default implementation of the interface
func (c *AddressDeleteCall) Do(call *v1.AddressesDeleteCall, opts ...googleapi.CallOption) (*v1.Operation, error) {
	return call.Do(opts...)
}
