package main

import (
	"github.com/felipeversiane/donation-server/internal/config"
	"github.com/felipeversiane/donation-server/internal/infra/database"
	"github.com/felipeversiane/donation-server/internal/infra/http"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		config.Module,
		database.Module,
		http.Module,
		fx.NopLogger,
	)

	app.Run()
}
