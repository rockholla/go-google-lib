// Package google is the root of the custom google api library
package google

import (
	"github.com/rockholla/go-google-lib/admin"
	"github.com/rockholla/go-google-lib/cloudbilling"
	"github.com/rockholla/go-google-lib/cloudidentity"
	"github.com/rockholla/go-google-lib/cloudresourcemanager"
	"github.com/rockholla/go-google-lib/compute"
	"github.com/rockholla/go-google-lib/deploymentmanager"
	"github.com/rockholla/go-google-lib/dns"
	"github.com/rockholla/go-google-lib/iam"
	"github.com/rockholla/go-google-lib/oauth"
	"github.com/rockholla/go-google-lib/storage"
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
	GetCloudIdentity(impersonateServiceAccountEmail string) (cloudidentity.Interface, error)
	GetAdmin(credentialsJSON string, domain string, adminUsername string) (admin.Interface, error)
	GetOAuth(scopes []string) (oauth.Interface, error)
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
	cloudIdentity        cloudidentity.Interface
	admin                admin.Interface
	oauth                oauth.Interface
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

// GetCloudIdentity will get the cloud identity library
func (google *Google) GetCloudIdentity(impersonateServiceAccountEmail string) (cloudidentity.Interface, error) {
	var err error
	if google.cloudIdentity == nil {
		google.cloudIdentity = &cloudidentity.CloudIdentity{}
		err = google.cloudIdentity.Initialize(impersonateServiceAccountEmail, google.log)
	}
	return google.cloudIdentity, err
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

// GetOAuth will get the oauth library
func (google *Google) GetOAuth(scopes []string) (oauth.Interface, error) {
	var err error
	if google.oauth == nil {
		google.oauth = &oauth.OAuth{}
		err = google.oauth.Initialize(google.credentials, google.log, scopes)
	}
	return google.oauth, err
}
