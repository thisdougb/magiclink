//go:build dev || test

package config

import (
	"os"
)

func GetTemplatePath(fileName string) string {
	return os.Getenv("GOPATH") + "/src/github.com/idthings/notify/pkg/templates/" + fileName
}
