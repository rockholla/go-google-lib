// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	gax "github.com/googleapis/gax-go/v2"
	kms "google.golang.org/genproto/googleapis/cloud/kms/v1"

	mock "github.com/stretchr/testify/mock"
)

// ClientInterface is an autogenerated mock type for the ClientInterface type
type ClientInterface struct {
	mock.Mock
}

// Decrypt provides a mock function with given fields: ctx, req, opts
func (_m *ClientInterface) Decrypt(ctx context.Context, req *kms.DecryptRequest, opts ...gax.CallOption) (*kms.DecryptResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, req)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *kms.DecryptResponse
	if rf, ok := ret.Get(0).(func(context.Context, *kms.DecryptRequest, ...gax.CallOption) *kms.DecryptResponse); ok {
		r0 = rf(ctx, req, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kms.DecryptResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *kms.DecryptRequest, ...gax.CallOption) error); ok {
		r1 = rf(ctx, req, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Encrypt provides a mock function with given fields: ctx, req, opts
func (_m *ClientInterface) Encrypt(ctx context.Context, req *kms.EncryptRequest, opts ...gax.CallOption) (*kms.EncryptResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, req)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *kms.EncryptResponse
	if rf, ok := ret.Get(0).(func(context.Context, *kms.EncryptRequest, ...gax.CallOption) *kms.EncryptResponse); ok {
		r0 = rf(ctx, req, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kms.EncryptResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *kms.EncryptRequest, ...gax.CallOption) error); ok {
		r1 = rf(ctx, req, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
