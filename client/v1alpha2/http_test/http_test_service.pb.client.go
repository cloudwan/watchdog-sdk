// Code generated by protoc-gen-goten-client
// API: HTTPTestService
// DO NOT EDIT!!!

package http_test_client

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/runtime/protoimpl"
)

// proto imports
import (
	probe "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/probe"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = protoimpl.DescBuilder{}
	_ = new(context.Context)
	_ = grpc.ClientConn{}
)

// make sure we're using proto imports
var (
	_ = &probe.Probe{}
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// HTTPTestServiceClient is the client API for HTTPTestService.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HTTPTestServiceClient interface {
	RunHTTPTest(ctx context.Context, in *RunHTTPTestRequest, opts ...grpc.CallOption) (*RunHTTPTestResponse, error)
}

type client struct {
	cc grpc.ClientConnInterface
}

func NewHTTPTestServiceClient(cc grpc.ClientConnInterface) HTTPTestServiceClient {
	return &client{cc}
}

func (c *client) RunHTTPTest(ctx context.Context, in *RunHTTPTestRequest, opts ...grpc.CallOption) (*RunHTTPTestResponse, error) {
	out := new(RunHTTPTestResponse)
	err := c.cc.Invoke(ctx, "/ntt.watchdog.v1alpha2.HTTPTestService/RunHTTPTest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
