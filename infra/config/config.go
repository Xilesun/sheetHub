package config

import (
	"sync/atomic"

	"github.com/Xilesun/sheethub/infra/constants"
	"github.com/adrg/xdg"
	"github.com/spf13/viper"
)

// AppConfig is the configuration for the application.
type AppConfig struct {
}

// DBConfig is the configuration for the database.
type DBConfig struct {
	DSN      string            `yaml:"dsn"`
	Dialect  constants.Dialect `yaml:"dialect"`
	Host     string            `yaml:"host"`
	Port     int               `yaml:"port"`
	Username string            `yaml:"username"`
	Password string            `yaml:"password"`
	Database string            `yaml:"database"`
}

// LogConfig is the configuration for the logger.
type LogConfig struct {
}

// Config is the configuration for the application.
type Config struct {
	App AppConfig `yaml:"app"`
	DB  DBConfig  `yaml:"db"`
	Log LogConfig `yaml:"log"`
}

var config *atomic.Value

func setDefaultConfig() error {
	dbFilePath, err := xdg.DataFile(constants.DefaultDBPath)
	if err != nil {
		return err
	}
	viper.SetDefault("db", map[string]interface{}{
		"dialect": constants.DialectSQLite,
		"dsn":     dbFilePath,
	})
	return nil
}

// reads the configuration.
func read() (*Config, error) {
	cfg := &Config{}
	config = new(atomic.Value)
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	if err := viper.Unmarshal(cfg); err != nil {
		return nil, err
	}
	config.Store(cfg)
	return config.Load().(*Config), nil
}

// Init initializes the configuration.
func Init() (*Config, error) {
	configFilePath, err := xdg.SearchConfigFile(constants.ConfigPath)
	if err == nil {
		viper.SetConfigFile(configFilePath)
		return read()
	}
	err = setDefaultConfig()
	if err != nil {
		return nil, err
	}
	configFilePath, err = xdg.ConfigFile(constants.AppName)
	if err != nil {
		return nil, err
	}
	err = viper.SafeWriteConfigAs(configFilePath)
	if err != nil {
		return nil, err
	}
	viper.SetConfigFile(configFilePath)
	return read()
}

// Set sets the configuration.
func Set(cfg *Config) {
	config.Store(cfg)
}

// Get returns the configuration.
func Get() *Config {
	return config.Load().(*Config)
}
