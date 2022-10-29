// Code generated by protoc-gen-goten-access
// Resource: QualityProfile
// DO NOT EDIT!!!

package quality_profile_access

import (
	"context"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"

	gotenaccess "github.com/cloudwan/goten-sdk/runtime/access"
	"github.com/cloudwan/goten-sdk/runtime/api/view"
	"github.com/cloudwan/goten-sdk/runtime/api/watch_type"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"

	quality_profile_client "github.com/cloudwan/watchdog-sdk/client/v1alpha2/quality_profile"
	quality_profile "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/quality_profile"
)

// QueryWatcher is a low-level, stateless watcher.
// Initial updates are sent in chunks. Once snapshot is complete, further
// changes are incremental - unless Reset flag is set, in which case another
// snapshot is received.
type QueryWatcher struct {
	client       quality_profile_client.QualityProfileServiceClient
	params       *QueryWatcherParams
	syncDeadline time.Time
	identifier   int
	evtsChan     chan *QueryWatcherEvent
	iEvtsChan    chan gotenaccess.QueryWatcherEvent
	resumeToken  string
	startingTime *timestamppb.Timestamp
}

type QueryWatcherParams struct {
	Parent       *quality_profile.ParentReference
	Filter       *quality_profile.Filter
	View         view.View
	FieldMask    *quality_profile.QualityProfile_FieldMask
	OrderBy      *quality_profile.OrderBy
	Cursor       *quality_profile.PagerCursor
	ChunkSize    int
	PageSize     int
	WatchType    watch_type.WatchType
	StartingTime *timestamppb.Timestamp

	RecoveryDeadline time.Duration
	RetryTimeout     time.Duration
}

type QueryWatcherEvent struct {
	Identifier   int
	Changes      []*quality_profile.QualityProfileChange
	Reset        bool
	LostSync     bool
	InSync       bool
	SnapshotSize int64
	CheckSize    bool
}

func (e *QueryWatcherEvent) GetWatcherIdentifier() int {
	return e.Identifier
}

func (e *QueryWatcherEvent) GetChanges() gotenresource.ResourceChangeList {
	return quality_profile.QualityProfileChangeList(e.Changes)
}

func (e *QueryWatcherEvent) IsReset() bool {
	return e.Reset
}

func (e *QueryWatcherEvent) IsLostSync() bool {
	return e.LostSync
}

func (e *QueryWatcherEvent) IsSync() bool {
	return e.InSync
}

func (e *QueryWatcherEvent) GetSnapshotSize() int64 {
	return e.SnapshotSize
}

func (e *QueryWatcherEvent) HasSnapshotSize() bool {
	return e.CheckSize
}

func NewQueryWatcher(id int, client quality_profile_client.QualityProfileServiceClient,
	params *QueryWatcherParams, evtsChan chan *QueryWatcherEvent) *QueryWatcher {
	return &QueryWatcher{
		client:       client,
		params:       params,
		identifier:   id,
		evtsChan:     evtsChan,
		startingTime: params.StartingTime,
	}
}

func NewQueryWatcherWithIChan(id int, client quality_profile_client.QualityProfileServiceClient,
	params *QueryWatcherParams, evtsChan chan gotenaccess.QueryWatcherEvent) *QueryWatcher {
	return &QueryWatcher{
		client:       client,
		params:       params,
		identifier:   id,
		iEvtsChan:    evtsChan,
		startingTime: params.StartingTime,
	}
}

func (qw *QueryWatcher) QueryWatcher() {}

func (qw *QueryWatcher) Run(ctx context.Context) error {
	log := ctxlogrus.Extract(ctx).
		WithField("query-watcher", "qualityProfile-query-watcher").
		WithField("query-parent", qw.params.Parent.String()).
		WithField("query-filter", qw.params.Filter.String()).
		WithField("query-order-by", qw.params.OrderBy.String()).
		WithField("query-cursor", qw.params.Cursor.String())
	ctx = ctxlogrus.ToContext(ctx, log)

	log.Infof("Running new query")
	inSync := false
	skipErrorBackoff := false
	for {
		stream, err := qw.client.WatchQualityProfiles(ctx, &quality_profile_client.WatchQualityProfilesRequest{
			Type:         qw.params.WatchType,
			Parent:       qw.params.Parent,
			Filter:       qw.params.Filter,
			View:         qw.params.View,
			FieldMask:    qw.params.FieldMask,
			MaxChunkSize: int32(qw.params.ChunkSize),
			OrderBy:      qw.params.OrderBy,
			ResumeToken:  qw.resumeToken,
			PageSize:     int32(qw.params.PageSize),
			PageToken:    qw.params.Cursor,
			StartingTime: qw.startingTime,
		})

		if err != nil {
			if ctx.Err() == nil {
				log.WithError(err).Warnf("watch initialization error")
			}
		} else {
			pending := make([]*quality_profile.QualityProfileChange, 0)
			for {
				resp, err := stream.Recv()
				if err != nil {
					if ctx.Err() == nil {
						log.WithError(err).Warnf("watch error")
					}
					break
				} else {
					var outputEvt *QueryWatcherEvent

					// until we reach initial sync, we will send all the data as we get to minimize
					// potential impact on memory (if receiver does not need state). Later on, we will
					// collect changes and send once IsCurrent flag is sent. This is to handle soft reset
					// flag. Changes after initial sync are however practically always small.
					skipErrorBackoff = true
					if inSync {
						pending = append(pending, resp.GetQualityProfileChanges()...)
						if resp.IsSoftReset {
							log.Debugf("received soft reset after %d changes", len(pending))
							pending = nil
						} else if resp.IsHardReset {
							log.Warnf("received hard reset after %d changes", len(pending))

							qw.resumeToken = ""
							inSync = false
							pending = nil
							outputEvt = &QueryWatcherEvent{
								Identifier: qw.identifier,
								Reset:      true,
							}
						} else if resp.GetSnapshotSize() >= 0 {
							log.Debugf("received snapshot size info: %d", resp.GetSnapshotSize())

							outputEvt = &QueryWatcherEvent{
								Identifier:   qw.identifier,
								SnapshotSize: resp.GetSnapshotSize(),
								CheckSize:    true,
								Changes:      pending,
								InSync:       true,
							}
						} else if resp.GetIsCurrent() {
							qw.syncDeadline = time.Time{}
							if resp.GetResumeToken() != "" {
								qw.resumeToken = resp.GetResumeToken()
								qw.startingTime = nil
							}
							if len(pending) > 0 {
								outputEvt = &QueryWatcherEvent{
									Identifier: qw.identifier,
									Changes:    pending,
									InSync:     true,
								}
							}
							pending = nil
						}
					} else {
						if resp.IsCurrent {
							log.Infof("query synchronized")
							inSync = true
							qw.syncDeadline = time.Time{}
							if resp.GetResumeToken() != "" {
								qw.resumeToken = resp.GetResumeToken()
								qw.startingTime = nil
							}
						}
						outputEvt = &QueryWatcherEvent{
							Identifier:   qw.identifier,
							Changes:      resp.GetQualityProfileChanges(),
							SnapshotSize: resp.SnapshotSize,
							Reset:        resp.IsHardReset || resp.IsSoftReset,
							InSync:       inSync,
							CheckSize:    resp.SnapshotSize >= 0,
						}
					}
					if outputEvt != nil {
						qw.sendEvt(ctx, outputEvt)
					}
				}
			}
		}

		if ctx.Err() != nil {
			return ctx.Err()
		}

		// if we disconnected during initial snapshot (we were not in sync), send a message to cancel all data
		if !inSync {
			evt := &QueryWatcherEvent{
				Identifier: qw.identifier,
				Reset:      true,
			}
			qw.sendEvt(ctx, evt)
		}
		if qw.syncDeadline.IsZero() && qw.params.RecoveryDeadline > 0 {
			qw.syncDeadline = time.Now().UTC().Add(qw.params.RecoveryDeadline)
			log.Infof("lost sync, scheduling recovery with timeout %s", qw.syncDeadline)
		} else if !qw.syncDeadline.IsZero() && time.Now().UTC().After(qw.syncDeadline) {
			log.Errorf("could not recover within %s, reporting lost sync", qw.syncDeadline)
			evt := &QueryWatcherEvent{
				Identifier: qw.identifier,
				LostSync:   true,
				Reset:      true,
			}
			qw.resumeToken = ""
			inSync = false
			qw.sendEvt(ctx, evt)
		}

		// If we had working watch, dont sleep on first disconnection, we are likely to be able to
		// reconnect quickly and then we dont want to miss updates
		if !skipErrorBackoff {
			backoff := time.After(qw.params.RetryTimeout)
			select {
			case <-backoff:
				log.Debugf("after backoff %s", qw.params.RetryTimeout)
			case <-ctx.Done():
				log.Debugf("context done, reason: %s", ctx.Err())
				return ctx.Err()
			}
		} else {
			skipErrorBackoff = false
		}
	}
}

func (qw *QueryWatcher) sendEvt(ctx context.Context, evt *QueryWatcherEvent) {
	if qw.evtsChan != nil {
		select {
		case <-ctx.Done():
		case qw.evtsChan <- evt:
		}
	} else {
		select {
		case <-ctx.Done():
		case qw.iEvtsChan <- evt:
		}
	}
}

func init() {
	gotenaccess.GetRegistry().RegisterQueryWatcherConstructor(quality_profile.GetDescriptor(), func(id int, cc grpc.ClientConnInterface,
		params *gotenaccess.QueryWatcherConfigParams, ch chan gotenaccess.QueryWatcherEvent) gotenaccess.QueryWatcher {
		cfg := &QueryWatcherParams{
			WatchType:        params.WatchType,
			View:             params.View,
			ChunkSize:        params.ChunkSize,
			PageSize:         params.PageSize,
			StartingTime:     params.StartingTime,
			RecoveryDeadline: params.RecoveryDeadline,
			RetryTimeout:     params.RetryTimeout,
		}
		if params.FieldMask != nil {
			cfg.FieldMask = params.FieldMask.(*quality_profile.QualityProfile_FieldMask)
		}
		if params.OrderBy != nil {
			cfg.OrderBy = params.OrderBy.(*quality_profile.OrderBy)
		}
		if params.Cursor != nil {
			cfg.Cursor = params.Cursor.(*quality_profile.PagerCursor)
		}
		if params.Parent != nil {
			cfg.Parent = params.Parent.(*quality_profile.ParentReference)
		}
		if params.Filter != nil {
			cfg.Filter = params.Filter.(*quality_profile.Filter)
		}
		return NewQueryWatcherWithIChan(id, quality_profile_client.NewQualityProfileServiceClient(cc), cfg, ch)
	})
}
