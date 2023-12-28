package snek

import (
	"time"

	"github.com/spf13/pflag"
)

// FlagInitializer is a function that initializes a flag on a command.
type FlagInitializer func(*pflag.FlagSet) error

// WithBoolVar adds a bool flag to the command with the specified name, value,
// and usage and uses the specified variable to store the value of the flag.
func WithBoolVar(variable *bool, name string, value bool, usage string) FlagInitializer {
	return func(flags *pflag.FlagSet) error {
		flags.BoolVar(variable, name, value, usage)
		return nil
	}
}

// WithBoolVarP adds a bool flag to the command with the specified name,
// shorthand, value, and usage and uses the specified variable to store the
// value of the flag.
func WithBoolVarP(variable *bool, name, shorthand string, value bool, usage string) FlagInitializer {
	return func(flags *pflag.FlagSet) error {
		flags.BoolVarP(variable, name, shorthand, value, usage)
		return nil
	}
}

// WithDurationVar adds a duration flag to the command with the specified name,
// value, and usage and uses the specified variable to store the value of the
// flag.
func WithDurationVar(variable *time.Duration, name string, value time.Duration, usage string) FlagInitializer {
	return func(flags *pflag.FlagSet) error {
		flags.DurationVar(variable, name, value, usage)
		return nil
	}
}

// WithDurationVarP adds a duration flag to the command with the specified
// name, shorthand, value, and usage and uses the specified variable to store
// the value of the flag.
func WithDurationVarP(variable *time.Duration, name, shorthand string, value time.Duration, usage string) FlagInitializer {
	return func(flags *pflag.FlagSet) error {
		flags.DurationVarP(variable, name, shorthand, value, usage)
		return nil
	}
}

// WithFloat32Var adds a float32 flag to the command with the specified name,
// value, and usage and uses the specified variable to store the value of the
// flag.
func WithFloat32Var(variable *float32, name string, value float32, usage string) FlagInitializer {
	return func(flags *pflag.FlagSet) error {
		flags.Float32Var(variable, name, value, usage)
		return nil
	}
}

// WithFloat32VarP adds a float32 flag to the command with the specified name,
// shorthand, value, and usage and uses the specified variable to store the
// value of the flag.
func WithFloat32VarP(variable *float32, name, shorthand string, value float32, usage string) FlagInitializer {
	return func(flags *pflag.FlagSet) error {
		flags.Float32VarP(variable, name, shorthand, value, usage)
		return nil
	}
}

// WithFloat64Var adds a float64 flag to the command with the specified name,
// value, and usage and uses the specified variable to store the value of the
// flag.
func WithFloat64Var(variable *float64, name string, value float64, usage string) FlagInitializer {
	return func(flags *pflag.FlagSet) error {
		flags.Float64Var(variable, name, value, usage)
		return nil
	}
}

// WithFloat64VarP adds a float64 flag to the command with the specified name,
// shorthand, value, and usage and uses the specified variable to store the
// value of the flag.
func WithFloat64VarP(variable *float64, name, shorthand string, value float64, usage string) FlagInitializer {
	return func(flags *pflag.FlagSet) error {
		flags.Float64VarP(variable, name, shorthand, value, usage)
		return nil
	}
}

// WithIntVar adds an int flag to the command with the specified name, value,
// and usage and uses the specified variable to store the value of the flag.
func WithIntVar(variable *int, name string, value int, usage string) FlagInitializer {
	return func(flags *pflag.FlagSet) error {
		flags.IntVar(variable, name, value, usage)
		return nil
	}
}

// WithIntVarP adds an int flag to the command with the specified name,
// shorthand, value, and usage and uses the specified variable to store the
// value of the flag.
func WithIntVarP(variable *int, name, shorthand string, value int, usage string) FlagInitializer {
	return func(flags *pflag.FlagSet) error {
		flags.IntVarP(variable, name, shorthand, value, usage)
		return nil
	}
}

// WithStringVar adds a string flag to the command with the specified name,
// value, and usage and uses the specified variable to store the value of the
// flag.
func WithStringVar(variable *string, name, value, usage string) FlagInitializer {
	return func(flags *pflag.FlagSet) error {
		flags.StringVar(variable, name, value, usage)
		return nil
	}
}

// WithStringVarP adds a string flag to the command with the specified name,
// shorthand, value, and usage and uses the specified variable to store the
// value of the flag.
func WithStringVarP(variable *string, name, shorthand, value, usage string) FlagInitializer {
	return func(flags *pflag.FlagSet) error {
		flags.StringVarP(variable, name, shorthand, value, usage)
		return nil
	}
}
