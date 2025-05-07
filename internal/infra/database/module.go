package database

import (
	"context"

	"go.uber.org/fx"

	"github.com/felipeversiane/donation-server/config"
)

var Module = fx.Options(
	fx.Provide(
		func(config config.DatabaseConfig) (Interface, error) {
			return New(config)
		},
	),
	fx.Invoke(func(lc fx.Lifecycle, db Interface) {
		lc.Append(fx.Hook{
			OnStop: func(ctx context.Context) error {
				db.Close()
				return nil
			},
		})
	}),
)
