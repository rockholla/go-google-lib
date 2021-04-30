package cloudidentity

import (
	"errors"
	"testing"

	loggermock "github.com/rockholla/go-lib/mocks/custom-mocks/logger"
	v1beta1 "google.golang.org/api/cloudidentity/v1beta1"
	googleapi "google.golang.org/api/googleapi"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var triggerGroupNotFound = false
var triggerGroupNotFoundRaw = false
var groupNotFoundErrorRaw = `Error: googleapi: got HTTP response code 404 with body: <!DOCTYPE html>
<html lang=en>
  <meta charset=utf-8>
  <meta name=viewport content="initial-scale=1, minimum-scale=1, width=device-width">
  <title>Error 404 (Not Found)!!1</title>
  <style>
    *{margin:0;padding:0}html,code{font:15px/22px arial,sans-serif}html{background:#fff;color:#222;padding:15px}body{margin:7% auto 0;max-width:390px;min-height:180px;padding:30px 0 15px}* > body{background:url(//www.google.com/images/errors/robot.png) 100% 5px no-repeat;padding-right:205px}p{margin:11px 0 22px;overflow:hidden}ins{color:#777;text-decoration:none}a img{border:0}@media screen and (max-width:772px){body{background:none;margin-top:0;max-width:none;padding-right:0}}#logo{background:url(//www.google.com/images/branding/googlelogo/1x/googlelogo_color_150x54dp.png) no-repeat;margin-left:-5px}@media only screen and (min-resolution:192dpi){#logo{background:url(//www.google.com/images/branding/googlelogo/2x/googlelogo_color_150x54dp.png) no-repeat 0% 0%/100% 100%;-moz-border-image:url(//www.google.com/images/branding/googlelogo/2x/googlelogo_color_150x54dp.png) 0}}@media only screen and (-webkit-min-device-pixel-ratio:2){#logo{background:url(//www.google.com/images/branding/googlelogo/2x/googlelogo_color_150x54dp.png) no-repeat;-webkit-background-size:100% 100%}}#logo{display:inline-block;height:54px;width:150px}
  </style>
  <a href=//www.google.com/><span id=logo aria-label=Google></span></a>
  <p><b>404.</b> <ins>That’s an error.</ins>
  <p>The requested URL <code>/v1beta1/group?alt=json&amp;prettyPrint=false</code> was not found on this server.  <ins>That’s all we know.</ins>
`

type groupGetMock struct{}
type groupCreateMock struct{}

func (c *groupGetMock) Do(call *v1beta1.GroupsGetCall, opts ...googleapi.CallOption) (*v1beta1.Group, error) {
	if triggerGroupNotFound {
		triggerGroupNotFound = false
		st := status.New(codes.NotFound, "not found")
		return nil, st.Err()
	}
	if triggerGroupNotFoundRaw {
		triggerGroupNotFoundRaw = false
		return nil, errors.New(groupNotFoundErrorRaw)
	}
	return &v1beta1.Group{}, nil
}

func (c *groupCreateMock) Do(call *v1beta1.GroupsCreateCall, opts ...googleapi.CallOption) (*v1beta1.Operation, error) {
	return &v1beta1.Operation{}, nil
}

func setCallMockDefaults(ci *CloudIdentity) {
	ci.Calls = &Calls{
		GroupGet:    &groupGetMock{},
		GroupCreate: &groupCreateMock{},
	}
}

func TestInitialize(t *testing.T) {
	ci := &CloudIdentity{}
	err := ci.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during cloudidentity.Initialize() with blank credentials: %s", err)
	}
	err = ci.Initialize("impersonate@sa", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during cloudidentity.Initialize() with explicit credentials: %s", err)
	}
}

func TestEnsureGroupNotExists(t *testing.T) {
	ci := &CloudIdentity{}
	err := ci.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error for cloudidentity.Initialize() with blank credentials: %s", err)
	}
	setCallMockDefaults(ci)
	triggerGroupNotFound = true
	err = ci.EnsureGroup("name", "domain", "customer-id")
	if err != nil {
		t.Errorf("Got unexpected error during cloudidentity.TestEnsureGroupExists(): %s", err)
	}
}

func TestEnsureGroupNotExistsRawError(t *testing.T) {
	ci := &CloudIdentity{}
	err := ci.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error for cloudidentity.Initialize() with blank credentials: %s", err)
	}
	setCallMockDefaults(ci)
	triggerGroupNotFoundRaw = true
	err = ci.EnsureGroup("name", "domain", "customer-id")
	if err != nil {
		t.Errorf("Got unexpected error during cloudidentity.TestEnsureGroupExists(): %s", err)
	}
}

func TestEnsureGroupExists(t *testing.T) {
	ci := &CloudIdentity{}
	err := ci.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error for cloudidentity.Initialize() with blank credentials: %s", err)
	}
	setCallMockDefaults(ci)
	err = ci.EnsureGroup("name", "domain", "customer-id")
	if err != nil {
		t.Errorf("Got unexpected error during cloudidentity.TestEnsureGroupExists(): %s", err)
	}
}
