// Package admin is the library for google admin/directory operations
package admin

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/rockholla/go-google-lib/admin/calls"
	"github.com/rockholla/go-lib/logger"
	"golang.org/x/oauth2/google"
	dirv1 "google.golang.org/api/admin/directory/v1"
	"google.golang.org/api/option"
)

// Interface represents functionality for Admin
type Interface interface {
	Initialize(credentialsJSON string, domain string, adminUsername string, log logger.Interface) error
	EnsureGroup(name string, description string) (*dirv1.Group, error)
	DeleteGroup(name string) error
	EnsureMembership(group string, member string) (*dirv1.Member, error)
	DeleteMembership(group string, member string) error
}

// Admin wraps google-provided apis for interacting with google.golang.org/api/admin/*
type Admin struct {
	log    logger.Interface
	DirV1  *dirv1.Service
	Calls  *Calls
	domain string
}

// Calls are interfaces for making the actual calls to various underlying apis
type Calls struct {
	GroupsInsert  calls.GroupsInsertCallInterface
	GroupsUpdate  calls.GroupsUpdateCallInterface
	GroupsGet     calls.GroupsGetCallInterface
	GroupsDelete  calls.GroupsDeleteCallInterface
	MembersGet    calls.MembersGetCallInterface
	MembersInsert calls.MembersInsertCallInterface
	MembersDelete calls.MembersDeleteCallInterface
}

// Initialize sets up necessary google-provided sdks and other local data
func (a *Admin) Initialize(credentialsJSON string, domain string, adminUsername string, log logger.Interface) error {
	var err error
	ctx := context.Background()
	a.log = log
	a.Calls = &Calls{
		GroupsInsert:  &calls.GroupsInsertCall{},
		GroupsUpdate:  &calls.GroupsUpdateCall{},
		GroupsGet:     &calls.GroupsGetCall{},
		GroupsDelete:  &calls.GroupsDeleteCall{},
		MembersGet:    &calls.MembersGetCall{},
		MembersInsert: &calls.MembersInsertCall{},
		MembersDelete: &calls.MembersDeleteCall{},
	}
	a.domain = domain
	if credentialsJSON == "" {
		return errors.New("credentials must be explicitly provided to the google admin package")
	}
	config, err := google.JWTConfigFromJSON([]byte(credentialsJSON), dirv1.AdminDirectoryGroupScope)
	if err != nil {
		return err
	}
	config.Subject = fmt.Sprintf("%s@%s", adminUsername, domain)
	a.log.Info("For Google admin and directory operations: impersonating %s", config.Subject)
	client := config.Client(ctx)
	if a.DirV1, err = dirv1.NewService(ctx, option.WithHTTPClient(client)); err != nil {
		return err
	}
	return nil
}

// EnsureGroup will make sure that a particular group exists in Google admin
func (a *Admin) EnsureGroup(name string, description string) (*dirv1.Group, error) {
	ctx := context.Background()
	email := name
	if !strings.Contains(email, "@") {
		email = fmt.Sprintf("%s@%s", name, a.domain)
	}
	apiGroup := &dirv1.Group{
		Email:       email,
		Name:        name,
		Description: description,
	}
	a.log.InfoPart("Ensuring that Google group %s exists...", email)
	groupsService := dirv1.NewGroupsService(a.DirV1)
	groupsGetCall := groupsService.Get(email).Context(ctx)
	existingGroup, err := a.Calls.GroupsGet.Do(groupsGetCall)
	if err != nil && !strings.Contains(strings.ToLower(err.Error()), "notfound") {
		a.log.InfoPart("error\n")
		return nil, err
	}
	if existingGroup == nil {
		a.log.InfoPart("creating...")
		groupsInsertCall := groupsService.Insert(apiGroup)
		_, err = a.Calls.GroupsInsert.Do(groupsInsertCall)
		if err != nil {
			a.log.InfoPart("error\n")
			return nil, err
		}
	} else {
		a.log.InfoPart("updating...")
		groupsUpdateCall := groupsService.Update(email, apiGroup)
		_, err = a.Calls.GroupsUpdate.Do(groupsUpdateCall)
		if err != nil {
			a.log.InfoPart("error\n")
			return nil, err
		}
	}
	a.log.InfoPart("done\n")
	return apiGroup, nil
}

// EnsureMembership will make sure that a member is part of a group in Google admin
func (a *Admin) EnsureMembership(group string, member string) (*dirv1.Member, error) {
	ctx := context.Background()
	groupEmail := group
	if !strings.Contains(groupEmail, "@") {
		groupEmail = fmt.Sprintf("%s@%s", group, a.domain)
	}
	memberEmail := member
	if !strings.Contains(memberEmail, "@") {
		memberEmail = fmt.Sprintf("%s@%s", member, a.domain)
	}
	a.log.InfoPart("Ensuring that %s is a member of Google group %s...", memberEmail, groupEmail)
	membersService := dirv1.NewMembersService(a.DirV1)
	membersGetCall := membersService.Get(groupEmail, memberEmail).Context(ctx)
	existingMember, err := a.Calls.MembersGet.Do(membersGetCall)
	if err != nil && !strings.Contains(strings.ToLower(err.Error()), "notfound") {
		a.log.InfoPart("error\n")
		return nil, err
	}
	if existingMember == nil {
		a.log.InfoPart("adding...")
		membersInsertCall := membersService.Insert(groupEmail, &dirv1.Member{
			Email: memberEmail,
		}).Context(ctx)
		newMember, err := a.Calls.MembersInsert.Do(membersInsertCall)
		if err != nil {
			a.log.InfoPart("error\n")
			return nil, err
		}
		a.log.InfoPart("done\n")
		return newMember, nil
	}
	a.log.InfoPart("already a member\n")
	return existingMember, nil
}

// DeleteGroup will delete a Google group
func (a *Admin) DeleteGroup(name string) error {
	ctx := context.Background()
	email := name
	if !strings.Contains(email, "@") {
		email = fmt.Sprintf("%s@%s", name, a.domain)
	}
	a.log.Info("Ensuring that group %s is deleted", email)
	groupsService := dirv1.NewGroupsService(a.DirV1)
	groupsDeleteCall := groupsService.Delete(email).Context(ctx)
	err := a.Calls.GroupsDelete.Do(groupsDeleteCall)
	if err != nil && !strings.Contains(strings.ToLower(err.Error()), "notfound") {
		return err
	}
	return nil
}

// DeleteMembership will remove a member from a Google group
func (a *Admin) DeleteMembership(group string, member string) error {
	ctx := context.Background()
	groupEmail := group
	if !strings.Contains(groupEmail, "@") {
		groupEmail = fmt.Sprintf("%s@%s", group, a.domain)
	}
	memberEmail := member
	if !strings.Contains(memberEmail, "@") {
		memberEmail = fmt.Sprintf("%s@%s", member, a.domain)
	}
	a.log.InfoPart("Ensuring member %s is removed from group %s...", memberEmail, groupEmail)
	membersService := dirv1.NewMembersService(a.DirV1)
	membersGetCall := membersService.Get(groupEmail, memberEmail).Context(ctx)
	existingMember, err := a.Calls.MembersGet.Do(membersGetCall)
	if err != nil && !strings.Contains(strings.ToLower(err.Error()), "notfound") {
		a.log.InfoPart("error\n")
		return nil
	}
	if existingMember != nil {
		a.log.InfoPart("removing...")
		membersDeleteCall := membersService.Delete(groupEmail, memberEmail).Context(ctx)
		err := a.Calls.MembersDelete.Do(membersDeleteCall)
		if err != nil {
			a.log.InfoPart("error\n")
			return err
		}
		a.log.InfoPart("done\n")
		return nil
	}
	a.log.InfoPart("not a member\n")
	return nil
}
