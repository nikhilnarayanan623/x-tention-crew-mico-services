package interfaces

import (
	"context"

	"github.com/nikhilnarayanan623/x-tention-crew/service2/pkg/utils/models"
)

type UserServiceClient interface {
	GetAllUsers(ctx context.Context) (models.AllUserDetails, error)
}
