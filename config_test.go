package snek_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ronelliott/snek"
)

func TestNewConfig(t *testing.T) {
	called := false
	cfg := snek.NewConfig(func(*snek.Config) {
		called = true
	})
	assert.NotNil(t, cfg, "NewConfig should return a Config")
	assert.True(t, called, "NewConfig should call the configurator function")
}

func TestWithDefaultLogFormat(t *testing.T) {
	cfg := snek.NewConfig()
	require.NotNil(t, cfg, "NewConfig should return a Config")
	require.Equal(t, "formatted", cfg.DefaultLogFormat, "The default log format should be formatted")
	cfg = snek.NewConfig(snek.WithDefaultLogFormat("json"))
	require.NotNil(t, cfg, "NewConfig should return a Config")
	assert.Equal(t, "json", cfg.DefaultLogFormat, "DefaultLogFormat should be set to json")
}

func TestWithDefaultLogLevel(t *testing.T) {
	cfg := snek.NewConfig()
	require.NotNil(t, cfg, "NewConfig should return a Config")
	require.Equal(t, "info", cfg.DefaultLogLevel, "The default log level should be info")
	cfg = snek.NewConfig(snek.WithDefaultLogLevel("debug"))
	require.NotNil(t, cfg, "NewConfig should return a Config")
	assert.Equal(t, "debug", cfg.DefaultLogLevel, "LogLevel should be debug")
}

func TestWithEnvironmentVariablePrefix(t *testing.T) {
	cfg := snek.NewConfig()
	require.NotNil(t, cfg, "NewConfig should return a Config")
	require.Equal(t, "", cfg.EnvironmentVariablePrefix, "The default environment variable prefix should be empty")
	cfg = snek.NewConfig(snek.WithEnvironmentVariablePrefix("TEST_"))
	require.NotNil(t, cfg, "NewConfig should return a Config")
	assert.Equal(t, "TEST_", cfg.EnvironmentVariablePrefix, "EnvironmentVariablePrefix should be TEST_")
}

func TestWithLogFormatCommandLineVariableHelp(t *testing.T) {
	cfg := snek.NewConfig()
	require.NotNil(t, cfg, "NewConfig should return a Config")
	require.NotEqual(t, "The log format.", cfg.LogFormatCommandLineVariableHelp,
		"The default log format command line variable help should not be The log format.")
	cfg = snek.NewConfig(snek.WithLogFormatCommandLineVariableHelp("The log format."))
	require.NotNil(t, cfg, "NewConfig should return a Config")
	assert.Equal(t, "The log format.", cfg.LogFormatCommandLineVariableHelp,
		"LogFormatCommandLineVariableHelp should be The log format.")
}

func TestWithLogFormatCommandLineVariableLongName(t *testing.T) {
	cfg := snek.NewConfig()
	require.NotNil(t, cfg, "NewConfig should return a Config")
	require.Equal(t, "log-format", cfg.LogFormatCommandLineVariableLongName,
		"The default log format command line variable long name should be log-format")
	cfg = snek.NewConfig(snek.WithLogFormatCommandLineVariableLongName("logformat"))
	require.NotNil(t, cfg, "NewConfig should return a Config")
	assert.Equal(t, "logformat", cfg.LogFormatCommandLineVariableLongName,
		"LogFormatCommandLineVariableLongName should be logformat")
}

func TestWithLogFormatCommandLineVariableShortName(t *testing.T) {
	cfg := snek.NewConfig()
	require.NotNil(t, cfg, "NewConfig should return a Config")
	require.Equal(t, "", cfg.LogFormatCommandLineVariableShortName,
		"LogFormatCommandLineVariableShortName should be f")
	cfg = snek.NewConfig(snek.WithLogFormatCommandLineVariableShortName("F"))
	require.NotNil(t, cfg, "NewConfig should return a Config")
	assert.Equal(t, "F", cfg.LogFormatCommandLineVariableShortName,
		"LogFormatCommandLineVariableShortName should be F")
}

func TestWithLogFormatEnvironmentVariableName(t *testing.T) {
	cfg := snek.NewConfig()
	require.NotNil(t, cfg, "NewConfig should return a Config")
	require.Equal(t, "LOG_FORMAT", cfg.LogFormatEnvironmentVariableName,
		"LogFormatEnvironmentVariableName should be LOG_FORMAT")
	cfg = snek.NewConfig(snek.WithLogFormatEnvironmentVariableName("LOGFORMAT"))
	require.NotNil(t, cfg, "NewConfig should return a Config")
	assert.Equal(t, "LOGFORMAT", cfg.LogFormatEnvironmentVariableName,
		"LogFormatEnvironmentVariableName should be LOGFORMAT")
}

func TestWithLogLevelCommandLineVariableHelp(t *testing.T) {
	cfg := snek.NewConfig()
	require.NotNil(t, cfg, "NewConfig should return a Config")
	require.NotEqual(t, "The log level.", cfg.LogLevelCommandLineVariableHelp,
		"The default log level command line variable help should not be The log level.")
	cfg = snek.NewConfig(snek.WithLogLevelCommandLineVariableHelp("The log level."))
	require.NotNil(t, cfg, "NewConfig should return a Config")
	assert.Equal(t, "The log level.", cfg.LogLevelCommandLineVariableHelp,
		"LogLevelCommandLineVariableHelp should be The log level.")
}

func TestWithLogLevelCommandLineVariableLongName(t *testing.T) {
	cfg := snek.NewConfig()
	require.NotNil(t, cfg, "NewConfig should return a Config")
	require.Equal(t, "log-level", cfg.LogLevelCommandLineVariableLongName,
		"LogLevelCommandLineVariableLongName should be log-level")
	cfg = snek.NewConfig(snek.WithLogLevelCommandLineVariableLongName("loglevel"))
	require.NotNil(t, cfg, "NewConfig should return a Config")
	assert.Equal(t, "loglevel", cfg.LogLevelCommandLineVariableLongName,
		"LogLevelCommandLineVariableLongName should be loglevel")
}

func TestWithLogLevelCommandLineVariableShortName(t *testing.T) {
	cfg := snek.NewConfig()
	require.NotNil(t, cfg, "NewConfig should return a Config")
	require.Equal(t, "", cfg.LogLevelCommandLineVariableShortName,
		"LogLevelCommandLineVariableShortName should be l")
	cfg = snek.NewConfig(snek.WithLogLevelCommandLineVariableShortName("L"))
	require.NotNil(t, cfg, "NewConfig should return a Config")
	assert.Equal(t, "L", cfg.LogLevelCommandLineVariableShortName,
		"LogLevelCommandLineVariableShortName should be L")
}

func TestWithLogLevelEnvironmentVariableName(t *testing.T) {
	cfg := snek.NewConfig()
	require.NotNil(t, cfg, "NewConfig should return a Config")
	require.Equal(t, "LOG_LEVEL", cfg.LogLevelEnvironmentVariableName,
		"LogLevelEnvironmentVariableName should be LOG_LEVEL")
	cfg = snek.NewConfig(snek.WithLogLevelEnvironmentVariableName("LOGLEVEL"))
	require.NotNil(t, cfg, "NewConfig should return a Config")
	assert.Equal(t, "LOGLEVEL", cfg.LogLevelEnvironmentVariableName,
		"LogLevelEnvironmentVariableName should be LOGLEVEL")
}
