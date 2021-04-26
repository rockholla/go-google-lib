package compute

import (
	"errors"
	"fmt"
	"testing"

	loggermock "github.com/rockholla/go-lib/mocks/custom-mocks/logger"
	v1 "google.golang.org/api/compute/v1"
	googleapi "google.golang.org/api/googleapi"
)

var (
	triggerRegionNotFound = false
	testProjectID         = "project-11111111111"
	testCredentials       = `{
  "client_id": "xxxxxxx.apps.googleusercontent.com",
  "client_secret": "xxxxxxxxxxxxxxx",
  "refresh_token": "xxxxxxxxx",
  "type": "authorized_user"
}`
)

type regionsGetMock struct{}
type instancesAggregatedListMock struct{}
type instancesStopMock struct{}
type instancesStartMock struct{}
type projectsSetCommonInstanceMetadataMock struct{}
type projectsGetMock struct{}
type targetPoolsListMock struct{}
type targetPoolDeleteMock struct{}
type backendServicesListMock struct{}
type backendServiceDeleteMock struct{}
type regionBackendServiceDeleteMock struct{}
type forwardingRulesListMock struct{}
type forwardingRuleDeleteMock struct{}
type healthChecksListMock struct{}
type healthCheckDeleteMock struct{}
type httpHealthChecksListMock struct{}
type httpHealthCheckDeleteMock struct{}
type disksListMock struct{}
type diskDeleteMock struct{}
type addressesListMock struct{}
type addressDeleteMock struct{}
type firewallsListMock struct{}
type firewallDeleteMock struct{}
type instanceGroupsListMock struct{}
type instanceGroupDeleteMock struct{}
type networkGetMock struct{}
type networkDeleteMock struct{}

// Do is the mock for default regionsGetMock
func (c *regionsGetMock) Do(call *v1.RegionsGetCall, opts ...googleapi.CallOption) (*v1.Region, error) {
	if triggerRegionNotFound {
		triggerRegionNotFound = false
		return &v1.Region{}, errors.New("NotFound")
	}
	return &v1.Region{
		Zones: []string{
			fmt.Sprintf("https://www.googleapis.com/compute/v1/projects/%s/zones/us-central1-a", testProjectID),
			fmt.Sprintf("https://www.googleapis.com/compute/v1/projects/%s/zones/us-central1-b", testProjectID),
			fmt.Sprintf("https://www.googleapis.com/compute/v1/projects/%s/zones/us-central1-c", testProjectID),
			fmt.Sprintf("https://www.googleapis.com/compute/v1/projects/%s/zones/us-central1-f", testProjectID),
		},
	}, nil
}

// Do is the mock for default instancesAggregatedListMock
func (c *instancesAggregatedListMock) Do(call *v1.InstancesAggregatedListCall, opts ...googleapi.CallOption) (*v1.InstanceAggregatedList, error) {
	return &v1.InstanceAggregatedList{
		Items: map[string]v1.InstancesScopedList{
			"us-central1-a": v1.InstancesScopedList{
				Instances: []*v1.Instance{
					&v1.Instance{
						Name: "instance-1",
						NetworkInterfaces: []*v1.NetworkInterface{
							&v1.NetworkInterface{
								Network:   fmt.Sprintf("/projects/%s/global/networks/one", testProjectID),
								NetworkIP: "10.1.0.10",
							},
							&v1.NetworkInterface{
								Network:   fmt.Sprintf("/projects/%s/global/networks/two", testProjectID),
								NetworkIP: "10.2.0.10",
							},
							&v1.NetworkInterface{
								Network:   fmt.Sprintf("/projects/%s/global/networks/three", testProjectID),
								NetworkIP: "10.3.0.10",
							},
						},
					},
				},
			},
			"us-central1-b": v1.InstancesScopedList{
				Instances: []*v1.Instance{
					&v1.Instance{
						Name: "instance-2",
						NetworkInterfaces: []*v1.NetworkInterface{
							&v1.NetworkInterface{
								Network:   fmt.Sprintf("/projects/%s/global/networks/one", testProjectID),
								NetworkIP: "10.1.0.11",
							},
							&v1.NetworkInterface{
								Network:   fmt.Sprintf("/projects/%s/global/networks/two", testProjectID),
								NetworkIP: "10.2.0.11",
							},
							&v1.NetworkInterface{
								Network:   fmt.Sprintf("/projects/%s/global/networks/three", testProjectID),
								NetworkIP: "10.3.0.11",
							},
						},
					},
				},
			},
		},
	}, nil
}

// Do is the mock for default instancesStopMock
func (c *instancesStopMock) Do(call *v1.InstancesStopCall, opts ...googleapi.CallOption) (*v1.Operation, error) {
	return &v1.Operation{}, nil
}

// Do is the mock for default instancesStartMock
func (c *instancesStartMock) Do(call *v1.InstancesStartCall, opts ...googleapi.CallOption) (*v1.Operation, error) {
	return &v1.Operation{}, nil
}

// Do is the mock for default regionsGetMock
func (c *projectsSetCommonInstanceMetadataMock) Do(call *v1.ProjectsSetCommonInstanceMetadataCall, opts ...googleapi.CallOption) (*v1.Operation, error) {
	return &v1.Operation{}, nil
}

// Do is the mock for default projectsGetMock
func (c *projectsGetMock) Do(call *v1.ProjectsGetCall, opts ...googleapi.CallOption) (*v1.Project, error) {
	return &v1.Project{
		CommonInstanceMetadata: &v1.Metadata{
			Items: []*v1.MetadataItems{},
		},
	}, nil
}

// Do is the mock for default targetPoolsListMock
func (c *targetPoolsListMock) Do(call *v1.TargetPoolsAggregatedListCall, opts ...googleapi.CallOption) (*v1.TargetPoolAggregatedList, error) {
	return &v1.TargetPoolAggregatedList{
		Items: map[string]v1.TargetPoolsScopedList{
			"item": v1.TargetPoolsScopedList{
				TargetPools: []*v1.TargetPool{
					&v1.TargetPool{
						Name:   "pool",
						Region: "us-central1",
					},
				},
			},
		},
	}, nil
}

// Do is the mock for default targetPoolDeleteMock
func (c *targetPoolDeleteMock) Do(call *v1.TargetPoolsDeleteCall, opts ...googleapi.CallOption) (*v1.Operation, error) {
	return &v1.Operation{}, nil
}

// Do is the mock for default backendServicesListMock
func (c *backendServicesListMock) Do(call *v1.BackendServicesAggregatedListCall, opts ...googleapi.CallOption) (*v1.BackendServiceAggregatedList, error) {
	return &v1.BackendServiceAggregatedList{
		Items: map[string]v1.BackendServicesScopedList{
			"item": v1.BackendServicesScopedList{
				BackendServices: []*v1.BackendService{
					&v1.BackendService{
						Name: "backend-service",
					},
				},
			},
		},
	}, nil
}

// Do is the mock for default backendServiceDeleteMock
func (c *backendServiceDeleteMock) Do(call *v1.BackendServicesDeleteCall, opts ...googleapi.CallOption) (*v1.Operation, error) {
	return &v1.Operation{}, nil
}

// Do is the mock for default regionBackendServiceDeleteMock
func (c *regionBackendServiceDeleteMock) Do(call *v1.RegionBackendServicesDeleteCall, opts ...googleapi.CallOption) (*v1.Operation, error) {
	return &v1.Operation{}, nil
}

// Do is the mock for default forwardingRulesListMock
func (c *forwardingRulesListMock) Do(call *v1.ForwardingRulesAggregatedListCall, opts ...googleapi.CallOption) (*v1.ForwardingRuleAggregatedList, error) {
	return &v1.ForwardingRuleAggregatedList{
		Items: map[string]v1.ForwardingRulesScopedList{
			"item": v1.ForwardingRulesScopedList{
				ForwardingRules: []*v1.ForwardingRule{
					&v1.ForwardingRule{
						Name:   "forwarding-rule",
						Region: "us-central1",
					},
				},
			},
		},
	}, nil
}

// Do is the mock for default forwardingRuleDeleteMock
func (c *forwardingRuleDeleteMock) Do(call *v1.ForwardingRulesDeleteCall, opts ...googleapi.CallOption) (*v1.Operation, error) {
	return &v1.Operation{}, nil
}

// Do is the mock for default healthChecksListMock
func (c *healthChecksListMock) Do(call *v1.HealthChecksAggregatedListCall, opts ...googleapi.CallOption) (*v1.HealthChecksAggregatedList, error) {
	return &v1.HealthChecksAggregatedList{
		Items: map[string]v1.HealthChecksScopedList{
			"item": v1.HealthChecksScopedList{
				HealthChecks: []*v1.HealthCheck{
					&v1.HealthCheck{
						Name: "healthCheck",
					},
				},
			},
		},
	}, nil
}

// Do is the mock for default healthCheckDeleteMock
func (c *healthCheckDeleteMock) Do(call *v1.HealthChecksDeleteCall, opts ...googleapi.CallOption) (*v1.Operation, error) {
	return &v1.Operation{}, nil
}

// Do is the mock for default httpHealthChecksListMock
func (c *httpHealthChecksListMock) Do(call *v1.HttpHealthChecksListCall, opts ...googleapi.CallOption) (*v1.HttpHealthCheckList, error) {
	return &v1.HttpHealthCheckList{
		Items: []*v1.HttpHealthCheck{
			&v1.HttpHealthCheck{
				Name: "healthCheck",
			},
		},
	}, nil
}

// Do is the mock for default httpHealthCheckDeleteMock
func (c *httpHealthCheckDeleteMock) Do(call *v1.HttpHealthChecksDeleteCall, opts ...googleapi.CallOption) (*v1.Operation, error) {
	return &v1.Operation{}, nil
}

// Do is the mock for default disksListMock
func (c *disksListMock) Do(call *v1.DisksAggregatedListCall, opts ...googleapi.CallOption) (*v1.DiskAggregatedList, error) {
	return &v1.DiskAggregatedList{
		Items: map[string]v1.DisksScopedList{
			"item": v1.DisksScopedList{
				Disks: []*v1.Disk{
					&v1.Disk{
						Name: "disk",
						Zone: "us-central1-a",
					},
				},
			},
		},
	}, nil
}

// Do is the mock for default diskDeleteMock
func (c *diskDeleteMock) Do(call *v1.DisksDeleteCall, opts ...googleapi.CallOption) (*v1.Operation, error) {
	return &v1.Operation{}, nil
}

// Do is the mock for default addressesListMock
func (c *addressesListMock) Do(call *v1.AddressesAggregatedListCall, opts ...googleapi.CallOption) (*v1.AddressAggregatedList, error) {
	return &v1.AddressAggregatedList{
		Items: map[string]v1.AddressesScopedList{
			"item": v1.AddressesScopedList{
				Addresses: []*v1.Address{
					&v1.Address{
						Name:   "address",
						Region: "us-central1",
					},
				},
			},
		},
	}, nil
}

// Do is the mock for default addressDeleteMock
func (c *addressDeleteMock) Do(call *v1.AddressesDeleteCall, opts ...googleapi.CallOption) (*v1.Operation, error) {
	return &v1.Operation{}, nil
}

// Do is the mock for default firewallsListMock
func (c *firewallsListMock) Do(call *v1.FirewallsListCall, opts ...googleapi.CallOption) (*v1.FirewallList, error) {
	return &v1.FirewallList{
		Items: []*v1.Firewall{
			&v1.Firewall{
				Name: "firewall",
			},
		},
	}, nil
}

// Do is the mock for default firewallDeleteMock
func (c *firewallDeleteMock) Do(call *v1.FirewallsDeleteCall, opts ...googleapi.CallOption) (*v1.Operation, error) {
	return &v1.Operation{}, nil
}

// Do is the mock for default instanceGroupsListMock
func (c *instanceGroupsListMock) Do(call *v1.InstanceGroupsAggregatedListCall, opts ...googleapi.CallOption) (*v1.InstanceGroupAggregatedList, error) {
	return &v1.InstanceGroupAggregatedList{
		Items: map[string]v1.InstanceGroupsScopedList{
			"item": v1.InstanceGroupsScopedList{
				InstanceGroups: []*v1.InstanceGroup{
					&v1.InstanceGroup{
						Name: "instance-group",
						Zone: "us-central1-a",
					},
				},
			},
		},
	}, nil
}

// Do is the mock for default instanceGroupDeleteMock
func (c *instanceGroupDeleteMock) Do(call *v1.InstanceGroupsDeleteCall, opts ...googleapi.CallOption) (*v1.Operation, error) {
	return &v1.Operation{}, nil
}

func (c *networkGetMock) Do(call *v1.NetworksGetCall, opts ...googleapi.CallOption) (*v1.Network, error) {
	return &v1.Network{}, nil
}

func (c *networkDeleteMock) Do(call *v1.NetworksDeleteCall, opts ...googleapi.CallOption) (*v1.Operation, error) {
	return &v1.Operation{}, nil
}

func setCallMockDefaults(c *Compute) {
	c.Calls = &Calls{
		RegionsGet:                        &regionsGetMock{},
		InstancesAggregatedList:           &instancesAggregatedListMock{},
		InstancesStop:                     &instancesStopMock{},
		InstancesStart:                    &instancesStartMock{},
		ProjectsSetCommonInstanceMetadata: &projectsSetCommonInstanceMetadataMock{},
		ProjectsGet:                       &projectsGetMock{},
		TargetPoolsList:                   &targetPoolsListMock{},
		TargetPoolDelete:                  &targetPoolDeleteMock{},
		BackendServicesList:               &backendServicesListMock{},
		BackendServiceDelete:              &backendServiceDeleteMock{},
		RegionBackendServiceDelete:        &regionBackendServiceDeleteMock{},
		ForwardingRulesList:               &forwardingRulesListMock{},
		ForwardingRuleDelete:              &forwardingRuleDeleteMock{},
		HealthChecksList:                  &healthChecksListMock{},
		HealthCheckDelete:                 &healthCheckDeleteMock{},
		HTTPHealthChecksList:              &httpHealthChecksListMock{},
		HTTPHealthCheckDelete:             &httpHealthCheckDeleteMock{},
		DisksList:                         &disksListMock{},
		DiskDelete:                        &diskDeleteMock{},
		AddressesList:                     &addressesListMock{},
		AddressDelete:                     &addressDeleteMock{},
		FirewallsList:                     &firewallsListMock{},
		FirewallDelete:                    &firewallDeleteMock{},
		InstanceGroupsList:                &instanceGroupsListMock{},
		InstanceGroupDelete:               &instanceGroupDeleteMock{},
		NetworkGet:                        &networkGetMock{},
		NetworkDelete:                     &networkDeleteMock{},
	}
}

func TestInitialize(t *testing.T) {
	c := &Compute{}
	err := c.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during compute.Initialize() with blank credentials: %s", err)
	}
	err = c.Initialize(testCredentials, loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during compute.Initialize() with explicit credentials: %s", err)
	}
}

func TestGetRegionZones(t *testing.T) {
	c := &Compute{}
	err := c.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during compute.Initialize() with blank credentials: %s", err)
	}
	setCallMockDefaults(c)
	zones, err := c.GetRegionZones(testProjectID, "us-central1")
	if err != nil {
		t.Errorf("Got unexpected error during compute.GetRegionZones(): %s", err)
	}
	if len(zones) != 4 {
		t.Errorf("Got unexpected result/value from compute.GetRegionZones(), expecting length of \"4\", but got: %d", len(zones))
	}
}

func TestGetInternalIPs(t *testing.T) {
	c := &Compute{}
	err := c.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during compute.Initialize() with blank credentials: %s", err)
	}
	setCallMockDefaults(c)
	instanceIPs, err := c.GetInternalIPs(testProjectID, "three")
	if err != nil {
		t.Errorf("Got unexpected error during compute.GetInternalIPs(): %s", err)
	}
	if len(instanceIPs) != 2 {
		t.Errorf("Got unexpected result/value from compute.GetInternalIPs(), expecting length of \"2\", but got: %d", len(instanceIPs))
	}
}

func TestPowerOff(t *testing.T) {
	c := &Compute{}
	err := c.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during compute.Initialize() with blank credentials: %s", err)
	}
	setCallMockDefaults(c)
	err = c.PowerOff(testProjectID)
	if err != nil {
		t.Errorf("Got unexpected error testing PowerOff: %s", err)
	}
}

func TestPowerOn(t *testing.T) {
	c := &Compute{}
	err := c.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during compute.Initialize() with blank credentials: %s", err)
	}
	setCallMockDefaults(c)
	err = c.PowerOn(testProjectID)
	if err != nil {
		t.Errorf("Got unexpected error testing PowerOn: %s", err)
	}
}

func TestSetCommonInstanceMetadata(t *testing.T) {
	c := &Compute{}
	err := c.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during compute.Initialize() with blank credentials: %s", err)
	}
	setCallMockDefaults(c)
	value := "value"
	metadataItems := []*v1.MetadataItems{
		&v1.MetadataItems{
			Key:   "key",
			Value: &value,
		},
	}
	err = c.SetCommonInstanceMetadata("project", metadataItems)
	if err != nil {
		t.Errorf("Got unexpected error during compute.SetCommonInstanceMetadata(): %s", err)
	}
}

func TestGetCommonInstanceMetadata(t *testing.T) {
	c := &Compute{}
	err := c.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during compute.Initialize() with blank credentials: %s", err)
	}
	setCallMockDefaults(c)
	_, err = c.GetCommonInstanceMetadata("project")
	if err != nil {
		t.Errorf("Got unexpected error during compute.GetCommonInstanceMetadata(): %s", err)
	}
}

func TestGetTargetPools(t *testing.T) {
	c := &Compute{}
	err := c.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during compute.Initialize() with blank credentials: %s", err)
	}
	setCallMockDefaults(c)
	_, err = c.GetTargetPools("project")
	if err != nil {
		t.Errorf("Got unexpected error during compute.GetTargetPools(): %s", err)
	}
}

func TestDeleteTargetPool(t *testing.T) {
	c := &Compute{}
	err := c.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during compute.Initialize() with blank credentials: %s", err)
	}
	setCallMockDefaults(c)
	err = c.DeleteTargetPool("project", "us-central1", "pool")
	if err != nil {
		t.Errorf("Got unexpected error during compute.DeleteTargetPool(): %s", err)
	}
}

func TestDeleteForwardingRule(t *testing.T) {
	c := &Compute{}
	err := c.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during compute.Initialize() with blank credentials: %s", err)
	}
	setCallMockDefaults(c)
	err = c.DeleteForwardingRule("project", "us-central1", "rule")
	if err != nil {
		t.Errorf("Got unexpected error during compute.DeleteForwardingRule(): %s", err)
	}
}

func TestGetBackendServices(t *testing.T) {
	c := &Compute{}
	err := c.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during compute.Initialize() with blank credentials: %s", err)
	}
	setCallMockDefaults(c)
	_, err = c.GetBackendServices("project")
	if err != nil {
		t.Errorf("Got unexpected error during compute.GetBackendServices(): %s", err)
	}
}

func TestDeleteBackendService(t *testing.T) {
	c := &Compute{}
	err := c.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during compute.Initialize() with blank credentials: %s", err)
	}
	setCallMockDefaults(c)
	err = c.DeleteBackendService("project", "backend-service")
	if err != nil {
		t.Errorf("Got unexpected error during compute.DeleteBackendService(): %s", err)
	}
}

func TestDeleteRegionBackendService(t *testing.T) {
	c := &Compute{}
	err := c.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during compute.Initialize() with blank credentials: %s", err)
	}
	setCallMockDefaults(c)
	err = c.DeleteRegionBackendService("project", "region", "backend-service")
	if err != nil {
		t.Errorf("Got unexpected error during compute.DeleteRegionBackendService(): %s", err)
	}
}

func TestGetHealthChecks(t *testing.T) {
	c := &Compute{}
	err := c.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during compute.Initialize() with blank credentials: %s", err)
	}
	setCallMockDefaults(c)
	_, err = c.GetHealthChecks("project")
	if err != nil {
		t.Errorf("Got unexpected error during compute.GetHealthChecks(): %s", err)
	}
}

func TestDeleteHealthCheck(t *testing.T) {
	c := &Compute{}
	err := c.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during compute.Initialize() with blank credentials: %s", err)
	}
	setCallMockDefaults(c)
	err = c.DeleteHealthCheck("project", "health-check")
	if err != nil {
		t.Errorf("Got unexpected error during compute.DeleteHealthCheck(): %s", err)
	}
}

func TestGetHTTPHealthChecks(t *testing.T) {
	c := &Compute{}
	err := c.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during compute.Initialize() with blank credentials: %s", err)
	}
	setCallMockDefaults(c)
	_, err = c.GetHTTPHealthChecks("project")
	if err != nil {
		t.Errorf("Got unexpected error during compute.GetHTTPHealthChecks(): %s", err)
	}
}

func TestDeleteHTTPHealthCheck(t *testing.T) {
	c := &Compute{}
	err := c.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during compute.Initialize() with blank credentials: %s", err)
	}
	setCallMockDefaults(c)
	err = c.DeleteHTTPHealthCheck("project", "health-check")
	if err != nil {
		t.Errorf("Got unexpected error during compute.DeleteHTTPHealthCheck(): %s", err)
	}
}

func TestGetDisks(t *testing.T) {
	c := &Compute{}
	err := c.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during compute.Initialize() with blank credentials: %s", err)
	}
	setCallMockDefaults(c)
	_, err = c.GetDisks("project")
	if err != nil {
		t.Errorf("Got unexpected error during compute.GetDisks(): %s", err)
	}
}

func TestDeleteDisk(t *testing.T) {
	c := &Compute{}
	err := c.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during compute.Initialize() with blank credentials: %s", err)
	}
	setCallMockDefaults(c)
	err = c.DeleteDisk("project", "us-central1-a", "disk")
	if err != nil {
		t.Errorf("Got unexpected error during compute.DeleteDisk(): %s", err)
	}
}

func TestGetAddresses(t *testing.T) {
	c := &Compute{}
	err := c.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during compute.Initialize() with blank credentials: %s", err)
	}
	setCallMockDefaults(c)
	_, err = c.GetAddresses("project")
	if err != nil {
		t.Errorf("Got unexpected error during compute.GetAddresses(): %s", err)
	}
}

func TestDeleteAddress(t *testing.T) {
	c := &Compute{}
	err := c.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during compute.Initialize() with blank credentials: %s", err)
	}
	setCallMockDefaults(c)
	err = c.DeleteAddress("project", "us-central1", "address")
	if err != nil {
		t.Errorf("Got unexpected error during compute.DeleteAddress(): %s", err)
	}
}

func TestGetFirewalls(t *testing.T) {
	c := &Compute{}
	err := c.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during compute.Initialize() with blank credentials: %s", err)
	}
	setCallMockDefaults(c)
	_, err = c.GetFirewalls("project")
	if err != nil {
		t.Errorf("Got unexpected error during compute.GetFirewalls(): %s", err)
	}
}

func TestDeleteFirewall(t *testing.T) {
	c := &Compute{}
	err := c.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during compute.Initialize() with blank credentials: %s", err)
	}
	setCallMockDefaults(c)
	err = c.DeleteFirewall("project", "firewall")
	if err != nil {
		t.Errorf("Got unexpected error during compute.DeleteFirewall(): %s", err)
	}
}

func TestGetInstanceGroups(t *testing.T) {
	c := &Compute{}
	err := c.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during compute.Initialize() with blank credentials: %s", err)
	}
	setCallMockDefaults(c)
	_, err = c.GetInstanceGroups("project")
	if err != nil {
		t.Errorf("Got unexpected error during compute.GetInstanceGroups(): %s", err)
	}
}

func TestDeleteInstanceGroup(t *testing.T) {
	c := &Compute{}
	err := c.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during compute.Initialize() with blank credentials: %s", err)
	}
	setCallMockDefaults(c)
	err = c.DeleteInstanceGroup("project", "zone", "instance-group")
	if err != nil {
		t.Errorf("Got unexpected error during compute.DeleteInstanceGroup(): %s", err)
	}
}

func TestGetNetwork(t *testing.T) {
	c := &Compute{}
	err := c.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during compute.Initialize() with blank credentials: %s", err)
	}
	setCallMockDefaults(c)
	_, err = c.GetNetwork("project", "network")
	if err != nil {
		t.Errorf("Got unexpected error during compute.TestGetNetwork(): %s", err)
	}
}

func TestDeleteNetwork(t *testing.T) {
	c := &Compute{}
	err := c.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during compute.Initialize() with blank credentials: %s", err)
	}
	setCallMockDefaults(c)
	err = c.DeleteNetwork("project", "network")
	if err != nil {
		t.Errorf("Got unexpected error during compute.TestDeleteNetwork(): %s", err)
	}
}
