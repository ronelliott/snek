package snek

import "errors"

var (
	// ErrLogFormatInvalid is returned when the log format is invalid.
	ErrLogFormatInvalid = errors.New("invalid log format")

	// ErrLogLevelInvalid is returned when the log level is invalid.
	ErrLogLevelInvalid = errors.New("invalid log level")

	// ErrLogFormatCommandLineVariableNameEmpty is returned when both the long and
	// short names of the log format command line variable are empty.
	ErrLogFormatCommandLineVariableNameEmpty = errors.New("log format command line variable name is empty")

	// ErrLogFormatEnvironmentVariableNameEmpty is returned when the log format
	// environment variable name is empty.
	ErrLogFormatEnvironmentVariableNameEmpty = errors.New("log format environment variable name is empty")

	// ErrLogLevelCommandLineVariableNameEmpty is returned when both the long and
	// short names of the log level command line variable are empty.
	ErrLogLevelCommandLineVariableNameEmpty = errors.New("log level command line variable name is empty")

	// ErrLogLevelEnvironmentVariableNameEmpty is returned when the log level
	// environment variable name is empty.
	ErrLogLevelEnvironmentVariableNameEmpty = errors.New("log level environment variable name is empty")

	// ErrLogOutputEmpty is returned when the log output is empty or nil.
	ErrLogOutputEmpty = errors.New("log output is empty")
)
