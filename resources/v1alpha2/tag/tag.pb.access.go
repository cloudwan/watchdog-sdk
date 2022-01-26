// Code generated by protoc-gen-goten-resource
// Resource: Tag
// DO NOT EDIT!!!

package tag

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
	probing_distribution "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/probing_distribution"
	project "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/project"
)

// ensure the imports are used
var (
	_ = context.Context(nil)

	_ = codes.Internal
	_ = status.Status{}

	_ = watch_type.WatchType_STATEFUL
	_ = gotenobject.FieldPath(nil)
	_ = gotenresource.ListQuery(nil)
)

// make sure we're using proto imports
var (
	_ = &ntt_meta.Meta{}
	_ = &probing_distribution.ProbingDistribution{}
	_ = &project.Project{}
)

type TagAccess interface {
	GetTag(context.Context, *GetQuery) (*Tag, error)
	BatchGetTags(context.Context, []*Reference, ...gotenresource.BatchGetOption) error
	QueryTags(context.Context, *ListQuery) (*QueryResultSnapshot, error)
	WatchTag(context.Context, *GetQuery, func(*TagChange) error) error
	WatchTags(context.Context, *WatchQuery, func(*QueryResultChange) error) error
	SaveTag(context.Context, *Tag, ...gotenresource.SaveOption) error
	DeleteTag(context.Context, *Reference, ...gotenresource.DeleteOption) error
}

type anyCastAccess struct {
	TagAccess
}

func AsAnyCastAccess(access TagAccess) gotenresource.Access {
	return &anyCastAccess{TagAccess: access}
}

func (a *anyCastAccess) Get(ctx context.Context, q gotenresource.GetQuery) (gotenresource.Resource, error) {
	if asTagQuery, ok := q.(*GetQuery); ok {
		return a.GetTag(ctx, asTagQuery)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Tag, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Query(ctx context.Context, q gotenresource.ListQuery) (gotenresource.QueryResultSnapshot, error) {
	if asTagQuery, ok := q.(*ListQuery); ok {
		return a.QueryTags(ctx, asTagQuery)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Tag, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Search(ctx context.Context, q gotenresource.SearchQuery) (gotenresource.SearchQueryResultSnapshot, error) {
	return nil, status.Errorf(codes.Internal, "Search is not available for Tag")
}

func (a *anyCastAccess) Watch(ctx context.Context, q gotenresource.GetQuery, cb func(ch gotenresource.ResourceChange) error) error {
	if asTagQuery, ok := q.(*GetQuery); ok {
		return a.WatchTag(ctx, asTagQuery, func(change *TagChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Tag, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) WatchQuery(ctx context.Context, q gotenresource.WatchQuery, cb func(ch gotenresource.QueryResultChange) error) error {
	if asTagQuery, ok := q.(*WatchQuery); ok {
		return a.WatchTags(ctx, asTagQuery, func(change *QueryResultChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Tag, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Save(ctx context.Context, res gotenresource.Resource, opts ...gotenresource.SaveOption) error {
	if asTagRes, ok := res.(*Tag); ok {
		return a.SaveTag(ctx, asTagRes, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Tag, got: %s",
		res.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Delete(ctx context.Context, ref gotenresource.Reference, opts ...gotenresource.DeleteOption) error {
	if asTagRef, ok := ref.(*Reference); ok {
		return a.DeleteTag(ctx, asTagRef, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Tag, got: %s",
		ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) BatchGet(ctx context.Context, toGet []gotenresource.Reference, opts ...gotenresource.BatchGetOption) error {
	tagRefs := make([]*Reference, 0, len(toGet))
	for _, ref := range toGet {
		if asTagRef, ok := ref.(*Reference); !ok {
			return status.Errorf(codes.Internal,
				"Unrecognized descriptor, expected Tag, got: %s",
				ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
		} else {
			tagRefs = append(tagRefs, asTagRef)
		}
	}
	return a.BatchGetTags(ctx, tagRefs, opts...)
}

func (a *anyCastAccess) GetResourceDescriptors() []gotenresource.Descriptor {
	return []gotenresource.Descriptor{
		GetDescriptor(),
	}
}
