package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct{}

// All of the env vars are prefixed with envVarPrefix, to try and avoid naming clashes.
const (
	envVarPrefix = "MAGICLINK_"
)

// Default values can be override with env vars, eg 'export MAGICLINK_API_PORT=80'
// We don't use the envVarPrefix internally, to stay portable.
var defaultValues = map[string]interface{}{
	"API_PORT":                     "8080",             // api listens on this port
	"URL_PREFIX":                   "/magiclink/",      // all inbound urls will be prefixed with this value
	"REDIS_HOST":                   "localhost",        // redis host name
	"REDIS_PORT":                   "6379",             // redis port
	"REDIS_KEY_PREFIX":             "magiclink:",       // all redis keys will be prefixed with this value
	"MAGICLINK_EXPIRES_MINS":       15,                 // auto-expire magic link id's
	"MAGICLINK_LENGTH":             64,                 // length of the id string
	"SESSION_NAME":                 "MagicLinkSession", // cookie session name after auth
	"SESSION_ID_LENGTH":            64,                 // our cookie session id length
	"SESSION_EXPIRES_MINS":         10080,              // our cookie expires after this many minutes (1 week)
	"RATE_LIMIT_MAX_SEND_REQUESTS": 3,                  // max requests for a magic link within TIME_PERIOD
	"RATE_LIMIT_TIME_PERIOD_MINS":  15,                 // time period within which we rate limit
}

// Public methods here.
// Use typed methods so we avoid type assertions at point of use.
func (c *Config) ValueAsStr(key string) string {

	defaultValue := defaultValues[key].(string)
	return c.getEnvVar(key, defaultValue).(string)
}

func (c *Config) ValueAsInt(key string) int {

	defaultValue := defaultValues[key].(int)
	return c.getEnvVar(key, defaultValue).(int)
}

// Private methods here
func (c *Config) getEnvVar(key string, fallback interface{}) interface{} {

	fullEnvVarName := fmt.Sprintf("%s%s", envVarPrefix, key)
	value, exists := os.LookupEnv(fullEnvVarName)
	if !exists {
		return fallback
	}

	switch fallback.(type) {
	case string:
		return value
	case int:
		valueAsInt, err := strconv.Atoi(value)
		if err != nil {
			return fallback
		}
		return valueAsInt
	}
	return fallback
}
