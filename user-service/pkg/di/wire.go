//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/nikhilnarayanan623/x-tention-crew/user-servcie/pkg/api"
	"github.com/nikhilnarayanan623/x-tention-crew/user-servcie/pkg/api/handler"
	"github.com/nikhilnarayanan623/x-tention-crew/user-servcie/pkg/api/service"
	"github.com/nikhilnarayanan623/x-tention-crew/user-servcie/pkg/config"
	"github.com/nikhilnarayanan623/x-tention-crew/user-servcie/pkg/db"
	"github.com/nikhilnarayanan623/x-tention-crew/user-servcie/pkg/repository"
	"github.com/nikhilnarayanan623/x-tention-crew/user-servcie/pkg/usecase"
)

func InitializeAPI(cfg config.Config) (*api.Server, error) {

	wire.Build(
		db.ConnectToDatabase,
		repository.NewUserRepo,
		repository.NewCacheRepo,
		usecase.NewAuthUseCase,
		service.NewUserService,
		handler.NewUserHandler,
		api.NewServer,
	)

	return &api.Server{}, nil
}
