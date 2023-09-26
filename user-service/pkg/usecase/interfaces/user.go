package interfaces

import (
	"context"

	"github.com/nikhilnarayanan623/x-tention-crew/pkg/utils/models/request"
)

type UserUseCase interface {
	CreateAccount(ctx context.Context, signUpDetails request.User) (uint32, error)
}
