package calls

import (
	v2beta "google.golang.org/api/deploymentmanager/v2beta"
	googleapi "google.golang.org/api/googleapi"
)

// ResourcesGetCallInterface is an interface to a call to get a resource from a deployment
type ResourcesGetCallInterface interface {
	Do(call *v2beta.ResourcesGetCall, opts ...googleapi.CallOption) (*v2beta.Resource, error)
}

// ResourcesGetCall is the default implementation for ResourcesGetCallInterface
type ResourcesGetCall struct{}

// Do performs the call, the default implementation of the interface
func (c *ResourcesGetCall) Do(call *v2beta.ResourcesGetCall, opts ...googleapi.CallOption) (*v2beta.Resource, error) {
	return call.Do(opts...)
}
