package calls

import (
	v1 "google.golang.org/api/compute/v1"
	googleapi "google.golang.org/api/googleapi"
)

// DisksListCallInterface is an interface to list all disks
type DisksListCallInterface interface {
	Do(call *v1.DisksAggregatedListCall, opts ...googleapi.CallOption) (*v1.DiskAggregatedList, error)
}

// DiskDeleteCallInterface is an interface to delete a disk
type DiskDeleteCallInterface interface {
	Do(call *v1.DisksDeleteCall, opts ...googleapi.CallOption) (*v1.Operation, error)
}

// DisksListCall is the default implementation for DisksListCallInterface
type DisksListCall struct{}

// DiskDeleteCall is the default implementation for DiskDeleteCallInterface
type DiskDeleteCall struct{}

// Do performs the call, the default implementation of the interface
func (c *DisksListCall) Do(call *v1.DisksAggregatedListCall, opts ...googleapi.CallOption) (*v1.DiskAggregatedList, error) {
	return call.Do(opts...)
}

// Do performs the call, the default implementation of the interface
func (c *DiskDeleteCall) Do(call *v1.DisksDeleteCall, opts ...googleapi.CallOption) (*v1.Operation, error) {
	return call.Do(opts...)
}
