package config_test

import (
	"os"
	"testing"

	"github.com/Wilo/climate-change/config"
	"gotest.tools/assert"
)

// fixtures
var (
	GITHUB_API_KEY  = "123456"
	WEATHER_API_KEY = "123456"
)

func TestGetEnv(t *testing.T) {

	t.Run("When env vars doesn't loaded and don't set default vars", func(t *testing.T) {
		_, err := config.GetEnv("GITHUB_API_KEY", "")
		assert.Error(t, err, "Github Api Key doesn't found!")
		_, err = config.GetEnv("WEATHER_API_KEY", "")
		assert.Error(t, err, "Weather Api Key doesn't found!")
	})

	t.Run("When env vars doesn't loaded, but use default vars", func(t *testing.T) {
		gak, err := config.GetEnv("GITHUB_API_KEY", GITHUB_API_KEY)
		assert.NilError(t, err)
		assert.Equal(t, gak, GITHUB_API_KEY)
		wak, err := config.GetEnv("WEATHER_API_KEY", WEATHER_API_KEY)
		assert.NilError(t, err)
		assert.Equal(t, wak, WEATHER_API_KEY)
	})

	t.Run("When the env vars doesn't exist", func(t *testing.T) {
		_, err := config.GetEnv("AWS_API_KEY", "")
		assert.Error(t, err, "Unknow")
	})

	t.Run("when the envs are loaded into the OS", func(t *testing.T) {

		// Put the variables into the system
		os.Setenv("GITHUB_API_KEY", GITHUB_API_KEY)
		os.Setenv("WEATHER_API_KEY", WEATHER_API_KEY)

		gak, err := config.GetEnv("GITHUB_API_KEY", "")
		assert.NilError(t, err)
		assert.Equal(t, gak, GITHUB_API_KEY)
		wak, err := config.GetEnv("WEATHER_API_KEY", "")
		assert.NilError(t, err)
		assert.Equal(t, wak, WEATHER_API_KEY)
	})
}

func BenchmarkGetEnv(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = config.GetEnv("GITHUB_API_KEY", "")
		_, _ = config.GetEnv("WEATHER_API_KEY", "")
	}
}
