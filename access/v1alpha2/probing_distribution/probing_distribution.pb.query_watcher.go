// Code generated by protoc-gen-goten-access
// Resource: ProbingDistribution
// DO NOT EDIT!!!

package probing_distribution_access

import (
	"context"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"

	"github.com/cloudwan/goten-sdk/runtime/api/view"
	"github.com/cloudwan/goten-sdk/runtime/api/watch_type"

	probing_distribution_client "github.com/cloudwan/watchdog-sdk/client/v1alpha2/probing_distribution"
	probing_distribution "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/probing_distribution"
)

// QueryWatcher is a low-level, stateless watcher.
// Initial updates are sent in chunks. Once snapshot is complete, further
// changes are incremental - unless Reset flag is set, in which case another
// snapshot is received.
type QueryWatcher struct {
	client       probing_distribution_client.ProbingDistributionServiceClient
	params       *QueryWatcherParams
	syncDeadline time.Time
	identifier   int
	evtsChan     chan *QueryWatcherEvent
	resumeToken  string
}

type QueryWatcherParams struct {
	Parent    *probing_distribution.ParentReference
	Filter    *probing_distribution.Filter
	View      view.View
	FieldMask *probing_distribution.ProbingDistribution_FieldMask
	OrderBy   *probing_distribution.OrderBy
	Cursor    *probing_distribution.PagerCursor
	ChunkSize int
	PageSize  int
	WatchType watch_type.WatchType

	RecoveryDeadline time.Duration
	RetryTimeout     time.Duration
}

type QueryWatcherEvent struct {
	Identifier   int
	Changes      []*probing_distribution.ProbingDistributionChange
	Reset        bool
	LostSync     bool
	InSync       bool
	SnapshotSize int64
	CheckSize    bool
}

func NewQueryWatcher(id int, client probing_distribution_client.ProbingDistributionServiceClient,
	params *QueryWatcherParams, evtsChan chan *QueryWatcherEvent) *QueryWatcher {
	return &QueryWatcher{
		client:     client,
		params:     params,
		identifier: id,
		evtsChan:   evtsChan,
	}
}

func (qw *QueryWatcher) Run(ctx context.Context) error {
	log := ctxlogrus.Extract(ctx).
		WithField("query-watcher", "probingDistribution-query-watcher").
		WithField("query-parent", qw.params.Parent.String()).
		WithField("query-filter", qw.params.Filter.String()).
		WithField("query-order-by", qw.params.OrderBy.String()).
		WithField("query-cursor", qw.params.Cursor.String())
	ctx = ctxlogrus.ToContext(ctx, log)

	log.Infof("Running new query")
	inSync := false
	for {
		stream, err := qw.client.WatchProbingDistributions(ctx, &probing_distribution_client.WatchProbingDistributionsRequest{
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
		})

		if err != nil {
			log.WithError(err).Errorf("watch initialization error")
		} else {
			pending := make([]*probing_distribution.ProbingDistributionChange, 0)
			for {
				resp, err := stream.Recv()
				if err != nil {
					log.WithError(err).Errorf("watch error")
					break
				} else {
					var outputEvt *QueryWatcherEvent

					// until we reach initial sync, we will send all the data as we get to minimize
					// potential impact on memory (if receiver does not need state). Later on, we will
					// collect changes and send once IsCurrent flag is sent. This is to handle soft reset
					// flag. Changes after initial sync are however practically always small.
					if inSync {
						pending = append(pending, resp.GetProbingDistributionChanges()...)
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
							}
						}
						outputEvt = &QueryWatcherEvent{
							Identifier:   qw.identifier,
							Changes:      resp.GetProbingDistributionChanges(),
							SnapshotSize: resp.SnapshotSize,
							Reset:        resp.IsHardReset || resp.IsSoftReset,
							InSync:       inSync,
							CheckSize:    resp.SnapshotSize >= 0,
						}
					}
					if outputEvt != nil {
						select {
						case <-ctx.Done():
							return ctx.Err()
						case qw.evtsChan <- outputEvt:
						}
					}
				}
			}
		}

		if ctx.Err() != nil {
			return ctx.Err()
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
			select {
			case <-ctx.Done():
				return ctx.Err()
			case qw.evtsChan <- evt:
			}
		}

		backoff := time.After(qw.params.RetryTimeout)
		select {
		case <-backoff:
			log.Debugf("after backoff %s", qw.params.RetryTimeout)
		case <-ctx.Done():
			log.Debugf("context done, reason: %s", ctx.Err())
			return ctx.Err()
		}
	}
}
