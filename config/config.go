/*
   A config system, with Go + Kubernetes in mind.

   Override the default using an env var (including envVarPrefix), such as:

       $ export MAGICLINK_SMTP_ENABLED=true

   In code we can safely call for the config value, which returns the env var value or the default:

   	    if cfg.ValueAsBool("SMTP_ENABLED") {
            // do stuff
        }
*/
package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type Config struct{}

// All of the env vars are prefixed with envVarPrefix, to try and avoid naming clashes.
const (
	envVarPrefix = "MAGICLINK_"
)

// We don't use the envVarPrefix internally, to stay portable.
var defaultValues = map[string]interface{}{
	"API_PORT":                     "8080",                                // api listens on this port
	"URL_PREFIX":                   "/magiclink/",                         // all inbound urls will be prefixed with this value
	"REDIS_HOST":                   "localhost",                           // redis host name
	"REDIS_PORT":                   "6379",                                // redis port
	"REDIS_KEY_PREFIX":             "magiclink:",                          // all redis keys will be prefixed with this value
	"MAGICLINK_URL":                "https://override.me/magiclink/auth/", // this is prepended to the magic link id
	"MAGICLINK_EXPIRES_MINS":       15,                                    // auto-expire magic link id's
	"MAGICLINK_LENGTH":             64,                                    // length of the id string
	"SESSION_NAME":                 "MagicLinkSession",                    // cookie session name after auth
	"SESSION_ID_LENGTH":            64,                                    // our cookie session id length
	"SESSION_EXPIRES_MINS":         10080,                                 // our cookie expires after this many minutes (1 week)
	"RATE_LIMIT_MAX_SEND_REQUESTS": 3,                                     // max requests for a magic link within TIME_PERIOD
	"RATE_LIMIT_TIME_PERIOD_MINS":  15,                                    // time period within which we rate limit
	"SESSION_OWNER_PROTECTED_URL":  "",                                    // we can expose a protected URL to return session owner
	"SESSION_OWNER_ACCESS_TOKENS":  "",                                    // session owner access tokens - off by default
	"SMTP_ENABLED":                 false,                                 // experimental
	"SMTP_HOST":                    "",                                    // experimental
	"SMTP_PORT":                    "25",                                  // experimental
	"SMTP_USER":                    "",                                    // experimental
	"SMTP_PASSWORD":                "",                                    // experimental
	"SMTP_ALLOWED_RECIPIENTS":      "",                                    // experimental
}

// Public methods here.
// Use typed methods so we avoid type assertions at point of use.
func (c *Config) ValueAsStr(key string) string {

	c.mapHasKey(key) // check key exists
	defaultValue := defaultValues[key].(string)
	return c.getEnvVar(key, defaultValue).(string)
}

func (c *Config) ValueAsInt(key string) int {

	c.mapHasKey(key) // check key exists
	defaultValue := defaultValues[key].(int)
	return c.getEnvVar(key, defaultValue).(int)
}

func (c *Config) ValueAsBool(key string) bool {

	c.mapHasKey(key) // check key exists
	defaultValue := defaultValues[key].(bool)
	return c.getEnvVar(key, defaultValue).(bool)
}

// Private methods here
func (c *Config) mapHasKey(key string) bool {
	if _, ok := defaultValues[key]; ok {
		return true
	}
	log.Fatal("Config.mapHasKey: trying to access non-existant config key: ", key)
	return false
}
func (c *Config) getEnvVar(key string, fallback interface{}) interface{} {

	fullEnvVarName := fmt.Sprintf("%s%s", envVarPrefix, key)
	value, exists := os.LookupEnv(fullEnvVarName)
	if !exists {
		return fallback
	}

	switch fallback.(type) {
	case bool:
		valueAsInt, err := strconv.ParseBool(value)
		if err != nil {
			return fallback
		}
		return valueAsInt
	case int:
		valueAsInt, err := strconv.Atoi(value)
		if err != nil {
			return fallback
		}
		return valueAsInt
	case string:
		return value
	}
	return fallback
}
