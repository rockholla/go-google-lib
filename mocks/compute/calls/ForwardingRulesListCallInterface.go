// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	compute "google.golang.org/api/compute/v1"
	googleapi "google.golang.org/api/googleapi"

	mock "github.com/stretchr/testify/mock"
)

// ForwardingRulesListCallInterface is an autogenerated mock type for the ForwardingRulesListCallInterface type
type ForwardingRulesListCallInterface struct {
	mock.Mock
}

// Do provides a mock function with given fields: call, opts
func (_m *ForwardingRulesListCallInterface) Do(call *compute.ForwardingRulesAggregatedListCall, opts ...googleapi.CallOption) (*compute.ForwardingRuleAggregatedList, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, call)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *compute.ForwardingRuleAggregatedList
	if rf, ok := ret.Get(0).(func(*compute.ForwardingRulesAggregatedListCall, ...googleapi.CallOption) *compute.ForwardingRuleAggregatedList); ok {
		r0 = rf(call, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*compute.ForwardingRuleAggregatedList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*compute.ForwardingRulesAggregatedListCall, ...googleapi.CallOption) error); ok {
		r1 = rf(call, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
