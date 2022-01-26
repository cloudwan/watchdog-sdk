// Code generated by protoc-gen-goten-resource
// Resource: ProbingTarget
// DO NOT EDIT!!!

package probing_target

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
	common "github.com/cloudwan/watchdog-sdk/resources/v1alpha/common"
	probe "github.com/cloudwan/watchdog-sdk/resources/v1alpha/probe"
	probing_target_group "github.com/cloudwan/watchdog-sdk/resources/v1alpha/probing_target_group"
	project "github.com/cloudwan/watchdog-sdk/resources/v1alpha/project"
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
	_ = &common.SoftwareVersion{}
	_ = &probe.Probe{}
	_ = &probing_target_group.ProbingTargetGroup{}
	_ = &project.Project{}
)

type ProbingTargetAccess interface {
	GetProbingTarget(context.Context, *GetQuery) (*ProbingTarget, error)
	BatchGetProbingTargets(context.Context, []*Reference, ...gotenresource.BatchGetOption) error
	QueryProbingTargets(context.Context, *ListQuery) (*QueryResultSnapshot, error)
	WatchProbingTarget(context.Context, *GetQuery, func(*ProbingTargetChange) error) error
	WatchProbingTargets(context.Context, *WatchQuery, func(*QueryResultChange) error) error
	SaveProbingTarget(context.Context, *ProbingTarget, ...gotenresource.SaveOption) error
	DeleteProbingTarget(context.Context, *Reference, ...gotenresource.DeleteOption) error
}

type anyCastAccess struct {
	ProbingTargetAccess
}

func AsAnyCastAccess(access ProbingTargetAccess) gotenresource.Access {
	return &anyCastAccess{ProbingTargetAccess: access}
}

func (a *anyCastAccess) Get(ctx context.Context, q gotenresource.GetQuery) (gotenresource.Resource, error) {
	if asProbingTargetQuery, ok := q.(*GetQuery); ok {
		return a.GetProbingTarget(ctx, asProbingTargetQuery)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected ProbingTarget, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Query(ctx context.Context, q gotenresource.ListQuery) (gotenresource.QueryResultSnapshot, error) {
	if asProbingTargetQuery, ok := q.(*ListQuery); ok {
		return a.QueryProbingTargets(ctx, asProbingTargetQuery)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected ProbingTarget, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Search(ctx context.Context, q gotenresource.SearchQuery) (gotenresource.SearchQueryResultSnapshot, error) {
	return nil, status.Errorf(codes.Internal, "Search is not available for ProbingTarget")
}

func (a *anyCastAccess) Watch(ctx context.Context, q gotenresource.GetQuery, cb func(ch gotenresource.ResourceChange) error) error {
	if asProbingTargetQuery, ok := q.(*GetQuery); ok {
		return a.WatchProbingTarget(ctx, asProbingTargetQuery, func(change *ProbingTargetChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected ProbingTarget, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) WatchQuery(ctx context.Context, q gotenresource.WatchQuery, cb func(ch gotenresource.QueryResultChange) error) error {
	if asProbingTargetQuery, ok := q.(*WatchQuery); ok {
		return a.WatchProbingTargets(ctx, asProbingTargetQuery, func(change *QueryResultChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected ProbingTarget, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Save(ctx context.Context, res gotenresource.Resource, opts ...gotenresource.SaveOption) error {
	if asProbingTargetRes, ok := res.(*ProbingTarget); ok {
		return a.SaveProbingTarget(ctx, asProbingTargetRes, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected ProbingTarget, got: %s",
		res.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Delete(ctx context.Context, ref gotenresource.Reference, opts ...gotenresource.DeleteOption) error {
	if asProbingTargetRef, ok := ref.(*Reference); ok {
		return a.DeleteProbingTarget(ctx, asProbingTargetRef, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected ProbingTarget, got: %s",
		ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) BatchGet(ctx context.Context, toGet []gotenresource.Reference, opts ...gotenresource.BatchGetOption) error {
	probingTargetRefs := make([]*Reference, 0, len(toGet))
	for _, ref := range toGet {
		if asProbingTargetRef, ok := ref.(*Reference); !ok {
			return status.Errorf(codes.Internal,
				"Unrecognized descriptor, expected ProbingTarget, got: %s",
				ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
		} else {
			probingTargetRefs = append(probingTargetRefs, asProbingTargetRef)
		}
	}
	return a.BatchGetProbingTargets(ctx, probingTargetRefs, opts...)
}

func (a *anyCastAccess) GetResourceDescriptors() []gotenresource.Descriptor {
	return []gotenresource.Descriptor{
		GetDescriptor(),
	}
}
