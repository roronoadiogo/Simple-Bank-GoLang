package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

type ConfigParamsDatabase struct {
	DBHost string
	DBPort int
	DBUser string
	DBPass string
	DBName string
}

var logger *zap.Logger

func init() {
	var err error
	logger, err = zap.NewProduction()
	if err != nil {
		panic(err)
	}
}

// Return the ENV that should be set in the APP_ENV, the Default is development environment
func configEnv() {

	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev"
	}

	err := godotenv.Load(fmt.Sprintf("config/.env.%s", env))
	if err != nil {
		logger.Error("Error loading .env file", zap.Error(err))
	}
}

// Return main ENVS for DatabaseConfiguration
func ConfigEnvDatabase() ConfigParamsDatabase {

	configEnv()

	databaseParams := ConfigParamsDatabase{

		DBHost: os.Getenv("DB_HOST"),
		DBPort: convertPortDatabase(os.Getenv("DB_PORT")),
		DBUser: os.Getenv("DB_USER"),
		DBPass: os.Getenv("DB_PASS"),
		DBName: os.Getenv("DB_NAME"),
	}

	return databaseParams
}

func convertPortDatabase(portInput string) int {
	port, err := strconv.Atoi(portInput)
	if err != nil {
		logger.Error("Error in set env params for database")
		return port
	}

	return port
}
