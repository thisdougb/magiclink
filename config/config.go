package config

import (
	"fmt"
	"os"
	"strconv"
)

const (
	envVarPrefix = "MAGICLINK_"
)

type Config struct{}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) getEnv(key, fallback string) string {

	fullEnvVarName := fmt.Sprintf("%s%s", envVarPrefix, key)
	value, exists := os.LookupEnv(fullEnvVarName)
	if !exists {
		value = fallback
	}
	return value
}
func (c *Config) getEnvAsInt(key string, fallback int) int {

	// we use fallback as str so we don't lost its value while
	// using the central getEnv() method
	fallbackAsStr := strconv.Itoa(fallback)

	value := c.getEnv(key, fallbackAsStr)
	valueAsInt, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}
	return valueAsInt
}

func (c *Config) API_PORT() string {
	return c.getEnv("API_PORT", "8080")
}
func (c *Config) REDIS_HOST() string {
	return c.getEnv("REDIS_HOST", "127.0.0.1")
}
func (c *Config) REDIS_PORT() string {
	return c.getEnv("REDIS_PORT", "6379")
}
func (c *Config) REDIS_KEY_PREFIX() string {
	return c.getEnv("REDIS_KEY_PREFIX", "magiclink:")
}

func (c *Config) MAGICLINK_LENGTH() int {
	return c.getEnvAsInt("MAGICLINK_LENGTH", 64)
}
func (c *Config) MAGICLINK_EXPIRES_MINS() int {
	return c.getEnvAsInt("MAGICLINK_EXPIRES_MINS", 15)
}

func (c *Config) SESSION_NAME() string {
	return c.getEnv("SESSION_NAME", "MagicLinkSession")
}
func (c *Config) SESSION_ID_LENGTH() int {
	return c.getEnvAsInt("SESSION_ID_LENGTH", 64)
}
func (c *Config) SESSION_EXPIRES_MINS() int {
	return c.getEnvAsInt("SESSION_EXPIRES_MINS", 10080) // 1 week
}
