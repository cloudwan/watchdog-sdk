// Code generated by protoc-gen-goten-access
// Resource: ProbingSession
// DO NOT EDIT!!!

package probing_session_access

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	gotenaccess "github.com/cloudwan/goten-sdk/runtime/access"
	"github.com/cloudwan/goten-sdk/runtime/api/watch_type"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"

	probing_session_client "github.com/cloudwan/watchdog-sdk/client/v1alpha2/probing_session"
	probing_session "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/probing_session"
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

type apiProbingSessionAccess struct {
	client probing_session_client.ProbingSessionServiceClient
}

func NewApiProbingSessionAccess(client probing_session_client.ProbingSessionServiceClient) probing_session.ProbingSessionAccess {
	return &apiProbingSessionAccess{client: client}
}

func (a *apiProbingSessionAccess) GetProbingSession(ctx context.Context, query *probing_session.GetQuery) (*probing_session.ProbingSession, error) {
	request := &probing_session_client.GetProbingSessionRequest{
		Name:      query.Reference,
		FieldMask: query.Mask,
	}
	res, err := a.client.GetProbingSession(ctx, request)
	if err != nil {
		return nil, err
	}
	query.Reference.Resolve(res)
	return res, nil
}

func (a *apiProbingSessionAccess) BatchGetProbingSessions(ctx context.Context, refs []*probing_session.Reference, opts ...gotenresource.BatchGetOption) error {
	batchGetOpts := gotenresource.MakeBatchGetOptions(opts)
	request := &probing_session_client.BatchGetProbingSessionsRequest{
		Names: refs,
	}
	resp, err := a.client.BatchGetProbingSessions(ctx, request)
	if err != nil {
		return err
	}
	resultMap := make(map[probing_session.Name]*probing_session.ProbingSession, len(refs))
	for _, resolvedRes := range resp.GetProbingSessions() {
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

func (a *apiProbingSessionAccess) QueryProbingSessions(ctx context.Context, query *probing_session.ListQuery) (*probing_session.QueryResultSnapshot, error) {
	request := &probing_session_client.ListProbingSessionsRequest{
		Filter:            query.Filter,
		FieldMask:         query.Mask,
		IncludePagingInfo: query.WithPagingInfo,
	}
	if query.Pager != nil {
		request.PageSize = int32(query.Pager.Limit)
		request.OrderBy = query.Pager.OrderBy
		request.PageToken = query.Pager.Cursor
	}
	resp, err := a.client.ListProbingSessions(ctx, request)
	if err != nil {
		return nil, err
	}
	return &probing_session.QueryResultSnapshot{
		ProbingSessions:   resp.ProbingSessions,
		NextPageCursor:    resp.NextPageToken,
		PrevPageCursor:    resp.PrevPageToken,
		TotalResultsCount: resp.TotalResultsCount,
		CurrentOffset:     resp.CurrentOffset,
	}, nil
}

func (a *apiProbingSessionAccess) SearchProbingSessions(ctx context.Context, query *probing_session.SearchQuery) (*probing_session.QueryResultSnapshot, error) {
	request := &probing_session_client.SearchProbingSessionsRequest{
		Phrase:    query.Phrase,
		Filter:    query.Filter,
		FieldMask: query.Mask,
	}
	if query.Pager != nil {
		request.PageSize = int32(query.Pager.Limit)
		request.OrderBy = query.Pager.OrderBy
		request.PageToken = query.Pager.Cursor
	}
	resp, err := a.client.SearchProbingSessions(ctx, request)
	if err != nil {
		return nil, err
	}
	return &probing_session.QueryResultSnapshot{
		ProbingSessions:   resp.ProbingSessions,
		NextPageCursor:    resp.NextPageToken,
		PrevPageCursor:    resp.PrevPageToken,
		CurrentOffset:     resp.CurrentOffset,
		TotalResultsCount: resp.TotalResultsCount,
	}, nil
}

func (a *apiProbingSessionAccess) WatchProbingSession(ctx context.Context, query *probing_session.GetQuery, observerCb func(*probing_session.ProbingSessionChange) error) error {
	request := &probing_session_client.WatchProbingSessionRequest{
		Name:      query.Reference,
		FieldMask: query.Mask,
	}
	changesStream, initErr := a.client.WatchProbingSession(ctx, request)
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

func (a *apiProbingSessionAccess) WatchProbingSessions(ctx context.Context, query *probing_session.WatchQuery, observerCb func(*probing_session.QueryResultChange) error) error {
	request := &probing_session_client.WatchProbingSessionsRequest{
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
	changesStream, initErr := a.client.WatchProbingSessions(ctx, request)
	if initErr != nil {
		return initErr
	}
	for {
		respChange, err := changesStream.Recv()
		if err != nil {
			return fmt.Errorf("watch recv error: %w", err)
		}
		changesWithPaging := &probing_session.QueryResultChange{
			Changes:      respChange.ProbingSessionChanges,
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

func (a *apiProbingSessionAccess) SaveProbingSession(ctx context.Context, res *probing_session.ProbingSession, opts ...gotenresource.SaveOption) error {
	saveOpts := gotenresource.MakeSaveOptions(opts)
	previousRes := saveOpts.GetPreviousResource()

	if previousRes == nil && !saveOpts.OnlyUpdate() && !saveOpts.OnlyCreate() {
		var err error
		previousRes, err = a.GetProbingSession(ctx, &probing_session.GetQuery{Reference: res.Name.AsReference()})
		if err != nil {
			if statusErr, ok := status.FromError(err); !ok || statusErr.Code() != codes.NotFound {
				return err
			}
		}
	}

	if saveOpts.OnlyUpdate() || previousRes != nil {
		updateRequest := &probing_session_client.UpdateProbingSessionRequest{
			ProbingSession: res,
		}
		if updateMask := saveOpts.GetUpdateMask(); updateMask != nil {
			updateRequest.UpdateMask = updateMask.(*probing_session.ProbingSession_FieldMask)
		}
		if mask, conditionalState := saveOpts.GetCAS(); mask != nil && conditionalState != nil {
			updateRequest.Cas = &probing_session_client.UpdateProbingSessionRequest_CAS{
				ConditionalState: conditionalState.(*probing_session.ProbingSession),
				FieldMask:        mask.(*probing_session.ProbingSession_FieldMask),
			}
		}
		_, err := a.client.UpdateProbingSession(ctx, updateRequest)
		if err != nil {
			return err
		}
		return nil
	} else {
		createRequest := &probing_session_client.CreateProbingSessionRequest{
			ProbingSession: res,
		}
		_, err := a.client.CreateProbingSession(ctx, createRequest)
		if err != nil {
			return err
		}
		return nil
	}
}

func (a *apiProbingSessionAccess) DeleteProbingSession(ctx context.Context, ref *probing_session.Reference, opts ...gotenresource.DeleteOption) error {
	request := &probing_session_client.DeleteProbingSessionRequest{
		Name: ref,
	}
	_, err := a.client.DeleteProbingSession(ctx, request)
	return err
}

func init() {
	gotenaccess.GetRegistry().RegisterApiAccessConstructor(probing_session.GetDescriptor(), func(cc grpc.ClientConnInterface) gotenresource.Access {
		return probing_session.AsAnyCastAccess(NewApiProbingSessionAccess(probing_session_client.NewProbingSessionServiceClient(cc)))
	})
}
