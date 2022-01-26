// Code generated by protoc-gen-goten-client
// API: ActivationService
// DO NOT EDIT!!!

package activation_client

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/runtime/protoimpl"
)

// proto imports
import (
	empty "github.com/golang/protobuf/ptypes/empty"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = protoimpl.DescBuilder{}
	_ = context.Context(nil)
	_ = grpc.ClientConn{}
)

// make sure we're using proto imports
var (
	_ = &empty.Empty{}
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

// ActivationServiceClient is the client API for ActivationService.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ActivationServiceClient interface {
	Activation(ctx context.Context, opts ...grpc.CallOption) (ActivationClientStream, error)
	SendActivationInvitation(ctx context.Context, in *SendActivationInvitationRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	ResetActivation(ctx context.Context, in *ResetActivationRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type client struct {
	cc grpc.ClientConnInterface
}

func NewActivationServiceClient(cc grpc.ClientConnInterface) ActivationServiceClient {
	return &client{cc}
}

func (c *client) Activation(ctx context.Context, opts ...grpc.CallOption) (ActivationClientStream, error) {
	stream, err := c.cc.NewStream(ctx,
		&grpc.StreamDesc{
			StreamName:    "Activation",
			ServerStreams: true,
			ClientStreams: true,
		},
		"/ntt.watchdog.v1alpha.ActivationService/Activation", opts...)
	if err != nil {
		return nil, err
	}
	x := &activationActivationClient{stream}
	return x, nil
}

type ActivationClientStream interface {
	Send(*ActivationRequest) error
	Recv() (*ActivationResponse, error)
	grpc.ClientStream
}

type activationActivationClient struct {
	grpc.ClientStream
}

func (x *activationActivationClient) Send(m *ActivationRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *activationActivationClient) Recv() (*ActivationResponse, error) {
	m := new(ActivationResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *client) SendActivationInvitation(ctx context.Context, in *SendActivationInvitationRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/ntt.watchdog.v1alpha.ActivationService/SendActivationInvitation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) ResetActivation(ctx context.Context, in *ResetActivationRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/ntt.watchdog.v1alpha.ActivationService/ResetActivation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
