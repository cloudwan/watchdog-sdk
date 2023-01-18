// Code generated by protoc-gen-goten-access
// Resource: SharedToken
// DO NOT EDIT!!!

package shared_token_access

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	gotenaccess "github.com/cloudwan/goten-sdk/runtime/access"
	"github.com/cloudwan/goten-sdk/runtime/api/watch_type"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"

	shared_token_client "github.com/cloudwan/watchdog-sdk/client/v1alpha2/shared_token"
	shared_token "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/shared_token"
)

var (
	_ = new(context.Context)
	_ = new(fmt.GoStringer)

	_ = new(grpc.ClientConnInterface)
	_ = codes.NotFound
	_ = status.Status{}

	_ = new(gotenaccess.Watcher)
	_ = watch_type.WatchType_STATEFUL
	_ = new(gotenresource.ListQuery)
)

type apiSharedTokenAccess struct {
	client shared_token_client.SharedTokenServiceClient
}

func NewApiSharedTokenAccess(client shared_token_client.SharedTokenServiceClient) shared_token.SharedTokenAccess {
	return &apiSharedTokenAccess{client: client}
}

func (a *apiSharedTokenAccess) GetSharedToken(ctx context.Context, query *shared_token.GetQuery) (*shared_token.SharedToken, error) {
	request := &shared_token_client.GetSharedTokenRequest{
		Name:      query.Reference,
		FieldMask: query.Mask,
	}
	res, err := a.client.GetSharedToken(ctx, request)
	if err != nil {
		return nil, err
	}
	query.Reference.Resolve(res)
	return res, nil
}

func (a *apiSharedTokenAccess) BatchGetSharedTokens(ctx context.Context, refs []*shared_token.Reference, opts ...gotenresource.BatchGetOption) error {
	batchGetOpts := gotenresource.MakeBatchGetOptions(opts)
	request := &shared_token_client.BatchGetSharedTokensRequest{
		Names: refs,
	}
	resp, err := a.client.BatchGetSharedTokens(ctx, request)
	if err != nil {
		return err
	}
	resultMap := make(map[shared_token.Name]*shared_token.SharedToken, len(refs))
	for _, resolvedRes := range resp.GetSharedTokens() {
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

func (a *apiSharedTokenAccess) QuerySharedTokens(ctx context.Context, query *shared_token.ListQuery) (*shared_token.QueryResultSnapshot, error) {
	request := &shared_token_client.ListSharedTokensRequest{
		Filter:            query.Filter,
		FieldMask:         query.Mask,
		IncludePagingInfo: query.WithPagingInfo,
	}
	if query.Pager != nil {
		request.PageSize = int32(query.Pager.Limit)
		request.OrderBy = query.Pager.OrderBy
		request.PageToken = query.Pager.Cursor
	}
	resp, err := a.client.ListSharedTokens(ctx, request)
	if err != nil {
		return nil, err
	}
	return &shared_token.QueryResultSnapshot{
		SharedTokens:      resp.SharedTokens,
		NextPageCursor:    resp.NextPageToken,
		PrevPageCursor:    resp.PrevPageToken,
		TotalResultsCount: resp.TotalResultsCount,
		CurrentOffset:     resp.CurrentOffset,
	}, nil
}

func (a *apiSharedTokenAccess) WatchSharedToken(ctx context.Context, query *shared_token.GetQuery, observerCb func(*shared_token.SharedTokenChange) error) error {
	request := &shared_token_client.WatchSharedTokenRequest{
		Name:      query.Reference,
		FieldMask: query.Mask,
	}
	changesStream, initErr := a.client.WatchSharedToken(ctx, request)
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

func (a *apiSharedTokenAccess) WatchSharedTokens(ctx context.Context, query *shared_token.WatchQuery, observerCb func(*shared_token.QueryResultChange) error) error {
	request := &shared_token_client.WatchSharedTokensRequest{
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
	changesStream, initErr := a.client.WatchSharedTokens(ctx, request)
	if initErr != nil {
		return initErr
	}
	for {
		respChange, err := changesStream.Recv()
		if err != nil {
			return fmt.Errorf("watch recv error: %w", err)
		}
		changesWithPaging := &shared_token.QueryResultChange{
			Changes:      respChange.SharedTokenChanges,
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

func (a *apiSharedTokenAccess) SaveSharedToken(ctx context.Context, res *shared_token.SharedToken, opts ...gotenresource.SaveOption) error {
	saveOpts := gotenresource.MakeSaveOptions(opts)
	previousRes := saveOpts.GetPreviousResource()

	if previousRes == nil && !saveOpts.OnlyUpdate() && !saveOpts.OnlyCreate() {
		var err error
		previousRes, err = a.GetSharedToken(ctx, &shared_token.GetQuery{Reference: res.Name.AsReference()})
		if err != nil {
			if statusErr, ok := status.FromError(err); !ok || statusErr.Code() != codes.NotFound {
				return err
			}
		}
	}

	if saveOpts.OnlyUpdate() || previousRes != nil {
		updateRequest := &shared_token_client.UpdateSharedTokenRequest{
			SharedToken: res,
		}
		if updateMask := saveOpts.GetUpdateMask(); updateMask != nil {
			updateRequest.UpdateMask = updateMask.(*shared_token.SharedToken_FieldMask)
		}
		if mask, conditionalState := saveOpts.GetCAS(); mask != nil && conditionalState != nil {
			updateRequest.Cas = &shared_token_client.UpdateSharedTokenRequest_CAS{
				ConditionalState: conditionalState.(*shared_token.SharedToken),
				FieldMask:        mask.(*shared_token.SharedToken_FieldMask),
			}
		}
		_, err := a.client.UpdateSharedToken(ctx, updateRequest)
		if err != nil {
			return err
		}
		return nil
	} else {
		createRequest := &shared_token_client.CreateSharedTokenRequest{
			SharedToken: res,
		}
		_, err := a.client.CreateSharedToken(ctx, createRequest)
		if err != nil {
			return err
		}
		return nil
	}
}

func (a *apiSharedTokenAccess) DeleteSharedToken(ctx context.Context, ref *shared_token.Reference, opts ...gotenresource.DeleteOption) error {
	request := &shared_token_client.DeleteSharedTokenRequest{
		Name: ref,
	}
	_, err := a.client.DeleteSharedToken(ctx, request)
	return err
}

func init() {
	gotenaccess.GetRegistry().RegisterApiAccessConstructor(shared_token.GetDescriptor(), func(cc grpc.ClientConnInterface) gotenresource.Access {
		return shared_token.AsAnyCastAccess(NewApiSharedTokenAccess(shared_token_client.NewSharedTokenServiceClient(cc)))
	})
}
