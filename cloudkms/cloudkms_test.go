package cloudkms

import (
	"context"
	"testing"

	gax "github.com/googleapis/gax-go/v2"
	loggermock "github.com/rockholla/go-lib/mocks/custom-mocks/logger"
	v1objects "google.golang.org/genproto/googleapis/cloud/kms/v1"
)

type ClientMock struct{}

func (c *ClientMock) Encrypt(ctx context.Context, req *v1objects.EncryptRequest, opts ...gax.CallOption) (*v1objects.EncryptResponse, error) {
	return &v1objects.EncryptResponse{}, nil
}
func (c *ClientMock) Decrypt(ctx context.Context, req *v1objects.DecryptRequest, opts ...gax.CallOption) (*v1objects.DecryptResponse, error) {
	return &v1objects.DecryptResponse{}, nil
}

func TestEncryptDecrypt(t *testing.T) {
	kms := &CloudKMS{}
	err := kms.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during cloudkms.Initialize(): %s", err)
	}
	kms.V1 = &ClientMock{}
	cryptoKey := &CryptoKey{
		ProjectID: "test-project",
		Location:  "us-central1",
		KeyRing:   "test-keyring",
		Name:      "test-encryption-key",
	}
	encrypted, err := kms.Encrypt(cryptoKey, "some data")
	if err != nil {
		t.Errorf("Got unexpected error from kms.Encrypt(): %s", err)
	}
	_, err = kms.Decrypt(cryptoKey, encrypted)
	if err != nil {
		t.Errorf("Got unexpected error from kms.Decrypt(): %s", err)
	}
}
