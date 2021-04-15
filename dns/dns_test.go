package dns

import (
	"testing"

	loggermock "github.com/rockholla/go-lib/mocks/custom-mocks/logger"
	v1 "google.golang.org/api/dns/v1"
	googleapi "google.golang.org/api/googleapi"
)

var (
	triggerExistingResourceRecordSet = false
	triggerPendingChangeStatus       = false
	testProjectID                    = "project-11111111111"
	testManagedZone                  = "go-google-lib.io"
	testName                         = "tests.go-google-lib.io."
	testResourceRecordSet            = &v1.ResourceRecordSet{
		Name: testName,
		Ttl:  3600,
		Type: "A",
		Rrdatas: []string{
			"1.2.3.4",
		},
	}
	testCredentials = `{
  "client_id": "xxxxxxx.apps.googleusercontent.com",
  "client_secret": "xxxxxxxxxxxxxxx",
  "refresh_token": "xxxxxxxxx",
  "type": "authorized_user"
}`
)

type rrsListMock struct{}
type changesCreateMock struct{}

// Do is the mock for default rrsListMock
func (r *rrsListMock) Do(call *v1.ResourceRecordSetsListCall, opts ...googleapi.CallOption) (*v1.ResourceRecordSetsListResponse, error) {
	if triggerExistingResourceRecordSet {
		triggerExistingResourceRecordSet = false
		return &v1.ResourceRecordSetsListResponse{
			Rrsets: []*v1.ResourceRecordSet{
				testResourceRecordSet,
			},
		}, nil
	}
	return &v1.ResourceRecordSetsListResponse{
		Rrsets: []*v1.ResourceRecordSet{},
	}, nil
}

// Do is the mock for default changesCreateMock
func (cc *changesCreateMock) Do(call *v1.ChangesCreateCall, opts ...googleapi.CallOption) (*v1.Change, error) {
	status := "done"
	if triggerPendingChangeStatus {
		triggerPendingChangeStatus = false
		status = "pending"
	}
	return &v1.Change{
		Status: status,
	}, nil
}

func setCallMockDefaults(d *DNS) {
	d.Calls = &Calls{
		ResourceRecordSetsList: &rrsListMock{},
		ChangesCreate:          &changesCreateMock{},
	}
}

func TestInitialize(t *testing.T) {
	d := &DNS{}
	err := d.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during dns.Initialize() with blank credentials: %s", err)
	}
	err = d.Initialize(testCredentials, loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during dns.Initialize() with explicit credentials: %s", err)
	}
}

func TestGetResourceRecordSet(t *testing.T) {
	d := &DNS{}
	err := d.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during dns.Initialize() with blank credentials: %s", err)
	}
	setCallMockDefaults(d)
	existing, err := d.GetResourceRecordSet(testProjectID, testManagedZone, testName)
	if err != nil {
		t.Errorf("Got unexpected error during dns.GetResourceRecordSet(): %s", err)
	}
	if existing != nil {
		t.Error("Got unexpected result/value from compute.GetResourceRecordSet(), expecting nil, but got an existing resource record set")
	}
}

func TestGetResourceRecordSetExisting(t *testing.T) {
	d := &DNS{}
	err := d.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during dns.Initialize() with blank credentials: %s", err)
	}
	triggerExistingResourceRecordSet = true
	setCallMockDefaults(d)
	existing, err := d.GetResourceRecordSet(testProjectID, testManagedZone, testName)
	if err != nil {
		t.Errorf("Got unexpected error during dns.GetResourceRecordSet(): %s", err)
	}
	if existing == nil {
		t.Error("Got unexpected result/value from compute.GetResourceRecordSet(), expecting a resource record set, but got nil")
	}
}

func TestSetResourceRecordSetsNew(t *testing.T) {
	d := &DNS{}
	err := d.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during dns.Initialize() with blank credentials: %s", err)
	}
	setCallMockDefaults(d)
	err = d.SetResourceRecordSets(testProjectID, testManagedZone, []*v1.ResourceRecordSet{testResourceRecordSet})
	if err != nil {
		t.Errorf("Got unexpected error during dns.GetResourceRecordSet(): %s", err)
	}
}

func TestSetResourceRecordSetsNewPending(t *testing.T) {
	d := &DNS{}
	err := d.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during dns.Initialize() with blank credentials: %s", err)
	}
	d.PendingWaitSeconds = 0
	triggerPendingChangeStatus = true
	setCallMockDefaults(d)
	err = d.SetResourceRecordSets(testProjectID, testManagedZone, []*v1.ResourceRecordSet{testResourceRecordSet})
	if err != nil {
		t.Errorf("Got unexpected error during dns.GetResourceRecordSet(): %s", err)
	}
}

func TestSetResourceRecordSetsUpdate(t *testing.T) {
	d := &DNS{}
	err := d.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during dns.Initialize() with blank credentials: %s", err)
	}
	triggerExistingResourceRecordSet = true
	setCallMockDefaults(d)
	err = d.SetResourceRecordSets(testProjectID, testManagedZone, []*v1.ResourceRecordSet{testResourceRecordSet})
	if err != nil {
		t.Errorf("Got unexpected error during dns.GetResourceRecordSet(): %s", err)
	}
}

func TestGetResourceRecordSets(t *testing.T) {
	d := &DNS{}
	err := d.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during dns.Initialize() with blank credentials: %s", err)
	}
	triggerExistingResourceRecordSet = true
	setCallMockDefaults(d)
	result, err := d.GetResourceRecordSets(testProjectID, testManagedZone)
	if err != nil {
		t.Errorf("Got unexpected error during dns.GetResourceRecordSets(): %s", err)
	}
	if len(result) != 1 {
		t.Errorf("Got unexpected result from dns.GetResourceRecordSets(), expecting list of length \"1\", but got %d", len(result))
	}
}

func TestDeleteResourceRecordSets(t *testing.T) {
	d := &DNS{}
	err := d.Initialize("", loggermock.GetLogMock())
	if err != nil {
		t.Errorf("Got unexpected error during dns.Initialize() with blank credentials: %s", err)
	}
	triggerExistingResourceRecordSet = true
	setCallMockDefaults(d)
	err = d.DeleteResourceRecordSets(testProjectID, testManagedZone)
	if err != nil {
		t.Errorf("Got unexpected error during dns.DeleteResourceRecordSets(): %s", err)
	}
}
