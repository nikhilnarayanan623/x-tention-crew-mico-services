package interfaces

import "github.com/nikhilnarayanan623/x-tention-crew/service2/pkg/utils/models"

type UseCase interface {
	GetSequentially(waitSecond int) (models.AllUserDetails, error)
	GetParallel(waitSecond int) (models.AllUserDetails, error)
}
