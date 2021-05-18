package config

import (
	"os"
)

func GetEnv(name, defaultvalue string) (string, error) {
	env := os.Getenv(name)
	if len(env) != 0 {
		return env, nil
	}
	if defaultvalue != "" {
		return defaultvalue, nil
	}
	return "", EnvError(name)
}
