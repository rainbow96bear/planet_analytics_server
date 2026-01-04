package server

import (
	"net"

	"github.com/rainbow96bear/planet_analytics_server/config"
	"github.com/rainbow96bear/planet_analytics_server/internal/bootstrap"
	pb "github.com/rainbow96bear/planet_utils/pb"
	"github.com/rainbow96bear/planet_utils/pkg/logger"
	"google.golang.org/grpc"
)

func RunGrpcServer(deps *bootstrap.Dependencies) error {
	listener, err := net.Listen("tcp", ":"+config.ANALYTICS_GRPC_PORT)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()

	analyticsServer := NewAnalyticsGrpcServer(deps.Services.Analytic)
	pb.RegisterAnalyticsServiceServer(grpcServer, analyticsServer)

	logger.Infof(
		"User gRPC Server running on :%s",
		config.ANALYTICS_GRPC_PORT,
	)

	return grpcServer.Serve(listener)
}
