package calls

import (
	v1 "google.golang.org/api/cloudbilling/v1"
	googleapi "google.golang.org/api/googleapi"
)

// ProjectsUpdateBillingInfoCallInterface is an interface to a call to update project billing info
type ProjectsUpdateBillingInfoCallInterface interface {
	Do(call *v1.ProjectsUpdateBillingInfoCall, opts ...googleapi.CallOption) (*v1.ProjectBillingInfo, error)
}

// ProjectsUpdateBillingInfoCall is the default implementation for ProjectsUpdateBillingInfoCallInterface
type ProjectsUpdateBillingInfoCall struct{}

// Do performs the call, the default implementation of the interface
func (c *ProjectsUpdateBillingInfoCall) Do(call *v1.ProjectsUpdateBillingInfoCall, opts ...googleapi.CallOption) (*v1.ProjectBillingInfo, error) {
	return call.Do(opts...)
}
