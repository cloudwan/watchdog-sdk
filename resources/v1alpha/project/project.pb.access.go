// Code generated by protoc-gen-goten-resource
// Resource: Project
// DO NOT EDIT!!!

package project

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
)

type ProjectAccess interface {
	GetProject(context.Context, *GetQuery) (*Project, error)
	BatchGetProjects(context.Context, []*Reference, ...gotenresource.BatchGetOption) error
	QueryProjects(context.Context, *ListQuery) (*QueryResultSnapshot, error)
	WatchProject(context.Context, *GetQuery, func(*ProjectChange) error) error
	WatchProjects(context.Context, *WatchQuery, func(*QueryResultChange) error) error
	SaveProject(context.Context, *Project, ...gotenresource.SaveOption) error
	DeleteProject(context.Context, *Reference, ...gotenresource.DeleteOption) error
}

type anyCastAccess struct {
	ProjectAccess
}

func AsAnyCastAccess(access ProjectAccess) gotenresource.Access {
	return &anyCastAccess{ProjectAccess: access}
}

func (a *anyCastAccess) Get(ctx context.Context, q gotenresource.GetQuery) (gotenresource.Resource, error) {
	if asProjectQuery, ok := q.(*GetQuery); ok {
		return a.GetProject(ctx, asProjectQuery)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Project, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Query(ctx context.Context, q gotenresource.ListQuery) (gotenresource.QueryResultSnapshot, error) {
	if asProjectQuery, ok := q.(*ListQuery); ok {
		return a.QueryProjects(ctx, asProjectQuery)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Project, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Search(ctx context.Context, q gotenresource.SearchQuery) (gotenresource.SearchQueryResultSnapshot, error) {
	return nil, status.Errorf(codes.Internal, "Search is not available for Project")
}

func (a *anyCastAccess) Watch(ctx context.Context, q gotenresource.GetQuery, cb func(ch gotenresource.ResourceChange) error) error {
	if asProjectQuery, ok := q.(*GetQuery); ok {
		return a.WatchProject(ctx, asProjectQuery, func(change *ProjectChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Project, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) WatchQuery(ctx context.Context, q gotenresource.WatchQuery, cb func(ch gotenresource.QueryResultChange) error) error {
	if asProjectQuery, ok := q.(*WatchQuery); ok {
		return a.WatchProjects(ctx, asProjectQuery, func(change *QueryResultChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Project, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Save(ctx context.Context, res gotenresource.Resource, opts ...gotenresource.SaveOption) error {
	if asProjectRes, ok := res.(*Project); ok {
		return a.SaveProject(ctx, asProjectRes, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Project, got: %s",
		res.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Delete(ctx context.Context, ref gotenresource.Reference, opts ...gotenresource.DeleteOption) error {
	if asProjectRef, ok := ref.(*Reference); ok {
		return a.DeleteProject(ctx, asProjectRef, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Project, got: %s",
		ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) BatchGet(ctx context.Context, toGet []gotenresource.Reference, opts ...gotenresource.BatchGetOption) error {
	projectRefs := make([]*Reference, 0, len(toGet))
	for _, ref := range toGet {
		if asProjectRef, ok := ref.(*Reference); !ok {
			return status.Errorf(codes.Internal,
				"Unrecognized descriptor, expected Project, got: %s",
				ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
		} else {
			projectRefs = append(projectRefs, asProjectRef)
		}
	}
	return a.BatchGetProjects(ctx, projectRefs, opts...)
}

func (a *anyCastAccess) GetResourceDescriptors() []gotenresource.Descriptor {
	return []gotenresource.Descriptor{
		GetDescriptor(),
	}
}
