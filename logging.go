package snek

import (
	"io"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	// LogFormatFormatted is the formatted log format that prints log lines in a
	// human-readable format.
	LogFormatFormatted = "formatted"

	// LogFormatJson is the JSON log format that prints log lines in a JSON format.
	LogFormatJson = "json"
)

// setupLogging sets the logging level and format to the specified level and
// format. If either the level or format is invalid, then an ErrorLogFormatInvalid
// or ErrorLogLevelInvalid error is returned.
//
// If the level is debug, then a debug log line is written confirming that debug
// logging is enabled.
func setupLogging(level, format string, out io.Writer) error {
	if err := setupLoggingFormat(format, out); err != nil {
		return err
	}

	if err := setupLoggingLevel(level); err != nil {
		return err
	}

	log.Info().Str("level", level).Str("format", format).Msg("Logging initialized.")
	return nil
}

// setupLoggingFormat sets the logging format to the specified format. If the
// format is invalid, then an ErrorLogFormatInvalid error is returned.
func setupLoggingFormat(format string, out io.Writer) error {
	switch format {
	case LogFormatFormatted:
		log.Logger = log.Output(zerolog.ConsoleWriter{
			Out:        out,
			TimeFormat: time.DateTime,
		})
	case LogFormatJson:
		log.Logger = log.Output(out)
	default:
		log.Error().Str("format", format).Msg("Invalid log format")
		return ErrLogFormatInvalid
	}
	return nil
}

// setupLoggingLevel sets the logging level to the specified level. If the level
// is invalid, then an ErrorLogLevelInvalid error is returned.
//
// If the level is debug, then a debug log line is written confirming that debug
// logging is enabled.
func setupLoggingLevel(level string) error {
	actual, err := zerolog.ParseLevel(level)
	if err != nil {
		log.Error().Err(err).Msg("Error parsing log level")
		return ErrLogLevelInvalid
	}
	zerolog.SetGlobalLevel(actual)
	log.Debug().Msg("Debug logging enabled.")
	return nil
}
