package database

import (
	"context"

	"go.uber.org/fx"

	"github.com/felipeversiane/donation-server/config"
	"github.com/felipeversiane/donation-server/pkg/logger"
)

var Module = fx.Options(
	fx.Provide(
		func(cfg config.Database, log logger.Interface) (Interface, error) {
			return New(cfg, log)
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
