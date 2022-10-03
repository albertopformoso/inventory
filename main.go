package main

import (
	"context"
	"fmt"
	"log"

	"github.com/albertopformoso/inventory/database"
	"github.com/albertopformoso/inventory/internal/controller"
	"github.com/albertopformoso/inventory/internal/repository"
	"github.com/albertopformoso/inventory/internal/routes"
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
			controller.New,
			routes.New,
			echo.New,
		),
		fx.Invoke(
			setLifeCycle,
		),
	)

	app.Run()
}

func setLifeCycle(lc fx.Lifecycle, e *echo.Echo, r *routes.Routes, s *settings.Settings) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			address := fmt.Sprintf(":%s", s.Port)
			go r.Start(e, address)

			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("Stopping server...")
			return nil
		},
	})
}
