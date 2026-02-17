package configs

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	// Set required env vars
	os.Setenv("LOGGER_STDOUT", "true")
	os.Setenv("LOGGER_FILE_LOCATION", "./logs")
	os.Setenv("LOGGER_FILE_MAX_AGE", "7")
	os.Setenv("LOGGER_LEVEL", "2")
	os.Setenv("LOGGER_MASKING", "true")
	os.Setenv("LOGGER_MASKING_PARAMS", "password,token")

	t.Run("LoadDefault", func(t *testing.T) {
		cfg := LoadDefault()
		assert.NotNil(t, cfg)
		assert.True(t, cfg.Logger.Stdout)
		assert.Equal(t, "./logs", cfg.Logger.FileLocation)
		assert.Equal(t, 7, cfg.Logger.FileMaxAge)
		assert.Equal(t, int8(2), cfg.Logger.Level)
		assert.True(t, cfg.Logger.Masking)
		assert.Equal(t, []string{"password", "token"}, cfg.Logger.MaskingParams)
	})

	t.Run("LoadTest", func(t *testing.T) {
		cfg := LoadTest()
		assert.NotNil(t, cfg)
		assert.Equal(t, 7, cfg.Logger.FileMaxAge)
	})
}
