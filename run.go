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
// Logging is setup using a PersistentPreRunE hook on the root command. If an
// error occurs while setting up logging, it is returned from Run.
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

	// When the root command is runnable (has Run/RunE) and also has subcommands,
	// cobra's legacyArgs validation in Find() rejects any positional arguments
	// with an "unknown command" error. Auto-apply ArbitraryArgs when no custom
	// Args validator has been set so the root Run handler receives them normally.
	if rootCmd.Args == nil && rootCmd.Runnable() && rootCmd.HasSubCommands() {
		rootCmd.Args = cobra.ArbitraryArgs
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

	// Use PersistentPreRunE instead of cobra.OnInitialize to scope logging setup
	// to this command tree rather than the package-level global, which accumulates
	// across multiple Run() calls (e.g. in tests). Chain any hooks the caller may
	// have already set via initializers so neither is lost.
	existingPreRunE := rootCmd.PersistentPreRunE
	existingPreRun := rootCmd.PersistentPreRun
	rootCmd.PersistentPreRunE = func(cmd *Command, args []string) error {
		if err := setupLogging(logLevel, logFormat, cfg.LogOutput); err != nil {
			return err
		}
		if existingPreRunE != nil {
			return existingPreRunE(cmd, args)
		}
		if existingPreRun != nil {
			existingPreRun(cmd, args)
		}
		return nil
	}

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
