package main

import (
	"github.com/felipeversiane/donation-server/internal/infrastructure/api"
	"github.com/felipeversiane/donation-server/internal/infrastructure/config"
	"github.com/felipeversiane/donation-server/internal/infrastructure/database"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		config.Module,
		database.Module,
		api.Module,
		fx.NopLogger,
	)

	app.Run()
}
