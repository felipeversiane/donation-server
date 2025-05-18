// @title Donation Server
// @version 1.0
// @description RESTful API to receive donations via pix or credit/debit card.
// @host localhost:8080
// @BasePath /api/v1
// @schemes http
package main

import (
	"github.com/felipeversiane/donation-server/config"
	"github.com/felipeversiane/donation-server/internal/adapter/in/http"
	"github.com/felipeversiane/donation-server/internal/provider/database"

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
