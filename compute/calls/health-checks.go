package calls

import (
	v1 "google.golang.org/api/compute/v1"
	googleapi "google.golang.org/api/googleapi"
)

// HealthChecksListCallInterface is an interface to list all health checks
type HealthChecksListCallInterface interface {
	Do(call *v1.HealthChecksAggregatedListCall, opts ...googleapi.CallOption) (*v1.HealthChecksAggregatedList, error)
}

// HealthCheckDeleteCallInterface is an interface to delete a health check
type HealthCheckDeleteCallInterface interface {
	Do(call *v1.HealthChecksDeleteCall, opts ...googleapi.CallOption) (*v1.Operation, error)
}

// HTTPHealthChecksListCallInterface is an interface to list all http (legacy) health checks
type HTTPHealthChecksListCallInterface interface {
	Do(call *v1.HttpHealthChecksListCall, opts ...googleapi.CallOption) (*v1.HttpHealthCheckList, error)
}

// HTTPHealthCheckDeleteCallInterface is an interface to delete a http (legacy) health check
type HTTPHealthCheckDeleteCallInterface interface {
	Do(call *v1.HttpHealthChecksDeleteCall, opts ...googleapi.CallOption) (*v1.Operation, error)
}

// HealthChecksListCall is the default implementation for HealthChecksListCallInterface
type HealthChecksListCall struct{}

// HealthCheckDeleteCall is the default implementation for HealthCheckDeleteCallInterface
type HealthCheckDeleteCall struct{}

// HTTPHealthChecksListCall is the default implementation for HTTPHealthChecksListCallInterface
type HTTPHealthChecksListCall struct{}

// HTTPHealthCheckDeleteCall is the default implementation for HTTPHealthCheckDeleteCallInterface
type HTTPHealthCheckDeleteCall struct{}

// Do performs the call, the default implementation of the interface
func (c *HealthChecksListCall) Do(call *v1.HealthChecksAggregatedListCall, opts ...googleapi.CallOption) (*v1.HealthChecksAggregatedList, error) {
	return call.Do(opts...)
}

// Do performs the call, the default implementation of the interface
func (c *HealthCheckDeleteCall) Do(call *v1.HealthChecksDeleteCall, opts ...googleapi.CallOption) (*v1.Operation, error) {
	return call.Do(opts...)
}

// Do performs the call, the default implementation of the interface
func (c *HTTPHealthChecksListCall) Do(call *v1.HttpHealthChecksListCall, opts ...googleapi.CallOption) (*v1.HttpHealthCheckList, error) {
	return call.Do(opts...)
}

// Do performs the call, the default implementation of the interface
func (c *HTTPHealthCheckDeleteCall) Do(call *v1.HttpHealthChecksDeleteCall, opts ...googleapi.CallOption) (*v1.Operation, error) {
	return call.Do(opts...)
}
