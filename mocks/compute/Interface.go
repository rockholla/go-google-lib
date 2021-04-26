// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	compute "github.com/rockholla/go-google-lib/compute"
	logger "github.com/rockholla/go-lib/logger"

	mock "github.com/stretchr/testify/mock"

	v1 "google.golang.org/api/compute/v1"
)

// Interface is an autogenerated mock type for the Interface type
type Interface struct {
	mock.Mock
}

// DeleteAddress provides a mock function with given fields: projectID, region, name
func (_m *Interface) DeleteAddress(projectID string, region string, name string) error {
	ret := _m.Called(projectID, region, name)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(projectID, region, name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteBackendService provides a mock function with given fields: projectID, name
func (_m *Interface) DeleteBackendService(projectID string, name string) error {
	ret := _m.Called(projectID, name)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(projectID, name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteDisk provides a mock function with given fields: projectID, zone, name
func (_m *Interface) DeleteDisk(projectID string, zone string, name string) error {
	ret := _m.Called(projectID, zone, name)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(projectID, zone, name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteFirewall provides a mock function with given fields: projectID, name
func (_m *Interface) DeleteFirewall(projectID string, name string) error {
	ret := _m.Called(projectID, name)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(projectID, name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteForwardingRule provides a mock function with given fields: projectID, region, name
func (_m *Interface) DeleteForwardingRule(projectID string, region string, name string) error {
	ret := _m.Called(projectID, region, name)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(projectID, region, name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteHTTPHealthCheck provides a mock function with given fields: projectID, name
func (_m *Interface) DeleteHTTPHealthCheck(projectID string, name string) error {
	ret := _m.Called(projectID, name)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(projectID, name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteHealthCheck provides a mock function with given fields: projectID, name
func (_m *Interface) DeleteHealthCheck(projectID string, name string) error {
	ret := _m.Called(projectID, name)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(projectID, name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteInstanceGroup provides a mock function with given fields: projectID, zone, name
func (_m *Interface) DeleteInstanceGroup(projectID string, zone string, name string) error {
	ret := _m.Called(projectID, zone, name)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(projectID, zone, name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteNetwork provides a mock function with given fields: projectID, name
func (_m *Interface) DeleteNetwork(projectID string, name string) error {
	ret := _m.Called(projectID, name)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(projectID, name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteRegionBackendService provides a mock function with given fields: projectID, region, name
func (_m *Interface) DeleteRegionBackendService(projectID string, region string, name string) error {
	ret := _m.Called(projectID, region, name)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(projectID, region, name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteTargetPool provides a mock function with given fields: projectID, region, name
func (_m *Interface) DeleteTargetPool(projectID string, region string, name string) error {
	ret := _m.Called(projectID, region, name)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(projectID, region, name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAddresses provides a mock function with given fields: projectID
func (_m *Interface) GetAddresses(projectID string) ([]*v1.Address, error) {
	ret := _m.Called(projectID)

	var r0 []*v1.Address
	if rf, ok := ret.Get(0).(func(string) []*v1.Address); ok {
		r0 = rf(projectID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v1.Address)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(projectID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBackendServices provides a mock function with given fields: projectID
func (_m *Interface) GetBackendServices(projectID string) ([]*v1.BackendService, error) {
	ret := _m.Called(projectID)

	var r0 []*v1.BackendService
	if rf, ok := ret.Get(0).(func(string) []*v1.BackendService); ok {
		r0 = rf(projectID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v1.BackendService)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(projectID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCommonInstanceMetadata provides a mock function with given fields: projectID
func (_m *Interface) GetCommonInstanceMetadata(projectID string) ([]*v1.MetadataItems, error) {
	ret := _m.Called(projectID)

	var r0 []*v1.MetadataItems
	if rf, ok := ret.Get(0).(func(string) []*v1.MetadataItems); ok {
		r0 = rf(projectID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v1.MetadataItems)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(projectID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDisks provides a mock function with given fields: projectID
func (_m *Interface) GetDisks(projectID string) ([]*v1.Disk, error) {
	ret := _m.Called(projectID)

	var r0 []*v1.Disk
	if rf, ok := ret.Get(0).(func(string) []*v1.Disk); ok {
		r0 = rf(projectID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v1.Disk)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(projectID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFirewalls provides a mock function with given fields: projectID
func (_m *Interface) GetFirewalls(projectID string) ([]*v1.Firewall, error) {
	ret := _m.Called(projectID)

	var r0 []*v1.Firewall
	if rf, ok := ret.Get(0).(func(string) []*v1.Firewall); ok {
		r0 = rf(projectID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v1.Firewall)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(projectID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetHTTPHealthChecks provides a mock function with given fields: projectID
func (_m *Interface) GetHTTPHealthChecks(projectID string) ([]*v1.HttpHealthCheck, error) {
	ret := _m.Called(projectID)

	var r0 []*v1.HttpHealthCheck
	if rf, ok := ret.Get(0).(func(string) []*v1.HttpHealthCheck); ok {
		r0 = rf(projectID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v1.HttpHealthCheck)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(projectID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetHealthChecks provides a mock function with given fields: projectID
func (_m *Interface) GetHealthChecks(projectID string) ([]*v1.HealthCheck, error) {
	ret := _m.Called(projectID)

	var r0 []*v1.HealthCheck
	if rf, ok := ret.Get(0).(func(string) []*v1.HealthCheck); ok {
		r0 = rf(projectID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v1.HealthCheck)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(projectID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetInstanceGroups provides a mock function with given fields: projectID
func (_m *Interface) GetInstanceGroups(projectID string) ([]*v1.InstanceGroup, error) {
	ret := _m.Called(projectID)

	var r0 []*v1.InstanceGroup
	if rf, ok := ret.Get(0).(func(string) []*v1.InstanceGroup); ok {
		r0 = rf(projectID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v1.InstanceGroup)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(projectID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetInternalIPs provides a mock function with given fields: projectID, network
func (_m *Interface) GetInternalIPs(projectID string, network string) ([]*compute.InstanceIP, error) {
	ret := _m.Called(projectID, network)

	var r0 []*compute.InstanceIP
	if rf, ok := ret.Get(0).(func(string, string) []*compute.InstanceIP); ok {
		r0 = rf(projectID, network)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*compute.InstanceIP)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(projectID, network)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNetwork provides a mock function with given fields: projectID, name
func (_m *Interface) GetNetwork(projectID string, name string) (*v1.Network, error) {
	ret := _m.Called(projectID, name)

	var r0 *v1.Network
	if rf, ok := ret.Get(0).(func(string, string) *v1.Network); ok {
		r0 = rf(projectID, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Network)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(projectID, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRegionZones provides a mock function with given fields: projectID, region
func (_m *Interface) GetRegionZones(projectID string, region string) ([]string, error) {
	ret := _m.Called(projectID, region)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string, string) []string); ok {
		r0 = rf(projectID, region)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(projectID, region)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTargetPools provides a mock function with given fields: projectID
func (_m *Interface) GetTargetPools(projectID string) ([]*v1.TargetPool, error) {
	ret := _m.Called(projectID)

	var r0 []*v1.TargetPool
	if rf, ok := ret.Get(0).(func(string) []*v1.TargetPool); ok {
		r0 = rf(projectID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v1.TargetPool)
		}
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

// PowerOff provides a mock function with given fields: projectID
func (_m *Interface) PowerOff(projectID string) error {
	ret := _m.Called(projectID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(projectID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PowerOn provides a mock function with given fields: projectID
func (_m *Interface) PowerOn(projectID string) error {
	ret := _m.Called(projectID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(projectID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetCommonInstanceMetadata provides a mock function with given fields: projectID, metadataItems
func (_m *Interface) SetCommonInstanceMetadata(projectID string, metadataItems []*v1.MetadataItems) error {
	ret := _m.Called(projectID, metadataItems)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []*v1.MetadataItems) error); ok {
		r0 = rf(projectID, metadataItems)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
