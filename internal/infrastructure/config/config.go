package config

import (
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

var (
	once sync.Once
)

type config struct {
	Database   DatabaseConfig
	HttpServer HttpServerConfig
	JwtToken   JwtTokenConfig
}

type ConfigInterface interface {
	GetDatabaseConfig() DatabaseConfig
	GetHttpServerConfig() HttpServerConfig
	GetJwtTokenConfig() JwtTokenConfig
}

type DatabaseConfig struct {
	Host            string
	Port            string
	User            string
	Password        string
	Name            string
	SslMode         string
	MaxConnections  int
	MinConnections  int
	ConnMaxLifetime int
}

type HttpServerConfig struct {
	Port         string
	ReadTimeout  int
	WriteTimeout int
	IdleTimeout  int
	Environment  string
}

type JwtTokenConfig struct {
	SecretKey        string
	SecretRefreshKey string
	Environment      string
}

func New() ConfigInterface {
	var cfg *config
	once.Do(func() {
		_ = godotenv.Load()

		cfg = &config{
			Database: DatabaseConfig{
				Host:            getEnv("DB_HOST", "localhost"),
				Port:            getEnv("DB_PORT", "5432"),
				User:            getEnv("DB_USER", "user"),
				Password:        getEnv("DB_PASSWORD", ""),
				Name:            getEnv("DB_NAME", "db"),
				SslMode:         getEnv("DB_SSL", "disable"),
				MaxConnections:  getEnvInt("DB_MAX_CONNECTIONS", 20),
				MinConnections:  getEnvInt("DB_MIN_CONNECTIONS", 1),
				ConnMaxLifetime: getEnvInt("DB_CONN_MAX_LIFETIME", 300),
			},
			HttpServer: HttpServerConfig{
				Port:         getEnv("HTTP_SERVER_PORT", "8000"),
				ReadTimeout:  getEnvInt("HTTP_SERVER_READ_TIMEOUT", 15),
				WriteTimeout: getEnvInt("HTTP_SERVER_WRITE_TIMEOUT", 15),
				IdleTimeout:  getEnvInt("HTTP_SERVER_IDLE_TIMEOUT", 60),
				Environment:  getEnv("ENVIRONMENT", "development"),
			},
			JwtToken: JwtTokenConfig{
				SecretKey:        getEnv("JWT_SECRET_KEY", "3891aSDk23aSDa3j#@sd"),
				SecretRefreshKey: getEnv("JWT_REFRESH_SECRET_KEY", "h3i12iaSD32u98da@#%aisd"),
				Environment:      getEnv("ENVIRONMENT", "development"),
			},
		}
	})

	return cfg
}

func (c *config) GetDatabaseConfig() DatabaseConfig {
	return c.Database
}

func (c *config) GetHttpServerConfig() HttpServerConfig {
	return c.HttpServer
}

func (c *config) GetJwtTokenConfig() JwtTokenConfig {
	return c.JwtToken
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		if defaultValue != "" {
			return defaultValue
		}
	}
	return value
}

func getEnvInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	parsedValue, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return parsedValue
}
