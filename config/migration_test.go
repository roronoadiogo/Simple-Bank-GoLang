package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestPerformDatabaseMigrations(t *testing.T) {
	// Mock logger
	logger := zap.NewNop()

	// Mock configDB
	configDB := ConfigParamsDatabase{
		DBHost: "localhost",
		DBPort: 5432,
		DBUser: "user",
		DBPass: "pass",
		DBName: "testdb",
	}

	t.Run("should return nil if migrations are already up to date", func(t *testing.T) {
		err := PerformDatabaseMigrations(logger, configDB)
		assert.NoError(t, err)
	})

	t.Run("should return an error if logger is nil", func(t *testing.T) {
		err := PerformDatabaseMigrations(nil, configDB)
		assert.EqualError(t, err, "logger cannot be nil")
	})

	t.Run("should return an error if database connection fails", func(t *testing.T) {
		configDB := ConfigParamsDatabase{
			DBHost: "invalidhost",
			DBPort: 5432,
			DBUser: "user",
			DBPass: "pass",
			DBName: "testdb",
		}
		err := PerformDatabaseMigrations(logger, configDB)
		assert.Error(t, err)
	})

	// Note: In the following tests, we assume that the migrations have not been applied yet
	// and will try to apply them, which may take some time depending on the size of the migrations.

	t.Run("should apply database migrations successfully", func(t *testing.T) {
		err := PerformDatabaseMigrations(logger, configDB)
		assert.NoError(t, err)
	})

}
