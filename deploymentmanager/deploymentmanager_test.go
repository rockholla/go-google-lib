package deploymentmanager

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	loggermock "github.com/rockholla/go-lib/mocks/custom-mocks/logger/logger"
	v2beta "google.golang.org/api/deploymentmanager/v2beta"
	googleapi "google.golang.org/api/googleapi"
)

const (
	testProjectID      = "0000000000"
	testFolderName     = "folders/1111111111"
	testDeploymentName = "test-deployment"
	testResourceName   = "test-resource"
	testOperationName  = "test-operation"
	testManifestURL    = "http://deployments/manifest-928372722"
	testManifestConfig = `---
resources:
  - name: one
    properties:
      id: one-id
  - name: two
    properties: {}
outputs:
  - name: oneID
    value: id`
	testManifestLayout = `---
outputs:
  - name: oneID
    value: id
    finalValue: one-id`
	testProperties = `---
one: "1"
two: "2"`
	testProjectProperties = `---
projectId: "test-project-id"`
	testCredentials = `{
  "client_id": "xxxxxxx.apps.googleusercontent.com",
  "client_secret": "xxxxxxxxxxxxxxx",
  "refresh_token": "xxxxxxxxx",
  "type": "authorized_user"
}`
)

var (
	triggerResourceNotFound          = false
	triggerDeploymentNotFound        = false
	triggerGetDeploymentRetry        = false
	triggerCreateOperationError      = false
	triggerUpdateOperationError      = false
	triggerDeleteOperationError      = false
	triggerPostGetOperationError     = false
	triggerPostOperationProgressLoop = true
	triggerProjectProperties         = false
	defaultDeployment                = &Deployment{
		Imports: []*Import{
			&Import{
				Name: "import",
				Path: "./.test-fixtures/import.yaml",
			},
		},
		Resources: []*Resource{
			&Resource{
				Name: fmt.Sprintf("%s-parent", testResourceName),
				Type: "cloudresourcemanager.v1.project",
				Properties: map[string]interface{}{
					"name":      fmt.Sprintf("%s-parent", testResourceName),
					"projectId": fmt.Sprintf("%s-parent", testProjectID),
					"parent": map[string]string{
						"type": "folder",
						"id":   strings.Replace(testFolderName, "folders/", "", 1),
					},
				},
			},
			&Resource{
				Name: testResourceName,
				Type: "cloudresourcemanager.v1.project",
				Properties: map[string]interface{}{
					"name":      testResourceName,
					"projectId": testProjectID,
					"parent": map[string]string{
						"type": "folder",
						"id":   strings.Replace(testFolderName, "folders/", "", 1),
					},
				},
				Metadata: &ResourceMetadata{
					DependsOn: []string{
						fmt.Sprintf("%s-parent", testResourceName),
					},
				},
			},
		},
		Outputs: []*Output{
			&Output{
				Name:  "projectId",
				Value: "projectId",
			},
		},
	}
)

type resourcesGetMock struct{}
type resourcesGetNotFoundMock struct{}
type deploymentsGetMock struct{}
type deploymentsInsertMock struct{}
type deploymentsUpdateMock struct{}
type deploymentsDeleteMock struct{}
type operationsGetMock struct{}
type manifestsGetMock struct{}

// Do is the mock for default resourcesGetMock
func (c *resourcesGetMock) Do(call *v2beta.ResourcesGetCall, opts ...googleapi.CallOption) (*v2beta.Resource, error) {
	if triggerResourceNotFound {
		triggerResourceNotFound = false
		return &v2beta.Resource{}, errors.New("NotFound")
	}
	properties := testProperties
	if triggerProjectProperties {
		triggerProjectProperties = false
		properties = testProjectProperties
	}
	return &v2beta.Resource{
		Name:       testResourceName,
		Properties: properties,
	}, nil
}

// Do is the mock for default deploymentsGetMock
func (c *deploymentsGetMock) Do(call *v2beta.DeploymentsGetCall, opts ...googleapi.CallOption) (*v2beta.Deployment, error) {
	if triggerDeploymentNotFound {
		triggerDeploymentNotFound = false
		return &v2beta.Deployment{}, errors.New("NotFound")
	}
	if triggerGetDeploymentRetry {
		triggerGetDeploymentRetry = false
		return &v2beta.Deployment{}, errors.New("wait a few minutes")
	}
	return &v2beta.Deployment{
		Name:     testDeploymentName,
		Manifest: testManifestURL,
	}, nil
}

// Do is the mock for default deploymentsInsertMock
func (c *deploymentsInsertMock) Do(call *v2beta.DeploymentsInsertCall, opts ...googleapi.CallOption) (*v2beta.Operation, error) {
	if triggerCreateOperationError {
		triggerCreateOperationError = false
		return &v2beta.Operation{
			Name:     testOperationName,
			Progress: 100,
			Error: &v2beta.OperationError{
				Errors: []*v2beta.OperationErrorErrors{
					&v2beta.OperationErrorErrors{
						Code:    "200",
						Message: "Error in operation",
					},
				},
			},
		}, nil
	}
	return &v2beta.Operation{
		Name:     testOperationName,
		Progress: 100,
		Error:    nil,
	}, nil
}

// Do is the mock for default deploymentsUpdateMock
func (c *deploymentsUpdateMock) Do(call *v2beta.DeploymentsUpdateCall, opts ...googleapi.CallOption) (*v2beta.Operation, error) {
	if triggerUpdateOperationError {
		triggerUpdateOperationError = false
		return &v2beta.Operation{
			Name:     testOperationName,
			Progress: 100,
			Error: &v2beta.OperationError{
				Errors: []*v2beta.OperationErrorErrors{
					&v2beta.OperationErrorErrors{
						Code:    "200",
						Message: "Error in operation",
					},
				},
			},
		}, nil
	}
	return &v2beta.Operation{
		Name:     testOperationName,
		Progress: 100,
		Error:    nil,
	}, nil
}

// Do is the mock for default deploymentsDeleteMock
func (c *deploymentsDeleteMock) Do(call *v2beta.DeploymentsDeleteCall, opts ...googleapi.CallOption) (*v2beta.Operation, error) {
	if triggerDeleteOperationError {
		triggerDeleteOperationError = false
		return &v2beta.Operation{
			Name:     testOperationName,
			Progress: 100,
			Error: &v2beta.OperationError{
				Errors: []*v2beta.OperationErrorErrors{
					&v2beta.OperationErrorErrors{
						Code:    "200",
						Message: "Error in operation",
					},
				},
			},
		}, nil
	}
	return &v2beta.Operation{
		Name:     testOperationName,
		Progress: 100,
		Error:    nil,
	}, nil
}

// Do is the mock for default operationsGetMock
func (c *operationsGetMock) Do(call *v2beta.OperationsGetCall, opts ...googleapi.CallOption) (*v2beta.Operation, error) {
	if triggerPostGetOperationError {
		triggerPostGetOperationError = false
		return &v2beta.Operation{
			Name:     testOperationName,
			Progress: 100,
			Error: &v2beta.OperationError{
				Errors: []*v2beta.OperationErrorErrors{
					&v2beta.OperationErrorErrors{
						Code:    "200",
						Message: "Error in operation",
					},
				},
			},
		}, nil
	}
	if triggerPostOperationProgressLoop {
		triggerPostOperationProgressLoop = false
		return &v2beta.Operation{
			Name:     testOperationName,
			Progress: 90,
			Error:    nil,
		}, nil
	}
	return &v2beta.Operation{
		Name:     testOperationName,
		Progress: 100,
		Error:    nil,
	}, nil
}

// Do is the mock for default manifestsGetMock
func (c *manifestsGetMock) Do(call *v2beta.ManifestsGetCall, opts ...googleapi.CallOption) (*v2beta.Manifest, error) {
	return &v2beta.Manifest{
		Config: &v2beta.ConfigFile{
			Content: testManifestConfig,
		},
		Layout: testManifestLayout,
	}, nil
}

func setCallMockDefaults(dm *DeploymentManager) {
	dm.ProgressWaitSeconds = 1
	dm.RetryWaitSeconds = 1
	dm.Calls = &Calls{
		ResourcesGet:      &resourcesGetMock{},
		DeploymentsGet:    &deploymentsGetMock{},
		DeploymentsInsert: &deploymentsInsertMock{},
		DeploymentsUpdate: &deploymentsUpdateMock{},
		DeploymentsDelete: &deploymentsDeleteMock{},
		OperationsGet:     &operationsGetMock{},
		ManifestsGet:      &manifestsGetMock{},
	}
}

func TestInitialize(t *testing.T) {
	dm := &DeploymentManager{}
	err := dm.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during deploymentmanager.Initialize() with blank credentials: %s", err)
	}
	err = dm.Initialize(testCredentials, loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during deploymentmanager.Initialize() with explicit credentials: %s", err)
	}
}

func TestGetResourcePropertyValue(t *testing.T) {
	dm := &DeploymentManager{}
	err := dm.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during deploymentmanager.Initialize() with blank credentials: %s", err)
	}
	setCallMockDefaults(dm)
	value, err := dm.GetResourcePropertyValue(testDeploymentName, testProjectID, testResourceName, "one")
	if err != nil {
		t.Errorf("Got unexpected error during deploymentmanager.GetResourcePropertyValue(): %s", err)
	}
	if value != "1" {
		t.Errorf("Got unexpected result/value from deploymentmanager.GetResourcePropertyValue(), expecting \"1\", but got: %s", value)
	}
}

func TestGetResourcePropertyValueResourceNotFound(t *testing.T) {
	dm := &DeploymentManager{}
	err := dm.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during deploymentmanager.Initialize() with blank credentials: %s", err)
	}
	setCallMockDefaults(dm)
	triggerResourceNotFound = true
	value, err := dm.GetResourcePropertyValue(testDeploymentName, testProjectID, "missing-resource", "one")
	if err != nil {
		t.Errorf("Got unexpected error during deploymentmanager.GetResourcePropertyValue() for resource not found: %s", err)
	}
	if value != "" {
		t.Errorf("Got unexpected result/value from deploymentmanager.GetResourcePropertyValue() for resource not found, expecting empty, but got: %s",
			value)
	}
}

func TestGetDeploymentWithoutParsedManifest(t *testing.T) {
	dm := &DeploymentManager{}
	err := dm.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during deploymentmanager.Initialize() with blank credentials: %s", err)
	}
	setCallMockDefaults(dm)
	deployment, err := dm.GetDeployment(testDeploymentName, testProjectID, false)
	if err != nil {
		t.Errorf("Got unexpected error during deploymentmanager.GetDeployment() without parsing manifest: %s", err)
	}
	if len(deployment.Outputs) > 0 {
		t.Errorf("expected outputs to be length 0 during deploymentmanager.GetDeployment() without parsing manifest, but got: %d", len(deployment.Outputs))
	}
}

func TestGetDeploymentWithParsedManifest(t *testing.T) {
	dm := &DeploymentManager{}
	err := dm.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during deploymentmanager.Initialize() with blank credentials: %s", err)
	}
	setCallMockDefaults(dm)
	deployment, err := dm.GetDeployment(testDeploymentName, testProjectID, true)
	if err != nil {
		t.Errorf("Got unexpected error during deploymentmanager.GetDeployment() with parsing manifest: %s", err)
	}
	if len(deployment.Outputs) <= 0 {
		t.Errorf("expected outputs length to be greater than zero during deploymentmanager.GetDeployment() with parsing manifest, but got: %d", len(deployment.Outputs))
	}
}

func TestEnsureDeploymentCreate(t *testing.T) {
	dm := &DeploymentManager{}
	err := dm.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during deploymentmanager.Initialize() with blank credentials: %s", err)
	}
	setCallMockDefaults(dm)
	triggerDeploymentNotFound = true
	outputs, err := dm.EnsureDeployment(testDeploymentName, "desc", testProjectID, defaultDeployment)
	if err != nil {
		t.Errorf("Got unexpected error during deploymentmanager.EnsureDeployment() for creating: %s", err)
	}
	if len(outputs) != 1 {
		t.Errorf("expected outputs to be length 1 during deploymentmanager.EnsureDeployment() for creating, but got: %d", len(outputs))
		return
	}
	if outputs[0].FinalValue != "one-id" {
		t.Errorf("expected first output finalValue to be \"one-id\" during deploymentmanager.EnsureDeployment() for creating, but got: %s", outputs[0].FinalValue)
	}
}

func TestEnsureDeploymentCreateWithRetry(t *testing.T) {
	dm := &DeploymentManager{}
	err := dm.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during deploymentmanager.Initialize() with blank credentials: %s", err)
	}
	setCallMockDefaults(dm)
	triggerGetDeploymentRetry = true
	outputs, err := dm.EnsureDeployment(testDeploymentName, "desc", testProjectID, defaultDeployment)
	if err != nil {
		t.Errorf("Got unexpected error during deploymentmanager.EnsureDeployment() for creating w/ retry: %s", err)
	}
	if len(outputs) != 1 {
		t.Errorf("expected outputs to be length 1 during deploymentmanager.EnsureDeployment() for creating w/ retry, but got: %d", len(outputs))
		return
	}
	if outputs[0].FinalValue != "one-id" {
		t.Errorf("expected first output finalValue to be \"one-id\" during deploymentmanager.EnsureDeployment() for creating w/ retry, but got: %s",
			outputs[0].FinalValue)
	}
}

func TestEnsureDeploymentCreateOperationError(t *testing.T) {
	dm := &DeploymentManager{}
	err := dm.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during deploymentmanager.Initialize() with blank credentials: %s", err)
	}
	setCallMockDefaults(dm)
	triggerDeploymentNotFound = true
	triggerCreateOperationError = true
	_, err = dm.EnsureDeployment(testDeploymentName, "desc", testProjectID, defaultDeployment)
	if err == nil {
		t.Error("Expected error during deploymentmanager.EnsureDeployment() when triggering create operation error, but got no error")
	}
}

func TestEnsureDeploymentCreateImportFileNotFound(t *testing.T) {
	dm := &DeploymentManager{}
	err := dm.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during deploymentmanager.Initialize() with blank credentials: %s", err)
	}
	setCallMockDefaults(dm)
	triggerDeploymentNotFound = true
	deployment := defaultDeployment
	deployment.Imports[0].Path = "/does/not/exist"
	_, err = dm.EnsureDeployment(testDeploymentName, "desc", testProjectID, deployment)
	if err == nil {
		t.Error("Expected error during deploymentmanager.EnsureDeployment() with invalid import file path, but got no error")
	}
	deployment.Imports[0].Path = "./.test-fixtures/import.yaml"
}

func TestEnsureDeploymentCreateWithProgressLoop(t *testing.T) {
	dm := &DeploymentManager{}
	err := dm.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during deploymentmanager.Initialize() with blank credentials: %s", err)
	}
	setCallMockDefaults(dm)
	triggerDeploymentNotFound = true
	triggerPostOperationProgressLoop = true
	outputs, err := dm.EnsureDeployment(testDeploymentName, "desc", testProjectID, defaultDeployment)
	if err != nil {
		t.Errorf("Got unexpected error during deploymentmanager.EnsureDeployment() for creating w/ progress loop: %s", err)
	}
	if len(outputs) != 1 {
		t.Errorf("expected outputs to be length 1 during deploymentmanager.EnsureDeployment() for creating w/ progress loop, but got: %d", len(outputs))
		return
	}
	if outputs[0].FinalValue != "one-id" {
		t.Errorf("expected first output finalValue to be \"one-id\" during deploymentmanager.EnsureDeployment() for creating w/ progress loop, but got: %s",
			outputs[0].FinalValue)
	}
}

func TestDeleteDeploymentDoesntExist(t *testing.T) {
	dm := &DeploymentManager{}
	err := dm.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during deploymentmanager.Initialize() with blank credentials: %s", err)
	}
	setCallMockDefaults(dm)
	triggerDeploymentNotFound = true
	err = dm.DeleteDeployment(testDeploymentName, testProjectID, false)
	if err != nil {
		t.Errorf("Got unexpected error during deploymentmanager.DeleteDeployment() with triggered deployment not found: %s", err)
	}
}

func TestDeleteDeploymentExists(t *testing.T) {
	dm := &DeploymentManager{}
	err := dm.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during deploymentmanager.Initialize() with blank credentials: %s", err)
	}
	setCallMockDefaults(dm)
	err = dm.DeleteDeployment(testDeploymentName, testProjectID, false)
	if err != nil {
		t.Errorf("Got unexpected error during deploymentmanager.DeleteDeployment() with existing deployment: %s", err)
	}
}

func TestDeleteDeploymentExistsAbandon(t *testing.T) {
	dm := &DeploymentManager{}
	err := dm.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during deploymentmanager.Initialize() with blank credentials: %s", err)
	}
	setCallMockDefaults(dm)
	err = dm.DeleteDeployment(testDeploymentName, testProjectID, true)
	if err != nil {
		t.Errorf("Got unexpected error during deploymentmanager.DeleteDeployment() with existing deployment, instructing to abandon: %s", err)
	}
}
