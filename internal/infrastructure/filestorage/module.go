package filestorage

import (
	"context"

	"github.com/felipeversiane/donation-server/internal/infrastructure/config"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		func(config config.FileStorageConfig) (FileStorageInterface, error) {
			return New(config)
		},
	),
	fx.Invoke(func(lc fx.Lifecycle, filestorage FileStorageInterface) {
		lc.Append(fx.Hook{
			OnStop: func(ctx context.Context) error {
				filestorage.Close()
				return nil
			},
		})
	}),
)
