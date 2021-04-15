package calls

import (
	v1 "google.golang.org/api/compute/v1"
	googleapi "google.golang.org/api/googleapi"
)

// TargetPoolsListCallInterface is an interface to list all target pools
type TargetPoolsListCallInterface interface {
	Do(call *v1.TargetPoolsAggregatedListCall, opts ...googleapi.CallOption) (*v1.TargetPoolAggregatedList, error)
}

// TargetPoolDeleteCallInterface is an interface to delete a target pool
type TargetPoolDeleteCallInterface interface {
	Do(call *v1.TargetPoolsDeleteCall, opts ...googleapi.CallOption) (*v1.Operation, error)
}

// BackendServicesListCallInterface is an interface to list all backend services
type BackendServicesListCallInterface interface {
	Do(call *v1.BackendServicesAggregatedListCall, opts ...googleapi.CallOption) (*v1.BackendServiceAggregatedList, error)
}

// BackendServiceDeleteCallInterface is an interface to delete a backend service
type BackendServiceDeleteCallInterface interface {
	Do(call *v1.BackendServicesDeleteCall, opts ...googleapi.CallOption) (*v1.Operation, error)
}

// RegionBackendServiceDeleteCallInterface is an interface to delete a backend service in a region
type RegionBackendServiceDeleteCallInterface interface {
	Do(call *v1.RegionBackendServicesDeleteCall, opts ...googleapi.CallOption) (*v1.Operation, error)
}

// ForwardingRulesListCallInterface is an interface to list all forwarding rules
type ForwardingRulesListCallInterface interface {
	Do(call *v1.ForwardingRulesAggregatedListCall, opts ...googleapi.CallOption) (*v1.ForwardingRuleAggregatedList, error)
}

// ForwardingRuleDeleteCallInterface is an interface to delete a forwarding rule
type ForwardingRuleDeleteCallInterface interface {
	Do(call *v1.ForwardingRulesDeleteCall, opts ...googleapi.CallOption) (*v1.Operation, error)
}

// InstanceGroupsListCallInterface is an interface to list all instance groups
type InstanceGroupsListCallInterface interface {
	Do(call *v1.InstanceGroupsAggregatedListCall, opts ...googleapi.CallOption) (*v1.InstanceGroupAggregatedList, error)
}

// InstanceGroupDeleteCallInterface is an interface to delete an instance group
type InstanceGroupDeleteCallInterface interface {
	Do(call *v1.InstanceGroupsDeleteCall, opts ...googleapi.CallOption) (*v1.Operation, error)
}

// TargetPoolsListCall is the default implementation for TargetPoolsListCallInterface
type TargetPoolsListCall struct{}

// TargetPoolDeleteCall is the default implementation for TargetPoolDeleteCallInterface
type TargetPoolDeleteCall struct{}

// BackendServicesListCall is the default implementation for BackendServicesListCallInterface
type BackendServicesListCall struct{}

// BackendServiceDeleteCall is the default implementation for BackendServiceDeleteCallInterface
type BackendServiceDeleteCall struct{}

// RegionBackendServiceDeleteCall is the default implementation for BackendServiceDeleteCallInterface
type RegionBackendServiceDeleteCall struct{}

// ForwardingRulesListCall is the default implementation for ForwardingRulesListCallInterface
type ForwardingRulesListCall struct{}

// ForwardingRuleDeleteCall is the default implementation for ForwardingRuleDeleteCallInterface
type ForwardingRuleDeleteCall struct{}

// InstanceGroupsListCall is the default implementation for InstanceGroupsListCallInterface
type InstanceGroupsListCall struct{}

// InstanceGroupDeleteCall is the default implementation for InstanceGroupDeleteCallInterface
type InstanceGroupDeleteCall struct{}

// Do performs the call, the default implementation of the interface
func (c *TargetPoolsListCall) Do(call *v1.TargetPoolsAggregatedListCall, opts ...googleapi.CallOption) (*v1.TargetPoolAggregatedList, error) {
	return call.Do(opts...)
}

// Do performs the call, the default implementation of the interface
func (c *TargetPoolDeleteCall) Do(call *v1.TargetPoolsDeleteCall, opts ...googleapi.CallOption) (*v1.Operation, error) {
	return call.Do(opts...)
}

// Do performs the call, the default implementation of the interface
func (c *BackendServicesListCall) Do(call *v1.BackendServicesAggregatedListCall, opts ...googleapi.CallOption) (*v1.BackendServiceAggregatedList, error) {
	return call.Do(opts...)
}

// Do performs the call, the default implementation of the interface
func (c *BackendServiceDeleteCall) Do(call *v1.BackendServicesDeleteCall, opts ...googleapi.CallOption) (*v1.Operation, error) {
	return call.Do(opts...)
}

// Do performs the call, the default implementation of the interface
func (c *RegionBackendServiceDeleteCall) Do(call *v1.RegionBackendServicesDeleteCall, opts ...googleapi.CallOption) (*v1.Operation, error) {
	return call.Do(opts...)
}

// Do performs the call, the default implementation of the interface
func (c *ForwardingRulesListCall) Do(call *v1.ForwardingRulesAggregatedListCall, opts ...googleapi.CallOption) (*v1.ForwardingRuleAggregatedList, error) {
	return call.Do(opts...)
}

// Do performs the call, the default implementation of the interface
func (c *ForwardingRuleDeleteCall) Do(call *v1.ForwardingRulesDeleteCall, opts ...googleapi.CallOption) (*v1.Operation, error) {
	return call.Do(opts...)
}

// Do performs the call, the default implementation of the interface
func (c *InstanceGroupsListCall) Do(call *v1.InstanceGroupsAggregatedListCall, opts ...googleapi.CallOption) (*v1.InstanceGroupAggregatedList, error) {
	return call.Do(opts...)
}

// Do performs the call, the default implementation of the interface
func (c *InstanceGroupDeleteCall) Do(call *v1.InstanceGroupsDeleteCall, opts ...googleapi.CallOption) (*v1.Operation, error) {
	return call.Do(opts...)
}
