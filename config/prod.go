//go:build prod

package config

func GetTemplatePath(fileName string) string {
	return "/app/pkg/templates/" + fileName
}
