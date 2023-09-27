package interfaces

import (
	"context"

	"github.com/nikhilnarayanan623/x-tention-crew/pkg/domain"
)

type UserRepo interface {
	IsUserAlreadyExistWithThisEmail(ctx context.Context, email string) (exist bool, err error)
	SaveUser(ctx context.Context, user domain.User) (domain.User, error)
	FindUserByID(ctx context.Context, id uint32) (domain.User, error)
}
