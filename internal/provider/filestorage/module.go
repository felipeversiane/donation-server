package filestorage

import (
	"context"

	"go.uber.org/fx"

	"github.com/felipeversiane/donation-server/config"
	"github.com/felipeversiane/donation-server/pkg/logger"
)

var Module = fx.Options(
	fx.Provide(func(cfg config.FileStorageConfig, log logger.Interface) (Interface, error) {
		return New(cfg, log)
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
