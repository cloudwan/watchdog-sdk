// Code generated by protoc-gen-goten-client
// API: TagService
// DO NOT EDIT!!!

package tag_client

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/runtime/protoimpl"
)

// proto imports
import (
	tag "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/tag"
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
	_ = &tag.Tag{}
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

// TagServiceClient is the client API for TagService.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TagServiceClient interface {
	GetTag(ctx context.Context, in *GetTagRequest, opts ...grpc.CallOption) (*tag.Tag, error)
	BatchGetTags(ctx context.Context, in *BatchGetTagsRequest, opts ...grpc.CallOption) (*BatchGetTagsResponse, error)
	ListTags(ctx context.Context, in *ListTagsRequest, opts ...grpc.CallOption) (*ListTagsResponse, error)
	WatchTag(ctx context.Context, in *WatchTagRequest, opts ...grpc.CallOption) (WatchTagClientStream, error)
	WatchTags(ctx context.Context, in *WatchTagsRequest, opts ...grpc.CallOption) (WatchTagsClientStream, error)
	CreateTag(ctx context.Context, in *CreateTagRequest, opts ...grpc.CallOption) (*tag.Tag, error)
	UpdateTag(ctx context.Context, in *UpdateTagRequest, opts ...grpc.CallOption) (*tag.Tag, error)
	DeleteTag(ctx context.Context, in *DeleteTagRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type client struct {
	cc grpc.ClientConnInterface
}

func NewTagServiceClient(cc grpc.ClientConnInterface) TagServiceClient {
	return &client{cc}
}

func (c *client) GetTag(ctx context.Context, in *GetTagRequest, opts ...grpc.CallOption) (*tag.Tag, error) {
	out := new(tag.Tag)
	err := c.cc.Invoke(ctx, "/ntt.watchdog.v1alpha2.TagService/GetTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) BatchGetTags(ctx context.Context, in *BatchGetTagsRequest, opts ...grpc.CallOption) (*BatchGetTagsResponse, error) {
	out := new(BatchGetTagsResponse)
	err := c.cc.Invoke(ctx, "/ntt.watchdog.v1alpha2.TagService/BatchGetTags", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) ListTags(ctx context.Context, in *ListTagsRequest, opts ...grpc.CallOption) (*ListTagsResponse, error) {
	out := new(ListTagsResponse)
	err := c.cc.Invoke(ctx, "/ntt.watchdog.v1alpha2.TagService/ListTags", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) WatchTag(ctx context.Context, in *WatchTagRequest, opts ...grpc.CallOption) (WatchTagClientStream, error) {
	stream, err := c.cc.NewStream(ctx,
		&grpc.StreamDesc{
			StreamName:    "WatchTag",
			ServerStreams: true,
		},
		"/ntt.watchdog.v1alpha2.TagService/WatchTag", opts...)
	if err != nil {
		return nil, err
	}
	x := &watchTagWatchTagClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WatchTagClientStream interface {
	Recv() (*WatchTagResponse, error)
	grpc.ClientStream
}

type watchTagWatchTagClient struct {
	grpc.ClientStream
}

func (x *watchTagWatchTagClient) Recv() (*WatchTagResponse, error) {
	m := new(WatchTagResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *client) WatchTags(ctx context.Context, in *WatchTagsRequest, opts ...grpc.CallOption) (WatchTagsClientStream, error) {
	stream, err := c.cc.NewStream(ctx,
		&grpc.StreamDesc{
			StreamName:    "WatchTags",
			ServerStreams: true,
		},
		"/ntt.watchdog.v1alpha2.TagService/WatchTags", opts...)
	if err != nil {
		return nil, err
	}
	x := &watchTagsWatchTagsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WatchTagsClientStream interface {
	Recv() (*WatchTagsResponse, error)
	grpc.ClientStream
}

type watchTagsWatchTagsClient struct {
	grpc.ClientStream
}

func (x *watchTagsWatchTagsClient) Recv() (*WatchTagsResponse, error) {
	m := new(WatchTagsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *client) CreateTag(ctx context.Context, in *CreateTagRequest, opts ...grpc.CallOption) (*tag.Tag, error) {
	out := new(tag.Tag)
	err := c.cc.Invoke(ctx, "/ntt.watchdog.v1alpha2.TagService/CreateTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) UpdateTag(ctx context.Context, in *UpdateTagRequest, opts ...grpc.CallOption) (*tag.Tag, error) {
	out := new(tag.Tag)
	err := c.cc.Invoke(ctx, "/ntt.watchdog.v1alpha2.TagService/UpdateTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) DeleteTag(ctx context.Context, in *DeleteTagRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/ntt.watchdog.v1alpha2.TagService/DeleteTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
