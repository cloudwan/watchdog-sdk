// Code generated by protoc-gen-goten-resource
// Resource: SharedToken
// DO NOT EDIT!!!

package shared_token

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cloudwan/goten-sdk/runtime/api/watch_type"
	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	common "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/common"
	probe "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/probe"
	probe_group "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/probe_group"
	project "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/project"
)

// ensure the imports are used
var (
	_ = new(context.Context)

	_ = codes.Internal
	_ = status.Status{}

	_ = watch_type.WatchType_STATEFUL
	_ = new(gotenobject.FieldPath)
	_ = new(gotenresource.ListQuery)
)

// make sure we're using proto imports
var (
	_ = &ntt_meta.Meta{}
	_ = &common.SoftwareVersion{}
	_ = &probe.Probe{}
	_ = &probe_group.ProbeGroup{}
	_ = &project.Project{}
)

type SharedTokenAccess interface {
	GetSharedToken(context.Context, *GetQuery) (*SharedToken, error)
	BatchGetSharedTokens(context.Context, []*Reference, ...gotenresource.BatchGetOption) error
	QuerySharedTokens(context.Context, *ListQuery) (*QueryResultSnapshot, error)
	WatchSharedToken(context.Context, *GetQuery, func(*SharedTokenChange) error) error
	WatchSharedTokens(context.Context, *WatchQuery, func(*QueryResultChange) error) error
	SaveSharedToken(context.Context, *SharedToken, ...gotenresource.SaveOption) error
	DeleteSharedToken(context.Context, *Reference, ...gotenresource.DeleteOption) error
}

type anyCastAccess struct {
	SharedTokenAccess
}

func AsAnyCastAccess(access SharedTokenAccess) gotenresource.Access {
	return &anyCastAccess{SharedTokenAccess: access}
}

func (a *anyCastAccess) Get(ctx context.Context, q gotenresource.GetQuery) (gotenresource.Resource, error) {
	if asSharedTokenQuery, ok := q.(*GetQuery); ok {
		return a.GetSharedToken(ctx, asSharedTokenQuery)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected SharedToken, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Query(ctx context.Context, q gotenresource.ListQuery) (gotenresource.QueryResultSnapshot, error) {
	if asSharedTokenQuery, ok := q.(*ListQuery); ok {
		return a.QuerySharedTokens(ctx, asSharedTokenQuery)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected SharedToken, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Search(ctx context.Context, q gotenresource.SearchQuery) (gotenresource.QueryResultSnapshot, error) {
	return nil, status.Errorf(codes.Internal, "Search is not available for SharedToken")
}

func (a *anyCastAccess) Watch(ctx context.Context, q gotenresource.GetQuery, cb func(ch gotenresource.ResourceChange) error) error {
	if asSharedTokenQuery, ok := q.(*GetQuery); ok {
		return a.WatchSharedToken(ctx, asSharedTokenQuery, func(change *SharedTokenChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected SharedToken, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) WatchQuery(ctx context.Context, q gotenresource.WatchQuery, cb func(ch gotenresource.QueryResultChange) error) error {
	if asSharedTokenQuery, ok := q.(*WatchQuery); ok {
		return a.WatchSharedTokens(ctx, asSharedTokenQuery, func(change *QueryResultChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected SharedToken, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Save(ctx context.Context, res gotenresource.Resource, opts ...gotenresource.SaveOption) error {
	if asSharedTokenRes, ok := res.(*SharedToken); ok {
		return a.SaveSharedToken(ctx, asSharedTokenRes, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected SharedToken, got: %s",
		res.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Delete(ctx context.Context, ref gotenresource.Reference, opts ...gotenresource.DeleteOption) error {
	if asSharedTokenRef, ok := ref.(*Reference); ok {
		return a.DeleteSharedToken(ctx, asSharedTokenRef, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected SharedToken, got: %s",
		ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) BatchGet(ctx context.Context, toGet []gotenresource.Reference, opts ...gotenresource.BatchGetOption) error {
	sharedTokenRefs := make([]*Reference, 0, len(toGet))
	for _, ref := range toGet {
		if asSharedTokenRef, ok := ref.(*Reference); !ok {
			return status.Errorf(codes.Internal,
				"Unrecognized descriptor, expected SharedToken, got: %s",
				ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
		} else {
			sharedTokenRefs = append(sharedTokenRefs, asSharedTokenRef)
		}
	}
	return a.BatchGetSharedTokens(ctx, sharedTokenRefs, opts...)
}

func (a *anyCastAccess) GetResourceDescriptors() []gotenresource.Descriptor {
	return []gotenresource.Descriptor{
		GetDescriptor(),
	}
}
