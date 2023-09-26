//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/nikhilnarayanan623/x-tention-crew/pkg/api"
	"github.com/nikhilnarayanan623/x-tention-crew/pkg/api/handler"
	"github.com/nikhilnarayanan623/x-tention-crew/pkg/api/service"
	"github.com/nikhilnarayanan623/x-tention-crew/pkg/config"
	"github.com/nikhilnarayanan623/x-tention-crew/pkg/db"
	"github.com/nikhilnarayanan623/x-tention-crew/pkg/repository"
	"github.com/nikhilnarayanan623/x-tention-crew/pkg/usecase"
)

func InitializeAPI(cfg config.Config) (*api.Server, error) {

	wire.Build(
		db.ConnectToDatabase,
		repository.NewUserRepo,
		usecase.NewAuthUseCase,
		service.NewUserService,
		handler.NewUserHandler,
		api.NewServer,
	)

	return &api.Server{}, nil
}
