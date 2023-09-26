package interfaces

import (
	"context"

	"github.com/nikhilnarayanan623/x-tention-crew/pkg/domain"
)

type UserRepo interface {
	IsUserAlreadyExistWithThisEmail(ctx context.Context, email string) (exist bool, err error)
	SaveUser(ctx context.Context, user domain.User) (uint32, error)
	FindUserByID(ctx context.Context, id uint) (domain.User, error)
}
