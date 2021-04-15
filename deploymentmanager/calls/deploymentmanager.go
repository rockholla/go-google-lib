// Package calls are mockable remote calls for operations
package calls

import (
	v2beta "google.golang.org/api/deploymentmanager/v2beta"
	googleapi "google.golang.org/api/googleapi"
)

// DeploymentsGetCallInterface is an interface to a call to get a deployment
type DeploymentsGetCallInterface interface {
	Do(call *v2beta.DeploymentsGetCall, opts ...googleapi.CallOption) (*v2beta.Deployment, error)
}

// DeploymentsInsertCallInterface is an interface to a call to create/insert a deployment
type DeploymentsInsertCallInterface interface {
	Do(call *v2beta.DeploymentsInsertCall, opts ...googleapi.CallOption) (*v2beta.Operation, error)
}

// DeploymentsUpdateCallInterface is an interface to a call to update a deployment
type DeploymentsUpdateCallInterface interface {
	Do(call *v2beta.DeploymentsUpdateCall, opts ...googleapi.CallOption) (*v2beta.Operation, error)
}

// DeploymentsDeleteCallInterface is an interface to a call to update a deployment
type DeploymentsDeleteCallInterface interface {
	Do(call *v2beta.DeploymentsDeleteCall, opts ...googleapi.CallOption) (*v2beta.Operation, error)
}

// DeploymentsGetCall is the default implementation for DeploymentsGetCallInterface
type DeploymentsGetCall struct{}

// DeploymentsInsertCall is the default implementation for DeploymentsInsertCallInterface
type DeploymentsInsertCall struct{}

// DeploymentsUpdateCall is the default implementation for DeploymentsUpdateCallInterface
type DeploymentsUpdateCall struct{}

// DeploymentsDeleteCall is the default implementation for DeploymentsDeleteCallInterface
type DeploymentsDeleteCall struct{}

// Do performs the call, the default implementation of the interface
func (c *DeploymentsGetCall) Do(call *v2beta.DeploymentsGetCall, opts ...googleapi.CallOption) (*v2beta.Deployment, error) {
	return call.Do(opts...)
}

// Do performs the call, the default implementation of the interface
func (c *DeploymentsInsertCall) Do(call *v2beta.DeploymentsInsertCall, opts ...googleapi.CallOption) (*v2beta.Operation, error) {
	return call.Do(opts...)
}

// Do performs the call, the default implementation of the interface
func (c *DeploymentsUpdateCall) Do(call *v2beta.DeploymentsUpdateCall, opts ...googleapi.CallOption) (*v2beta.Operation, error) {
	return call.Do(opts...)
}

// Do performs the call, the default implementation of the interface
func (c *DeploymentsDeleteCall) Do(call *v2beta.DeploymentsDeleteCall, opts ...googleapi.CallOption) (*v2beta.Operation, error) {
	return call.Do(opts...)
}
