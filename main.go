package main

import (
	"context"

	"github.com/albertopformoso/inventory/database"
	"github.com/albertopformoso/inventory/settings"
	"github.com/jmoiron/sqlx"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			context.Background,
			settings.New,
			database.New,
		),
		fx.Invoke(
			func(db *sqlx.DB){
				_, err := db.Query("SELECT * FROM user")
				if err != nil {
					panic(err)
				}
			},
		),
	)

	app.Run()
}
