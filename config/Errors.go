package config

import (
	"errors"
)

type LoadEnvVarsError struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	Message string `json:"message"`
}

var (
	ErrGithubApiKeyEnvVarsDoesnotLoaded  = LoadEnvVarsError{Name: "GITHUB_API_KEY", Type: "config_error", Message: "Github Api Key doesn't found!"}
	ErrWeatherApiKeyEnvVarsDoesnotLoaded = LoadEnvVarsError{Name: "WEATHER_API_KEY", Type: "config_error", Message: "Weather Api Key doesn't found!"}
)

func EnvError(env string) error {
	switch env {
	case ErrGithubApiKeyEnvVarsDoesnotLoaded.Name:
		return errors.New(ErrGithubApiKeyEnvVarsDoesnotLoaded.Message)

	case ErrWeatherApiKeyEnvVarsDoesnotLoaded.Name:
		return errors.New(ErrWeatherApiKeyEnvVarsDoesnotLoaded.Message)
	}
	return errors.New("Unknow")
}
