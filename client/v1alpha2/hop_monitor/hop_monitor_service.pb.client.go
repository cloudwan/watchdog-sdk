// Code generated by protoc-gen-goten-client
// API: HopMonitorService
// DO NOT EDIT!!!

package hop_monitor_client

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

// HopMonitorServiceClient is the client API for HopMonitorService.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HopMonitorServiceClient interface {
	RunHopMonitor(ctx context.Context, in *RunHopMonitorRequest, opts ...grpc.CallOption) (RunHopMonitorClientStream, error)
}

type client struct {
	cc grpc.ClientConnInterface
}

func NewHopMonitorServiceClient(cc grpc.ClientConnInterface) HopMonitorServiceClient {
	return &client{cc}
}

func (c *client) RunHopMonitor(ctx context.Context, in *RunHopMonitorRequest, opts ...grpc.CallOption) (RunHopMonitorClientStream, error) {
	stream, err := c.cc.NewStream(ctx,
		&grpc.StreamDesc{
			StreamName:    "RunHopMonitor",
			ServerStreams: true,
		},
		"/ntt.watchdog.v1alpha2.HopMonitorService/RunHopMonitor", opts...)
	if err != nil {
		return nil, err
	}
	x := &runHopMonitorRunHopMonitorClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RunHopMonitorClientStream interface {
	Recv() (*RunHopMonitorResponse, error)
	grpc.ClientStream
}

type runHopMonitorRunHopMonitorClient struct {
	grpc.ClientStream
}

func (x *runHopMonitorRunHopMonitorClient) Recv() (*RunHopMonitorResponse, error) {
	m := new(RunHopMonitorResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}
