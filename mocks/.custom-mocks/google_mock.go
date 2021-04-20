package googlemock

import (
	"github.com/rockholla/go-google-lib/admin"
	"github.com/rockholla/go-google-lib/cloudbilling"
	"github.com/rockholla/go-google-lib/cloudresourcemanager"
	"github.com/rockholla/go-google-lib/compute"
	"github.com/rockholla/go-google-lib/deploymentmanager"
	"github.com/rockholla/go-google-lib/dns"
	"github.com/rockholla/go-google-lib/iam"
	adminmock "github.com/rockholla/go-google-lib/mocks/admin"
	cloudbillingmock "github.com/rockholla/go-google-lib/mocks/cloudbilling"
	cloudresourcemanagermock "github.com/rockholla/go-google-lib/mocks/cloudresourcemanager"
	computemock "github.com/rockholla/go-google-lib/mocks/compute"
	deploymentmanagermock "github.com/rockholla/go-google-lib/mocks/deploymentmanager"
	dnsmock "github.com/rockholla/go-google-lib/mocks/dns"
	iammock "github.com/rockholla/go-google-lib/mocks/iam"
	storagemock "github.com/rockholla/go-google-lib/mocks/storage"
	"github.com/rockholla/go-google-lib/storage"
	"github.com/rockholla/go-lib/logger"
)

// GoogleMock is a mock of our root level access to different APIs
type GoogleMock struct {
	cloudResourceManager cloudresourcemanager.Interface
	cloudBilling         cloudbilling.Interface
	iam                  iam.Interface
	deploymentManager    deploymentmanager.Interface
	storage              storage.Interface
	compute              compute.Interface
	dns                  dns.Interface
	admin                admin.Interface
}

// Initialize is a no-op in the mock
func (google *GoogleMock) Initialize(credentials string, log logger.Interface) {}

// GetCloudResourceManager mock
func (google *GoogleMock) GetCloudResourceManager() (cloudresourcemanager.Interface, error) {
	return &cloudresourcemanagermock.Interface{}, nil
}

// GetCloudBilling mock
func (google *GoogleMock) GetCloudBilling() (cloudbilling.Interface, error) {
	return &cloudbillingmock.Interface{}, nil
}

// GetIAM mock
func (google *GoogleMock) GetIAM() (iam.Interface, error) {
	return &iammock.Interface{}, nil
}

// GetDeploymentManager mock
func (google *GoogleMock) GetDeploymentManager() (deploymentmanager.Interface, error) {
	return &deploymentmanagermock.Interface{}, nil
}

// GetStorage mock
func (google *GoogleMock) GetStorage() (storage.Interface, error) {
	return &storagemock.Interface{}, nil
}

// GetCompute mock
func (google *GoogleMock) GetCompute() (compute.Interface, error) {
	return &computemock.Interface{}, nil
}

// GetDNS mock
func (google *GoogleMock) GetDNS() (dns.Interface, error) {
	return &dnsmock.Interface{}, nil
}

// GetAdmin mock
func (google *GoogleMock) GetAdmin(credentialsJSON string, domain string, adminUsername string) (admin.Interface, error) {
	return &adminmock.Interface{}, nil
}
