// Code generated by protoc-gen-goten-access
// Resource: InternetQualityRating
// DO NOT EDIT!!!

package internet_quality_rating_access

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cloudwan/goten-sdk/runtime/api/watch_type"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"

	internet_quality_rating_client "github.com/cloudwan/watchdog-sdk/client/v1alpha2/internet_quality_rating"
	internet_quality_rating "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/internet_quality_rating"
)

var (
	_ = context.Context(nil)
	_ = fmt.GoStringer(nil)

	_ = codes.NotFound
	_ = status.Status{}

	_ = watch_type.WatchType_STATEFUL
	_ = gotenresource.ListQuery(nil)
)

type apiInternetQualityRatingAccess struct {
	client internet_quality_rating_client.InternetQualityRatingServiceClient
}

func NewApiInternetQualityRatingAccess(client internet_quality_rating_client.InternetQualityRatingServiceClient) internet_quality_rating.InternetQualityRatingAccess {
	return &apiInternetQualityRatingAccess{client: client}
}

func (a *apiInternetQualityRatingAccess) GetInternetQualityRating(ctx context.Context, query *internet_quality_rating.GetQuery) (*internet_quality_rating.InternetQualityRating, error) {
	request := &internet_quality_rating_client.GetInternetQualityRatingRequest{
		Name:      query.Reference,
		FieldMask: query.Mask,
	}
	return a.client.GetInternetQualityRating(ctx, request)
}

func (a *apiInternetQualityRatingAccess) BatchGetInternetQualityRatings(ctx context.Context, refs []*internet_quality_rating.Reference, opts ...gotenresource.BatchGetOption) error {
	batchGetOpts := gotenresource.MakeBatchGetOptions(opts)
	request := &internet_quality_rating_client.BatchGetInternetQualityRatingsRequest{
		Names: refs,
	}
	resp, err := a.client.BatchGetInternetQualityRatings(ctx, request)
	if err != nil {
		return err
	}
	resultMap := make(map[internet_quality_rating.Name]*internet_quality_rating.InternetQualityRating, len(refs))
	for _, resolvedRes := range resp.GetInternetQualityRatings() {
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

func (a *apiInternetQualityRatingAccess) QueryInternetQualityRatings(ctx context.Context, query *internet_quality_rating.ListQuery) (*internet_quality_rating.QueryResultSnapshot, error) {
	request := &internet_quality_rating_client.ListInternetQualityRatingsRequest{
		Filter:    query.Filter,
		FieldMask: query.Mask,
	}
	if query.Pager != nil {
		request.PageSize = int32(query.Pager.Limit)
		request.OrderBy = query.Pager.OrderBy
		request.PageToken = query.Pager.Cursor
	}
	resp, err := a.client.ListInternetQualityRatings(ctx, request)
	if err != nil {
		return nil, err
	}
	return &internet_quality_rating.QueryResultSnapshot{
		InternetQualityRatings: resp.InternetQualityRatings,
		NextPageCursor:         resp.NextPageToken,
		PrevPageCursor:         resp.PrevPageToken,
	}, nil
}

func (a *apiInternetQualityRatingAccess) WatchInternetQualityRating(ctx context.Context, query *internet_quality_rating.GetQuery, observerCb func(*internet_quality_rating.InternetQualityRatingChange) error) error {
	request := &internet_quality_rating_client.WatchInternetQualityRatingRequest{
		Name:      query.Reference,
		FieldMask: query.Mask,
	}
	changesStream, initErr := a.client.WatchInternetQualityRating(ctx, request)
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

func (a *apiInternetQualityRatingAccess) WatchInternetQualityRatings(ctx context.Context, query *internet_quality_rating.WatchQuery, observerCb func(*internet_quality_rating.QueryResultChange) error) error {
	request := &internet_quality_rating_client.WatchInternetQualityRatingsRequest{
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
	changesStream, initErr := a.client.WatchInternetQualityRatings(ctx, request)
	if initErr != nil {
		return initErr
	}
	for {
		respChange, err := changesStream.Recv()
		if err != nil {
			return fmt.Errorf("watch recv error: %w", err)
		}
		changesWithPaging := &internet_quality_rating.QueryResultChange{
			Changes:      respChange.InternetQualityRatingChanges,
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

func (a *apiInternetQualityRatingAccess) SaveInternetQualityRating(ctx context.Context, res *internet_quality_rating.InternetQualityRating, opts ...gotenresource.SaveOption) error {
	saveOpts := gotenresource.MakeSaveOptions(opts)
	previousRes := saveOpts.GetPreviousResource()

	if previousRes == nil {
		var err error
		previousRes, err = a.GetInternetQualityRating(ctx, &internet_quality_rating.GetQuery{Reference: res.Name.AsReference()})
		if err != nil {
			if statusErr, ok := status.FromError(err); !ok || statusErr.Code() != codes.NotFound {
				return err
			}
		}
	}

	if previousRes != nil {
		updateRequest := &internet_quality_rating_client.UpdateInternetQualityRatingRequest{
			InternetQualityRating: res,
		}
		_, err := a.client.UpdateInternetQualityRating(ctx, updateRequest)
		if err != nil {
			return err
		}
		return nil
	} else {
		createRequest := &internet_quality_rating_client.CreateInternetQualityRatingRequest{
			InternetQualityRating: res,
		}
		_, err := a.client.CreateInternetQualityRating(ctx, createRequest)
		if err != nil {
			return err
		}
		return nil
	}
}

func (a *apiInternetQualityRatingAccess) DeleteInternetQualityRating(ctx context.Context, ref *internet_quality_rating.Reference, opts ...gotenresource.DeleteOption) error {
	request := &internet_quality_rating_client.DeleteInternetQualityRatingRequest{
		Name: ref,
	}
	_, err := a.client.DeleteInternetQualityRating(ctx, request)
	return err
}
