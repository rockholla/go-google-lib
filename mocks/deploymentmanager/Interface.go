// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	deploymentmanager "github.com/rockholla/go-google-lib/deploymentmanager"
	logger "github.com/rockholla/go-lib/logger"

	mock "github.com/stretchr/testify/mock"
)

// Interface is an autogenerated mock type for the Interface type
type Interface struct {
	mock.Mock
}

// DeleteDeployment provides a mock function with given fields: deploymentName, inProject, abandon
func (_m *Interface) DeleteDeployment(deploymentName string, inProject string, abandon bool) error {
	ret := _m.Called(deploymentName, inProject, abandon)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, bool) error); ok {
		r0 = rf(deploymentName, inProject, abandon)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EnsureDeployment provides a mock function with given fields: deploymentName, description, inProject, deployment
func (_m *Interface) EnsureDeployment(deploymentName string, description string, inProject string, deployment *deploymentmanager.Deployment) ([]*deploymentmanager.Output, error) {
	ret := _m.Called(deploymentName, description, inProject, deployment)

	var r0 []*deploymentmanager.Output
	if rf, ok := ret.Get(0).(func(string, string, string, *deploymentmanager.Deployment) []*deploymentmanager.Output); ok {
		r0 = rf(deploymentName, description, inProject, deployment)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*deploymentmanager.Output)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string, *deploymentmanager.Deployment) error); ok {
		r1 = rf(deploymentName, description, inProject, deployment)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDeployment provides a mock function with given fields: deploymentName, inProject, parseManifest
func (_m *Interface) GetDeployment(deploymentName string, inProject string, parseManifest bool) (*deploymentmanager.Deployment, error) {
	ret := _m.Called(deploymentName, inProject, parseManifest)

	var r0 *deploymentmanager.Deployment
	if rf, ok := ret.Get(0).(func(string, string, bool) *deploymentmanager.Deployment); ok {
		r0 = rf(deploymentName, inProject, parseManifest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*deploymentmanager.Deployment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, bool) error); ok {
		r1 = rf(deploymentName, inProject, parseManifest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetResourcePropertyValue provides a mock function with given fields: deploymentName, inProject, resourceName, propertyName
func (_m *Interface) GetResourcePropertyValue(deploymentName string, inProject string, resourceName string, propertyName string) (string, error) {
	ret := _m.Called(deploymentName, inProject, resourceName, propertyName)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string, string, string) string); ok {
		r0 = rf(deploymentName, inProject, resourceName, propertyName)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string, string) error); ok {
		r1 = rf(deploymentName, inProject, resourceName, propertyName)
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
