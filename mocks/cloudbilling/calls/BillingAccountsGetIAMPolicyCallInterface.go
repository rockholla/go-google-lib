// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	cloudbilling "google.golang.org/api/cloudbilling/v1"
	googleapi "google.golang.org/api/googleapi"

	mock "github.com/stretchr/testify/mock"
)

// BillingAccountsGetIAMPolicyCallInterface is an autogenerated mock type for the BillingAccountsGetIAMPolicyCallInterface type
type BillingAccountsGetIAMPolicyCallInterface struct {
	mock.Mock
}

// Do provides a mock function with given fields: call, opts
func (_m *BillingAccountsGetIAMPolicyCallInterface) Do(call *cloudbilling.BillingAccountsGetIamPolicyCall, opts ...googleapi.CallOption) (*cloudbilling.Policy, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, call)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *cloudbilling.Policy
	if rf, ok := ret.Get(0).(func(*cloudbilling.BillingAccountsGetIamPolicyCall, ...googleapi.CallOption) *cloudbilling.Policy); ok {
		r0 = rf(call, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cloudbilling.Policy)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*cloudbilling.BillingAccountsGetIamPolicyCall, ...googleapi.CallOption) error); ok {
		r1 = rf(call, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
