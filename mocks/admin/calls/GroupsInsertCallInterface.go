// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	admin "google.golang.org/api/admin/directory/v1"

	googleapi "google.golang.org/api/googleapi"

	mock "github.com/stretchr/testify/mock"
)

// GroupsInsertCallInterface is an autogenerated mock type for the GroupsInsertCallInterface type
type GroupsInsertCallInterface struct {
	mock.Mock
}

// Do provides a mock function with given fields: call, opts
func (_m *GroupsInsertCallInterface) Do(call *admin.GroupsInsertCall, opts ...googleapi.CallOption) (*admin.Group, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, call)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *admin.Group
	if rf, ok := ret.Get(0).(func(*admin.GroupsInsertCall, ...googleapi.CallOption) *admin.Group); ok {
		r0 = rf(call, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.Group)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*admin.GroupsInsertCall, ...googleapi.CallOption) error); ok {
		r1 = rf(call, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
