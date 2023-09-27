//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/nikhilnarayanan623/x-tention-crew/service2/pkg/api"
	"github.com/nikhilnarayanan623/x-tention-crew/service2/pkg/api/handler"
	"github.com/nikhilnarayanan623/x-tention-crew/service2/pkg/client"
	"github.com/nikhilnarayanan623/x-tention-crew/service2/pkg/config"
	"github.com/nikhilnarayanan623/x-tention-crew/service2/pkg/usecase"
)

func InitializeAPI(cfg config.Config) (*api.Server, error) {

	wire.Build(
		client.NewUserServiceClient,
		usecase.NewUseCase,
		handler.NewHandler,
		api.NewServer,
	)

	return &api.Server{}, nil
}
