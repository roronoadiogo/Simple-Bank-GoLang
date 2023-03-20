package main

import (
	"github.com/roronoadiogo/Simple-Bank-GoLang/config"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	configEnvs := config.ConfigEnvDatabase()
	config.PerformDatabaseMigrations(logger, configEnvs)

}
