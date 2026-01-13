package bootstrap

import (
	"time"

	"github.com/rainbow96bear/planet_analytics_server/internal/grpc/client"
	"github.com/rainbow96bear/planet_analytics_server/internal/repository"
	"github.com/rainbow96bear/planet_analytics_server/internal/service"
	"gorm.io/gorm"
)

type Dependencies struct {
	Repos       *Repositories
	GrpcClients *client.GrpcClients
	Services    *Services
}

type Repositories struct {
	Analytics *repository.AnalyticsRepository
	// OauthAccounts *repository.OauthAccountsRepository
	// Users         *repository.UsersRepository
}

type Services struct {
	// Kakao service.KakaoOauthServiceInterface
	Analytic service.AnalyticsServiceInterface
}

func InitDependencies(db *gorm.DB) (*Dependencies, error) {
	repos := initRepositories(db)

	deps := &Dependencies{
		Repos:       repos,
		GrpcClients: nil, // 처음엔 없음
	}

	// Services는 "빈 gRPC" 상태로 먼저 생성
	// deps.Services = initServices(db, repos, deps.GrpcClients)
	deps.Services = initServices(db, repos)

	// background reconnect
	go func() {
		for {
			grpcClients, err := client.NewGrpcClients()
			if err != nil {
				time.Sleep(5 * time.Second)
				continue
			}

			deps.GrpcClients = grpcClients
			break
		}
	}()

	return deps, nil
}

func initRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		Analytics: repository.NewAnalyticsRepository(db),
		// Profile:  repository.NewProfilesRepository(db),
		// Calendar: repository.NewCalendarEventsRepository(db),
		// Todo:     repository.NewTodosRepository(db),
		// Follow:   repository.NewFollowsRepository(db),
	}
}

func initServices(
	db *gorm.DB,
	repos *Repositories,
) *Services {
	analyticsSvc := service.NewAnalyticsService(
		db,
		repos.Analytics,
	)

	// profileSvc := service.NewProfileService(db, repos.Profile)

	// calendarSvc := service.NewCalendarService(
	// 	db,
	// 	repos.Analytics,
	// )

	// todoSvc := service.NewTodoService(
	// 	db,
	// 	repos.Todo,
	// 	grpcClients.Analytics, // ✅ port만 주입
	// )

	return &Services{
		Analytic: analyticsSvc,
		// Profile:  profileSvc,
		// Calendar: calendarSvc,
		// Todo:     todoSvc,
	}
}
