package config

import (
	"flag"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type (
	Configuration struct {
		API      API      `yaml:"api"`
		Database Database `yaml:"database"`
	}

	API struct {
		Host           string        `yaml:"host"`
		Addr           string        `yaml:"addr"`
		MaxHeaderBytes int           `yaml:"maxHeaderBytes"`
		ReadTimeout    time.Duration `yaml:"readTimeout"`
		WriteTimeout   time.Duration `yaml:"writeTimeout"`
		IdleTimeout    time.Duration `yaml:"idleTimeout"`
	}

	Database struct {
		DatabaseURL string        `yaml:"databaseUrl"`
		Driver      string        `yaml:"driver"`
		DBName      string        `yaml:"dbname"`
		SchemesPath string        `yaml:"schemesPath"`
		CtxTimeout  time.Duration `yaml:"ctxTimeout"`
		IdleTimeout time.Duration `yaml:"idleTimeout"`
	}
)

func initConfig() error {
	var configPath = flag.String("config-path", "configs/", "path to config file")

	flag.Parse()

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(*configPath)

	return viper.ReadInConfig()
}

func NewConfig() (*Configuration, error) {
	config := new(Configuration)

	if err := initConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return &Configuration{}, err
	}

	logrus.Info("Configs are initialized")

	return config, nil
}
