// Package compute is the library for google cloud compute operations
package compute

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/rockholla/go-google-lib/compute/calls"
	"github.com/rockholla/go-lib/logger"
	v1 "google.golang.org/api/compute/v1"
	"google.golang.org/api/option"
)

// Local utility function to extract a zone string from a zone URL
func urlZone(url string) string {
	re := regexp.MustCompile(`[^/]*$`)
	return string(re.Find([]byte(url)))
}

// Interface represents functionality for DeploymentManager
type Interface interface {
	Initialize(credentials string, log logger.Interface) error
	GetRegionZones(projectID string, region string) ([]string, error)
	GetInternalIPs(projectID string, network string) ([]*InstanceIP, error)
	PowerOff(projectID string) error
	PowerOn(projectID string) error
	SetCommonInstanceMetadata(projectID string, metadataItems []*v1.MetadataItems) error
	GetCommonInstanceMetadata(projectID string) ([]*v1.MetadataItems, error)
	GetTargetPools(projectID string) ([]*v1.TargetPool, error)
	DeleteTargetPool(projectID string, region string, name string) error
	DeleteForwardingRule(projectID string, region string, name string) error
	GetBackendServices(projectID string) ([]*v1.BackendService, error)
	DeleteBackendService(projectID string, name string) error
	DeleteRegionBackendService(projectID string, region string, name string) error
	GetHealthChecks(projectID string) ([]*v1.HealthCheck, error)
	DeleteHealthCheck(projectID string, name string) error
	GetHTTPHealthChecks(projectID string) ([]*v1.HttpHealthCheck, error)
	DeleteHTTPHealthCheck(projectID string, name string) error
	GetDisks(projectID string) ([]*v1.Disk, error)
	DeleteDisk(projectID string, zone string, name string) error
	GetAddresses(projectID string) ([]*v1.Address, error)
	DeleteAddress(projectID string, region string, name string) error
	GetFirewalls(projectID string) ([]*v1.Firewall, error)
	DeleteFirewall(projectID string, name string) error
	GetInstanceGroups(projectID string) ([]*v1.InstanceGroup, error)
	DeleteInstanceGroup(projectID string, zone string, name string) error
	GetNetwork(projectID string, name string) (*v1.Network, error)
	DeleteNetwork(projectID string, name string) error
}

// InstanceIP is an IP for a VM instance
type InstanceIP struct {
	VMName string
	IP     string
}

// Compute is a wrapper around the google-provided sdks/apis for google.golang.org/api/compute/*
type Compute struct {
	log   logger.Interface
	V1    *v1.Service
	Calls *Calls
}

type cachedForwardingRule struct {
	Name   string
	Region string
}

// Calls are interfaces for making the actual calls to various underlying apis
type Calls struct {
	RegionsGet                        calls.RegionsGetCallInterface
	InstancesAggregatedList           calls.InstancesAggregatedListCallInterface
	InstancesStop                     calls.InstancesStopCallInterface
	InstancesStart                    calls.InstancesStartCallInterface
	ProjectsSetCommonInstanceMetadata calls.ProjectsSetCommonInstanceMetadataCallInterface
	ProjectsGet                       calls.ProjectsGetCallInterface
	TargetPoolsList                   calls.TargetPoolsListCallInterface
	TargetPoolDelete                  calls.TargetPoolDeleteCallInterface
	BackendServicesList               calls.BackendServicesListCallInterface
	BackendServiceDelete              calls.BackendServiceDeleteCallInterface
	RegionBackendServiceDelete        calls.RegionBackendServiceDeleteCallInterface
	ForwardingRulesList               calls.ForwardingRulesListCallInterface
	ForwardingRuleDelete              calls.ForwardingRuleDeleteCallInterface
	HealthChecksList                  calls.HealthChecksListCallInterface
	HealthCheckDelete                 calls.HealthCheckDeleteCallInterface
	HTTPHealthChecksList              calls.HTTPHealthChecksListCallInterface
	HTTPHealthCheckDelete             calls.HTTPHealthCheckDeleteCallInterface
	DisksList                         calls.DisksListCallInterface
	DiskDelete                        calls.DiskDeleteCallInterface
	AddressesList                     calls.AddressesListCallInterface
	AddressDelete                     calls.AddressDeleteCallInterface
	FirewallsList                     calls.FirewallsListCallInterface
	FirewallDelete                    calls.FirewallDeleteCallInterface
	InstanceGroupsList                calls.InstanceGroupsListCallInterface
	InstanceGroupDelete               calls.InstanceGroupDeleteCallInterface
	NetworkGet                        calls.NetworkGetCallInterface
	NetworkDelete                     calls.NetworkDeleteCallInterface
}

// Initialize sets up necessary google-provided sdks and other local data
func (c *Compute) Initialize(credentials string, log logger.Interface) error {
	var err error
	ctx := context.Background()
	c.log = log
	c.Calls = &Calls{
		RegionsGet:                        &calls.RegionsGetCall{},
		InstancesAggregatedList:           &calls.InstancesAggregatedListCall{},
		InstancesStop:                     &calls.InstancesStopCall{},
		InstancesStart:                    &calls.InstancesStartCall{},
		ProjectsSetCommonInstanceMetadata: &calls.ProjectsSetCommonInstanceMetadataCall{},
		ProjectsGet:                       &calls.ProjectsGetCall{},
		TargetPoolsList:                   &calls.TargetPoolsListCall{},
		TargetPoolDelete:                  &calls.TargetPoolDeleteCall{},
		BackendServicesList:               &calls.BackendServicesListCall{},
		BackendServiceDelete:              &calls.BackendServiceDeleteCall{},
		RegionBackendServiceDelete:        &calls.RegionBackendServiceDeleteCall{},
		ForwardingRulesList:               &calls.ForwardingRulesListCall{},
		ForwardingRuleDelete:              &calls.ForwardingRuleDeleteCall{},
		HealthChecksList:                  &calls.HealthChecksListCall{},
		HealthCheckDelete:                 &calls.HealthCheckDeleteCall{},
		HTTPHealthChecksList:              &calls.HTTPHealthChecksListCall{},
		HTTPHealthCheckDelete:             &calls.HTTPHealthCheckDeleteCall{},
		DisksList:                         &calls.DisksListCall{},
		DiskDelete:                        &calls.DiskDeleteCall{},
		AddressesList:                     &calls.AddressesListCall{},
		AddressDelete:                     &calls.AddressDeleteCall{},
		FirewallsList:                     &calls.FirewallsListCall{},
		FirewallDelete:                    &calls.FirewallDeleteCall{},
		InstanceGroupsList:                &calls.InstanceGroupsListCall{},
		InstanceGroupDelete:               &calls.InstanceGroupDeleteCall{},
		NetworkGet:                        &calls.NetworkGetCall{},
		NetworkDelete:                     &calls.NetworkDeleteCall{},
	}
	if credentials != "" {
		if c.V1, err = v1.NewService(ctx, option.WithCredentialsJSON([]byte(credentials))); err != nil {
			return err
		}
	} else {
		if c.V1, err = v1.NewService(ctx); err != nil {
			return err
		}
	}
	return nil
}

// GetRegionZones will return a list of zone names available in a region
func (c *Compute) GetRegionZones(projectID string, region string) ([]string, error) {
	ctx := context.Background()
	regionsService := v1.NewRegionsService(c.V1)
	regionsGetCall := regionsService.Get(projectID, region).Context(ctx)
	r, err := c.Calls.RegionsGet.Do(regionsGetCall)
	if err != nil {
		return []string{}, err
	}
	return r.Zones, nil
}

// GetInternalIPs will return a list of InstanceIP objects, which includes the name and internal
// IP for the VMName on the network interface attached to the specified network name
func (c *Compute) GetInternalIPs(projectID string, network string) ([]*InstanceIP, error) {
	ctx := context.Background()
	var result []*InstanceIP
	instancesService := v1.NewInstancesService(c.V1)
	instancesListCall := instancesService.AggregatedList(projectID).Context(ctx).MaxResults(1000)
	instancesListResult, err := c.Calls.InstancesAggregatedList.Do(instancesListCall)
	if err != nil {
		return result, err
	}
	for _, items := range instancesListResult.Items {
		for _, instance := range items.Instances {
			ip := ""
			for _, networkInterface := range instance.NetworkInterfaces {
				if strings.Contains(networkInterface.Network, fmt.Sprintf("projects/%s/global/networks/%s", projectID, network)) {
					ip = networkInterface.NetworkIP
				}
			}
			if ip != "" {
				result = append(result, &InstanceIP{
					VMName: instance.Name,
					IP:     ip,
				})
			}
		}
	}
	return result, nil
}

// PowerOff will shut down all instances
func (c Compute) PowerOff(projectID string) error {
	ctx := context.Background()
	instancesService := v1.NewInstancesService(c.V1)
	instancesListCall := instancesService.AggregatedList(projectID).Context(ctx).MaxResults(1000)
	instancesListResult, err := c.Calls.InstancesAggregatedList.Do(instancesListCall)
	if err != nil {
		return err
	}
	// Go through the instances and stop them
	for _, items := range instancesListResult.Items {
		for _, instance := range items.Instances {
			zone := urlZone(instance.Zone)
			instancesStopCall := instancesService.Stop(projectID, zone, instance.Name).Context(ctx)
			_, err = c.Calls.InstancesStop.Do(instancesStopCall)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// PowerOn will start all instances in a project.
func (c Compute) PowerOn(projectID string) error {
	ctx := context.Background()
	instancesService := v1.NewInstancesService(c.V1)
	instancesListCall := instancesService.AggregatedList(projectID).Context(ctx).MaxResults(1000)
	instancesListResult, err := c.Calls.InstancesAggregatedList.Do(instancesListCall)
	if err != nil {
		return err
	}
	// Go through the instances and start them
	for _, items := range instancesListResult.Items {
		for _, instance := range items.Instances {
			zone := urlZone(instance.Zone)
			instancesStartCall := instancesService.Start(projectID, zone, instance.Name).Context(ctx)
			_, err = c.Calls.InstancesStart.Do(instancesStartCall)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// SetCommonInstanceMetadata will set project-level metadata to be used by any compute instance
func (c *Compute) SetCommonInstanceMetadata(projectID string, metadataItems []*v1.MetadataItems) error {
	ctx := context.Background()
	projectsService := v1.NewProjectsService(c.V1)
	metadata := &v1.Metadata{
		Items: metadataItems,
	}
	setCommonInstanceMetadataCall := projectsService.SetCommonInstanceMetadata(projectID, metadata).Context(ctx)
	_, err := c.Calls.ProjectsSetCommonInstanceMetadata.Do(setCommonInstanceMetadataCall)
	return err
}

// GetCommonInstanceMetadata will get project-level compute metadata
func (c *Compute) GetCommonInstanceMetadata(projectID string) ([]*v1.MetadataItems, error) {
	ctx := context.Background()
	projectsService := v1.NewProjectsService(c.V1)
	getProjectCall := projectsService.Get(projectID).Context(ctx)
	project, err := c.Calls.ProjectsGet.Do(getProjectCall)
	if err != nil {
		return []*v1.MetadataItems{}, err
	}
	return project.CommonInstanceMetadata.Items, nil
}

// GetTargetPools will return a list of all target pools/load balancer using target pools
func (c *Compute) GetTargetPools(projectID string) ([]*v1.TargetPool, error) {
	var list []*v1.TargetPool
	var err error
	ctx := context.Background()
	targetPoolsService := v1.NewTargetPoolsService(c.V1)
	targetPoolsListCall := targetPoolsService.AggregatedList(projectID).Context(ctx)
	result, err := c.Calls.TargetPoolsList.Do(targetPoolsListCall)
	if err != nil {
		return list, err
	}
	for _, item := range result.Items {
		for _, pool := range item.TargetPools {
			list = append(list, pool)
		}
	}
	return list, nil
}

// GetBackendServices will return a list of all load balancer backend services
func (c *Compute) GetBackendServices(projectID string) ([]*v1.BackendService, error) {
	var list []*v1.BackendService
	var err error
	ctx := context.Background()
	backendServicesService := v1.NewBackendServicesService(c.V1)
	backendServicesListCall := backendServicesService.AggregatedList(projectID).Context(ctx)
	result, err := c.Calls.BackendServicesList.Do(backendServicesListCall)
	if err != nil {
		return list, err
	}
	for _, item := range result.Items {
		for _, backendService := range item.BackendServices {
			list = append(list, backendService)
		}
	}
	return list, nil
}

// DeleteTargetPool will delete a single load balancer/target pool
func (c *Compute) DeleteTargetPool(projectID string, region string, name string) error {
	var err error
	ctx := context.Background()
	targetPoolsService := v1.NewTargetPoolsService(c.V1)
	targetPoolsDeleteCall := targetPoolsService.Delete(projectID, c.getResourceNameFromURL(region), name).Context(ctx)
	_, err = c.Calls.TargetPoolDelete.Do(targetPoolsDeleteCall)
	return err
}

// DeleteForwardingRule will delete an LB forwarding rule
func (c *Compute) DeleteForwardingRule(projectID string, region string, name string) error {
	ctx := context.Background()
	forwardingRulesService := v1.NewForwardingRulesService(c.V1)
	forwardingRulesDeleteCall := forwardingRulesService.Delete(projectID, c.getResourceNameFromURL(region), name).Context(ctx)
	if _, err := c.Calls.ForwardingRuleDelete.Do(forwardingRulesDeleteCall); err != nil {
		return err
	}
	return nil
}

// DeleteBackendService will delete an LB backend service
func (c *Compute) DeleteBackendService(projectID string, name string) error {
	ctx := context.Background()
	backendServicesService := v1.NewBackendServicesService(c.V1)
	backendServiceDeleteCall := backendServicesService.Delete(projectID, name).Context(ctx)
	if _, err := c.Calls.BackendServiceDelete.Do(backendServiceDeleteCall); err != nil {
		return err
	}
	return nil
}

// DeleteRegionBackendService will delete an LB backend service in a region
func (c *Compute) DeleteRegionBackendService(projectID string, region string, name string) error {
	ctx := context.Background()
	region = c.getResourceNameFromURL(region)
	name = c.getResourceNameFromURL(name)
	backendServicesService := v1.NewRegionBackendServicesService(c.V1)
	backendServiceDeleteCall := backendServicesService.Delete(projectID, region, name).Context(ctx)
	if _, err := c.Calls.RegionBackendServiceDelete.Do(backendServiceDeleteCall); err != nil {
		return err
	}
	return nil
}

// GetHealthChecks will return a list of all health checks in a project
func (c *Compute) GetHealthChecks(projectID string) ([]*v1.HealthCheck, error) {
	var list []*v1.HealthCheck
	var err error
	ctx := context.Background()
	healthChecksService := v1.NewHealthChecksService(c.V1)
	healthChecksListCall := healthChecksService.AggregatedList(projectID).Context(ctx)
	result, err := c.Calls.HealthChecksList.Do(healthChecksListCall)
	if err != nil {
		return list, err
	}
	for _, item := range result.Items {
		for _, healthCheck := range item.HealthChecks {
			list = append(list, healthCheck)
		}
	}
	return list, nil
}

// DeleteHealthCheck will delete a compute health check
func (c *Compute) DeleteHealthCheck(projectID string, name string) error {
	var err error
	ctx := context.Background()
	healthChecksService := v1.NewHealthChecksService(c.V1)
	healthCheckDeleteCall := healthChecksService.Delete(projectID, c.getResourceNameFromURL(name)).Context(ctx)
	if _, err = c.Calls.HealthCheckDelete.Do(healthCheckDeleteCall); err != nil {
		return err
	}
	return nil
}

// GetHTTPHealthChecks will return a list of all http (legacy) health checks in a project
func (c *Compute) GetHTTPHealthChecks(projectID string) ([]*v1.HttpHealthCheck, error) {
	var list []*v1.HttpHealthCheck
	var err error
	ctx := context.Background()
	healthChecksService := v1.NewHttpHealthChecksService(c.V1)
	healthChecksListCall := healthChecksService.List(projectID).Context(ctx)
	result, err := c.Calls.HTTPHealthChecksList.Do(healthChecksListCall)
	if err != nil {
		return list, err
	}
	for _, item := range result.Items {
		list = append(list, item)
	}
	return list, nil
}

// DeleteHTTPHealthCheck will delete a compute http (legacy) health check
func (c *Compute) DeleteHTTPHealthCheck(projectID string, name string) error {
	var err error
	ctx := context.Background()
	healthChecksService := v1.NewHttpHealthChecksService(c.V1)
	healthCheckDeleteCall := healthChecksService.Delete(projectID, c.getResourceNameFromURL(name)).Context(ctx)
	if _, err = c.Calls.HTTPHealthCheckDelete.Do(healthCheckDeleteCall); err != nil {
		return err
	}
	return nil
}

// GetDisks will return a list of all disks
func (c *Compute) GetDisks(projectID string) ([]*v1.Disk, error) {
	var list []*v1.Disk
	ctx := context.Background()
	disksService := v1.NewDisksService(c.V1)
	disksListCall := disksService.AggregatedList(projectID).Context(ctx)
	result, err := c.Calls.DisksList.Do(disksListCall)
	if err != nil {
		return list, err
	}
	for _, item := range result.Items {
		for _, pool := range item.Disks {
			list = append(list, pool)
		}
	}
	return list, nil
}

// DeleteDisk will delete a single disk
func (c *Compute) DeleteDisk(projectID string, zone string, name string) error {
	ctx := context.Background()
	zone = c.getResourceNameFromURL(zone)
	disksService := v1.NewDisksService(c.V1)
	disksDeleteCall := disksService.Delete(projectID, zone, name).Context(ctx)
	_, err := c.Calls.DiskDelete.Do(disksDeleteCall)
	return err
}

// GetAddresses will return a list of all compute addresses
func (c *Compute) GetAddresses(projectID string) ([]*v1.Address, error) {
	var list []*v1.Address
	ctx := context.Background()
	addressesService := v1.NewAddressesService(c.V1)
	addressesListCall := addressesService.AggregatedList(projectID).Context(ctx)
	result, err := c.Calls.AddressesList.Do(addressesListCall)
	if err != nil {
		return list, err
	}
	for _, item := range result.Items {
		for _, address := range item.Addresses {
			list = append(list, address)
		}
	}
	return list, nil
}

// DeleteAddress will delete a single disk
func (c *Compute) DeleteAddress(projectID string, region string, name string) error {
	ctx := context.Background()
	region = c.getResourceNameFromURL(region)
	addressesService := v1.NewAddressesService(c.V1)
	addressesDeleteCall := addressesService.Delete(projectID, region, name).Context(ctx)
	_, err := c.Calls.AddressDelete.Do(addressesDeleteCall)
	return err
}

// GetFirewalls will return a list of all compute firewall rules
func (c *Compute) GetFirewalls(projectID string) ([]*v1.Firewall, error) {
	var list []*v1.Firewall
	ctx := context.Background()
	firewallsService := v1.NewFirewallsService(c.V1)
	firewallsListCall := firewallsService.List(projectID).Context(ctx)
	result, err := c.Calls.FirewallsList.Do(firewallsListCall)
	if err != nil {
		return list, err
	}
	for _, item := range result.Items {
		list = append(list, item)
	}
	return list, nil
}

// DeleteFirewall will delete a single firewall
func (c *Compute) DeleteFirewall(projectID string, name string) error {
	ctx := context.Background()
	name = c.getResourceNameFromURL(name)
	firewallsService := v1.NewFirewallsService(c.V1)
	firewallsDeleteCall := firewallsService.Delete(projectID, name).Context(ctx)
	_, err := c.Calls.FirewallDelete.Do(firewallsDeleteCall)
	return err
}

// GetInstanceGroups will return a list of all compute instance groups
func (c *Compute) GetInstanceGroups(projectID string) ([]*v1.InstanceGroup, error) {
	var list []*v1.InstanceGroup
	ctx := context.Background()
	instanceGroupsService := v1.NewInstanceGroupsService(c.V1)
	instanceGroupsListCall := instanceGroupsService.AggregatedList(projectID).Context(ctx)
	result, err := c.Calls.InstanceGroupsList.Do(instanceGroupsListCall)
	if err != nil {
		return list, err
	}
	for _, item := range result.Items {
		for _, instanceGroup := range item.InstanceGroups {
			list = append(list, instanceGroup)
		}
	}
	return list, nil
}

// DeleteInstanceGroup will delete a single instance group
func (c *Compute) DeleteInstanceGroup(projectID string, zone string, name string) error {
	ctx := context.Background()
	zone = c.getResourceNameFromURL(zone)
	name = c.getResourceNameFromURL(name)
	instanceGroupsService := v1.NewInstanceGroupsService(c.V1)
	instanceGroupsDeleteCall := instanceGroupsService.Delete(projectID, zone, name).Context(ctx)
	_, err := c.Calls.InstanceGroupDelete.Do(instanceGroupsDeleteCall)
	return err
}

// GetNetwork will retrieve an existing network in a project
func (c *Compute) GetNetwork(projectID string, name string) (*v1.Network, error) {
	ctx := context.Background()
	networksService := v1.NewNetworksService(c.V1)
	networkGetCall := networksService.Get(projectID, name).Context(ctx)
	return c.Calls.NetworkGet.Do(networkGetCall)
}

// DeleteNetwork will delete a network in a project
func (c *Compute) DeleteNetwork(projectID string, name string) error {
	ctx := context.Background()
	networksService := v1.NewNetworksService(c.V1)
	networkDeleteCall := networksService.Delete(projectID, name).Context(ctx)
	_, err := c.Calls.NetworkDelete.Do(networkDeleteCall)
	return err
}

func (c *Compute) getResourceNameFromURL(url string) string {
	parts := strings.Split(url, "/")
	return parts[len(parts)-1]
}
