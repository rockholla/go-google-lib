// Package google is the root of the custom google api library
package google

import (
	"github.com/rockholla/go-google-lib/google/admin"
	"github.com/rockholla/go-google-lib/google/cloudbilling"
	"github.com/rockholla/go-google-lib/google/cloudresourcemanager"
	"github.com/rockholla/go-google-lib/google/compute"
	"github.com/rockholla/go-google-lib/google/deploymentmanager"
	"github.com/rockholla/go-google-lib/google/dns"
	"github.com/rockholla/go-google-lib/google/iam"
	"github.com/rockholla/go-google-lib/google/storage"
	"github.com/rockholla/go-lib/logger"
)

// Interface is the interface for all google api/sdk libraries
type Interface interface {
	Initialize(credentials string, log logger.Interface)
	GetCloudResourceManager() (cloudresourcemanager.Interface, error)
	GetCloudBilling() (cloudbilling.Interface, error)
	GetIAM() (iam.Interface, error)
	GetDeploymentManager() (deploymentmanager.Interface, error)
	GetStorage() (storage.Interface, error)
	GetCompute() (compute.Interface, error)
	GetDNS() (dns.Interface, error)
	GetAdmin(credentialsJSON string, domain string, adminUsername string) (admin.Interface, error)
}

// Google is all related api/sdk libraries
type Google struct {
	credentials          string
	log                  logger.Interface
	cloudResourceManager cloudresourcemanager.Interface
	cloudBilling         cloudbilling.Interface
	iam                  iam.Interface
	deploymentManager    deploymentmanager.Interface
	storage              storage.Interface
	compute              compute.Interface
	dns                  dns.Interface
	admin                admin.Interface
}

// Initialize will set initial values for all libraries: credentials, logger
func (google *Google) Initialize(credentials string, log logger.Interface) {
	google.log = log
	google.credentials = credentials
	if google.credentials != "" {
		google.log.Info("Using provided Google credentials key")
	} else {
		google.log.Info("Using Google default application credentials")
	}
}

// GetCloudResourceManager will get the cloud resource manager library
func (google *Google) GetCloudResourceManager() (cloudresourcemanager.Interface, error) {
	var err error
	if google.cloudResourceManager == nil {
		google.cloudResourceManager = &cloudresourcemanager.CloudResourceManager{}
		err = google.cloudResourceManager.Initialize(google.credentials, google.log)
	}
	return google.cloudResourceManager, err
}

// GetCloudBilling will get the cloud billing library
func (google *Google) GetCloudBilling() (cloudbilling.Interface, error) {
	var err error
	if google.cloudBilling == nil {
		google.cloudBilling = &cloudbilling.CloudBilling{}
		err = google.cloudBilling.Initialize(google.credentials, google.log)
	}
	return google.cloudBilling, err
}

// GetIAM will get the IAM library
func (google *Google) GetIAM() (iam.Interface, error) {
	var err error
	if google.iam == nil {
		google.iam = &iam.IAM{}
		err = google.iam.Initialize(google.credentials, google.log)
	}
	return google.iam, err
}

// GetDeploymentManager will get the deployment manager library
func (google *Google) GetDeploymentManager() (deploymentmanager.Interface, error) {
	var err error
	if google.deploymentManager == nil {
		google.deploymentManager = &deploymentmanager.DeploymentManager{}
		err = google.deploymentManager.Initialize(google.credentials, google.log)
	}
	return google.deploymentManager, err
}

// GetStorage will get the storage library
func (google *Google) GetStorage() (storage.Interface, error) {
	var err error
	if google.storage == nil {
		google.storage = &storage.Storage{}
		err = google.storage.Initialize(google.credentials, google.log)
	}
	return google.storage, err
}

// GetCompute will get the compute library
func (google *Google) GetCompute() (compute.Interface, error) {
	var err error
	if google.compute == nil {
		google.compute = &compute.Compute{}
		err = google.compute.Initialize(google.credentials, google.log)
	}
	return google.compute, err
}

// GetDNS will get the dns library
func (google *Google) GetDNS() (dns.Interface, error) {
	var err error
	if google.dns == nil {
		google.dns = &dns.DNS{}
		err = google.dns.Initialize(google.credentials, google.log)
	}
	return google.dns, err
}

// GetAdmin will get the admin library
func (google *Google) GetAdmin(credentialsJSON string, domain string, adminUsername string) (admin.Interface, error) {
	var err error
	if google.admin == nil {
		google.admin = &admin.Admin{}
		err = google.admin.Initialize(credentialsJSON, domain, adminUsername, google.log)
	}
	return google.admin, err
}
