// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	logger "github.com/rockholla/go-lib/logger"
	mock "github.com/stretchr/testify/mock"
)

// Interface is an autogenerated mock type for the Interface type
type Interface struct {
	mock.Mock
}

// GetAccessToken provides a mock function with given fields:
func (_m *Interface) GetAccessToken() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Initialize provides a mock function with given fields: credentials, log, scopes
func (_m *Interface) Initialize(credentials string, log logger.Interface, scopes []string) error {
	ret := _m.Called(credentials, log, scopes)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, logger.Interface, []string) error); ok {
		r0 = rf(credentials, log, scopes)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
