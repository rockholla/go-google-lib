package calls

import (
	v1 "google.golang.org/api/dns/v1"
	googleapi "google.golang.org/api/googleapi"
)

// ResourceRecordSetsListCallInterface is an interface to a call to get a list of resource record sets in a managed zone
type ResourceRecordSetsListCallInterface interface {
	Do(call *v1.ResourceRecordSetsListCall, opts ...googleapi.CallOption) (*v1.ResourceRecordSetsListResponse, error)
}

// ResourceRecordSetsListCall is the default implementation for ResourceRecordSetsListCallInterface
type ResourceRecordSetsListCall struct{}

// Do performs the call, the default implementation of the interface
func (c *ResourceRecordSetsListCall) Do(call *v1.ResourceRecordSetsListCall, opts ...googleapi.CallOption) (*v1.ResourceRecordSetsListResponse, error) {
	return call.Do(opts...)
}
