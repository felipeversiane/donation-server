package config

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(
		New,
		func(cfg Interface) LogConfig {
			return cfg.GetLogConfig()
		},
		func(cfg Interface) DatabaseConfig {
			return cfg.GetDatabaseConfig()
		},
		func(cfg Interface) HTTPServerConfig {
			return cfg.GetHTTPServerConfig()
		},
		func(cfg Interface) JwtTokenConfig {
			return cfg.GetJwtTokenConfig()
		},
		func(cfg Interface) FileStorageConfig {
			return cfg.GetFileStorageConfig()
		},
		func(cfg Interface) SentryConfig {
			return cfg.GetSentryConfig()
		},
	),
)
