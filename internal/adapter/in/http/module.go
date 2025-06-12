package http

import (
	"context"

	"go.uber.org/fx"

	"github.com/felipeversiane/donation-server/config"
	"github.com/felipeversiane/donation-server/pkg/logger"
)

var Module = fx.Options(
	fx.Provide(
		func(cfg config.Interface, log logger.Interface) ServerInterface {
			return New(cfg.HTTPServer(), log)
		},
	),
	fx.Invoke(func(lc fx.Lifecycle, server ServerInterface) {
		lc.Append(fx.Hook{
			OnStart: func(ctx context.Context) error {
				go func() {
					server.InitRoutes()
					if err := server.Start(); err != nil {
						panic(err)
					}
				}()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				return server.Shutdown(ctx)
			},
		})
	}),
)
