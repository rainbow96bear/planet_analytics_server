package client

import (
	"github.com/rainbow96bear/planet_analytics_server/config"
)

type GrpcClients struct {
	User *UserClient
	// 앞으로 증가할 클라이언트들
	// FeedClient pb.FeedServiceClient
	// ChatClient pb.ChatServiceClient
}

func NewGrpcClients() (*GrpcClients, error) {
	conn, err := NewGrpcConn(config.USER_GRPC_SERVER_ADDR)
	if err != nil {
		return nil, err
	}
	userClient := NewUserClient(conn)
	return &GrpcClients{
		User: userClient,
	}, nil
}
