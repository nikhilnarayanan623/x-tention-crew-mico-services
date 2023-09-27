package service

import (
	"github.com/nikhilnarayanan623/x-tention-crew/pkg/pb"
	"github.com/nikhilnarayanan623/x-tention-crew/pkg/usecase/interfaces"
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
