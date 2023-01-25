package server

import (
	"context"

	log "github.com/sirupsen/logrus"
	epb "github.com/ukama/ukama/systems/common/pb/gen/events"
	"github.com/ukama/ukama/systems/subscriber/asr/pkg/db"
)

type HlrEventServer struct {
	hlrRepo  db.HlrRecordRepo
	gutiRepo db.GutiRepo
	epb.UnimplementedEventNotificationServiceServer
}

func NewHlrEventServer(hlrRepo db.HlrRecordRepo, gutiRepo db.GutiRepo) *HlrEventServer {
	return &HlrEventServer{
		hlrRepo:  hlrRepo,
		gutiRepo: gutiRepo,
	}
}

func (l *HlrEventServer) EventNotification(ctx context.Context, e *epb.Event) (*epb.EventResponse, error) {
	log.Infof("Received a message with Routing key %s and Message %+v", e.RoutingKey, e.Msg)
	switch e.RoutingKey {

	default:
		log.Errorf("No handler routing key %s", e.RoutingKey)
	}

	return &epb.EventResponse{}, nil
}
