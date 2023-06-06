// Code generated by mockery v2.10.4. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	storage_v1 "github.com/kjschnei001/jaeger/proto-gen/storage_v1"
)

// PluginCapabilitiesClient is an autogenerated mock type for the PluginCapabilitiesClient type
type PluginCapabilitiesClient struct {
	mock.Mock
}

// Capabilities provides a mock function with given fields: ctx, in, opts
func (_m *PluginCapabilitiesClient) Capabilities(ctx context.Context, in *storage_v1.CapabilitiesRequest, opts ...grpc.CallOption) (*storage_v1.CapabilitiesResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *storage_v1.CapabilitiesResponse
	if rf, ok := ret.Get(0).(func(context.Context, *storage_v1.CapabilitiesRequest, ...grpc.CallOption) *storage_v1.CapabilitiesResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storage_v1.CapabilitiesResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *storage_v1.CapabilitiesRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
