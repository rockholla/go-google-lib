// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	logger "github.com/rockholla/go-lib/logger"
	mock "github.com/stretchr/testify/mock"

	storage "github.com/rockholla/go-google-lib/storage"
)

// Interface is an autogenerated mock type for the Interface type
type Interface struct {
	mock.Mock
}

// EnsureBucket provides a mock function with given fields: name, projectID
func (_m *Interface) EnsureBucket(name string, projectID string) error {
	ret := _m.Called(name, projectID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(name, projectID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EnsureBucketRoles provides a mock function with given fields: bucket, member, roles
func (_m *Interface) EnsureBucketRoles(bucket string, member string, roles []string) error {
	ret := _m.Called(bucket, member, roles)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, []string) error); ok {
		r0 = rf(bucket, member, roles)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EnsureObject provides a mock function with given fields: bucket, path, object
func (_m *Interface) EnsureObject(bucket string, path string, object *storage.Object) error {
	ret := _m.Called(bucket, path, object)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, *storage.Object) error); ok {
		r0 = rf(bucket, path, object)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetObject provides a mock function with given fields: bucket, path
func (_m *Interface) GetObject(bucket string, path string) ([]byte, error) {
	ret := _m.Called(bucket, path)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(string, string) []byte); ok {
		r0 = rf(bucket, path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(bucket, path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetServiceAccount provides a mock function with given fields: projectID
func (_m *Interface) GetServiceAccount(projectID string) (string, error) {
	ret := _m.Called(projectID)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(projectID)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(projectID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Initialize provides a mock function with given fields: credentials, log
func (_m *Interface) Initialize(credentials string, log logger.Interface) error {
	ret := _m.Called(credentials, log)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, logger.Interface) error); ok {
		r0 = rf(credentials, log)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
