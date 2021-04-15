package calls

import (
	v2beta "google.golang.org/api/deploymentmanager/v2beta"
	googleapi "google.golang.org/api/googleapi"
)

// OperationsGetCallInterface is an interface to a call to get a deployment manager operation
type OperationsGetCallInterface interface {
	Do(call *v2beta.OperationsGetCall, opts ...googleapi.CallOption) (*v2beta.Operation, error)
}

// OperationsGetCall is the default implementation for OperationsGetCallInterface
type OperationsGetCall struct{}

// Do performs the call, the default implementation of the interface
func (c *OperationsGetCall) Do(call *v2beta.OperationsGetCall, opts ...googleapi.CallOption) (*v2beta.Operation, error) {
	return call.Do(opts...)
}
