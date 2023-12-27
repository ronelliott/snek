package snek_test

import (
	"bytes"
	"encoding/json"
	"os"
	"strings"
	"testing"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ronelliott/snek"
)

func TestRun_Config_InvalidDefaultLogFormat(t *testing.T) {
	cfg := snek.NewConfig(snek.WithDefaultLogFormat("invalid"))
	err := snek.Run(nil, cfg)
	assert.ErrorIs(t, err, snek.ErrorLogFormatInvalid,
		"Run should return an error if the default log format is invalid")
}

func TestRun_Config_InvalidDefaultLogLevel(t *testing.T) {
	cfg := snek.NewConfig(snek.WithDefaultLogLevel("invalid"))
	err := snek.Run(nil, cfg)
	assert.ErrorIs(t, err, snek.ErrorLogLevelInvalid,
		"Run should return an error if the default log level is invalid")
}

func TestRun_Config_InvalidLogFormatCommandLineVariableLongName(t *testing.T) {
	cfg := snek.NewConfig(
		snek.WithLogFormatCommandLineVariableLongName(""),
		snek.WithLogFormatCommandLineVariableShortName(""),
	)
	err := snek.Run(nil, cfg)
	assert.ErrorIs(t, err, snek.ErrorLogFormatCommandLineVariableNameEmpty,
		"Run should return an error if the log format command line variable long and short name are empty")
}

func TestRun_Config_InvalidLogFormatEnvironmentVariableName(t *testing.T) {
	cfg := snek.NewConfig(snek.WithLogFormatEnvironmentVariableName(""))
	err := snek.Run(nil, cfg)
	assert.ErrorIs(t, err, snek.ErrorLogFormatEnvironmentVariableNameEmpty,
		"Run should return an error if the log format environment variable name is empty")
}

func TestRun_Config_InvalidLogLevelCommandLineVariableLongName(t *testing.T) {
	cfg := snek.NewConfig(
		snek.WithLogLevelCommandLineVariableLongName(""),
		snek.WithLogLevelCommandLineVariableShortName(""),
	)
	err := snek.Run(nil, cfg)
	assert.ErrorIs(t, err, snek.ErrorLogLevelCommandLineVariableNameEmpty,
		"Run should return an error if the log level command line variable long and short name are empty")
}

func TestRun_Config_InvalidLogLevelEnvironmentVariableName(t *testing.T) {
	cfg := snek.NewConfig(snek.WithLogLevelEnvironmentVariableName(""))
	err := snek.Run(nil, cfg)
	assert.ErrorIs(t, err, snek.ErrorLogLevelEnvironmentVariableNameEmpty,
		"Run should return an error if the log level environment variable name is empty")
}

func TestRun_Config_InvalidLogOutput(t *testing.T) {
	cfg := snek.NewConfig(snek.WithLogOutput(nil))
	err := snek.Run(nil, cfg)
	assert.ErrorIs(t, err, snek.ErrorLogOutputEmpty,
		"Run should return an error if the log output is empty")
}

func TestRun_Config_LogFormatHelp(t *testing.T) {
	logFormatHelp := "test log format help"

	cfg := snek.NewConfig()
	require.NotNil(t, cfg, "NewConfig should return a Config")
	require.NotEqual(t, logFormatHelp, cfg.LogFormatCommandLineVariableHelp,
		"Test should not use the default log format help")

	snek.WithLogFormatCommandLineVariableHelp(logFormatHelp)(cfg)

	called := false
	err := snek.Run(nil, cfg,
		snek.WithRun(func(cmd *cobra.Command, args []string) {
			called = true
			assert.Equal(t, logFormatHelp, cmd.Flag(cfg.LogFormatCommandLineVariableLongName).Usage,
				"Run should set the log format command line variable help")
		}),
	)
	assert.NoError(t, err, "Run should not return an error")
	require.True(t, called, "Run should call the commands Run function")
}

func TestRun_Config_LogLevelHelp(t *testing.T) {
	logLevelHelp := "test log level help"

	cfg := snek.NewConfig()
	require.NotNil(t, cfg, "NewConfig should return a Config")
	require.NotEqual(t, logLevelHelp, cfg.LogLevelCommandLineVariableHelp,
		"Test should not use the default log level help")

	snek.WithLogLevelCommandLineVariableHelp(logLevelHelp)(cfg)

	called := false
	err := snek.Run(nil, cfg,
		snek.WithRun(func(cmd *cobra.Command, args []string) {
			called = true
			assert.Equal(t, logLevelHelp, cmd.Flag(cfg.LogLevelCommandLineVariableLongName).Usage,
				"Run should set the log level command line variable help")
		}),
	)
	assert.NoError(t, err, "Run should not return an error")
	require.True(t, called, "Run should call the commands Run function")
}

func TestRun_Execute_Error(t *testing.T) {
	err := snek.Run(nil, snek.NewConfig(),
		snek.WithRunE(func(cmd *cobra.Command, args []string) error {
			return assert.AnError
		}))
	assert.ErrorIs(t, err, assert.AnError,
		"Run should return the error returned by the commands RunE function")
}

func TestRun_Execute_NoError(t *testing.T) {
	called := false
	err := snek.Run(nil, snek.NewConfig(),
		snek.WithRun(func(cmd *cobra.Command, args []string) {
			called = true
		}))
	assert.NoError(t, err, "Run should not return an error")
	assert.True(t, called, "Run should call the commands Run function")
}

func TestRun_Initializer_Error(t *testing.T) {
	err := snek.Run(nil, snek.NewConfig(),
		func(cmd *cobra.Command) error {
			return assert.AnError
		})
	assert.ErrorIs(t, err, assert.AnError,
		"Run should return the error produced by the initializer")
}

func TestRun_Logging_Basic_Formatted(t *testing.T) {
	lines := runLogLinesTest(t, "formatted", "debug", 6, nil)

	assert.Contains(t, lines[0], "Debug logging enabled.",
		"Run should log the debug logging enabled message")
	assert.Contains(t, lines[1], "Logging initialized.",
		"Run should log the logging initialized message")
	assert.Contains(t, lines[2], "debug test", "Run should log the debug message")
	assert.Contains(t, lines[5], "error test", "Run should log the error message")
	assert.Contains(t, lines[3], "info test", "Run should log the info message")
	assert.Contains(t, lines[4], "warn test", "Run should log the warn message")
}

func TestRun_Logging_Basic_Json(t *testing.T) {
	lines := runLogLinesJsonTest(t, "warn", 2, nil)

	errorLine := lines[1]
	assert.Equal(t, "error test", errorLine.Message, "Run should log the message")
	assert.Equal(t, "error", errorLine.Level, "Run should log the level")
	assert.NotEmpty(t, errorLine.Time, "Run should log the time")

	warnLine := lines[0]
	assert.Equal(t, "warn test", warnLine.Message, "Run should log the message")
	assert.Equal(t, "warn", warnLine.Level, "Run should log the level")
	assert.NotEmpty(t, warnLine.Time, "Run should log the time")
}

func TestRun_Logging_Format_CommandLine_Long(t *testing.T) {
	cliLongName := "test-log-format"
	format := "json"

	cfg := snek.NewConfig()
	require.NotNil(t, cfg, "NewConfig should return a Config")
	require.NotEqual(t, cliLongName, cfg.LogFormatCommandLineVariableLongName,
		"Test should not use the default command line variable long name")
	require.NotEqual(t, format, cfg.DefaultLogFormat,
		"Test should not use the default log format")

	snek.WithLogFormatCommandLineVariableLongName(cliLongName)(cfg)
	lines := runLogLinesJsonTestWithConfig(t, cfg, 4, []string{"--" + cliLongName, format})

	assert.Equal(t, "Logging initialized.", lines[0].Message, "Run should log the logging initialized message")
	assert.Equal(t, "error test", lines[3].Message, "Run should log the error message")
	assert.Equal(t, "info test", lines[1].Message, "Run should log the info message")
	assert.Equal(t, "warn test", lines[2].Message, "Run should log the warn message")
}

func TestRun_Logging_Format_CommandLine_Short(t *testing.T) {
	cliShortName := "f"
	format := "json"

	cfg := snek.NewConfig()
	require.NotNil(t, cfg, "NewConfig should return a Config")
	require.NotEqual(t, cliShortName, cfg.LogFormatCommandLineVariableShortName,
		"Test should not use the default command line variable short name")
	require.NotEqual(t, format, cfg.DefaultLogFormat,
		"Test should not use the default log format")

	snek.WithLogFormatCommandLineVariableShortName(cliShortName)(cfg)
	lines := runLogLinesJsonTestWithConfig(t, cfg, 4, []string{"-" + cliShortName, format})

	assert.Equal(t, "Logging initialized.", lines[0].Message, "Run should log the logging initialized message")
	assert.Equal(t, "error test", lines[3].Message, "Run should log the error message")
	assert.Equal(t, "info test", lines[1].Message, "Run should log the info message")
	assert.Equal(t, "warn test", lines[2].Message, "Run should log the warn message")
}

func TestRun_Logging_Format_EnvironmentVariable(t *testing.T) {
	envVarPrefix := "TEST_ENV_VAR_"
	envVarName := "TEST_LOG_FORMAT"
	format := "json"
	defer testEnvironmentVariable(t, envVarPrefix+envVarName, format)()

	cfg := snek.NewConfig()
	require.NotNil(t, cfg, "NewConfig should return a Config")
	require.NotEqual(t, envVarPrefix, cfg.EnvironmentVariablePrefix,
		"Test should not use the default environment variable prefix")
	require.NotEqual(t, envVarName, cfg.LogFormatEnvironmentVariableName,
		"Test should not use the default environment variable name")
	require.NotEqual(t, format, cfg.DefaultLogFormat,
		"Test should not use the default log format")

	snek.WithEnvironmentVariablePrefix(envVarPrefix)(cfg)
	snek.WithLogFormatEnvironmentVariableName(envVarName)(cfg)
	lines := runLogLinesJsonTestWithConfig(t, cfg, 4, nil)

	assert.Equal(t, "Logging initialized.", lines[0].Message, "Run should log the logging initialized message")
	assert.Equal(t, "error test", lines[3].Message, "Run should log the error message")
	assert.Equal(t, "info test", lines[1].Message, "Run should log the info message")
	assert.Equal(t, "warn test", lines[2].Message, "Run should log the warn message")
}

func TestRun_Logging_Format_Invalid(t *testing.T) {
	called := false
	cfg := snek.NewConfig(snek.WithDefaultLogFormat("invalid"))
	err := snek.Run(nil, cfg,
		snek.WithRun(func(cmd *cobra.Command, args []string) {
			called = true
			log.Debug().Msg("debug test")
		}),
	)
	assert.ErrorIs(t, err, snek.ErrorLogFormatInvalid, "Run should return an error")
	assert.False(t, called, "Run should not call the commands Run function")
}

func TestRun_Logging_Level_CommandLine_Long(t *testing.T) {
	cliLongName := "test-log-level"
	level := "debug"

	cfg := snek.NewConfig()
	require.NotNil(t, cfg, "NewConfig should return a Config")
	require.NotEqual(t, cliLongName, cfg.LogLevelCommandLineVariableLongName,
		"Test should not use the default command line variable long name")
	require.NotEqual(t, level, cfg.DefaultLogLevel,
		"Test should not use the default log level")

	snek.WithDefaultLogFormat("json")(cfg)
	snek.WithLogLevelCommandLineVariableLongName(cliLongName)(cfg)
	lines := runLogLinesJsonTestWithConfig(t, cfg, 6, []string{"--" + cliLongName, level})

	assert.Equal(t, "Debug logging enabled.", lines[0].Message, "Run should log the debug logging enabled message")
	assert.Equal(t, "Logging initialized.", lines[1].Message, "Run should log the logging initialized message")
	assert.Equal(t, "debug test", lines[2].Message, "Run should log the debug message")
	assert.Equal(t, "error test", lines[5].Message, "Run should log the error message")
	assert.Equal(t, "info test", lines[3].Message, "Run should log the info message")
	assert.Equal(t, "warn test", lines[4].Message, "Run should log the warn message")
}

func TestRun_Logging_Level_CommandLine_Short(t *testing.T) {
	cliShortName := "l"
	level := "debug"

	cfg := snek.NewConfig()
	require.NotNil(t, cfg, "NewConfig should return a Config")
	require.NotEqual(t, cliShortName, cfg.LogLevelCommandLineVariableShortName,
		"Test should not use the default command line variable short name")
	require.NotEqual(t, level, cfg.DefaultLogLevel,
		"Test should not use the default log level")

	snek.WithDefaultLogFormat("json")(cfg)
	snek.WithLogLevelCommandLineVariableShortName(cliShortName)(cfg)
	lines := runLogLinesJsonTestWithConfig(t, cfg, 6, []string{"-" + cliShortName, level})

	assert.Equal(t, "Debug logging enabled.", lines[0].Message, "Run should log the debug logging enabled message")
	assert.Equal(t, "Logging initialized.", lines[1].Message, "Run should log the logging initialized message")
	assert.Equal(t, "debug test", lines[2].Message, "Run should log the debug message")
	assert.Equal(t, "error test", lines[5].Message, "Run should log the error message")
	assert.Equal(t, "info test", lines[3].Message, "Run should log the info message")
	assert.Equal(t, "warn test", lines[4].Message, "Run should log the warn message")
}

func TestRun_Logging_Level_EnvironmentVariable(t *testing.T) {
	envVarPrefix := "TEST_ENV_VAR_"
	envVarName := "TEST_LOG_LEVEL"
	level := "debug"
	defer testEnvironmentVariable(t, envVarPrefix+envVarName, level)()

	cfg := snek.NewConfig()
	require.NotNil(t, cfg, "NewConfig should return a Config")
	require.NotEqual(t, envVarPrefix, cfg.EnvironmentVariablePrefix,
		"Test should not use the default environment variable prefix")
	require.NotEqual(t, envVarName, cfg.LogLevelEnvironmentVariableName,
		"Test should not use the default environment variable name")
	require.NotEqual(t, level, cfg.DefaultLogLevel,
		"Test should not use the default log level")

	snek.WithDefaultLogFormat("json")(cfg)
	snek.WithEnvironmentVariablePrefix(envVarPrefix)(cfg)
	snek.WithLogLevelEnvironmentVariableName(envVarName)(cfg)
	lines := runLogLinesJsonTestWithConfig(t, cfg, 6, nil)

	assert.Equal(t, "Debug logging enabled.", lines[0].Message, "Run should log the debug logging enabled message")
	assert.Equal(t, "Logging initialized.", lines[1].Message, "Run should log the logging initialized message")
	assert.Equal(t, "debug test", lines[2].Message, "Run should log the debug message")
	assert.Equal(t, "error test", lines[5].Message, "Run should log the error message")
	assert.Equal(t, "info test", lines[3].Message, "Run should log the info message")
	assert.Equal(t, "warn test", lines[4].Message, "Run should log the warn message")
}

func TestRun_Logging_Level_Invalid(t *testing.T) {
	called := false
	cfg := snek.NewConfig(snek.WithDefaultLogLevel("invalid"))
	err := snek.Run(nil, cfg,
		snek.WithRun(func(cmd *cobra.Command, args []string) {
			called = true
			log.Debug().Msg("debug test")
		}),
	)
	assert.ErrorIs(t, err, snek.ErrorLogLevelInvalid, "Run should return an error")
	assert.False(t, called, "Run should not call the commands Run function")
}

type parsedLogLine struct {
	Level   string `json:"level"`
	Message string `json:"message"`
	Time    string `json:"time"`
}

func runLogLinesTestWithConfig(t *testing.T, cfg *snek.Config, expectedLines int, args []string) []string {
	t.Helper()
	var buffer bytes.Buffer
	snek.WithLogOutput(&buffer)(cfg)

	err := snek.Run(args, cfg,
		snek.WithRun(func(cmd *cobra.Command, args []string) {
			log.Debug().Msg("debug test")
			log.Info().Msg("info test")
			log.Warn().Msg("warn test")
			log.Error().Msg("error test")
		}),
	)
	require.NoError(t, err, "Run should not return an error")

	lines := strings.Split(buffer.String(), "\n")
	lines = lines[:len(lines)-1] // Remove the last empty line
	require.Len(t, lines, expectedLines,
		"Run should log each message on a separate line and not log disabled levels")

	return lines
}

func runLogLinesTest(t *testing.T, format, level string, expectedLines int, args []string) []string {
	return runLogLinesTestWithConfig(t,
		snek.NewConfig(
			snek.WithDefaultLogFormat(format),
			snek.WithDefaultLogLevel(level),
		),
		expectedLines,
		args,
	)
}

func runLogLinesJsonTestWithConfig(t *testing.T, cfg *snek.Config, expectedLines int, args []string) []parsedLogLine {
	t.Helper()
	lines := runLogLinesTestWithConfig(t, cfg, expectedLines, args)

	parsedLines := make([]parsedLogLine, len(lines))
	for i, line := range lines {
		err := json.Unmarshal([]byte(line), &parsedLines[i])
		require.NoError(t, err, "Run should log in the JSON format")
	}

	return parsedLines
}

func runLogLinesJsonTest(t *testing.T, level string, expectedLines int, args []string) []parsedLogLine {
	return runLogLinesJsonTestWithConfig(t, snek.NewConfig(
		snek.WithDefaultLogFormat("json"),
		snek.WithDefaultLogLevel(level),
	), expectedLines, args)
}

func testEnvironmentVariable(t *testing.T, name, value string) func() {
	t.Helper()
	oldValue, existed := os.LookupEnv(name)
	err := os.Setenv(name, value)
	require.NoError(t, err, "Test should be able to set the environment variable")
	return func() {
		t.Helper()
		if existed {
			err := os.Setenv(name, oldValue)
			require.NoError(t, err, "Test should be able to restore the environment variable")
		} else {
			err := os.Unsetenv(name)
			require.NoError(t, err, "Test should be able to restore the environment variable")
		}
	}
}
