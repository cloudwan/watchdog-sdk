// Code generated by protoc-gen-goten-client
// API: SpeedTestService
// DO NOT EDIT!!!

package speed_test_client

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

// SpeedTestServiceClient is the client API for SpeedTestService.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SpeedTestServiceClient interface {
	RunSpeedTest(ctx context.Context, in *RunSpeedTestRequest, opts ...grpc.CallOption) (RunSpeedTestClientStream, error)
}

type client struct {
	cc grpc.ClientConnInterface
}

func NewSpeedTestServiceClient(cc grpc.ClientConnInterface) SpeedTestServiceClient {
	return &client{cc}
}

func (c *client) RunSpeedTest(ctx context.Context, in *RunSpeedTestRequest, opts ...grpc.CallOption) (RunSpeedTestClientStream, error) {
	stream, err := c.cc.NewStream(ctx,
		&grpc.StreamDesc{
			StreamName:    "RunSpeedTest",
			ServerStreams: true,
		},
		"/ntt.watchdog.v1alpha2.SpeedTestService/RunSpeedTest", opts...)
	if err != nil {
		return nil, err
	}
	x := &runSpeedTestRunSpeedTestClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RunSpeedTestClientStream interface {
	Recv() (*RunSpeedTestResponse, error)
	grpc.ClientStream
}

type runSpeedTestRunSpeedTestClient struct {
	grpc.ClientStream
}

func (x *runSpeedTestRunSpeedTestClient) Recv() (*RunSpeedTestResponse, error) {
	m := new(RunSpeedTestResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}
