package client

import (
	"context"

	"github.com/rainbow96bear/planet_utils/pb"
	"google.golang.org/grpc"
)

type UserClient struct {
	client pb.UserServiceClient
}

func NewUserClient(conn *grpc.ClientConn) *UserClient {
	return &UserClient{
		client: pb.NewUserServiceClient(conn),
	}
}

func (u *UserClient) CreateProfile(ctx context.Context, req *pb.CreateProfileRequest) (*pb.CreateProfileResponse, error) {

	return u.client.CreateProfile(ctx, req)
}
