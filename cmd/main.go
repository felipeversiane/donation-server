// @title Donation API
// @version 1.0
// @description RESTful API to receive donations via pix or credit/debit card.
// @host localhost:8000
// @schemes http
package main

import (
	"github.com/felipeversiane/donation-server/config"
	"github.com/felipeversiane/donation-server/internal/adapter/in/http"
	"github.com/felipeversiane/donation-server/internal/provider/database"
	"github.com/felipeversiane/donation-server/pkg/logger"

	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		config.Module,
		logger.Module,
		database.Module,
		http.Module,
	)

	app.Run()
}
