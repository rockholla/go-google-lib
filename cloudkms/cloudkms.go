// Package cloudkms is the library for google cloud kms operations
package cloudkms

import (
	"context"
	"encoding/base64"
	"fmt"

	v1 "cloud.google.com/go/kms/apiv1"
	"github.com/rockholla/go-google-lib/logger"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	v1objects "google.golang.org/genproto/googleapis/cloud/kms/v1"
)

// Interface represents functionality for DeploymentManager
type Interface interface {
	Initialize(credentials string, log logger.Interface) error
	Encrypt(key *CryptoKey, data string) (string, error)
	Decrypt(key *CryptoKey, data string) (string, error)
}

// ClientInterface represents the underlying kms api client
type ClientInterface interface {
	Encrypt(ctx context.Context, req *v1objects.EncryptRequest, opts ...gax.CallOption) (*v1objects.EncryptResponse, error)
	Decrypt(ctx context.Context, req *v1objects.DecryptRequest, opts ...gax.CallOption) (*v1objects.DecryptResponse, error)
}

// CloudKMS is a wrapper around the google-provided sdks/apis for google.golang.org/api/cloudkms/*
type CloudKMS struct {
	log logger.Interface
	V1  ClientInterface
}

// CryptoKey represents an encryption key within a project, location, and key ring
type CryptoKey struct {
	ProjectID string
	Location  string
	KeyRing   string
	Name      string
}

// Initialize sets up necessary google-provided sdks and other local data
func (kms *CloudKMS) Initialize(credentials string, log logger.Interface) error {
	var err error
	ctx := context.Background()
	kms.log = log
	if credentials != "" {
		if kms.V1, err = v1.NewKeyManagementClient(ctx, option.WithCredentialsJSON([]byte(credentials))); err != nil {
			return err
		}
	} else {
		if kms.V1, err = v1.NewKeyManagementClient(ctx); err != nil {
			return err
		}
	}
	return nil
}

// Encrypt will receive data in and encrypt with a designated KMS crypto key
func (kms *CloudKMS) Encrypt(key *CryptoKey, data string) (string, error) {
	ctx := context.Background()
	request := &v1objects.EncryptRequest{
		Name:      fmt.Sprintf("projects/%s/locations/%s/keyRings/%s/cryptoKeys/%s", key.ProjectID, key.Location, key.KeyRing, key.Name),
		Plaintext: []byte(data),
	}
	response, err := kms.V1.Encrypt(ctx, request)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(response.Ciphertext), nil
}

// Decrypt will receive data in and decrypt with a designated KMS crypto key
func (kms *CloudKMS) Decrypt(key *CryptoKey, data string) (string, error) {
	ctx := context.Background()
	decoded, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}
	request := &v1objects.DecryptRequest{
		Name:       fmt.Sprintf("projects/%s/locations/%s/keyRings/%s/cryptoKeys/%s", key.ProjectID, key.Location, key.KeyRing, key.Name),
		Ciphertext: decoded,
	}
	response, err := kms.V1.Decrypt(ctx, request)
	if err != nil {
		return "", err
	}
	return string(response.Plaintext), nil
}
