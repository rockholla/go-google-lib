// Package oauth deals in Google-specific oauth APIs
package oauth

import (
	"context"

	"github.com/rockholla/go-lib/logger"
	googleoauth "golang.org/x/oauth2/google"
)

// will be used if no list of scopes is provided explicitly
var defaultScopes = []string{
	"https://www.googleapis.com/auth/compute",
	"https://www.googleapis.com/auth/cloud-platform",
	"https://www.googleapis.com/auth/cloud-identity",
	"https://www.googleapis.com/auth/ndev.clouddns.readwrite",
	"https://www.googleapis.com/auth/devstorage.full_control",
	"https://www.googleapis.com/auth/userinfo.email",
}

// Interface represents functionality for OAuth
type Interface interface {
	Initialize(credentials string, log logger.Interface, scopes []string) error
	GetAccessToken() (string, error)
}

// OAuth wraps google-provided apis for interacting with pkg.go.dev/golang.org/x/oauth2/google/*
type OAuth struct {
	log         logger.Interface
	Credentials *googleoauth.Credentials
}

// Initialize sets up necessary google-provided sdks and other local data
func (o *OAuth) Initialize(credentials string, log logger.Interface, scopes []string) error {
	var err error
	ctx := context.Background()
	o.log = log
	if len(scopes) == 0 {
		scopes = defaultScopes
	}
	if credentials != "" {
		o.Credentials, err = googleoauth.CredentialsFromJSON(ctx, []byte(credentials), scopes...)
		if err != nil {
			return err
		}
	} else {
		defaultTokenSource, err := googleoauth.DefaultTokenSource(ctx, scopes...)
		if err != nil {
			return err
		}
		o.Credentials = &googleoauth.Credentials{
			TokenSource: defaultTokenSource,
		}
	}
	return nil
}

// GetAccessToken will return the access token as a string
// TODO: unit test this
func (o *OAuth) GetAccessToken() (string, error) {
	token, err := o.Credentials.TokenSource.Token()
	if err != nil {
		return "", err
	}
	return token.AccessToken, nil
}
