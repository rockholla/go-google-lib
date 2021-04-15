// Package storage is the library for google cloud storage operations
package storage

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"

	"cloud.google.com/go/iam"
	api "cloud.google.com/go/storage"
	"github.com/rockholla/go-google-lib/logger"
	"google.golang.org/api/option"
)

// Interface represents functionality for storage
type Interface interface {
	Initialize(credentials string, log logger.Interface) error
	EnsureObject(bucket string, path string, object *Object) error
	GetObject(bucket string, path string) ([]byte, error)
	GetServiceAccount(projectID string) (string, error)
	EnsureBucketRoles(bucket string, member string, roles []string) error
}

// Storage wraps google-provided apis for interacting with cloud.google.com/go/storage/*
type Storage struct {
	log    logger.Interface
	Client *api.Client
}

// Object is a storage object
type Object struct {
	ContentType string
	Data        []byte
}

// Initialize sets up necessary google-provided sdks and other local data
func (storage *Storage) Initialize(credentials string, log logger.Interface) error {
	var err error
	ctx := context.Background()
	storage.log = log
	if credentials != "" {
		if storage.Client, err = api.NewClient(ctx, option.WithCredentialsJSON([]byte(credentials))); err != nil {
			return err
		}
	} else {
		if storage.Client, err = api.NewClient(ctx); err != nil {
			return err
		}
	}
	return nil
}

// EnsureObject will make sure that an object exists and is updated with provided data in a bucket
func (storage *Storage) EnsureObject(bucket string, path string, object *Object) error {
	ctx := context.Background()
	objectWriter := storage.Client.Bucket(bucket).Object(path).NewWriter(ctx)
	objectWriter.ContentType = object.ContentType
	errs := ""
	if _, err := objectWriter.Write(object.Data); err != nil {
		errs = err.Error()
	}
	if err := objectWriter.Close(); err != nil {
		errs = fmt.Sprintf("%s %s", errs, err)
	}
	if errs != "" {
		return errors.New(errs)
	}
	return nil
}

// GetObject will get a storage bucket object content bytes
func (storage *Storage) GetObject(bucket string, path string) ([]byte, error) {
	var content []byte
	ctx := context.Background()
	objectReader, err := storage.Client.Bucket(bucket).Object(path).NewReader(ctx)
	if err != nil {
		return content, err
	}
	content, err = ioutil.ReadAll(objectReader)
	objectReader.Close()
	if err != nil {
		return content, err
	}
	return content, nil
}

// GetServiceAccount will return the storage service account for a project
func (storage *Storage) GetServiceAccount(projectID string) (string, error) {
	ctx := context.Background()
	return storage.Client.ServiceAccount(ctx, projectID)
}

// EnsureBucketRoles makes sure that a particular member has the supplied roles on the bucket
func (storage *Storage) EnsureBucketRoles(bucket string, member string, roles []string) error {
	ctx := context.Background()
	storage.log.Info("Ensuring member %s has roles on gs://%s:", member, bucket)
	bucketIAMHandle := storage.Client.Bucket(bucket).IAM()
	policy, err := bucketIAMHandle.Policy(ctx)
	if err != nil {
		return err
	}
	for _, role := range roles {
		storage.log.ListItem(role)
		roleName := iam.RoleName(role)
		if !policy.HasRole(member, roleName) {
			policy.Add(member, roleName)
		}
	}
	return bucketIAMHandle.SetPolicy(ctx, policy)
}
