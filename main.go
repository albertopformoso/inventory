package main

import (
	"context"
	"fmt"

	"github.com/albertopformoso/inventory/cmd/server"
	"github.com/albertopformoso/inventory/database"
	"github.com/albertopformoso/inventory/internal/controller"
	"github.com/albertopformoso/inventory/internal/repository"
	"github.com/albertopformoso/inventory/internal/routes"
	"github.com/albertopformoso/inventory/internal/service"
	"github.com/albertopformoso/inventory/logger"
	"github.com/albertopformoso/inventory/settings"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			context.Background,
			settings.New,
			logger.New,
			database.New,
			repository.New,
			service.New,
			controller.New,
			routes.New,
			echo.New,
			server.New,
		),
		fx.Invoke(
			setLifeCycle,
		),
	)

	app.Run()
}

func setLifeCycle(lc fx.Lifecycle, e *echo.Echo, server *server.Server, s *settings.Settings, log *zerolog.Logger) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			address := fmt.Sprintf(":%s", s.Port)
			go server.Start(e, address)

			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Info().Msg("Stopping server...")
			return nil
		},
	})
}
