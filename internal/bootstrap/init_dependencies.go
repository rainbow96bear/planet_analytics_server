package bootstrap

import (
	"github.com/rainbow96bear/planet_analytics_server/internal/grpc/client"
	grpcclient "github.com/rainbow96bear/planet_analytics_server/internal/grpc/client"
	"github.com/rainbow96bear/planet_analytics_server/internal/repository"
	"github.com/rainbow96bear/planet_analytics_server/internal/service"
	"gorm.io/gorm"
)

type Dependencies struct {
	Repos       *Repositories
	GrpcClients *grpcclient.GrpcClients
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
	// --- Repository 초기화 ---
	repos := initRepositories(db)

	// // --- gRPC Clients 초기화 ---
	grpcClients, err := grpcclient.NewGrpcClients()
	if err != nil {
		return nil, err
	}

	// // --- Service 초기화 ---
	services := initServices(db, repos, grpcClients)

	// // DI Container 패턴
	return &Dependencies{
		Repos:    repos,
		Services: services,
	}, nil
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
	grpcClients *client.GrpcClients,
) *Services {
	analyticsSvc := service.NewAnalyticsService(db, repos.Analytics, &grpcClients.User)

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
