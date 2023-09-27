package client

import (
	"context"
	"fmt"

	"github.com/nikhilnarayanan623/x-tention-crew/service2/pkg/client/interfaces"
	"github.com/nikhilnarayanan623/x-tention-crew/service2/pkg/config"
	"github.com/nikhilnarayanan623/x-tention-crew/service2/pkg/pb"
	"github.com/nikhilnarayanan623/x-tention-crew/service2/pkg/utils/models"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type userServiceClient struct {
	client pb.UserServiceClient
}

func NewUserServiceClient(cfg config.Config) (interfaces.UserServiceClient, error) {

	addr := fmt.Sprintf("%s:%s", cfg.UserServiceGrpcHost, cfg.UserServiceGrpcPort)

	cc, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := pb.NewUserServiceClient(cc)

	return &userServiceClient{
		client: client,
	}, nil
}

func (u *userServiceClient) GetAllUsers(ctx context.Context) (models.AllUserDetails, error) {

	res, err := u.client.GetUsers(ctx, &pb.GetUsersRequest{})

	if err != nil {
		return models.AllUserDetails{}, err
	}

	return models.AllUserDetails{
		Count: res.GetUserCount(),
		Names: res.GetNames(),
	}, nil
}
