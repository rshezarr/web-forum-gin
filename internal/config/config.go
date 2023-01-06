package config

import (
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type (
	Config struct {
		API      API      `mapstructure:"api"`
		Database Database `mapstructure:"database"`
	}

	API struct {
		Host           string        `mapstructure:"host"`
		Addr           string        `mapstructure:"addr"`
		MaxHeaderBytes int           `mapstructure:"maxHeaderByte"`
		ReadTimeout    time.Duration `mapstructure:"readTimeout"`
		WriteTimeout   time.Duration `mapstructure:"writeTimeout"`
	}

	Database struct {
		Driver      string `mapstructure:"driver"`
		DBName      string `mapstructure:"dbname"`
		DatabaseURL string `mapstructure:"databaseUrl"`
		SchemePath  string `mapstructure:"schemesPath"`
	}
)

func NewConfig() (*Config, error) {
	logrus.Info("Configs are initializing...")

	var config *Config

	if err := viper.Unmarshal(&config); err != nil {
		return &Config{}, err
	}

	logrus.Info("Configs are initialized")

	return config, nil
}
