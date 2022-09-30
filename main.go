package main

import (
	"context"

	"github.com/albertopformoso/inventory/database"
	"github.com/albertopformoso/inventory/internal/repository"
	"github.com/albertopformoso/inventory/internal/service"
	"github.com/albertopformoso/inventory/settings"
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
		),
		fx.Invoke(
			func(ctx context.Context, svc service.Service) {
				err := svc.RegisterUser(ctx, "my@email.com", "myname", "mypassword");
				if err != nil {
					panic(err)
				}
				user, err := svc.LoginUser(ctx, "my@email.com", "mypassword")
				if err != nil { panic(err) }

				if user.Name != "myname" { panic("wrong name") }
			},
		),
	)

	app.Run()
}
