package filestorage

import (
	"context"

	"go.uber.org/fx"

	"github.com/felipeversiane/donation-server/config"
)

var Module = fx.Options(
	fx.Provide(
		func(config config.FileStorageConfig) (Interface, error) {
			return New(config)
		},
	),
	fx.Invoke(func(lc fx.Lifecycle, filestorage Interface) {
		lc.Append(fx.Hook{
			OnStop: func(ctx context.Context) error {
				filestorage.Close()
				return nil
			},
		})
	}),
)
