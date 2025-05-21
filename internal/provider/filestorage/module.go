package filestorage

import (
	"context"

	"go.uber.org/fx"

	"github.com/felipeversiane/donation-server/config"
)

var Module = fx.Options(
	fx.Provide(func(cfg config.FileStorageConfig) (Interface, error) {
		return New(cfg)
	}),
	fx.Invoke(func(lc fx.Lifecycle, fs Interface) {
		lc.Append(fx.Hook{
			OnStart: func(ctx context.Context) error {
				return fs.CreateBucket()
			},
			OnStop: func(ctx context.Context) error {
				return nil
			},
		})
	}),
)
