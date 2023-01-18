// Code generated by protoc-gen-goten-resource
// Resource: ProbingConfig
// DO NOT EDIT!!!

package probing_config

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
	probing_target_group "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/probing_target_group"
	project "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/project"
	duration "github.com/golang/protobuf/ptypes/duration"
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
	_ = &duration.Duration{}
	_ = &common.SoftwareVersion{}
	_ = &probe.Probe{}
	_ = &probing_target.ProbingTarget{}
	_ = &probing_target_group.ProbingTargetGroup{}
	_ = &project.Project{}
)

type ProbingConfigAccess interface {
	GetProbingConfig(context.Context, *GetQuery) (*ProbingConfig, error)
	BatchGetProbingConfigs(context.Context, []*Reference, ...gotenresource.BatchGetOption) error
	QueryProbingConfigs(context.Context, *ListQuery) (*QueryResultSnapshot, error)
	WatchProbingConfig(context.Context, *GetQuery, func(*ProbingConfigChange) error) error
	WatchProbingConfigs(context.Context, *WatchQuery, func(*QueryResultChange) error) error
	SaveProbingConfig(context.Context, *ProbingConfig, ...gotenresource.SaveOption) error
	DeleteProbingConfig(context.Context, *Reference, ...gotenresource.DeleteOption) error
}

type anyCastAccess struct {
	ProbingConfigAccess
}

func AsAnyCastAccess(access ProbingConfigAccess) gotenresource.Access {
	return &anyCastAccess{ProbingConfigAccess: access}
}

func (a *anyCastAccess) Get(ctx context.Context, q gotenresource.GetQuery) (gotenresource.Resource, error) {
	if asProbingConfigQuery, ok := q.(*GetQuery); ok {
		return a.GetProbingConfig(ctx, asProbingConfigQuery)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected ProbingConfig, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Query(ctx context.Context, q gotenresource.ListQuery) (gotenresource.QueryResultSnapshot, error) {
	if asProbingConfigQuery, ok := q.(*ListQuery); ok {
		return a.QueryProbingConfigs(ctx, asProbingConfigQuery)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected ProbingConfig, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Search(ctx context.Context, q gotenresource.SearchQuery) (gotenresource.QueryResultSnapshot, error) {
	return nil, status.Errorf(codes.Internal, "Search is not available for ProbingConfig")
}

func (a *anyCastAccess) Watch(ctx context.Context, q gotenresource.GetQuery, cb func(ch gotenresource.ResourceChange) error) error {
	if asProbingConfigQuery, ok := q.(*GetQuery); ok {
		return a.WatchProbingConfig(ctx, asProbingConfigQuery, func(change *ProbingConfigChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected ProbingConfig, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) WatchQuery(ctx context.Context, q gotenresource.WatchQuery, cb func(ch gotenresource.QueryResultChange) error) error {
	if asProbingConfigQuery, ok := q.(*WatchQuery); ok {
		return a.WatchProbingConfigs(ctx, asProbingConfigQuery, func(change *QueryResultChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected ProbingConfig, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Save(ctx context.Context, res gotenresource.Resource, opts ...gotenresource.SaveOption) error {
	if asProbingConfigRes, ok := res.(*ProbingConfig); ok {
		return a.SaveProbingConfig(ctx, asProbingConfigRes, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected ProbingConfig, got: %s",
		res.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Delete(ctx context.Context, ref gotenresource.Reference, opts ...gotenresource.DeleteOption) error {
	if asProbingConfigRef, ok := ref.(*Reference); ok {
		return a.DeleteProbingConfig(ctx, asProbingConfigRef, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected ProbingConfig, got: %s",
		ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) BatchGet(ctx context.Context, toGet []gotenresource.Reference, opts ...gotenresource.BatchGetOption) error {
	probingConfigRefs := make([]*Reference, 0, len(toGet))
	for _, ref := range toGet {
		if asProbingConfigRef, ok := ref.(*Reference); !ok {
			return status.Errorf(codes.Internal,
				"Unrecognized descriptor, expected ProbingConfig, got: %s",
				ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
		} else {
			probingConfigRefs = append(probingConfigRefs, asProbingConfigRef)
		}
	}
	return a.BatchGetProbingConfigs(ctx, probingConfigRefs, opts...)
}

func (a *anyCastAccess) GetResourceDescriptors() []gotenresource.Descriptor {
	return []gotenresource.Descriptor{
		GetDescriptor(),
	}
}
