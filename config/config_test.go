// +build dev test

package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestValueAsStr(t *testing.T) {

	var cfg *Config // dynamic config settings

	// test: no env var, should return default
	os.Unsetenv("MAGICLINK_API_PORT")
	assert.Equal(t, "8080", cfg.ValueAsStr("API_PORT"), "no env var set")

	// test: setting an env var override
	os.Setenv("MAGICLINK_API_PORT", "hello")
	assert.Equal(t, "hello", cfg.ValueAsStr("API_PORT"), "env var set")
	os.Unsetenv("MAGICLINK_API_PORT")
}

func TestValueAsInt(t *testing.T) {

	var cfg *Config // dynamic config settings

	// test: no env var, should return default
	os.Unsetenv("MAGICLINK_MAGICLINK_LENGTH")
	assert.Equal(t, 64, cfg.ValueAsInt("MAGICLINK_LENGTH"), "no env var set")

	// test: setting an env var override
	os.Setenv("MAGICLINK_MAGICLINK_LENGTH", "32")
	assert.Equal(t, 32, cfg.ValueAsInt("MAGICLINK_LENGTH"), "env var set")
	os.Unsetenv("MAGICLINK_MAGICLINK_LENGTH")

	// test: setting an env var override to non-int
	os.Setenv("MAGICLINK_MAGICLINK_LENGTH", ";")
	assert.Equal(t, 64, cfg.ValueAsInt("MAGICLINK_LENGTH"), "env var not int")
	os.Unsetenv("MAGICLINK_MAGICLINK_LENGTH")
}

func TestValueAsBool(t *testing.T) {

	var cfg *Config // dynamic config settings

	// test: no env var, should return default
	os.Unsetenv("MAGICLINK_SMTP_ENABLED")
	assert.Equal(t, false, cfg.ValueAsBool("SMTP_ENABLED"), "no env var set")

	// test: setting an env var override
	os.Setenv("MAGICLINK_SMTP_ENABLED", "true")
	assert.Equal(t, true, cfg.ValueAsBool("SMTP_ENABLED"), "env var set")
	os.Unsetenv("MAGICLINK_SMTP_ENABLED")

	// test: setting an env var override to non-int
	os.Setenv("MAGICLINK_SMTP_ENABLED", ";")
	assert.Equal(t, false, cfg.ValueAsBool("SMTP_ENABLED"), "env var not int")
	os.Unsetenv("MAGICLINK_SMTP_ENABLED")
}

func TestGetEnvVar(t *testing.T) {

	var cfg *Config // dynamic config settings

	// test: no env var, should return default
	os.Setenv("MAGICLINK_TEST_STR", "success")
	assert.Equal(t, "success", cfg.getEnvVar("TEST_STR", "success"), "TEST_STR")
	os.Unsetenv("MAGICLINK_TEST_STR")

	os.Setenv("MAGICLINKTEST_INT", "32")
	assert.Equal(t, 1, cfg.getEnvVar("TEST_INT", 1), "TEST_INT")
	os.Unsetenv("MAGICLINKTEST_INT")

	os.Setenv("MAGICLINK_TEST_UNKNOWN", "2.2")
	assert.Equal(t, 1.1, cfg.getEnvVar("TEST_UNKNOWN", 1.1), "TEST_UNKNOWN")
	os.Unsetenv("MAGICLINK_TEST_UNKNOWN")
}
