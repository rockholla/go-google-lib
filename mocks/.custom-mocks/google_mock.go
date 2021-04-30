package googlemock

import (
	"github.com/rockholla/go-google-lib/admin"
	"github.com/rockholla/go-google-lib/cloudbilling"
	"github.com/rockholla/go-google-lib/cloudidentity"
	"github.com/rockholla/go-google-lib/cloudresourcemanager"
	"github.com/rockholla/go-google-lib/compute"
	"github.com/rockholla/go-google-lib/deploymentmanager"
	"github.com/rockholla/go-google-lib/dns"
	"github.com/rockholla/go-google-lib/iam"
	adminmock "github.com/rockholla/go-google-lib/mocks/admin"
	cloudbillingmock "github.com/rockholla/go-google-lib/mocks/cloudbilling"
	cloudidentitymock "github.com/rockholla/go-google-lib/mocks/cloudidentity"
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
// A user of this mock is meant to set each of the needed underlying APIs
// and what each mocks on this object, e.g.
//
// crmMock := &crmmocks.Interface{}
// crmMock.On("EnsureFolder", rootFolderDisplayName, fmt.Sprintf("organizations/%s", testOrganizationID)).Return(testRootFolderName, nil)
//
// mock := &googlelibmocks.GoogleMock{
// 	CloudResourceManager: crmMock,
// }
type GoogleMock struct {
	CloudResourceManager *cloudresourcemanagermock.Interface
	CloudBilling         *cloudbillingmock.Interface
	CloudIdentity        *cloudidentitymock.Interface
	Admin                *adminmock.Interface
	Compute              *computemock.Interface
	DeploymentManager    *deploymentmanagermock.Interface
	DNS                  *dnsmock.Interface
	IAM                  *iammock.Interface
	Storage              *storagemock.Interface
}

// Initialize is a no-op in the mock
func (m *GoogleMock) Initialize(credentials string, log logger.Interface) {}

// GetCloudResourceManager mock
func (m *GoogleMock) GetCloudResourceManager() (cloudresourcemanager.Interface, error) {
	return m.CloudResourceManager, nil
}

// GetCloudBilling mock
func (m *GoogleMock) GetCloudBilling() (cloudbilling.Interface, error) {
	return m.CloudBilling, nil
}

// GetCloudIdentity mock
func (m *GoogleMock) GetCloudIdentity() (cloudidentity.Interface, error) {
	return m.CloudIdentity, nil
}

// GetIAM mock
func (m *GoogleMock) GetIAM() (iam.Interface, error) {
	return m.IAM, nil
}

// GetDeploymentManager mock
func (m *GoogleMock) GetDeploymentManager() (deploymentmanager.Interface, error) {
	return m.DeploymentManager, nil
}

// GetStorage mock
func (m *GoogleMock) GetStorage() (storage.Interface, error) {
	return m.Storage, nil
}

// GetCompute mock
func (m *GoogleMock) GetCompute() (compute.Interface, error) {
	return m.Compute, nil
}

// GetDNS mock
func (m *GoogleMock) GetDNS() (dns.Interface, error) {
	return m.DNS, nil
}

// GetAdmin mock
func (m *GoogleMock) GetAdmin(credentialsJSON string, domain string, adminUsername string) (admin.Interface, error) {
	return m.Admin, nil
}
