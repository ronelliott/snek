package snek

import "errors"

var (
	// ErrorLogFormatInvalid is returned when the log format is invalid.
	ErrorLogFormatInvalid = errors.New("invalid log format")

	// ErrorLogLevelInvalid is returned when the log level is invalid.
	ErrorLogLevelInvalid = errors.New("invalid log level")

	// ErrorLogFormatCommandLineVariableNameEmpty is returned when both the long and
	// short names of the log format command line variable are empty.
	ErrorLogFormatCommandLineVariableNameEmpty = errors.New("log format command line variable name is empty")

	// ErrorLogFormatEnvironmentVariableNameEmpty is returned when the log format
	// environment variable name is empty.
	ErrorLogFormatEnvironmentVariableNameEmpty = errors.New("log format environment variable name is empty")

	// ErrorLogLevelCommandLineVariableNameEmpty is returned when both the long and
	// short names of the log level command line variable are empty.
	ErrorLogLevelCommandLineVariableNameEmpty = errors.New("log level command line variable name is empty")

	// ErrorLogLevelEnvironmentVariableNameEmpty is returned when the log level
	// environment variable name is empty.
	ErrorLogLevelEnvironmentVariableNameEmpty = errors.New("log level environment variable name is empty")

	// ErrorLogOutputEmpty is returned when the log output is empty or nil.
	ErrorLogOutputEmpty = errors.New("log output is empty")
)
