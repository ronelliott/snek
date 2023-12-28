package main

import (
	"net/http"

	"github.com/ronelliott/snek"
	"github.com/spf13/cobra"
)

func main() {
	snek.RunExit(
		snek.NewConfig(
			snek.WithDefaultLogFormat("json"),
			snek.WithDefaultLogLevel("debug"),
			snek.WithEnvironmentVariablePrefix("MY_AWESOME_APP_"),
		),
		snek.WithUse("my-awesome-command"),
		snek.WithSubCommandGenerator(
			newRunCommand,
		),
	)
}

func newRunCommand() (*cobra.Command, error) {
	port := ":3000"
	return snek.NewCommand(
		snek.WithUse("run"),
		snek.WithShort("Run the application"),
		snek.WithFlag(
			snek.WithStringVarP(&port, "port", "p", port, "The port to bind to"),
		),
		snek.WithSimpleRunE(func(args []string) error {
			return http.ListenAndServe(port)
		}),
	)
}
