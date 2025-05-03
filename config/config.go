package config

import (
	"log/slog"
	"sync"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

var (
	once sync.Once
	cfg  *Config
)

type Config struct {
	Environment string            `env:"ENVIRONMENT" envDefault:"development"`
	Log         LogConfig         `envPrefix:"LOG_"`
	Database    DatabaseConfig    `envPrefix:"DB_"`
	HttpServer  HttpServerConfig  `envPrefix:"HTTP_SERVER_"`
	JwtToken    JwtTokenConfig    `envPrefix:"JWT_"`
	FileStorage FileStorageConfig `envPrefix:"FILE_STORAGE_"`
	Sentry      SentryConfig      `envPrefix:"SENTRY_"`
}

type ConfigInterface interface {
	GetEnvironment() string
	GetDatabaseConfig() DatabaseConfig
	GetHttpServerConfig() HttpServerConfig
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

type HttpServerConfig struct {
	Port         string `env:"PORT" envDefault:"8000"`
	ReadTimeout  int    `env:"READ_TIMEOUT" envDefault:"15"`
	WriteTimeout int    `env:"WRITE_TIMEOUT" envDefault:"15"`
	IdleTimeout  int    `env:"IDLE_TIMEOUT" envDefault:"60"`
	RateLimit    string `env:"RATE_LIMIT" envDefault:"100-S"`
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

func New() ConfigInterface {
	once.Do(func() {
		_ = godotenv.Load()

		cfg = &Config{}
		if err := env.Parse(cfg); err != nil {
			slog.Error("error parsing config", "error", err)
		}
	})
	return cfg
}

func (c *Config) GetEnvironment() string {
	return c.Environment
}

func (c *Config) GetLogConfig() LogConfig {
	return c.Log
}

func (c *Config) GetDatabaseConfig() DatabaseConfig {
	return c.Database
}

func (c *Config) GetHttpServerConfig() HttpServerConfig {
	return c.HttpServer
}

func (c *Config) GetJwtTokenConfig() JwtTokenConfig {
	return c.JwtToken
}

func (c *Config) GetFileStorageConfig() FileStorageConfig {
	return c.FileStorage
}

func (c *Config) GetSentryConfig() SentryConfig {
	return c.Sentry
}
