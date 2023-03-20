package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConfigEnvDatabase(t *testing.T) {
	os.Setenv("APP_ENV", "test")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_USER", "testuser")
	os.Setenv("DB_PASS", "testpass")
	os.Setenv("DB_NAME", "testdb")

	configParams := ConfigEnvDatabase()

	expectedParams := ConfigParamsDatabase{
		DBHost: "localhost",
		DBPort: 5432,
		DBUser: "testuser",
		DBPass: "testpass",
		DBName: "testdb",
	}

	t.Run("Verify the config database setting right", func(t *testing.T) {
		require.NotEmpty(t, configParams)
		require.Equal(t, configParams, expectedParams)
	})

}
