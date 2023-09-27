package service

import (
	"context"

	"github.com/nikhilnarayanan623/x-tention-crew/pkg/pb"
	"github.com/nikhilnarayanan623/x-tention-crew/pkg/usecase/interfaces"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService struct {
	usecase interfaces.UserUseCase
	pb.UnimplementedUserServiceServer
}

func NewUserService(usecase interfaces.UserUseCase) pb.UserServiceServer {
	return &UserService{
		usecase: usecase,
	}
}

func (u *UserService) GetUsers(ctx context.Context, req *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {

	allUsers, err := u.usecase.FindAllUsersNameAndCount(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	response := &pb.GetUsersResponse{
		UserCount: allUsers.Count,
		Names:     allUsers.Names,
	}

	return response, nil
}
