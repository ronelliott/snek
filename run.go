package snek

import (
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// RunExit calls Run with the given initializers and then exits with a status
// code of 1 if an error is returned from Run.
func RunExit(cfg *Config, initializers ...Initializer) {
	if err := Run(nil, cfg, initializers...); err != nil {
		os.Exit(1)
	}
}

// Run creates a new root command with the specified initializers and executes
// the command.
//
// If args is provided, it is used to parse command line arguments. If args is not
// provided, then os.Args[1:] will be used instead.
//
// Each initializer is called with the generated root command as an argument, and
// if an error is returned, then command generation stops and Run returns the error.
//
// This function will add two persistent flags to the root command:
//
//	--log-format: The log format to use.
//	--log-level: The logging level to use.
//
// The command line variable names can be changed by using the following
// initializers:
//
//	snek.WithLogFormatCommandLineVariableLongName
//	snek.WithLogLevelCommandLineVariableLongName
//
// Additionally, short names can be set using the following initializers:
//
//	snek.WithLogFormatCommandLineVariableShortName
//	snek.WithLogLevelCommandLineVariableShortName
//
// The log format and level can also be set using the following environment
// variables:
//
//	LOG_FORMAT: The log format to use.
//	LOG_LEVEL: The logging level to use.
//
// The environment variables used can be changed by using the following
// initializers:
//
//	snek.WithLogFormatEnvironmentVariableName
//	snek.WithLogLevelEnvironmentVariableName
//
// Note that the environment variable prefix will be prepended to the configured
// environment variable names. To change the environment variable prefix, use
// snek.WithEnvironmentVariablePrefix. The default environment variable prefix
// is an empty string.
//
// The default log format is `formatted` and the default log level is `info`. The
// default log output is `os.Stdout`. To change the default log format, level,
// or output, use the following initializers:
//
//	snek.WithDefaultLogFormat
//	snek.WithDefaultLogLevel
//	snek.WithLogOutput
//
// Logging is setup by using a cobra.OnInitialize function. This function will
// setup logging using the configured log format and level. If an error occurs
// while setting up logging, then the application will exit with a status code
// of 1.
//
// If the log level is set to `debug`, then a debug log line is written confirming
// that debug logging is enabled.
func Run(args []string, cfg *Config, initializers ...Initializer) error {
	// ---------------------------------------------------------------------------
	// Config
	// ---------------------------------------------------------------------------

	if cfg == nil {
		cfg = NewConfig()
	}

	if err := cfg.validate(); err != nil {
		log.Error().Err(err).Msg("Error validating config")
		return err
	}

	// ---------------------------------------------------------------------------
	// Root Command
	// ---------------------------------------------------------------------------

	rootCmd, err := NewCommand(initializers...)
	if err != nil {
		log.Error().Err(err).Msg("Error creating root command")
		return err
	}

	// ---------------------------------------------------------------------------
	// Logging
	// ---------------------------------------------------------------------------

	logFormat := getEnvOrDefault(
		cfg.EnvironmentVariablePrefix+cfg.LogFormatEnvironmentVariableName,
		cfg.DefaultLogFormat)
	logLevel := getEnvOrDefault(
		cfg.EnvironmentVariablePrefix+cfg.LogLevelEnvironmentVariableName,
		cfg.DefaultLogLevel)

	pflags := rootCmd.PersistentFlags()
	pflags.StringVarP(
		&logFormat,
		cfg.LogFormatCommandLineVariableLongName,
		cfg.LogFormatCommandLineVariableShortName,
		logFormat,
		cfg.LogFormatCommandLineVariableHelp)
	pflags.StringVarP(
		&logLevel,
		cfg.LogLevelCommandLineVariableLongName,
		cfg.LogLevelCommandLineVariableShortName,
		logLevel,
		cfg.LogLevelCommandLineVariableHelp)

	cobra.OnInitialize(func() {
		if err := setupLogging(logLevel, logFormat, cfg.LogOutput); err != nil {
			log.Error().Err(err).Msg("Error setting up logging")
			os.Exit(1)
		}
	})

	// ---------------------------------------------------------------------------
	// Execute
	// ---------------------------------------------------------------------------

	rootCmd.SetArgs(args)
	if err := rootCmd.Execute(); err != nil {
		log.Error().Err(err).Msg("Error executing command")
		return err
	}

	return nil
}

// getEnvOrDefault returns the value of the specified environment variable or
// the specified default value if the environment variable is not set.
func getEnvOrDefault(envVarName, defaultValue string) string {
	value, ok := os.LookupEnv(envVarName)
	if !ok {
		value = defaultValue
	}

	return value
}
