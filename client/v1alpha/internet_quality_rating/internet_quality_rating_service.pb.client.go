// Code generated by protoc-gen-goten-client
// API: InternetQualityRatingService
// DO NOT EDIT!!!

package internet_quality_rating_client

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/runtime/protoimpl"
)

// proto imports
import (
	internet_quality_rating "github.com/cloudwan/watchdog-sdk/resources/v1alpha/internet_quality_rating"
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
	_ = &internet_quality_rating.InternetQualityRating{}
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

// InternetQualityRatingServiceClient is the client API for InternetQualityRatingService.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type InternetQualityRatingServiceClient interface {
	GetInternetQualityRating(ctx context.Context, in *GetInternetQualityRatingRequest, opts ...grpc.CallOption) (*internet_quality_rating.InternetQualityRating, error)
	BatchGetInternetQualityRatings(ctx context.Context, in *BatchGetInternetQualityRatingsRequest, opts ...grpc.CallOption) (*BatchGetInternetQualityRatingsResponse, error)
	ListInternetQualityRatings(ctx context.Context, in *ListInternetQualityRatingsRequest, opts ...grpc.CallOption) (*ListInternetQualityRatingsResponse, error)
	WatchInternetQualityRating(ctx context.Context, in *WatchInternetQualityRatingRequest, opts ...grpc.CallOption) (WatchInternetQualityRatingClientStream, error)
	WatchInternetQualityRatings(ctx context.Context, in *WatchInternetQualityRatingsRequest, opts ...grpc.CallOption) (WatchInternetQualityRatingsClientStream, error)
	CreateInternetQualityRating(ctx context.Context, in *CreateInternetQualityRatingRequest, opts ...grpc.CallOption) (*internet_quality_rating.InternetQualityRating, error)
	UpdateInternetQualityRating(ctx context.Context, in *UpdateInternetQualityRatingRequest, opts ...grpc.CallOption) (*internet_quality_rating.InternetQualityRating, error)
	DeleteInternetQualityRating(ctx context.Context, in *DeleteInternetQualityRatingRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type client struct {
	cc grpc.ClientConnInterface
}

func NewInternetQualityRatingServiceClient(cc grpc.ClientConnInterface) InternetQualityRatingServiceClient {
	return &client{cc}
}

func (c *client) GetInternetQualityRating(ctx context.Context, in *GetInternetQualityRatingRequest, opts ...grpc.CallOption) (*internet_quality_rating.InternetQualityRating, error) {
	out := new(internet_quality_rating.InternetQualityRating)
	err := c.cc.Invoke(ctx, "/ntt.watchdog.v1alpha.InternetQualityRatingService/GetInternetQualityRating", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) BatchGetInternetQualityRatings(ctx context.Context, in *BatchGetInternetQualityRatingsRequest, opts ...grpc.CallOption) (*BatchGetInternetQualityRatingsResponse, error) {
	out := new(BatchGetInternetQualityRatingsResponse)
	err := c.cc.Invoke(ctx, "/ntt.watchdog.v1alpha.InternetQualityRatingService/BatchGetInternetQualityRatings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) ListInternetQualityRatings(ctx context.Context, in *ListInternetQualityRatingsRequest, opts ...grpc.CallOption) (*ListInternetQualityRatingsResponse, error) {
	out := new(ListInternetQualityRatingsResponse)
	err := c.cc.Invoke(ctx, "/ntt.watchdog.v1alpha.InternetQualityRatingService/ListInternetQualityRatings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) WatchInternetQualityRating(ctx context.Context, in *WatchInternetQualityRatingRequest, opts ...grpc.CallOption) (WatchInternetQualityRatingClientStream, error) {
	stream, err := c.cc.NewStream(ctx,
		&grpc.StreamDesc{
			StreamName:    "WatchInternetQualityRating",
			ServerStreams: true,
		},
		"/ntt.watchdog.v1alpha.InternetQualityRatingService/WatchInternetQualityRating", opts...)
	if err != nil {
		return nil, err
	}
	x := &watchInternetQualityRatingWatchInternetQualityRatingClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WatchInternetQualityRatingClientStream interface {
	Recv() (*WatchInternetQualityRatingResponse, error)
	grpc.ClientStream
}

type watchInternetQualityRatingWatchInternetQualityRatingClient struct {
	grpc.ClientStream
}

func (x *watchInternetQualityRatingWatchInternetQualityRatingClient) Recv() (*WatchInternetQualityRatingResponse, error) {
	m := new(WatchInternetQualityRatingResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *client) WatchInternetQualityRatings(ctx context.Context, in *WatchInternetQualityRatingsRequest, opts ...grpc.CallOption) (WatchInternetQualityRatingsClientStream, error) {
	stream, err := c.cc.NewStream(ctx,
		&grpc.StreamDesc{
			StreamName:    "WatchInternetQualityRatings",
			ServerStreams: true,
		},
		"/ntt.watchdog.v1alpha.InternetQualityRatingService/WatchInternetQualityRatings", opts...)
	if err != nil {
		return nil, err
	}
	x := &watchInternetQualityRatingsWatchInternetQualityRatingsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WatchInternetQualityRatingsClientStream interface {
	Recv() (*WatchInternetQualityRatingsResponse, error)
	grpc.ClientStream
}

type watchInternetQualityRatingsWatchInternetQualityRatingsClient struct {
	grpc.ClientStream
}

func (x *watchInternetQualityRatingsWatchInternetQualityRatingsClient) Recv() (*WatchInternetQualityRatingsResponse, error) {
	m := new(WatchInternetQualityRatingsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *client) CreateInternetQualityRating(ctx context.Context, in *CreateInternetQualityRatingRequest, opts ...grpc.CallOption) (*internet_quality_rating.InternetQualityRating, error) {
	out := new(internet_quality_rating.InternetQualityRating)
	err := c.cc.Invoke(ctx, "/ntt.watchdog.v1alpha.InternetQualityRatingService/CreateInternetQualityRating", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) UpdateInternetQualityRating(ctx context.Context, in *UpdateInternetQualityRatingRequest, opts ...grpc.CallOption) (*internet_quality_rating.InternetQualityRating, error) {
	out := new(internet_quality_rating.InternetQualityRating)
	err := c.cc.Invoke(ctx, "/ntt.watchdog.v1alpha.InternetQualityRatingService/UpdateInternetQualityRating", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) DeleteInternetQualityRating(ctx context.Context, in *DeleteInternetQualityRatingRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/ntt.watchdog.v1alpha.InternetQualityRatingService/DeleteInternetQualityRating", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
