// Package dns is the library for google cloud dns operations
package dns

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/rockholla/go-google-lib/dns/calls"
	"github.com/rockholla/go-lib/logger"
	v1 "google.golang.org/api/dns/v1"
	"google.golang.org/api/option"
)

// Interface represents functionality for DNS
type Interface interface {
	Initialize(credentials string, log logger.Interface) error
	GetResourceRecordSets(projectID string, managedZone string) ([]*v1.ResourceRecordSet, error)
	GetResourceRecordSet(projectID string, managedZone string, name string) (*v1.ResourceRecordSet, error)
	SetResourceRecordSets(projectID string, managedZone string, records []*v1.ResourceRecordSet) error
	DeleteResourceRecordSets(projectID string, managedZone string) error
}

// DNS is a wrapper around the google-provided sdks/apis for google.golang.org/api/dns/*
type DNS struct {
	log                logger.Interface
	V1                 *v1.Service
	Calls              *Calls
	PendingWaitSeconds int64
}

// Calls are interfaces for making the actual calls to various underlying apis
type Calls struct {
	ChangesCreate          calls.ChangesCreateCallInterface
	ResourceRecordSetsList calls.ResourceRecordSetsListCallInterface
}

// Initialize sets up necessary google-provided sdks and other local data
func (d *DNS) Initialize(credentials string, log logger.Interface) error {
	var err error
	ctx := context.Background()
	d.log = log
	d.PendingWaitSeconds = 5
	d.Calls = &Calls{
		ChangesCreate:          &calls.ChangesCreateCall{},
		ResourceRecordSetsList: &calls.ResourceRecordSetsListCall{},
	}
	if credentials != "" {
		if d.V1, err = v1.NewService(ctx, option.WithCredentialsJSON([]byte(credentials))); err != nil {
			return err
		}
	} else {
		if d.V1, err = v1.NewService(ctx); err != nil {
			return err
		}
	}
	return nil
}

// GetResourceRecordSets will return all resource record sets for a managed zone
func (d *DNS) GetResourceRecordSets(projectID string, managedZone string) ([]*v1.ResourceRecordSet, error) {
	ctx := context.Background()
	rrsService := v1.NewResourceRecordSetsService(d.V1)
	rrsListCall := rrsService.List(projectID, managedZone).Context(ctx)
	rrsList, err := d.Calls.ResourceRecordSetsList.Do(rrsListCall)
	if err != nil {
		return nil, err
	}
	return rrsList.Rrsets, nil
}

// GetResourceRecordSet will search for an existing record set by the resourcer record set name
func (d *DNS) GetResourceRecordSet(projectID string, managedZone string, name string) (*v1.ResourceRecordSet, error) {
	ctx := context.Background()
	rrsService := v1.NewResourceRecordSetsService(d.V1)
	rrsListCall := rrsService.List(projectID, managedZone).Context(ctx).Name(name)
	rrsList, err := d.Calls.ResourceRecordSetsList.Do(rrsListCall)
	if err != nil {
		return nil, err
	}
	if len(rrsList.Rrsets) == 0 {
		return nil, nil
	}
	return rrsList.Rrsets[0], nil
}

// SetResourceRecordSets will create or update a DNS zone with one or more record sets
func (d *DNS) SetResourceRecordSets(projectID string, managedZone string, records []*v1.ResourceRecordSet) error {
	var deletions []*v1.ResourceRecordSet
	var additions []*v1.ResourceRecordSet
	var change *v1.Change
	logItems := []string{}
	for _, record := range records {
		existing, err := d.GetResourceRecordSet(projectID, managedZone, record.Name)
		if err != nil {
			return fmt.Errorf("Error trying to get existing resource record set: %s", err)
		}
		action := "creating"
		if existing != nil {
			deletions = append(deletions, existing)
			action = "recreating"
		}
		logItems = append(logItems, fmt.Sprintf("====> %s %s => %s %s", action, record.Name, record.Type, strings.Join(record.Rrdatas, ",")))
		additions = append(additions, record)
	}
	d.log.Info("Ensuring the DNS zone %s has the following records:", managedZone)
	for _, item := range logItems {
		d.log.ListItem(item)
	}
	if len(deletions) > 0 {
		change = &v1.Change{
			Deletions: deletions,
		}
		if err := d.executeChange(projectID, managedZone, change); err != nil {
			return err
		}
	}
	change = &v1.Change{
		Additions: additions,
	}
	if err := d.executeChange(projectID, managedZone, change); err != nil {
		return err
	}
	return nil
}

// DeleteResourceRecordSets will remove all resource record sets from a managed zone
func (d *DNS) DeleteResourceRecordSets(projectID string, managedZone string) error {
	var deletions []*v1.ResourceRecordSet
	resourceRecordSets, err := d.GetResourceRecordSets(projectID, managedZone)
	if err != nil {
		return err
	}
	d.log.Info("Deleting all records from DNS zone %s:", managedZone)
	for _, resourceRecordSet := range resourceRecordSets {
		if resourceRecordSet.Type == "SOA" || resourceRecordSet.Type == "NS" {
			continue
		}
		deletions = append(deletions, resourceRecordSet)
		d.log.ListItem("%s %s", resourceRecordSet.Type, resourceRecordSet.Name)
	}
	change := &v1.Change{
		Deletions: deletions,
	}
	if err := d.executeChange(projectID, managedZone, change); err != nil {
		return err
	}
	return nil
}

func (d *DNS) executeChange(projectID string, managedZone string, change *v1.Change) error {
	ctx := context.Background()
	changesService := v1.NewChangesService(d.V1)
	var changesCreateCall *v1.ChangesCreateCall
	changesCreateCall = changesService.Create(projectID, managedZone, change).Context(ctx)
	processedChange, err := d.Calls.ChangesCreate.Do(changesCreateCall)
	if err != nil {
		return err
	}
	if processedChange.Status == "pending" {
		time.Sleep(time.Duration(d.PendingWaitSeconds) * time.Second)
	}
	return nil
}
