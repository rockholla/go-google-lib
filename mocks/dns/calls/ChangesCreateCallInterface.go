// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	dns "google.golang.org/api/dns/v1"
	googleapi "google.golang.org/api/googleapi"

	mock "github.com/stretchr/testify/mock"
)

// ChangesCreateCallInterface is an autogenerated mock type for the ChangesCreateCallInterface type
type ChangesCreateCallInterface struct {
	mock.Mock
}

// Do provides a mock function with given fields: call, opts
func (_m *ChangesCreateCallInterface) Do(call *dns.ChangesCreateCall, opts ...googleapi.CallOption) (*dns.Change, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, call)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *dns.Change
	if rf, ok := ret.Get(0).(func(*dns.ChangesCreateCall, ...googleapi.CallOption) *dns.Change); ok {
		r0 = rf(call, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dns.Change)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*dns.ChangesCreateCall, ...googleapi.CallOption) error); ok {
		r1 = rf(call, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
