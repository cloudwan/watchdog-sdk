// Code generated by protoc-gen-goten-access
// Resource: AdminArea
// DO NOT EDIT!!!

package admin_area_access

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cloudwan/goten-sdk/runtime/api/watch_type"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"

	admin_area_client "github.com/cloudwan/watchdog-sdk/client/v1alpha2/admin_area"
	admin_area "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/admin_area"
)

var (
	_ = context.Context(nil)
	_ = fmt.GoStringer(nil)

	_ = codes.NotFound
	_ = status.Status{}

	_ = watch_type.WatchType_STATEFUL
	_ = gotenresource.ListQuery(nil)
)

type apiAdminAreaAccess struct {
	client admin_area_client.AdminAreaServiceClient
}

func NewApiAdminAreaAccess(client admin_area_client.AdminAreaServiceClient) admin_area.AdminAreaAccess {
	return &apiAdminAreaAccess{client: client}
}

func (a *apiAdminAreaAccess) GetAdminArea(ctx context.Context, query *admin_area.GetQuery) (*admin_area.AdminArea, error) {
	request := &admin_area_client.GetAdminAreaRequest{
		Name:      query.Reference,
		FieldMask: query.Mask,
	}
	return a.client.GetAdminArea(ctx, request)
}

func (a *apiAdminAreaAccess) BatchGetAdminAreas(ctx context.Context, refs []*admin_area.Reference, opts ...gotenresource.BatchGetOption) error {
	batchGetOpts := gotenresource.MakeBatchGetOptions(opts)
	request := &admin_area_client.BatchGetAdminAreasRequest{
		Names: refs,
	}
	resp, err := a.client.BatchGetAdminAreas(ctx, request)
	if err != nil {
		return err
	}
	resultMap := make(map[admin_area.Name]*admin_area.AdminArea, len(refs))
	for _, resolvedRes := range resp.GetAdminAreas() {
		resultMap[*resolvedRes.GetName()] = resolvedRes
	}
	for _, ref := range refs {
		resolvedRes := resultMap[ref.Name]
		if resolvedRes != nil {
			ref.Resolve(resolvedRes)
		}
	}
	if batchGetOpts.MustResolveAll() && len(resp.GetMissing()) > 0 {
		return status.Errorf(codes.NotFound, "Number of references not found: %d", len(resp.GetMissing()))
	}
	return nil
}

func (a *apiAdminAreaAccess) QueryAdminAreas(ctx context.Context, query *admin_area.ListQuery) (*admin_area.QueryResultSnapshot, error) {
	request := &admin_area_client.ListAdminAreasRequest{
		Filter:    query.Filter,
		FieldMask: query.Mask,
	}
	if query.Pager != nil {
		request.PageSize = int32(query.Pager.Limit)
		request.OrderBy = query.Pager.OrderBy
		request.PageToken = query.Pager.Cursor
	}
	resp, err := a.client.ListAdminAreas(ctx, request)
	if err != nil {
		return nil, err
	}
	return &admin_area.QueryResultSnapshot{
		AdminAreas:     resp.AdminAreas,
		NextPageCursor: resp.NextPageToken,
		PrevPageCursor: resp.PrevPageToken,
	}, nil
}

func (a *apiAdminAreaAccess) WatchAdminArea(ctx context.Context, query *admin_area.GetQuery, observerCb func(*admin_area.AdminAreaChange) error) error {
	request := &admin_area_client.WatchAdminAreaRequest{
		Name:      query.Reference,
		FieldMask: query.Mask,
	}
	changesStream, initErr := a.client.WatchAdminArea(ctx, request)
	if initErr != nil {
		return initErr
	}
	for {
		resp, err := changesStream.Recv()
		if err != nil {
			return fmt.Errorf("watch recv error: %w", err)
		}
		change := resp.GetChange()
		if err := observerCb(change); err != nil {
			return err
		}
	}
}

func (a *apiAdminAreaAccess) WatchAdminAreas(ctx context.Context, query *admin_area.WatchQuery, observerCb func(*admin_area.QueryResultChange) error) error {
	request := &admin_area_client.WatchAdminAreasRequest{
		Filter:       query.Filter,
		FieldMask:    query.Mask,
		MaxChunkSize: int32(query.ChunkSize),
		Type:         query.WatchType,
		ResumeToken:  query.ResumeToken,
	}
	if query.Pager != nil {
		request.OrderBy = query.Pager.OrderBy
		request.PageSize = int32(query.Pager.Limit)
		request.PageToken = query.Pager.Cursor
	}
	changesStream, initErr := a.client.WatchAdminAreas(ctx, request)
	if initErr != nil {
		return initErr
	}
	for {
		respChange, err := changesStream.Recv()
		if err != nil {
			return fmt.Errorf("watch recv error: %w", err)
		}
		changesWithPaging := &admin_area.QueryResultChange{
			Changes:      respChange.AdminAreaChanges,
			IsCurrent:    respChange.IsCurrent,
			IsHardReset:  respChange.IsHardReset,
			IsSoftReset:  respChange.IsSoftReset,
			ResumeToken:  respChange.ResumeToken,
			SnapshotSize: respChange.SnapshotSize,
		}
		if respChange.PageTokenChange != nil {
			changesWithPaging.PrevPageCursor = respChange.PageTokenChange.PrevPageToken
			changesWithPaging.NextPageCursor = respChange.PageTokenChange.NextPageToken
		}
		if err := observerCb(changesWithPaging); err != nil {
			return err
		}
	}
}

func (a *apiAdminAreaAccess) SaveAdminArea(ctx context.Context, res *admin_area.AdminArea, opts ...gotenresource.SaveOption) error {
	saveOpts := gotenresource.MakeSaveOptions(opts)
	previousRes := saveOpts.GetPreviousResource()

	if previousRes == nil {
		var err error
		previousRes, err = a.GetAdminArea(ctx, &admin_area.GetQuery{Reference: res.Name.AsReference()})
		if err != nil {
			if statusErr, ok := status.FromError(err); !ok || statusErr.Code() != codes.NotFound {
				return err
			}
		}
	}

	if previousRes != nil {
		updateRequest := &admin_area_client.UpdateAdminAreaRequest{
			AdminArea: res,
		}
		_, err := a.client.UpdateAdminArea(ctx, updateRequest)
		if err != nil {
			return err
		}
		return nil
	} else {
		createRequest := &admin_area_client.CreateAdminAreaRequest{
			AdminArea: res,
		}
		_, err := a.client.CreateAdminArea(ctx, createRequest)
		if err != nil {
			return err
		}
		return nil
	}
}

func (a *apiAdminAreaAccess) DeleteAdminArea(ctx context.Context, ref *admin_area.Reference, opts ...gotenresource.DeleteOption) error {
	request := &admin_area_client.DeleteAdminAreaRequest{
		Name: ref,
	}
	_, err := a.client.DeleteAdminArea(ctx, request)
	return err
}
