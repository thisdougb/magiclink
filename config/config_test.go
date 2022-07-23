//go:build dev || test
// +build dev test

package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

// adds our test values
func init() {
	defaultValues["_TEST_INT_TEN"] = 10
	defaultValues["_TEST_STR_AAA"] = "AAA"
	defaultValues["_TEST_BOOL_TLS"] = false
}

func TestValueAsStr(t *testing.T) {

	var cfg *Config // dynamic config settings

	// test: unset potential env var, this should return the Str value in defaultValue[map]
	os.Unsetenv(envVarPrefix + "_TEST_STR_AAA")
	assert.Equal(t, "AAA", cfg.ValueAsStr("_TEST_STR_AAA"), "no env var set")

	// test: now override the defaultValue[map] using an env var value
	os.Setenv(envVarPrefix+"_TEST_STR_AAA", "hello")
	assert.Equal(t, "hello", cfg.ValueAsStr("_TEST_STR_AAA"), "env var set")
	os.Unsetenv(envVarPrefix + "_TEST_STR_AAA")
}

func TestValueAsInt(t *testing.T) {

	var cfg *Config // dynamic config settings

	// test: unset potential env var, this should return the int value in defaultValue[map]
	os.Unsetenv(envVarPrefix + "_TEST_INT_TEN")
	assert.Equal(t, 10, cfg.ValueAsInt("_TEST_INT_TEN"), "no env var set")

	// test: now override the defaultValue[map] using an env var value
	os.Setenv(envVarPrefix+"_TEST_INT_TEN", "20")
	assert.Equal(t, 20, cfg.ValueAsInt("_TEST_INT_TEN"), "env var set")
	os.Unsetenv(envVarPrefix + "_TEST_INT_TEN")

	// test: now we use a non-int env var, which should be ignored
	os.Setenv(envVarPrefix+"_TEST_INT_TEN", ";")
	assert.Equal(t, 10, cfg.ValueAsInt("_TEST_INT_TEN"), "env var not int")
	os.Unsetenv(envVarPrefix + "_TEST_INT_TEN")

}

func TestValueAsBool(t *testing.T) {

	var cfg *Config // dynamic config settings

	// test: unset potential env var, this should return the int value in defaultValue[map]
	os.Unsetenv(envVarPrefix + "_TEST_BOOL_TLS")
	assert.Equal(t, false, cfg.ValueAsBool("_TEST_BOOL_TLS"), "no env var set")

	// test: now override the defaultValue[map] using an env var value
	os.Setenv(envVarPrefix+"_TEST_BOOL_TLS", "true")
	assert.Equal(t, true, cfg.ValueAsBool("_TEST_BOOL_TLS"), "env var set")
	os.Unsetenv(envVarPrefix + "_TEST_BOOL_TLS")

	// test: now we use a non-int env var, which should be ignored
	os.Setenv(envVarPrefix+"_TEST_BOOL_TLS", "hello")
	assert.Equal(t, false, cfg.ValueAsBool("_TEST_BOOL_TLS"), "env var not int")
	os.Unsetenv(envVarPrefix + "_TEST_BOOL_TLS")
}
func TestGetEnvVar(t *testing.T) {

	var cfg *Config // dynamic config settings

	// test: when we set a str env var, should should get that value
	os.Setenv(envVarPrefix+"_TEST_STR_ISSET", "isset")
	assert.Equal(t, "isset", cfg.getEnvVar("_TEST_STR_ISSET", "isset"), "_TEST_STR_ISSET")
	os.Unsetenv(envVarPrefix + "_TEST_STR_ISSET")

	// test: when no env var exists we should use the fallback value in 2nd arg
	assert.Equal(t, "fallback", cfg.getEnvVar("_TEST_STR_ISSET", "fallback"), "_TEST_STR_ISSET")

	// test: when we set an int env var, should should get that value
	os.Setenv(envVarPrefix+"_TEST_INT", "32")
	assert.Equal(t, 32, cfg.getEnvVar("_TEST_INT", 1), "_TEST_INT")
	os.Unsetenv(envVarPrefix + "_TEST_INT")

	// test: when we set a non-int env var, should should get the fallback value
	// this is because we don't convert non-int automatically in getEnvVar
	os.Setenv(envVarPrefix+"TEST_UNKNOWN", "2.2")
	assert.Equal(t, 1.1, cfg.getEnvVar("TEST_UNKNOWN", 1.1), "TEST_UNKNOWN")
	os.Unsetenv(envVarPrefix + "TEST_UNKNOWN")
}
