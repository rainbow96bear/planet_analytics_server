package repository

import (
	"context"

	"github.com/rainbow96bear/planet_analytics_server/internal/tx"
	"github.com/rainbow96bear/planet_utils/pb"
	"gorm.io/gorm"
)

type AnalyticsRepository struct {
	db *gorm.DB
}

func NewAnalyticsRepository(db *gorm.DB) *AnalyticsRepository {
	if db == nil {
		panic("database connection is required")
	}

	return &AnalyticsRepository{
		db: db,
	}
}

func (r *AnalyticsRepository) getDB(ctx context.Context) *gorm.DB {
	if tx := tx.GetTx(ctx); tx != nil { // Context에서 트랜잭션 확인
		// 트랜잭션이 있을 경우 Context를 연결하여 반환 (gorm 권장 사항)
		return tx.WithContext(ctx)
	}
	// 트랜잭션이 없을 경우 기본 DB 연결을 반환
	return r.db.WithContext(ctx)
}

func (r *AnalyticsRepository) SaveEvent(
	ctx context.Context,
	e *pb.PublishEventRequest,
) error {

	return r.db.Exec(`
		INSERT INTO analytics_event
		(event_name, user_id, anonymous_id, session_id, occurred_at, properties)
		VALUES (?, ?, ?, ?, to_timestamp(?), ?)
	`,
		e.EventName,
		e.UserId,
		e.AnonymousId,
		e.SessionId,
		e.OccurredAt,
		e.Properties,
	).Error
}
