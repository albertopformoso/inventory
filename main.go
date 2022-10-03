package main

import (
	"context"
	"fmt"

	"github.com/albertopformoso/inventory/database"
	"github.com/albertopformoso/inventory/internal/api"
	"github.com/albertopformoso/inventory/internal/repository"
	"github.com/albertopformoso/inventory/internal/service"
	"github.com/albertopformoso/inventory/settings"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			context.Background,
			settings.New,
			database.New,
			repository.New,
			service.New,
			api.New,
			echo.New,
		),
		fx.Invoke(
			setLifeCycle,
		),
	)

	app.Run()
}

func setLifeCycle(lc fx.Lifecycle, e *echo.Echo, api *api.API, s *settings.Settings) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			address := fmt.Sprintf(":%s", s.Port)
			go api.Start(e, address)

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
}
