package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	api "cloud.google.com/go/storage"
	loggermock "github.com/rockholla/go-lib/mocks/custom-mocks/logger"
	"google.golang.org/api/option"
)

var (
	testCredentials = `{
  "client_id": "xxxxxxx.apps.googleusercontent.com",
  "client_secret": "xxxxxxxxxxxxxxx",
  "refresh_token": "xxxxxxxxx",
  "type": "authorized_user"
}`
)

type mockTransport struct {
	gotReq  *http.Request
	gotBody []byte
	results []transportResult
}

type transportResult struct {
	res *http.Response
	err error
}

func (t *mockTransport) addResult(res *http.Response, err error) {
	t.results = append(t.results, transportResult{res, err})
}

func (t *mockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	t.gotReq = req
	t.gotBody = nil
	if req.Body != nil {
		bytes, err := ioutil.ReadAll(req.Body)
		if err != nil {
			return nil, err
		}
		t.gotBody = bytes
	}
	if len(t.results) == 0 {
		return nil, fmt.Errorf("error handling request")
	}
	result := t.results[0]
	t.results = t.results[1:]
	return result.res, result.err
}

func (t *mockTransport) gotJSONBody() map[string]interface{} {
	m := map[string]interface{}{}
	if err := json.Unmarshal(t.gotBody, &m); err != nil {
		panic(err)
	}
	return m
}

func mockClient(t *testing.T, m *mockTransport) *api.Client {
	client, err := api.NewClient(context.Background(), option.WithHTTPClient(&http.Client{Transport: m}))
	if err != nil {
		t.Fatal(err)
	}
	return client
}

func bodyReader(s string) io.ReadCloser {
	return ioutil.NopCloser(strings.NewReader(s))
}

func TestInitialize(t *testing.T) {
	s := &Storage{}
	err := s.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during storage.Initialize() with blank credentials: %s", err)
	}
	err = s.Initialize(testCredentials, loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during storage.Initialize() with explicit credentials: %s", err)
	}
}

func TestEnsureBucketDoesntExist(t *testing.T) {
	s := &Storage{}
	err := s.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during storage.Initialize() with blank credentials: %s", err)
	}
	mt := &mockTransport{}
	// TODO: this won't test the create case block for now b/c we can't handle it with our mock transport as is, come back for it
	mt.addResult(&http.Response{StatusCode: 200, Body: bodyReader("{}")}, nil)
	s.Client = mockClient(t, mt)
	err = s.EnsureBucket("bucket", "project-id")
	if err != nil {
		t.Errorf("Got unexpected error for storage.EnsureBucket(): %s", err)
	}
}

func TestEnsureBucketExists(t *testing.T) {
	s := &Storage{}
	err := s.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during storage.Initialize() with blank credentials: %s", err)
	}
	mt := &mockTransport{}
	mt.addResult(&http.Response{StatusCode: 200, Body: bodyReader("{}")}, nil)
	s.Client = mockClient(t, mt)
	err = s.EnsureBucket("bucket", "project-id")
	if err != nil {
		t.Errorf("Got unexpected error for storage.EnsureBucket(): %s", err)
	}
}

func TestEnsureObject(t *testing.T) {
	s := &Storage{}
	err := s.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during storage.Initialize() with blank credentials: %s", err)
	}
	mt := &mockTransport{}
	mt.addResult(&http.Response{StatusCode: 200, Body: bodyReader("{}")}, nil)
	s.Client = mockClient(t, mt)
	err = s.EnsureObject("bucket", "path/to/object", &Object{
		Data:        []byte("file-data"),
		ContentType: "text/plain",
	})
	if err != nil {
		t.Errorf("Got unexpected error for storage.EnsureObject(): %s", err)
	}
}

func TestEnsureObjectError(t *testing.T) {
	s := &Storage{}
	err := s.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during storage.Initialize() with blank credentials: %s", err)
	}
	mt := &mockTransport{}
	s.Client = mockClient(t, mt)
	err = s.EnsureObject("bucket", "path/to/object", &Object{
		Data:        []byte("file-data"),
		ContentType: "text/plain",
	})
	if err == nil {
		t.Error("Expected error from storage.EnsureObject(), but got no error")
	}
}

func TestGetObject(t *testing.T) {
	s := &Storage{}
	err := s.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during storage.Initialize() with blank credentials: %s", err)
	}
	mt := &mockTransport{}
	mt.addResult(&http.Response{StatusCode: 200, Body: bodyReader("{}")}, nil)
	s.Client = mockClient(t, mt)
	_, err = s.GetObject("bucket", "path/to/object")
	if err != nil {
		t.Errorf("Got unexpected error for storage.GetObject(): %s", err)
	}
}

func TestGetObjectError(t *testing.T) {
	s := &Storage{}
	err := s.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during storage.Initialize() with blank credentials: %s", err)
	}
	mt := &mockTransport{}
	s.Client = mockClient(t, mt)
	_, err = s.GetObject("bucket", "path/to/object")
	if err == nil {
		t.Error("Expected error from storage.GetObject(), but got no error")
	}
}

func TestGetServiceAccount(t *testing.T) {
	s := &Storage{}
	err := s.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during storage.Initialize() with blank credentials: %s", err)
	}
	mt := &mockTransport{}
	mt.addResult(&http.Response{StatusCode: 200, Body: bodyReader("{\"email_address\": \"test-service-account\", \"kind\": \"storage#serviceAccount\"}")}, nil)
	s.Client = mockClient(t, mt)
	_, err = s.GetServiceAccount("test-project")
	if err != nil {
		t.Errorf("Got unexpected error for storage.GetServiceAccount(): %s", err)
	}
}

func TestEnsureBucketRoles(t *testing.T) {
	s := &Storage{}
	err := s.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during storage.Initialize() with blank credentials: %s", err)
	}
	mt := &mockTransport{}
	iamJSON := `{
  "kind": "storage#policy",
  "resourceId": "projects/test-project/buckets/test-bucket",
  "bindings": []
}`
	mt.addResult(&http.Response{StatusCode: 200, Body: bodyReader(iamJSON)}, nil)
	mt.addResult(&http.Response{StatusCode: 200, Body: bodyReader("{}")}, nil)
	s.Client = mockClient(t, mt)
	err = s.EnsureBucketRoles("test-bucket", "serviceAccount:test-sa@go-google-lib.tests", []string{"roles/storage.objectViewer"})
	if err != nil {
		t.Errorf("Got unexpected error for storage.EnsureBucketRoles(): %s", err)
	}
}
