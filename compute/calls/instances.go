// Package calls are mockable remote calls for operations
package calls

import (
	v1 "google.golang.org/api/compute/v1"
	googleapi "google.golang.org/api/googleapi"
)

// InstancesAggregatedListCallInterface is an interface to a call to list instances across all zones
type InstancesAggregatedListCallInterface interface {
	Do(call *v1.InstancesAggregatedListCall, opts ...googleapi.CallOption) (*v1.InstanceAggregatedList, error)
}

// InstancesStopCallInterface is an interface to a call to stop instances
type InstancesStopCallInterface interface {
	Do(call *v1.InstancesStopCall, opts ...googleapi.CallOption) (*v1.Operation, error)
}

// InstancesStartCallInterface is an interface to a call to start instances
type InstancesStartCallInterface interface {
	Do(call *v1.InstancesStartCall, opts ...googleapi.CallOption) (*v1.Operation, error)
}

// ProjectsSetCommonInstanceMetadataCallInterface is an interface to a call to set project-level metadata for instances
type ProjectsSetCommonInstanceMetadataCallInterface interface {
	Do(call *v1.ProjectsSetCommonInstanceMetadataCall, opts ...googleapi.CallOption) (*v1.Operation, error)
}

// ProjectsGetCallInterface is an interface to a call get a compute project object
type ProjectsGetCallInterface interface {
	Do(call *v1.ProjectsGetCall, opts ...googleapi.CallOption) (*v1.Project, error)
}

// InstancesAggregatedListCall is the default implementation for InstancesAggregatedListCallInterface
type InstancesAggregatedListCall struct{}

// InstancesStopCall is the default implementation for InstancesStopCallInterface
type InstancesStopCall struct{}

// InstancesStartCall is the default implementation for InstancesStartCallInterface
type InstancesStartCall struct{}

// ProjectsSetCommonInstanceMetadataCall is the default implementation for SetCommonInstanceMetadataCallInterface
type ProjectsSetCommonInstanceMetadataCall struct{}

// ProjectsGetCall is the default implementation for ProjectsGetCallInterface
type ProjectsGetCall struct{}

// Do performs the call, the default implementation of the interface
func (c *InstancesAggregatedListCall) Do(call *v1.InstancesAggregatedListCall, opts ...googleapi.CallOption) (*v1.InstanceAggregatedList, error) {
	return call.Do(opts...)
}

// Do performs the call, the default implementation of the interface
func (c *InstancesStopCall) Do(call *v1.InstancesStopCall, opts ...googleapi.CallOption) (*v1.Operation, error) {
	return call.Do(opts...)
}

// Do performs the call, the default implementation of the interface
func (c *InstancesStartCall) Do(call *v1.InstancesStartCall, opts ...googleapi.CallOption) (*v1.Operation, error) {
	return call.Do(opts...)
}

// Do performs the call, the default implementation of the interface
func (c *ProjectsSetCommonInstanceMetadataCall) Do(call *v1.ProjectsSetCommonInstanceMetadataCall, opts ...googleapi.CallOption) (*v1.Operation, error) {
	return call.Do(opts...)
}

// Do performs the call, the default implementation of the interface
func (c *ProjectsGetCall) Do(call *v1.ProjectsGetCall, opts ...googleapi.CallOption) (*v1.Project, error) {
	return call.Do(opts...)
}
