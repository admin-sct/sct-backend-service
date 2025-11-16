package config

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// Config holds application configuration
type Config struct {
	Server ServerConfig
	DB     DBConfig
	Log    LogConfig
}

// ServerConfig holds server configuration
type ServerConfig struct {
	Port     int
	Host     string
	LogLevel string
}

// DBConfig holds database configuration
type DBConfig struct {
	Host     string
	Port     int
	Name     string
	User     string
	Password string
}

// LogConfig holds logging configuration
type LogConfig struct {
	Level  string
	Format string
}

// ConfigFxOption provides configuration via fx
func ConfigFxOption(configFilePath string) fx.Option {
	return fx.Provide(func() (*Config, error) {
		// TODO: Load configuration from file
		// For now, return default configuration
		return &Config{
			Server: ServerConfig{
				Port:     8080,
				Host:     "0.0.0.0",
				LogLevel: "info",
			},
			DB: DBConfig{
				Host: "localhost",
				Port: 5432,
				Name: "sct_db",
				User: "postgres",
			},
			Log: LogConfig{
				Level:  "info",
				Format: "json",
			},
		}, nil
	})
}

// LoggerFxOption provides logger via fx
func LoggerFxOption() fx.Option {
	return fx.Provide(func(config *Config) (*zap.Logger, error) {
		var logger *zap.Logger
		var err error

		if config.Log.Level == "debug" {
			logger, err = zap.NewDevelopment()
		} else {
			logger, err = zap.NewProduction()
		}

		if err != nil {
			return nil, err
		}

		return logger, nil
	})
}

