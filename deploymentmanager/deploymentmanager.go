// Package deploymentmanager is the library for google cloud deployment manager operations
package deploymentmanager

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/rockholla/go-google-lib/deploymentmanager/calls"
	"github.com/rockholla/go-lib/logger"
	v2beta "google.golang.org/api/deploymentmanager/v2beta"
	"google.golang.org/api/option"
	yaml "gopkg.in/yaml.v2"
)

// Interface represents functionality for DeploymentManager
type Interface interface {
	Initialize(credentials string, log logger.Interface) error
	GetResourcePropertyValue(deploymentName string, inProject string, resourceName string, propertyName string) (string, error)
	GetDeployment(deploymentName string, inProject string, parseManifest bool) (*Deployment, error)
	EnsureDeployment(deploymentName string, description string, inProject string, deployment *Deployment) ([]*Output, error)
	DeleteDeployment(deploymentName string, inProject string, abandon bool) error
}

// DeploymentManager is a wrapper around the google-provided sdks/apis for google.golang.org/api/deploymentmanager/*
type DeploymentManager struct {
	log                 logger.Interface
	V2Beta              *v2beta.Service
	Calls               *Calls
	RetryWaitSeconds    int64
	ProgressWaitSeconds int64
}

// Calls are interfaces for making the actual calls to various underlying apis
type Calls struct {
	ResourcesGet      calls.ResourcesGetCallInterface
	DeploymentsGet    calls.DeploymentsGetCallInterface
	DeploymentsInsert calls.DeploymentsInsertCallInterface
	DeploymentsUpdate calls.DeploymentsUpdateCallInterface
	DeploymentsDelete calls.DeploymentsDeleteCallInterface
	OperationsGet     calls.OperationsGetCallInterface
	ManifestsGet      calls.ManifestsGetCallInterface
}

// Deployment is a generic object for a single deployment manager deployment
// see https://cloud.google.com/deployment-manager/docs/configuration/syntax-reference
type Deployment struct {
	Imports   []*Import          `yaml:"-"`
	Resources []*Resource        `yaml:"resources,omitempty"`
	Outputs   []*Output          `yaml:"outputs,omitempty"`
	Source    *v2beta.Deployment `yaml:"-"`
}

// Import is a deployment import file
type Import struct {
	Name string
	Path string
}

// Resource is a single resource in a single deployment
type Resource struct {
	Name       string                 `yaml:"name,omitempty"`
	Type       string                 `yaml:"type,omitempty"`
	Action     string                 `yaml:"action,omitempty"`
	Properties map[string]interface{} `yaml:"properties,omitempty"`
	Metadata   *ResourceMetadata      `yaml:"metadata,omitempty"`
}

// ResourceMetadata is a Deployment Manager single resource metadata
type ResourceMetadata struct {
	DependsOn     []string `yaml:"dependsOn,omitempty"`
	RuntimePolicy []string `yaml:"runtimePolicy,omitempty"`
}

// Output is a single deployment output
type Output struct {
	Name       string `yaml:"name,omitempty"`
	Value      string `yaml:"value,omitempty"`
	FinalValue string `yaml:"finalValue,omitempty"`
}

// Initialize sets up necessary google-provided sdks and other local data
func (dm *DeploymentManager) Initialize(credentials string, log logger.Interface) error {
	var err error
	ctx := context.Background()
	dm.log = log
	dm.RetryWaitSeconds = 60
	dm.ProgressWaitSeconds = 10
	dm.Calls = &Calls{
		ResourcesGet:      &calls.ResourcesGetCall{},
		DeploymentsGet:    &calls.DeploymentsGetCall{},
		DeploymentsInsert: &calls.DeploymentsInsertCall{},
		DeploymentsUpdate: &calls.DeploymentsUpdateCall{},
		DeploymentsDelete: &calls.DeploymentsDeleteCall{},
		OperationsGet:     &calls.OperationsGetCall{},
		ManifestsGet:      &calls.ManifestsGetCall{},
	}
	if credentials != "" {
		if dm.V2Beta, err = v2beta.NewService(ctx, option.WithCredentialsJSON([]byte(credentials))); err != nil {
			return err
		}
	} else {
		if dm.V2Beta, err = v2beta.NewService(ctx); err != nil {
			return err
		}
	}
	return nil
}

// GetResourcePropertyValue will get an existing resource property, if the resource exists, otherwise will return blank
func (dm *DeploymentManager) GetResourcePropertyValue(deploymentName string, inProject string, resourceName string, propertyName string) (string, error) {
	value := ""
	ctx := context.Background()
	resourcesService := v2beta.NewResourcesService(dm.V2Beta)
	resourceGetCall := resourcesService.Get(inProject, deploymentName, resourceName).Context(ctx)
	resource, err := dm.Calls.ResourcesGet.Do(resourceGetCall)
	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "notfound") {
			return value, nil
		}
		return value, err
	}
	var properties map[string]interface{}
	err = yaml.Unmarshal([]byte(resource.Properties), &properties)
	if err != nil {
		return value, err
	}
	if value, exists := properties[propertyName]; exists {
		return value.(string), nil
	}
	return value, nil
}

func (dm *DeploymentManager) isRetryError(err error) bool {
	if err == nil {
		return false
	}
	if strings.Contains(strings.ToLower(err.Error()), "wait a few minutes") {
		dm.log.Info("instructed to wait, retrying in %d seconds (this likely means that the deployment manager api is not enabled, yet)...\n", dm.RetryWaitSeconds)
		time.Sleep(time.Duration(dm.RetryWaitSeconds) * time.Second)
		return true
	}
	if strings.Contains(strings.ToLower(err.Error()), "conflicting operation") {
		dm.log.Info("conflicting operation ongoing, retrying in %d seconds...\n", dm.RetryWaitSeconds)
		time.Sleep(time.Duration(dm.RetryWaitSeconds) * time.Second)
		return true
	}
	return false
}

func (dm *DeploymentManager) processOperation(operation *v2beta.Operation) error {
	if operation.Error != nil {
		errorMessage := ""
		for _, e := range operation.Error.Errors {
			errorMessage = fmt.Sprintf("%s %s", errorMessage, e.Message)
		}
		return errors.New(strings.TrimSpace(errorMessage))
	}
	return nil
}

// GetDeployment will get an existing deployment if it exists
func (dm *DeploymentManager) GetDeployment(deploymentName string, inProject string, parseManifest bool) (*Deployment, error) {
	ctx := context.Background()
	deployment := &Deployment{}
	deploymentManagerService := v2beta.NewDeploymentsService(dm.V2Beta)
	deploymentGetCall := deploymentManagerService.Get(inProject, deploymentName).Context(ctx)
	existingDeployment, err := dm.Calls.DeploymentsGet.Do(deploymentGetCall)
	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "notfound") {
			return nil, nil
		} else if dm.isRetryError(err) {
			return dm.GetDeployment(deploymentName, inProject, parseManifest)
		} else {
			return deployment, err
		}
	}
	deployment.Source = existingDeployment
	if parseManifest {
		manifestsService := v2beta.NewManifestsService(dm.V2Beta)
		manifestGetCall := manifestsService.Get(inProject, deployment.Source.Name,
			deployment.Source.Manifest[strings.LastIndex(deployment.Source.Manifest, "/")+1:]).Context(ctx)
		manifest, err := dm.Calls.ManifestsGet.Do(manifestGetCall)
		if err != nil {
			return deployment, fmt.Errorf("error getting deployment manifest: %s", err.Error())
		}
		var layout map[string]interface{}
		if manifest.Layout != "" {
			if err = yaml.Unmarshal([]byte(manifest.Layout), &layout); err != nil {
				return deployment, fmt.Errorf("error parsing deployment manifest layout: %s", err.Error())
			}
		}
		var deploymentConfig *Deployment
		if manifest.Config != nil {
			if err = yaml.Unmarshal([]byte(manifest.Config.Content), &deploymentConfig); err != nil {
				return deployment, fmt.Errorf("error parsing deployment manifest config: %s", err.Error())
			}
			deployment.Resources = deploymentConfig.Resources
		}
		if _, exists := layout["outputs"]; exists {
			for _, output := range layout["outputs"].([]interface{}) {
				outputMap := output.(map[interface{}]interface{})
				deployment.Outputs = append(deployment.Outputs, &Output{
					Name:       outputMap["name"].(string),
					Value:      outputMap["value"].(string),
					FinalValue: outputMap["finalValue"].(string),
				})
			}
		}
	}

	return deployment, nil
}

// EnsureDeployment will make sure that a deployment exists
func (dm *DeploymentManager) EnsureDeployment(deploymentName string, description string, inProject string, deployment *Deployment) ([]*Output, error) {
	ctx := context.Background()
	var operation *v2beta.Operation
	var outputs []*Output
	targetConfiguration := &v2beta.TargetConfiguration{}
	dm.log.InfoPart("Ensuring deployment \"%s\" in project \"%s\"...", deploymentName, inProject)
	for _, deploymentImport := range deployment.Imports {
		importFileContent, err := ioutil.ReadFile(deploymentImport.Path)
		if err != nil {
			return outputs, fmt.Errorf("error reading import file %s: %s", deploymentImport.Path, err.Error())
		}
		targetConfiguration.Imports = append(targetConfiguration.Imports, &v2beta.ImportFile{
			Name:    deploymentImport.Name,
			Content: string(importFileContent),
		})
	}
	deploymentData, err := yaml.Marshal(deployment)
	if err != nil {
		return outputs, fmt.Errorf("error parsing deployment data: %s", err.Error())
	}
	targetConfiguration.Config = &v2beta.ConfigFile{
		Content: string(deploymentData),
	}

	deploymentManagerDeployment := &v2beta.Deployment{
		Name:        deploymentName,
		Description: description,
		Target:      targetConfiguration,
	}
	deploymentManagerService := v2beta.NewDeploymentsService(dm.V2Beta)
	existingDeployment, err := dm.GetDeployment(deploymentName, inProject, false)
	if err != nil {
		return outputs, fmt.Errorf("error trying to determine if deployment exists already: %s", err.Error())
	}
	if existingDeployment == nil {
		dm.log.SpinnerStart("creating")
		deploymentInsertCall := deploymentManagerService.Insert(inProject, deploymentManagerDeployment).Context(ctx)
		operation, err = dm.Calls.DeploymentsInsert.Do(deploymentInsertCall)
	} else {
		deploymentManagerDeployment.Fingerprint = existingDeployment.Source.Fingerprint
		dm.log.SpinnerStart("updating")
		deploymentUpdateCall := deploymentManagerService.Update(inProject, deploymentName, deploymentManagerDeployment).Context(ctx)
		operation, err = dm.Calls.DeploymentsUpdate.Do(deploymentUpdateCall)
	}
	if dm.isRetryError(err) {
		dm.log.SpinnerStop()
		return dm.EnsureDeployment(deploymentName, description, inProject, deployment)
	}
	if operationErr := dm.trackOperation(operation, inProject); err != nil || operationErr != nil {
		if err == nil {
			err = operationErr
		}
		dm.log.SpinnerStop()
		dm.log.InfoPart("error\n")
		return outputs, fmt.Errorf("error creating or updating the deployment: %s", err.Error())
	}
	dm.log.SpinnerStop()
	dm.log.InfoPart("done\n")
	existingDeployment, err = dm.GetDeployment(deploymentName, inProject, true)
	if err != nil {
		return outputs, fmt.Errorf("error getting updated deployment after operation: %s", err.Error())
	}
	outputs = existingDeployment.Outputs
	return outputs, nil
}

// DeleteDeployment will fully delete a deployment
func (dm *DeploymentManager) DeleteDeployment(deploymentName string, inProject string, abandon bool) error {
	var err error
	ctx := context.Background()
	dm.log.InfoPart("Deleting deployment \"%s\" in project \"%s\"...", deploymentName, inProject)
	existingDeployment, err := dm.GetDeployment(deploymentName, inProject, false)
	if err != nil {
		dm.log.InfoPart("\n")
		return fmt.Errorf("error trying to determine if deployment exists already: %s", err.Error())
	}
	if existingDeployment == nil {
		dm.log.InfoPart("doesn't exist\n")
		return nil
	}
	dm.log.SpinnerStart("deleting")
	deploymentManagerService := v2beta.NewDeploymentsService(dm.V2Beta)
	deploymentDeleteCall := deploymentManagerService.Delete(inProject, deploymentName).Context(ctx)
	if abandon {
		deploymentDeleteCall = deploymentDeleteCall.DeletePolicy("ABANDON")
	}
	operation, err := dm.Calls.DeploymentsDelete.Do(deploymentDeleteCall)
	if dm.isRetryError(err) {
		dm.log.SpinnerStop()
		return dm.DeleteDeployment(deploymentName, inProject, abandon)
	}

	if operationErr := dm.trackOperation(operation, inProject); err != nil || operationErr != nil {
		if err == nil {
			err = operationErr
		}
		dm.log.SpinnerStop()
		dm.log.InfoPart("error\n")
		return fmt.Errorf("error deleting deployment: %s", err.Error())
	}
	dm.log.SpinnerStop()
	dm.log.InfoPart("done\n")
	return nil
}

func (dm *DeploymentManager) trackOperation(operation *v2beta.Operation, inProject string) error {
	var err error
	ctx := context.Background()
	if operation == nil {
		return nil
	}
	if err = dm.processOperation(operation); err != nil {
		return err
	}
	operationsService := v2beta.NewOperationsService(dm.V2Beta)
	operationGetCall := operationsService.Get(inProject, operation.Name).Context(ctx)
	operation, err = dm.Calls.OperationsGet.Do(operationGetCall)
	for operation != nil && operation.Progress < 100 && operation.Error == nil {
		operation, err = dm.Calls.OperationsGet.Do(operationGetCall)
		time.Sleep(time.Duration(dm.ProgressWaitSeconds) * time.Second)
	}
	if operation == nil {
		return nil
	}
	if err = dm.processOperation(operation); err != nil {
		return err
	}
	return nil
}

// GetOutputValue is a helper for finding the value for a particular output name
func GetOutputValue(outputs []*Output, name string) string {
	for _, output := range outputs {
		if output.Name == name {
			return output.FinalValue
		}
	}
	return ""
}
