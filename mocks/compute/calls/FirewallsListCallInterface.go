// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	compute "google.golang.org/api/compute/v1"
	googleapi "google.golang.org/api/googleapi"

	mock "github.com/stretchr/testify/mock"
)

// FirewallsListCallInterface is an autogenerated mock type for the FirewallsListCallInterface type
type FirewallsListCallInterface struct {
	mock.Mock
}

// Do provides a mock function with given fields: call, opts
func (_m *FirewallsListCallInterface) Do(call *compute.FirewallsListCall, opts ...googleapi.CallOption) (*compute.FirewallList, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, call)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *compute.FirewallList
	if rf, ok := ret.Get(0).(func(*compute.FirewallsListCall, ...googleapi.CallOption) *compute.FirewallList); ok {
		r0 = rf(call, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*compute.FirewallList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*compute.FirewallsListCall, ...googleapi.CallOption) error); ok {
		r1 = rf(call, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
