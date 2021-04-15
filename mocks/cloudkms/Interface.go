// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	cloudkms "github.com/rockholla/go-google-lib/cloudkms"
	logger "github.com/rockholla/go-lib/logger"

	mock "github.com/stretchr/testify/mock"
)

// Interface is an autogenerated mock type for the Interface type
type Interface struct {
	mock.Mock
}

// Decrypt provides a mock function with given fields: key, data
func (_m *Interface) Decrypt(key *cloudkms.CryptoKey, data string) (string, error) {
	ret := _m.Called(key, data)

	var r0 string
	if rf, ok := ret.Get(0).(func(*cloudkms.CryptoKey, string) string); ok {
		r0 = rf(key, data)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*cloudkms.CryptoKey, string) error); ok {
		r1 = rf(key, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Encrypt provides a mock function with given fields: key, data
func (_m *Interface) Encrypt(key *cloudkms.CryptoKey, data string) (string, error) {
	ret := _m.Called(key, data)

	var r0 string
	if rf, ok := ret.Get(0).(func(*cloudkms.CryptoKey, string) string); ok {
		r0 = rf(key, data)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*cloudkms.CryptoKey, string) error); ok {
		r1 = rf(key, data)
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
