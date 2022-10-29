// Code generated by protoc-gen-goten-resource
// Resource: ProbingTargetGroup
// DO NOT EDIT!!!

package probing_target_group

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
	_ = &project.Project{}
)

type ProbingTargetGroupAccess interface {
	GetProbingTargetGroup(context.Context, *GetQuery) (*ProbingTargetGroup, error)
	BatchGetProbingTargetGroups(context.Context, []*Reference, ...gotenresource.BatchGetOption) error
	QueryProbingTargetGroups(context.Context, *ListQuery) (*QueryResultSnapshot, error)
	SearchProbingTargetGroups(context.Context, *SearchQuery) (*SearchQueryResultSnapshot, error)
	WatchProbingTargetGroup(context.Context, *GetQuery, func(*ProbingTargetGroupChange) error) error
	WatchProbingTargetGroups(context.Context, *WatchQuery, func(*QueryResultChange) error) error
	SaveProbingTargetGroup(context.Context, *ProbingTargetGroup, ...gotenresource.SaveOption) error
	DeleteProbingTargetGroup(context.Context, *Reference, ...gotenresource.DeleteOption) error
}

type anyCastAccess struct {
	ProbingTargetGroupAccess
}

func AsAnyCastAccess(access ProbingTargetGroupAccess) gotenresource.Access {
	return &anyCastAccess{ProbingTargetGroupAccess: access}
}

func (a *anyCastAccess) Get(ctx context.Context, q gotenresource.GetQuery) (gotenresource.Resource, error) {
	if asProbingTargetGroupQuery, ok := q.(*GetQuery); ok {
		return a.GetProbingTargetGroup(ctx, asProbingTargetGroupQuery)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected ProbingTargetGroup, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Query(ctx context.Context, q gotenresource.ListQuery) (gotenresource.QueryResultSnapshot, error) {
	if asProbingTargetGroupQuery, ok := q.(*ListQuery); ok {
		return a.QueryProbingTargetGroups(ctx, asProbingTargetGroupQuery)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected ProbingTargetGroup, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Search(ctx context.Context, q gotenresource.SearchQuery) (gotenresource.SearchQueryResultSnapshot, error) {
	if asProbingTargetGroupQuery, ok := q.(*SearchQuery); ok {
		return a.SearchProbingTargetGroups(ctx, asProbingTargetGroupQuery)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected ProbingTargetGroup, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Watch(ctx context.Context, q gotenresource.GetQuery, cb func(ch gotenresource.ResourceChange) error) error {
	if asProbingTargetGroupQuery, ok := q.(*GetQuery); ok {
		return a.WatchProbingTargetGroup(ctx, asProbingTargetGroupQuery, func(change *ProbingTargetGroupChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected ProbingTargetGroup, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) WatchQuery(ctx context.Context, q gotenresource.WatchQuery, cb func(ch gotenresource.QueryResultChange) error) error {
	if asProbingTargetGroupQuery, ok := q.(*WatchQuery); ok {
		return a.WatchProbingTargetGroups(ctx, asProbingTargetGroupQuery, func(change *QueryResultChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected ProbingTargetGroup, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Save(ctx context.Context, res gotenresource.Resource, opts ...gotenresource.SaveOption) error {
	if asProbingTargetGroupRes, ok := res.(*ProbingTargetGroup); ok {
		return a.SaveProbingTargetGroup(ctx, asProbingTargetGroupRes, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected ProbingTargetGroup, got: %s",
		res.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Delete(ctx context.Context, ref gotenresource.Reference, opts ...gotenresource.DeleteOption) error {
	if asProbingTargetGroupRef, ok := ref.(*Reference); ok {
		return a.DeleteProbingTargetGroup(ctx, asProbingTargetGroupRef, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected ProbingTargetGroup, got: %s",
		ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) BatchGet(ctx context.Context, toGet []gotenresource.Reference, opts ...gotenresource.BatchGetOption) error {
	probingTargetGroupRefs := make([]*Reference, 0, len(toGet))
	for _, ref := range toGet {
		if asProbingTargetGroupRef, ok := ref.(*Reference); !ok {
			return status.Errorf(codes.Internal,
				"Unrecognized descriptor, expected ProbingTargetGroup, got: %s",
				ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
		} else {
			probingTargetGroupRefs = append(probingTargetGroupRefs, asProbingTargetGroupRef)
		}
	}
	return a.BatchGetProbingTargetGroups(ctx, probingTargetGroupRefs, opts...)
}

func (a *anyCastAccess) GetResourceDescriptors() []gotenresource.Descriptor {
	return []gotenresource.Descriptor{
		GetDescriptor(),
	}
}
