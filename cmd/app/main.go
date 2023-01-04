package main

import (
	"flag"
	"forum/internal/app"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		logrus.Fatal(err)
	}

	app.Run()
}

func initConfig() error {
	var configPath = flag.String("config-path", "configs/", "path to config file")

	flag.Parse()

	viper.SetConfigName("forum")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(*configPath)

	return viper.ReadInConfig()
}
