package calls

import (
	v1 "google.golang.org/api/cloudresourcemanager/v1"
	googleapi "google.golang.org/api/googleapi"
	suv1 "google.golang.org/api/serviceusage/v1"
)

// ProjectsListCallInterface is an interface to a call to list projects
type ProjectsListCallInterface interface {
	Do(call *v1.ProjectsListCall, opts ...googleapi.CallOption) (*v1.ListProjectsResponse, error)
}

// ProjectsGetCallInterface is an interface to a call to get a single project
type ProjectsGetCallInterface interface {
	Do(call *v1.ProjectsGetCall, opts ...googleapi.CallOption) (*v1.Project, error)
}

// ProjectsCreateCallInterface is an interface to a call to create a project
type ProjectsCreateCallInterface interface {
	Do(call *v1.ProjectsCreateCall, opts ...googleapi.CallOption) (*v1.Operation, error)
}

// ProjectsDeleteCallInterface is an interface to a call to delete a project
type ProjectsDeleteCallInterface interface {
	Do(call *v1.ProjectsDeleteCall, opts ...googleapi.CallOption) (*v1.Empty, error)
}

// ProjectsGetIAMPolicyCallInterface is an interface to a call to get the iam policy for a project
type ProjectsGetIAMPolicyCallInterface interface {
	Do(call *v1.ProjectsGetIamPolicyCall, opts ...googleapi.CallOption) (*v1.Policy, error)
}

// ProjectsSetIAMPolicyCallInterface is an interface to a call to set the iam policy for a project
type ProjectsSetIAMPolicyCallInterface interface {
	Do(call *v1.ProjectsSetIamPolicyCall, opts ...googleapi.CallOption) (*v1.Policy, error)
}

// ServiceEnableCallInterface is an interface to call to enable a service/api on a project
type ServiceEnableCallInterface interface {
	Do(call *suv1.ServicesEnableCall, opts ...googleapi.CallOption) (*suv1.Operation, error)
}

// ProjectsListCall is the default implementation for ProjectsListCallInterface
type ProjectsListCall struct{}

// ProjectsGetCall is the default implementation for ProjectsGetCallInterface
type ProjectsGetCall struct{}

// ProjectsCreateCall is the default implementation for ProjectsCreateCallInterface
type ProjectsCreateCall struct{}

// ProjectsDeleteCall is the default implementation for ProjectsDeleteCallInterface
type ProjectsDeleteCall struct{}

// ProjectsGetIAMPolicyCall is the default implementation for ProjectsGetIAMPolicyCallInterface
type ProjectsGetIAMPolicyCall struct{}

// ProjectsSetIAMPolicyCall is the default implementation for ProjectsSetIAMPolicyCallInterface
type ProjectsSetIAMPolicyCall struct{}

// ServiceEnableCall is the default implementation for ServiceEnableCallInterface
type ServiceEnableCall struct{}

// Do performs the call, the default implementation of the interface
func (c *ProjectsListCall) Do(call *v1.ProjectsListCall, opts ...googleapi.CallOption) (*v1.ListProjectsResponse, error) {
	return call.Do(opts...)
}

// Do performs the call, the default implementation of the interface
func (c *ProjectsGetCall) Do(call *v1.ProjectsGetCall, opts ...googleapi.CallOption) (*v1.Project, error) {
	return call.Do(opts...)
}

// Do performs the call, the default implementation of the interface
func (c *ProjectsCreateCall) Do(call *v1.ProjectsCreateCall, opts ...googleapi.CallOption) (*v1.Operation, error) {
	return call.Do(opts...)
}

// Do performs the call, the default implementation of the interface
func (c *ProjectsDeleteCall) Do(call *v1.ProjectsDeleteCall, opts ...googleapi.CallOption) (*v1.Empty, error) {
	return call.Do(opts...)
}

// Do performs the call, the default implementation of the interface
func (c *ProjectsGetIAMPolicyCall) Do(call *v1.ProjectsGetIamPolicyCall, opts ...googleapi.CallOption) (*v1.Policy, error) {
	return call.Do(opts...)
}

// Do performs the call, the default implementation of the interface
func (c *ProjectsSetIAMPolicyCall) Do(call *v1.ProjectsSetIamPolicyCall, opts ...googleapi.CallOption) (*v1.Policy, error) {
	return call.Do(opts...)
}

// Do performs the call, the default implementation of the interface
func (c *ServiceEnableCall) Do(call *suv1.ServicesEnableCall, opts ...googleapi.CallOption) (*suv1.Operation, error) {
	return call.Do(opts...)
}
