//go:build dev || test

package config

import (
	"os"
)

func (c *Config) GetTemplatePath(fileName string) string {
	return os.Getenv("GOPATH") + "/src/github.com/idthings/notify/pkg/templates/" + fileName
}
