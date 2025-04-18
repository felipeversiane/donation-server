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
}

type ConfigInterface interface {
	GetDatabaseConfig() DatabaseConfig
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
		}
	})

	return cfg
}

func (c *config) GetDatabaseConfig() DatabaseConfig {
	return c.Database
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
