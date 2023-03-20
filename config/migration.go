package config

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func PerformDatabaseMigrations(logger *zap.Logger, configDB ConfigParamsDatabase) error {
	if logger == nil {
		return errors.New("logger cannot be nil")
	}

	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		configDB.DBHost, configDB.DBPort, configDB.DBUser, configDB.DBPass, configDB.DBName))

	if err != nil {
		logger.Error("failed to open database connection", zap.Error(err))
		return err
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		logger.Error("failed to create Postgres driver instance", zap.Error(err))
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		configDB.DBName,
		driver,
	)
	if err != nil {
		logger.Error("failed to create migrations instance", zap.Error(err))
		return err
	}

	if err != migrate.ErrNoChange {
		logger.Sugar().Warn("Not applied migrations, database don't have changes.")
		m.Close()

		return nil
	}

	if err := m.Up(); err != nil {
		logger.Error("failed to apply database migrations", zap.Error(err))
		return err
	}

	if err, _ := m.Close(); err != nil {
		logger.Error("failed to close migrations instance", zap.Error(err))
		return err
	}

	return nil
}
