package main

import (
	"github.com/felipeversiane/donation-server/internal/adapter/in/http"
	"github.com/felipeversiane/donation-server/internal/adapter/out/database"
	"github.com/felipeversiane/donation-server/internal/config"

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
