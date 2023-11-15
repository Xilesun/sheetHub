package config

import (
	"sync/atomic"

	"github.com/Xilesun/sheethub/server/infra/constants"
	"github.com/Xilesun/sheethub/server/infra/logger"
	"github.com/spf13/viper"
)

// AppConfig is the configuration for the application.
type AppConfig struct {
	Env  string `yaml:"env" env:"ENV"`
	Port int    `yaml:"port" env:"PORT"`
}

// DBConfig is the configuration for the database.
type DBConfig struct {
	DSN      string            `yaml:"dsn" env:"DSN"`
	Dialect  constants.Dialect `yaml:"dialect" env:"DIALECT"`
	Host     string            `yaml:"host" env:"HOST"`
	Port     int               `yaml:"port" env:"PORT"`
	Username string            `yaml:"username" env:"USERNAME"`
	Password string            `yaml:"password" env:"PASSWORD"`
	Database string            `yaml:"database" env:"DATABASE"`
}

// LogConfig is the configuration for the logger.
type LogConfig struct {
}

// Config is the configuration for the application.
type Config struct {
	App     AppConfig `yaml:"app" envPrefix:"APP_"`
	DB      DBConfig  `yaml:"db" envPrefix:"DB_"`
	Log     LogConfig `yaml:"log" envPrefix:"LOG_"`
	Storage string    `yaml:"storage" env:"STORAGE"`
}

var config *atomic.Value

// Init initializes viper.
func Init(file string) {
	if file != "" {
		viper.SetConfigFile(file)
		return
	}
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
}

// Read reads the configuration.
func Read() (*Config, error) {
	cfg := &Config{}
	config = new(atomic.Value)
	if err := viper.ReadInConfig(); err != nil {
		logger.Errorf("Failed to read config file: %v", err)
		return nil, err
	}
	if err := viper.Unmarshal(cfg); err != nil {
		logger.Errorf("Unable to decode config, %v", err)
		return nil, err
	}
	config.Store(cfg)
	return config.Load().(*Config), nil
}

// Get returns the configuration.
func Get() *Config {
	return config.Load().(*Config)
}
