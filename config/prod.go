//go:build prod

package config

func (c *Config) GetTemplatePath(fileName string) string {
	return "/app/pkg/templates/" + fileName
}
