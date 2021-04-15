// Package calls are mockable remote calls for operations
package calls

import (
	v1 "google.golang.org/api/dns/v1"
	googleapi "google.golang.org/api/googleapi"
)

// ChangesCreateCallInterface is an interface to a call to create a dns-related change
type ChangesCreateCallInterface interface {
	Do(call *v1.ChangesCreateCall, opts ...googleapi.CallOption) (*v1.Change, error)
}

// ChangesCreateCall is the default implementation for ChangesCreateCallInterface
type ChangesCreateCall struct{}

// Do performs the call, the default implementation of the interface
func (c *ChangesCreateCall) Do(call *v1.ChangesCreateCall, opts ...googleapi.CallOption) (*v1.Change, error) {
	return call.Do(opts...)
}
