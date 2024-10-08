package snek

import "github.com/spf13/cobra"

// Command is a wrapper around Command to allow easier imports.
type Command = cobra.Command

// Initializer is a function that initializes a cobra command by setting values
// on the command.
type Initializer func(*Command) error

// NewCommand creates a new cobra command with the specified initializers. Each
// initializer is called in order with the command to initialize passed as an argument.
// If an error is returned from an initializer, then the command is not created
// and the error is returned.
func NewCommand(initializers ...Initializer) (*Command, error) {
	cmd := &Command{}
	for _, initializer := range initializers {
		if err := initializer(cmd); err != nil {
			return nil, err
		}
	}
	return cmd, nil
}

// WithAliases sets the aliases on the command.
func WithAliases(aliases ...string) Initializer {
	return func(cmd *Command) error {
		cmd.Aliases = aliases
		return nil
	}
}

// WithDeprecated sets the deprecated message on the command.
func WithDeprecated(deprecated string) Initializer {
	return func(cmd *Command) error {
		cmd.Deprecated = deprecated
		return nil
	}
}

// WithExample sets the example on the command.
func WithExample(example string) Initializer {
	return func(cmd *Command) error {
		cmd.Example = example
		return nil
	}
}

// WithFlag adds the specified flags to the command.
func WithFlag(flags ...FlagInitializer) Initializer {
	return func(cmd *Command) error {
		for _, flag := range flags {
			if err := flag(cmd.Flags()); err != nil {
				return err
			}
		}

		return nil
	}
}

// WithLong sets the long description on the command.
func WithLong(long string) Initializer {
	return func(cmd *Command) error {
		cmd.Long = long
		return nil
	}
}

// WithRun sets the run function on the command.
func WithRun(run func(*Command, []string)) Initializer {
	return func(cmd *Command) error {
		cmd.Run = run
		return nil
	}
}

// WithRunE sets the error run function on the command.
func WithRunE(run func(*Command, []string) error) Initializer {
	return func(cmd *Command) error {
		cmd.RunE = run
		return nil
	}
}

// WithShort sets the short description on the command.
func WithShort(short string) Initializer {
	return func(cmd *Command) error {
		cmd.Short = short
		return nil
	}
}

// WithSimpleRun sets the run function on the command to a simple function that
// takes a slice of strings of the arguments and does not return an error.
func WithSimpleRun(run func([]string)) Initializer {
	return func(cmd *Command) error {
		cmd.Run = func(cmd *Command, args []string) {
			run(args)
		}
		return nil
	}
}

// WithSimpleRunE sets the run function on the command to a simple function that
// takes a slice of strings of the arguments and returns an error.
func WithSimpleRunE(run func([]string) error) Initializer {
	return func(cmd *Command) error {
		cmd.RunE = func(cmd *Command, args []string) error {
			return run(args)
		}
		return nil
	}
}

// WithSubCommand adds the specified subcommands to the command.
func WithSubCommand(subcommands ...*Command) Initializer {
	return func(cmd *Command) error {
		cmd.AddCommand(subcommands...)
		return nil
	}
}

// WithSubCommandGenerator adds the specified subcommand generators to the
// command. Each generator is called and the returned command is added to the
// command as a subcommand if there is no error. If an error is returned from
// any of the generators, then the command is not created and the error is
// returned.
func WithSubCommandGenerator(generators ...func() (*Command, error)) Initializer {
	return func(cmd *Command) error {
		for _, generator := range generators {
			generated, err := generator()
			if err != nil {
				return err
			}

			cmd.AddCommand(generated)
		}

		return nil
	}
}

// WithUse sets the name on the command.
func WithUse(name string) Initializer {
	return func(cmd *Command) error {
		cmd.Use = name
		return nil
	}
}

// WithValidArgs sets the valid args on the command.
func WithValidArgs(validArgs ...string) Initializer {
	return func(cmd *Command) error {
		cmd.ValidArgs = validArgs
		return nil
	}
}

// WithVersion sets the version on the command.
func WithVersion(version string) Initializer {
	return func(cmd *Command) error {
		cmd.Version = version
		return nil
	}
}
