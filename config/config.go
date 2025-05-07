package config

import (
	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type config struct {
	Log         LogConfig         `envPrefix:"LOG_"`
	Database    DatabaseConfig    `envPrefix:"DB_"`
	HTTPServer  HTTPServerConfig  `envPrefix:"HTTP_SERVER_"`
	JwtToken    JwtTokenConfig    `envPrefix:"JWT_"`
	FileStorage FileStorageConfig `envPrefix:"FILE_STORAGE_"`
	Sentry      SentryConfig      `envPrefix:"SENTRY_"`
}

type Interface interface {
	GetDatabaseConfig() DatabaseConfig
	GetHTTPServerConfig() HTTPServerConfig
	GetJwtTokenConfig() JwtTokenConfig
	GetFileStorageConfig() FileStorageConfig
	GetSentryConfig() SentryConfig
	GetLogConfig() LogConfig
}

type LogConfig struct {
	Level     string `env:"LEVEL" envDefault:"info"`
	AddSource bool   `env:"ADD_SOURCE" envDefault:"false"`
}

type DatabaseConfig struct {
	Host            string `env:"HOST" envDefault:"localhost"`
	Port            string `env:"PORT" envDefault:"5432"`
	User            string `env:"USER" envDefault:"user"`
	Password        string `env:"PASSWORD" envDefault:""`
	Name            string `env:"NAME" envDefault:"db"`
	SslMode         string `env:"SSL" envDefault:"disable"`
	MaxConnections  int    `env:"MAX_CONNECTIONS" envDefault:"20"`
	MinConnections  int    `env:"MIN_CONNECTIONS" envDefault:"1"`
	ConnMaxLifetime int    `env:"CONN_MAX_LIFETIME" envDefault:"300"`
}

type HTTPServerConfig struct {
	Port         string `env:"PORT" envDefault:"8000"`
	ReadTimeout  int    `env:"READ_TIMEOUT" envDefault:"15"`
	WriteTimeout int    `env:"WRITE_TIMEOUT" envDefault:"15"`
	IdleTimeout  int    `env:"IDLE_TIMEOUT" envDefault:"60"`
	RateLimit    string `env:"RATE_LIMIT" envDefault:"100-S"`
	Environment  string `env:"ENVIRONMENT" envDefault:"development"`
}

type JwtTokenConfig struct {
	SecretKey        string `env:"SECRET_KEY" envDefault:"3891aSDk23aSDa3j#@sd"`
	SecretRefreshKey string `env:"REFRESH_SECRET_KEY" envDefault:"h3i12iaSD32u98da@#%aisd"`
}

type FileStorageConfig struct {
	BasePath             string `env:"PATH" envDefault:"./uploads"`
	MaxSize              int64  `env:"MAX_SIZE" envDefault:"5242880"`
	FilePermissions      uint32 `env:"FILE_PERMISSIONS" envDefault:"0644"`
	DirectoryPermissions uint32 `env:"DIRECTORY_PERMISSIONS" envDefault:"0755"`
}

type SentryConfig struct {
	DSN              string  `env:"DSN"`
	TracesSampleRate float64 `env:"TRACES_SAMPLE_RATE" envDefault:"1.0"`
}

func New() (Interface, error) {
	_ = godotenv.Load()

	cfg := &config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

func (c *config) GetLogConfig() LogConfig {
	return c.Log
}

func (c *config) GetDatabaseConfig() DatabaseConfig {
	return c.Database
}

func (c *config) GetHTTPServerConfig() HTTPServerConfig {
	return c.HTTPServer
}

func (c *config) GetJwtTokenConfig() JwtTokenConfig {
	return c.JwtToken
}

func (c *config) GetFileStorageConfig() FileStorageConfig {
	return c.FileStorage
}

func (c *config) GetSentryConfig() SentryConfig {
	return c.Sentry
}
