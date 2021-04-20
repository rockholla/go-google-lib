// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	compute "google.golang.org/api/compute/v1"
	googleapi "google.golang.org/api/googleapi"

	mock "github.com/stretchr/testify/mock"
)

// DisksListCallInterface is an autogenerated mock type for the DisksListCallInterface type
type DisksListCallInterface struct {
	mock.Mock
}

// Do provides a mock function with given fields: call, opts
func (_m *DisksListCallInterface) Do(call *compute.DisksAggregatedListCall, opts ...googleapi.CallOption) (*compute.DiskAggregatedList, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, call)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *compute.DiskAggregatedList
	if rf, ok := ret.Get(0).(func(*compute.DisksAggregatedListCall, ...googleapi.CallOption) *compute.DiskAggregatedList); ok {
		r0 = rf(call, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*compute.DiskAggregatedList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*compute.DisksAggregatedListCall, ...googleapi.CallOption) error); ok {
		r1 = rf(call, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}