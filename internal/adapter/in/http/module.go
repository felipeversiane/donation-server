package http

import (
	"context"

	"github.com/felipeversiane/donation-server/config"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		func(cfg config.ConfigInterface) HttpServerInterface {
			return New(
				cfg.GetHttpServerConfig(),
				cfg.GetSentryConfig(),
				cfg.GetEnvironment(),
			)
		},
	),
	fx.Invoke(func(lc fx.Lifecycle, server HttpServerInterface) {
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
