package service

import (
	"context"

	"github.com/rainbow96bear/planet_analytics_server/internal/repository"
	"github.com/rainbow96bear/planet_utils/pb"
	"gorm.io/gorm"
)

type AnalyticsServiceInterface interface {
	PublishEvent(
		ctx context.Context,
		req *pb.PublishEventRequest,
	) (*pb.PublishEventResponse, error)
}

type AnalyticsService struct {
	db            *gorm.DB
	AnalyticsRepo *repository.AnalyticsRepository
}

func NewAnalyticsService(
	db *gorm.DB,
	analyticsRepo *repository.AnalyticsRepository,
) AnalyticsServiceInterface {
	return &AnalyticsService{
		db:            db,
		AnalyticsRepo: analyticsRepo,
	}
}

func (s *AnalyticsService) PublishEvent(
	ctx context.Context,
	req *pb.PublishEventRequest,
) (*pb.PublishEventResponse, error) {

	// 최소 검증
	if req.EventName == "" {
		return &pb.PublishEventResponse{Success: false}, nil
	}

	_ = s.AnalyticsRepo.SaveEvent(ctx, req)

	return &pb.PublishEventResponse{Success: true}, nil
}
