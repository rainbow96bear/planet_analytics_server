package server

import (
	"context"

	"github.com/rainbow96bear/planet_analytics_server/internal/service"
	pb "github.com/rainbow96bear/planet_utils/pb"
	"github.com/rainbow96bear/planet_utils/pkg/logger"
)

type AnalyticsGrpcServer struct {
	pb.UnimplementedAnalyticsServiceServer
	analyticseService service.AnalyticsServiceInterface
}

func NewAnalyticsGrpcServer(
	analyticsSvc service.AnalyticsServiceInterface,
) *AnalyticsGrpcServer {
	return &AnalyticsGrpcServer{
		analyticseService: analyticsSvc,
	}
}

func (s *AnalyticsGrpcServer) PublishEvent(
	ctx context.Context,
	req *pb.PublishEventRequest,
) (*pb.PublishEventResponse, error) {

	if req == nil || req.EventName == "" {
		logger.Warnf(
			"grpc PublishEvent invalid request req=%+v",
			req,
		)
		return &pb.PublishEventResponse{
			Success: false,
		}, nil
	}

	if _, err := s.analyticseService.PublishEvent(ctx, req); err != nil {
		logger.Errorf(
			"grpc PublishEvent failed event=%s userId=%s err=%v",
			req.EventName,
			req.UserId,
			err,
		)

		return &pb.PublishEventResponse{
			Success: false,
		}, nil
	}

	logger.Infof(
		"grpc PublishEvent success event=%s userId=%s",
		req.EventName,
		req.UserId,
	)

	return &pb.PublishEventResponse{
		Success: true,
	}, nil
}
