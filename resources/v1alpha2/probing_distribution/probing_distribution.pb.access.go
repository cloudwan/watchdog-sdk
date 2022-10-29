// Code generated by protoc-gen-goten-resource
// Resource: ProbingDistribution
// DO NOT EDIT!!!

package probing_distribution

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
	probing_target "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/probing_target"
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
	_ = &probing_target.ProbingTarget{}
	_ = &project.Project{}
)

type ProbingDistributionAccess interface {
	GetProbingDistribution(context.Context, *GetQuery) (*ProbingDistribution, error)
	BatchGetProbingDistributions(context.Context, []*Reference, ...gotenresource.BatchGetOption) error
	QueryProbingDistributions(context.Context, *ListQuery) (*QueryResultSnapshot, error)
	SearchProbingDistributions(context.Context, *SearchQuery) (*SearchQueryResultSnapshot, error)
	WatchProbingDistribution(context.Context, *GetQuery, func(*ProbingDistributionChange) error) error
	WatchProbingDistributions(context.Context, *WatchQuery, func(*QueryResultChange) error) error
	SaveProbingDistribution(context.Context, *ProbingDistribution, ...gotenresource.SaveOption) error
	DeleteProbingDistribution(context.Context, *Reference, ...gotenresource.DeleteOption) error
}

type anyCastAccess struct {
	ProbingDistributionAccess
}

func AsAnyCastAccess(access ProbingDistributionAccess) gotenresource.Access {
	return &anyCastAccess{ProbingDistributionAccess: access}
}

func (a *anyCastAccess) Get(ctx context.Context, q gotenresource.GetQuery) (gotenresource.Resource, error) {
	if asProbingDistributionQuery, ok := q.(*GetQuery); ok {
		return a.GetProbingDistribution(ctx, asProbingDistributionQuery)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected ProbingDistribution, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Query(ctx context.Context, q gotenresource.ListQuery) (gotenresource.QueryResultSnapshot, error) {
	if asProbingDistributionQuery, ok := q.(*ListQuery); ok {
		return a.QueryProbingDistributions(ctx, asProbingDistributionQuery)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected ProbingDistribution, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Search(ctx context.Context, q gotenresource.SearchQuery) (gotenresource.SearchQueryResultSnapshot, error) {
	if asProbingDistributionQuery, ok := q.(*SearchQuery); ok {
		return a.SearchProbingDistributions(ctx, asProbingDistributionQuery)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected ProbingDistribution, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Watch(ctx context.Context, q gotenresource.GetQuery, cb func(ch gotenresource.ResourceChange) error) error {
	if asProbingDistributionQuery, ok := q.(*GetQuery); ok {
		return a.WatchProbingDistribution(ctx, asProbingDistributionQuery, func(change *ProbingDistributionChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected ProbingDistribution, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) WatchQuery(ctx context.Context, q gotenresource.WatchQuery, cb func(ch gotenresource.QueryResultChange) error) error {
	if asProbingDistributionQuery, ok := q.(*WatchQuery); ok {
		return a.WatchProbingDistributions(ctx, asProbingDistributionQuery, func(change *QueryResultChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected ProbingDistribution, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Save(ctx context.Context, res gotenresource.Resource, opts ...gotenresource.SaveOption) error {
	if asProbingDistributionRes, ok := res.(*ProbingDistribution); ok {
		return a.SaveProbingDistribution(ctx, asProbingDistributionRes, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected ProbingDistribution, got: %s",
		res.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Delete(ctx context.Context, ref gotenresource.Reference, opts ...gotenresource.DeleteOption) error {
	if asProbingDistributionRef, ok := ref.(*Reference); ok {
		return a.DeleteProbingDistribution(ctx, asProbingDistributionRef, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected ProbingDistribution, got: %s",
		ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) BatchGet(ctx context.Context, toGet []gotenresource.Reference, opts ...gotenresource.BatchGetOption) error {
	probingDistributionRefs := make([]*Reference, 0, len(toGet))
	for _, ref := range toGet {
		if asProbingDistributionRef, ok := ref.(*Reference); !ok {
			return status.Errorf(codes.Internal,
				"Unrecognized descriptor, expected ProbingDistribution, got: %s",
				ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
		} else {
			probingDistributionRefs = append(probingDistributionRefs, asProbingDistributionRef)
		}
	}
	return a.BatchGetProbingDistributions(ctx, probingDistributionRefs, opts...)
}

func (a *anyCastAccess) GetResourceDescriptors() []gotenresource.Descriptor {
	return []gotenresource.Descriptor{
		GetDescriptor(),
	}
}
