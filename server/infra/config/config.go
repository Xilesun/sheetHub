package config

import (
	"os"
	"sync/atomic"

	"github.com/Xilesun/sheethub/server/infra/constants"
	"github.com/caarlos0/env/v9"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
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
	App AppConfig `yaml:"app" envPrefix:"APP_"`
	DB  DBConfig  `yaml:"db" envPrefix:"DB_"`
	Log LogConfig `yaml:"log" envPrefix:"LOG_"`
}

var config *atomic.Value

func readFromEnv() error {
	cfg := &Config{}
	var envMp map[string]string
	envMp, err := godotenv.Read()
	if err != nil {
		return err
	}
	opts := env.Options{
		Environment: envMp,
	}
	if err := env.ParseWithOptions(cfg, opts); err != nil {
		return err
	}
	config.Store(cfg)
	return nil
}

func readFromYAML(file string) error {
	cfg := &Config{}
	conf, err := os.ReadFile(file)
	if err != nil {
		return readFromEnv()
	}
	if err := yaml.Unmarshal(conf, cfg); err != nil {
		return readFromEnv()
	}
	config.Store(cfg)
	return nil
}

// Read reads the configuration from config YAML file or .env file or environment variables.
func Read(file string) (*Config, error) {
	config = new(atomic.Value)
	if file == "" {
		err := readFromEnv()
		if err != nil {
			return nil, err
		}
	} else {
		err := readFromYAML(file)
		if err != nil {
			return nil, err
		}
	}
	return config.Load().(*Config), nil
}

// Get returns the configuration.
func Get() *Config {
	return config.Load().(*Config)
}
