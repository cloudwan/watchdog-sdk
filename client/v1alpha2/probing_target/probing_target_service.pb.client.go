// Code generated by protoc-gen-goten-client
// API: ProbingTargetService
// DO NOT EDIT!!!

package probing_target_client

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/runtime/protoimpl"
)

// proto imports
import (
	probing_target "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/probing_target"
	empty "github.com/golang/protobuf/ptypes/empty"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = protoimpl.DescBuilder{}
	_ = new(context.Context)
	_ = grpc.ClientConn{}
)

// make sure we're using proto imports
var (
	_ = &empty.Empty{}
	_ = &probing_target.ProbingTarget{}
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

// ProbingTargetServiceClient is the client API for ProbingTargetService.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProbingTargetServiceClient interface {
	GetProbingTarget(ctx context.Context, in *GetProbingTargetRequest, opts ...grpc.CallOption) (*probing_target.ProbingTarget, error)
	BatchGetProbingTargets(ctx context.Context, in *BatchGetProbingTargetsRequest, opts ...grpc.CallOption) (*BatchGetProbingTargetsResponse, error)
	ListProbingTargets(ctx context.Context, in *ListProbingTargetsRequest, opts ...grpc.CallOption) (*ListProbingTargetsResponse, error)
	WatchProbingTarget(ctx context.Context, in *WatchProbingTargetRequest, opts ...grpc.CallOption) (WatchProbingTargetClientStream, error)
	WatchProbingTargets(ctx context.Context, in *WatchProbingTargetsRequest, opts ...grpc.CallOption) (WatchProbingTargetsClientStream, error)
	CreateProbingTarget(ctx context.Context, in *CreateProbingTargetRequest, opts ...grpc.CallOption) (*probing_target.ProbingTarget, error)
	UpdateProbingTarget(ctx context.Context, in *UpdateProbingTargetRequest, opts ...grpc.CallOption) (*probing_target.ProbingTarget, error)
	DeleteProbingTarget(ctx context.Context, in *DeleteProbingTargetRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	SearchProbingTargets(ctx context.Context, in *SearchProbingTargetsRequest, opts ...grpc.CallOption) (*SearchProbingTargetsResponse, error)
}

type client struct {
	cc grpc.ClientConnInterface
}

func NewProbingTargetServiceClient(cc grpc.ClientConnInterface) ProbingTargetServiceClient {
	return &client{cc}
}

func (c *client) GetProbingTarget(ctx context.Context, in *GetProbingTargetRequest, opts ...grpc.CallOption) (*probing_target.ProbingTarget, error) {
	out := new(probing_target.ProbingTarget)
	err := c.cc.Invoke(ctx, "/ntt.watchdog.v1alpha2.ProbingTargetService/GetProbingTarget", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) BatchGetProbingTargets(ctx context.Context, in *BatchGetProbingTargetsRequest, opts ...grpc.CallOption) (*BatchGetProbingTargetsResponse, error) {
	out := new(BatchGetProbingTargetsResponse)
	err := c.cc.Invoke(ctx, "/ntt.watchdog.v1alpha2.ProbingTargetService/BatchGetProbingTargets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) ListProbingTargets(ctx context.Context, in *ListProbingTargetsRequest, opts ...grpc.CallOption) (*ListProbingTargetsResponse, error) {
	out := new(ListProbingTargetsResponse)
	err := c.cc.Invoke(ctx, "/ntt.watchdog.v1alpha2.ProbingTargetService/ListProbingTargets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) WatchProbingTarget(ctx context.Context, in *WatchProbingTargetRequest, opts ...grpc.CallOption) (WatchProbingTargetClientStream, error) {
	stream, err := c.cc.NewStream(ctx,
		&grpc.StreamDesc{
			StreamName:    "WatchProbingTarget",
			ServerStreams: true,
		},
		"/ntt.watchdog.v1alpha2.ProbingTargetService/WatchProbingTarget", opts...)
	if err != nil {
		return nil, err
	}
	x := &watchProbingTargetWatchProbingTargetClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WatchProbingTargetClientStream interface {
	Recv() (*WatchProbingTargetResponse, error)
	grpc.ClientStream
}

type watchProbingTargetWatchProbingTargetClient struct {
	grpc.ClientStream
}

func (x *watchProbingTargetWatchProbingTargetClient) Recv() (*WatchProbingTargetResponse, error) {
	m := new(WatchProbingTargetResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *client) WatchProbingTargets(ctx context.Context, in *WatchProbingTargetsRequest, opts ...grpc.CallOption) (WatchProbingTargetsClientStream, error) {
	stream, err := c.cc.NewStream(ctx,
		&grpc.StreamDesc{
			StreamName:    "WatchProbingTargets",
			ServerStreams: true,
		},
		"/ntt.watchdog.v1alpha2.ProbingTargetService/WatchProbingTargets", opts...)
	if err != nil {
		return nil, err
	}
	x := &watchProbingTargetsWatchProbingTargetsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WatchProbingTargetsClientStream interface {
	Recv() (*WatchProbingTargetsResponse, error)
	grpc.ClientStream
}

type watchProbingTargetsWatchProbingTargetsClient struct {
	grpc.ClientStream
}

func (x *watchProbingTargetsWatchProbingTargetsClient) Recv() (*WatchProbingTargetsResponse, error) {
	m := new(WatchProbingTargetsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *client) CreateProbingTarget(ctx context.Context, in *CreateProbingTargetRequest, opts ...grpc.CallOption) (*probing_target.ProbingTarget, error) {
	out := new(probing_target.ProbingTarget)
	err := c.cc.Invoke(ctx, "/ntt.watchdog.v1alpha2.ProbingTargetService/CreateProbingTarget", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) UpdateProbingTarget(ctx context.Context, in *UpdateProbingTargetRequest, opts ...grpc.CallOption) (*probing_target.ProbingTarget, error) {
	out := new(probing_target.ProbingTarget)
	err := c.cc.Invoke(ctx, "/ntt.watchdog.v1alpha2.ProbingTargetService/UpdateProbingTarget", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) DeleteProbingTarget(ctx context.Context, in *DeleteProbingTargetRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/ntt.watchdog.v1alpha2.ProbingTargetService/DeleteProbingTarget", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) SearchProbingTargets(ctx context.Context, in *SearchProbingTargetsRequest, opts ...grpc.CallOption) (*SearchProbingTargetsResponse, error) {
	out := new(SearchProbingTargetsResponse)
	err := c.cc.Invoke(ctx, "/ntt.watchdog.v1alpha2.ProbingTargetService/SearchProbingTargets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
