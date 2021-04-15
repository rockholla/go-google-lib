package calls

import (
	v2beta "google.golang.org/api/deploymentmanager/v2beta"
	googleapi "google.golang.org/api/googleapi"
)

// ManifestsGetCallInterface is an interface to a call to get a deployment manifest/config
type ManifestsGetCallInterface interface {
	Do(call *v2beta.ManifestsGetCall, opts ...googleapi.CallOption) (*v2beta.Manifest, error)
}

// ManifestsGetCall is the default implementation for ManifestsGetCallInterface
type ManifestsGetCall struct{}

// Do performs the call, the default implementation of the interface
func (c *ManifestsGetCall) Do(call *v2beta.ManifestsGetCall, opts ...googleapi.CallOption) (*v2beta.Manifest, error) {
	return call.Do(opts...)
}
