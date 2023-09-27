package interfaces

import (
	"context"

	"github.com/nikhilnarayanan623/x-tention-crew/pkg/domain"
	"github.com/nikhilnarayanan623/x-tention-crew/pkg/utils/models/response"
)

type UserRepo interface {
	IsUserAlreadyExistWithThisEmail(ctx context.Context, email string) (exist bool, err error)
	IsUserExist(ctx context.Context, userID uint32) (bool, error)
	FindUserByID(ctx context.Context, id uint32) (domain.User, error)
	SaveUser(ctx context.Context, user domain.User) (domain.User, error)
	UpdateUser(ctx context.Context, user domain.User) (domain.User, error)
	DeleteUser(ctx context.Context, userID uint32) error
	FindAllUsersNameAndCount(ctx context.Context)(response.AllUsers,error)
}
