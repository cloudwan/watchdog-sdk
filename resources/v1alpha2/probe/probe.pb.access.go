// Code generated by protoc-gen-goten-resource
// Resource: Probe
// DO NOT EDIT!!!

package probe

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
	ntt_memo "github.com/cloudwan/edgelq-sdk/common/types/memo"
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	devices_device "github.com/cloudwan/edgelq-sdk/devices/resources/v1alpha2/device"
	iam_service_account "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/service_account"
	common "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/common"
	probe_group "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/probe_group"
	project "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/project"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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
	_ = &ntt_memo.Memo{}
	_ = &ntt_meta.Meta{}
	_ = &devices_device.Device{}
	_ = &iam_service_account.ServiceAccount{}
	_ = &timestamp.Timestamp{}
	_ = &common.SoftwareVersion{}
	_ = &probe_group.ProbeGroup{}
	_ = &project.Project{}
)

type ProbeAccess interface {
	GetProbe(context.Context, *GetQuery) (*Probe, error)
	BatchGetProbes(context.Context, []*Reference, ...gotenresource.BatchGetOption) error
	QueryProbes(context.Context, *ListQuery) (*QueryResultSnapshot, error)
	SearchProbes(context.Context, *SearchQuery) (*QueryResultSnapshot, error)
	WatchProbe(context.Context, *GetQuery, func(*ProbeChange) error) error
	WatchProbes(context.Context, *WatchQuery, func(*QueryResultChange) error) error
	SaveProbe(context.Context, *Probe, ...gotenresource.SaveOption) error
	DeleteProbe(context.Context, *Reference, ...gotenresource.DeleteOption) error
}

type anyCastAccess struct {
	ProbeAccess
}

func AsAnyCastAccess(access ProbeAccess) gotenresource.Access {
	return &anyCastAccess{ProbeAccess: access}
}

func (a *anyCastAccess) Get(ctx context.Context, q gotenresource.GetQuery) (gotenresource.Resource, error) {
	if asProbeQuery, ok := q.(*GetQuery); ok {
		return a.GetProbe(ctx, asProbeQuery)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Probe, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Query(ctx context.Context, q gotenresource.ListQuery) (gotenresource.QueryResultSnapshot, error) {
	if asProbeQuery, ok := q.(*ListQuery); ok {
		return a.QueryProbes(ctx, asProbeQuery)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Probe, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Search(ctx context.Context, q gotenresource.SearchQuery) (gotenresource.QueryResultSnapshot, error) {
	if asProbeQuery, ok := q.(*SearchQuery); ok {
		return a.SearchProbes(ctx, asProbeQuery)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Probe, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Watch(ctx context.Context, q gotenresource.GetQuery, cb func(ch gotenresource.ResourceChange) error) error {
	if asProbeQuery, ok := q.(*GetQuery); ok {
		return a.WatchProbe(ctx, asProbeQuery, func(change *ProbeChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Probe, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) WatchQuery(ctx context.Context, q gotenresource.WatchQuery, cb func(ch gotenresource.QueryResultChange) error) error {
	if asProbeQuery, ok := q.(*WatchQuery); ok {
		return a.WatchProbes(ctx, asProbeQuery, func(change *QueryResultChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Probe, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Save(ctx context.Context, res gotenresource.Resource, opts ...gotenresource.SaveOption) error {
	if asProbeRes, ok := res.(*Probe); ok {
		return a.SaveProbe(ctx, asProbeRes, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Probe, got: %s",
		res.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Delete(ctx context.Context, ref gotenresource.Reference, opts ...gotenresource.DeleteOption) error {
	if asProbeRef, ok := ref.(*Reference); ok {
		return a.DeleteProbe(ctx, asProbeRef, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Probe, got: %s",
		ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) BatchGet(ctx context.Context, toGet []gotenresource.Reference, opts ...gotenresource.BatchGetOption) error {
	probeRefs := make([]*Reference, 0, len(toGet))
	for _, ref := range toGet {
		if asProbeRef, ok := ref.(*Reference); !ok {
			return status.Errorf(codes.Internal,
				"Unrecognized descriptor, expected Probe, got: %s",
				ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
		} else {
			probeRefs = append(probeRefs, asProbeRef)
		}
	}
	return a.BatchGetProbes(ctx, probeRefs, opts...)
}

func (a *anyCastAccess) GetResourceDescriptors() []gotenresource.Descriptor {
	return []gotenresource.Descriptor{
		GetDescriptor(),
	}
}
