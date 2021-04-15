package calls

import (
	v1 "google.golang.org/api/compute/v1"
	googleapi "google.golang.org/api/googleapi"
)

// RegionsGetCallInterface is an interface to a call to get a compute region
type RegionsGetCallInterface interface {
	Do(call *v1.RegionsGetCall, opts ...googleapi.CallOption) (*v1.Region, error)
}

// RegionsGetCall is the default implementation for RegionsGetCallInterface
type RegionsGetCall struct{}

// Do performs the call, the default implementation of the interface
func (c *RegionsGetCall) Do(call *v1.RegionsGetCall, opts ...googleapi.CallOption) (*v1.Region, error) {
	return call.Do(opts...)
}
