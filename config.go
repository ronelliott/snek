package snek

import (
	"io"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Config is the configuration used by snek for configuring the generated root command.
type Config struct {
	// DefaultLogFormat is the default log format to use when logging.
	//
	// Valid values are `formatted` and `json`.
	//
	// The default value is `formatted`.
	DefaultLogFormat string

	// DefaultLogLevel is the default log level to use when logging.
	//
	// Valid values are `debug`, `error`, `fatal`, `info`, `panic`, `trace`, and `warn`.
	//
	// The default value is `info`.
	DefaultLogLevel string

	// EnvironmentVariablePrefix is the prefix that will be prepended to all
	// environment variables used by snek.
	//
	// The default value is an empty string.
	EnvironmentVariablePrefix string

	// LogFormatCommandLineVariableHelp is the help text for the command line
	// variable that will be used to set the log format.
	LogFormatCommandLineVariableHelp string

	// LogFormatCommandLineVariableLongName is the long name of the command line
	// variable that will be used to set the log format, i.e. `--log-format`.
	//
	// The long name will be prepended with a double dash (`--`), therefore one
	// should not be provided when configuring this value.
	//
	// If the long name is an empty string, then no long name will be used.
	//
	// The default value is `log-format`.
	LogFormatCommandLineVariableLongName string

	// LogFormatCommandLineVariableShortName is the short name of the command line
	// variable that will be used to set the log format, i.e. `-f`.
	//
	// The short name will be prepended with a single dash (`-`), therefore one
	// should not be provided when configuring this value.
	//
	// If the short name is an empty string, then no short name will be used.
	//
	// The default value is an empty string.
	LogFormatCommandLineVariableShortName string

	// LogFormatEnvironmentVariableName is the name of the environment variable
	// that will be used to set the log format.
	//
	// The configured environment variable prefix will be prepended to the
	// configured environment variable name.
	//
	// The default value is `LOG_FORMAT`.
	LogFormatEnvironmentVariableName string

	// LogLevelCommandLineVariableHelp is the help text for the command line
	// variable that will be used to set the log level.
	LogLevelCommandLineVariableHelp string

	// LogLevelCommandLineVariableLongName is the long name of the command line
	// variable that will be used to set the log level, i.e. `--log-level`.
	//
	// The long name will be prepended with a double dash (`--`), therefore one
	// should not be provided when configuring this value.
	//
	// If the long name is an empty string, then no long name will be used.
	//
	// The default value is `log-level`.
	LogLevelCommandLineVariableLongName string

	// LogLevelCommandLineVariableShortName is the short name of the command line
	// variable that will be used to set the log level, i.e. `-l`.
	//
	// The short name will be prepended with a single dash (`-`), therefore one
	// should not be provided when configuring this value.
	//
	// If the short name is an empty string, then no short name will be used.
	//
	// The default value is an empty string.
	LogLevelCommandLineVariableShortName string

	// LogLevelEnvironmentVariableName is the name of the environment variable
	// that will be used to set the log level.
	//
	// The configured environment variable prefix will be prepended to the
	// configured environment variable name.
	//
	// The default value is `LOG_LEVEL`.
	LogLevelEnvironmentVariableName string

	// LogOutput is the output that will be used for logging.
	//
	// The default value is `os.Stdout`.
	LogOutput io.Writer
}

// Configurator is a function that can be used to configure snek.
type Configurator func(*Config)

// NewConfig creates a new Config with the default values. Each initializer
// is called with the Config to initialize any values within the configuration.
func NewConfig(initializers ...Configurator) *Config {
	cfg := &Config{
		DefaultLogFormat:                      "formatted",
		DefaultLogLevel:                       "info",
		EnvironmentVariablePrefix:             "",
		LogFormatCommandLineVariableHelp:      "The log format to use when logging. Valid values are `formatted` and `json`.",
		LogFormatCommandLineVariableLongName:  "log-format",
		LogFormatCommandLineVariableShortName: "",
		LogFormatEnvironmentVariableName:      "LOG_FORMAT",
		LogLevelCommandLineVariableHelp:       "The logging level to use. Logs with a level greater than or equal to the specified level will be logged. Valid values are `debug`, `error`, `fatal`, `info`, `panic`, `trace`, and `warn`.",
		LogLevelCommandLineVariableLongName:   "log-level",
		LogLevelCommandLineVariableShortName:  "",
		LogLevelEnvironmentVariableName:       "LOG_LEVEL",
		LogOutput:                             os.Stdout,
	}

	for _, initializer := range initializers {
		initializer(cfg)
	}

	return cfg
}

// validate validates that the values in the Config are valid.
//
// If the values are not valid, then an error is returned. Otherwise, nil is
// returned.
//
// This function checks the following:
// - DefaultLogFormat is valid
// - DefaultLogLevel is valid
// - LogFormatCommandLineVariableLongName and LogFormatCommandLineVariableShortName are not both empty
// - LogFormatEnvironmentVariableName is not empty
// - LogLevelCommandLineVariableLongName and LogLevelCommandLineVariableShortName are not both empty
// - LogLevelEnvironmentVariableName is not empty
// - LogOutput is not nil
func (cfg *Config) validate() error {
	switch cfg.DefaultLogFormat {
	case LogFormatFormatted, LogFormatJson:
	default:
		log.Error().Str("format", cfg.DefaultLogFormat).Msg("Default log format is invalid")
		return ErrLogFormatInvalid
	}

	_, err := zerolog.ParseLevel(cfg.DefaultLogLevel)
	if err != nil {
		log.Error().Str("level", cfg.DefaultLogLevel).Msg("Default log level is invalid")
		return ErrLogLevelInvalid
	}

	if len(cfg.LogFormatCommandLineVariableLongName) == 0 && len(cfg.LogFormatCommandLineVariableShortName) == 0 {
		log.Error().Msg("Log format command line variable long name and short name are both empty")
		return ErrLogFormatCommandLineVariableNameEmpty
	}

	if len(cfg.LogFormatEnvironmentVariableName) == 0 {
		log.Error().Msg("Log format environment variable name is empty")
		return ErrLogFormatEnvironmentVariableNameEmpty
	}

	if len(cfg.LogLevelCommandLineVariableLongName) == 0 && len(cfg.LogLevelCommandLineVariableShortName) == 0 {
		log.Error().Msg("Log level command line variable long name and short name are both empty")
		return ErrLogLevelCommandLineVariableNameEmpty
	}

	if len(cfg.LogLevelEnvironmentVariableName) == 0 {
		log.Error().Msg("Log level environment variable name is empty")
		return ErrLogLevelEnvironmentVariableNameEmpty
	}

	if cfg.LogOutput == nil {
		log.Error().Msg("Log output is nil")
		return ErrLogOutputEmpty
	}

	return nil
}

// WithDefaultLogFormat sets the default log format to the provided value.
//
// The default value is `formatted`.
func WithDefaultLogFormat(format string) Configurator {
	return func(cfg *Config) {
		cfg.DefaultLogFormat = format
	}
}

// WithDefaultLogLevel sets the default log level to the provided value.
//
// Valid values are `debug`, `error`, `fatal`, `info`, `panic`, `trace`, and `warn`.
//
// The default value is `info`.
func WithDefaultLogLevel(level string) Configurator {
	return func(cfg *Config) {
		cfg.DefaultLogLevel = level
	}
}

// WithEnvironmentVariablePrefix sets the environment variable prefix to the provided value.
//
// The default value is an empty string.
func WithEnvironmentVariablePrefix(prefix string) Configurator {
	return func(cfg *Config) {
		cfg.EnvironmentVariablePrefix = prefix
	}
}

// WithLogLevelCommandLineVariableHelp sets the help text for the command line
// variable that will be used to set the log level.
func WithLogFormatCommandLineVariableHelp(help string) Configurator {
	return func(cfg *Config) {
		cfg.LogFormatCommandLineVariableHelp = help
	}
}

// WithLogFormatCommandLineVariableLongName sets the long name of the command line
// variable that will be used to set the log format, i.e. `--log-format`.
//
// The long name will be prepended with a double dash (`--`), therefore one
// should not be provided when configuring this value. If the long name
// is an empty string, then no long name will be used.
//
// The default value is `log-format`.
func WithLogFormatCommandLineVariableLongName(name string) Configurator {
	return func(cfg *Config) {
		cfg.LogFormatCommandLineVariableLongName = name
	}
}

// WithLogFormatCommandLineVariableShortName sets the short name of the command line
// variable that will be used to set the log format, i.e. `-f`.
//
// The short name will be prepended with a single dash (`-`), therefore one
// should not be provided when configuring this value. If the short name
// is an empty string, then no short name will be used.
//
// The default value is an empty string.
func WithLogFormatCommandLineVariableShortName(name string) Configurator {
	return func(cfg *Config) {
		cfg.LogFormatCommandLineVariableShortName = name
	}
}

// WithLogFormatEnvironmentVariableName sets the name of the environment variable
// that will be used to set the log format. The configured environment
// variable prefix will be prepended to the configured environment variable name.
//
// The default value is `LOG_FORMAT`.
func WithLogFormatEnvironmentVariableName(name string) Configurator {
	return func(cfg *Config) {
		cfg.LogFormatEnvironmentVariableName = name
	}
}

// WithLogLevelCommandLineVariableHelp sets the help text for the command line
// variable that will be used to set the log level.
func WithLogLevelCommandLineVariableHelp(help string) Configurator {
	return func(cfg *Config) {
		cfg.LogLevelCommandLineVariableHelp = help
	}
}

// WithLogLevelCommandLineVariableLongName sets the long name of the command line
// variable that will be used to set the log level, i.e. `--log-level`.
//
// The long name will be prepended with a double dash (`--`), therefore one
// should not be provided when configuring this value. If the long name
// is an empty string, then no long name will be used.
//
// The default value is `log-level`.
func WithLogLevelCommandLineVariableLongName(name string) Configurator {
	return func(cfg *Config) {
		cfg.LogLevelCommandLineVariableLongName = name
	}
}

// WithLogLevelCommandLineVariableShortName sets the short name of the command line
// variable that will be used to set the log level, i.e. `-l`.
//
// The short name will be prepended with a single dash (`-`), therefore one
// should not be provided when configuring this value. If the short name
// is an empty string, then no short name will be used.
//
// The default value is an empty string.
func WithLogLevelCommandLineVariableShortName(name string) Configurator {
	return func(cfg *Config) {
		cfg.LogLevelCommandLineVariableShortName = name
	}
}

// WithLogLevelEnvironmentVariableName sets the name of the environment variable
// that will be used to set the log level. The configured environment
// variable prefix will be prepended to the configured environment variable name.
//
// The default value is `LOG_LEVEL`.
func WithLogLevelEnvironmentVariableName(name string) Configurator {
	return func(cfg *Config) {
		cfg.LogLevelEnvironmentVariableName = name
	}
}

// WithLogOutput sets the log output to the provided value.
//
// The default value is `os.Stdout`.
func WithLogOutput(output io.Writer) Configurator {
	return func(cfg *Config) {
		cfg.LogOutput = output
	}
}
