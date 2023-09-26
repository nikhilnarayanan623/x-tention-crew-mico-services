package interfaces

import (
	"context"

	"github.com/nikhilnarayanan623/x-tention-crew/pkg/utils/models/request"
	"github.com/nikhilnarayanan623/x-tention-crew/pkg/utils/models/response"
)

type UserUseCase interface {
	CreateAccount(ctx context.Context, signUpDetails request.User) (response.User, error)
	GetAccount(ctx context.Context, userID uint32) (response.User, error)
}
