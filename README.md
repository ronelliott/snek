# snek

[![Go Reference](https://pkg.go.dev/badge/github.com/ronelliott/snek.svg)](https://pkg.go.dev/github.com/ronelliott/snek) [![Go Report Card](https://goreportcard.com/badge/github.com/ronelliott/snek)](https://goreportcard.com/report/github.com/ronelliott/snek)

Snek is a lite wrapper around both [cobra](https://github.com/spf13/cobra) and [zerolog](https://github.com/rs/zerolog). It generates a configurable root command, adds some command line arguments and environment variables to configure both the log format and level and then initialized logging prior to executing the command.

## Examples

```go
err := snek.Run(nil,
	snek.NewConfig(
		snek.WithDefaultLogFormat("json"),
		snek.WithDefaultLogLevel("debug"),
	),
	snek.WithUse("my-command"),
	snek.WithRun(func(cmd *cobra.Command, args []string) {
		// ...
	}),
)
```

The above example sets the following:
- The default log format to `json`
- The default log level to `debug`
- The command Use to `my-command`
- The command's run function
