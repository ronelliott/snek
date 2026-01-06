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

// WithInt8Var adds an int8 flag to the command with the specified name, value,
// and usage and uses the specified variable to store the value of the flag.
func WithInt8Var(variable *int8, name string, value int8, usage string) FlagInitializer {
	return func(flags *pflag.FlagSet) error {
		flags.Int8Var(variable, name, value, usage)
		return nil
	}
}

// WithInt8VarP adds an int8 flag to the command with the specified name,
// shorthand, value, and usage and uses the specified variable to store the
// value of the flag.
func WithInt8VarP(variable *int8, name, shorthand string, value int8, usage string) FlagInitializer {
	return func(flags *pflag.FlagSet) error {
		flags.Int8VarP(variable, name, shorthand, value, usage)
		return nil
	}
}

// WithInt16Var adds an int16 flag to the command with the specified name,
// value, and usage and uses the specified variable to store the value of the
// flag.
func WithInt16Var(variable *int16, name string, value int16, usage string) FlagInitializer {
	return func(flags *pflag.FlagSet) error {
		flags.Int16Var(variable, name, value, usage)
		return nil
	}
}

// WithInt16VarP adds an int16 flag to the command with the specified name,
// shorthand, value, and usage and uses the specified variable to store the
// value of the flag.
func WithInt16VarP(variable *int16, name, shorthand string, value int16, usage string) FlagInitializer {
	return func(flags *pflag.FlagSet) error {
		flags.Int16VarP(variable, name, shorthand, value, usage)
		return nil
	}
}

// WithInt32Var adds an int32 flag to the command with the specified name,
// value, and usage and uses the specified variable to store the value of the
// flag.
func WithInt32Var(variable *int32, name string, value int32, usage string) FlagInitializer {
	return func(flags *pflag.FlagSet) error {
		flags.Int32Var(variable, name, value, usage)
		return nil
	}
}

// WithInt32VarP adds an int32 flag to the command with the specified name,
// shorthand, value, and usage and uses the specified variable to store the
// value of the flag.
func WithInt32VarP(variable *int32, name, shorthand string, value int32, usage string) FlagInitializer {
	return func(flags *pflag.FlagSet) error {
		flags.Int32VarP(variable, name, shorthand, value, usage)
		return nil
	}
}

// WithInt64Var adds an int64 flag to the command with the specified name,
// value, and usage and uses the specified variable to store the value of the
// flag.
func WithInt64Var(variable *int64, name string, value int64, usage string) FlagInitializer {
	return func(flags *pflag.FlagSet) error {
		flags.Int64Var(variable, name, value, usage)
		return nil
	}
}

// WithInt64VarP adds an int64 flag to the command with the specified name,
// shorthand, value, and usage and uses the specified variable to store the
// value of the flag.
func WithInt64VarP(variable *int64, name, shorthand string, value int64, usage string) FlagInitializer {
	return func(flags *pflag.FlagSet) error {
		flags.Int64VarP(variable, name, shorthand, value, usage)
		return nil
	}
}

// WithUintVar adds a uint flag to the command with the specified name, value,
// and usage and uses the specified variable to store the value of the flag.
func WithUintVar(variable *uint, name string, value uint, usage string) FlagInitializer {
	return func(flags *pflag.FlagSet) error {
		flags.UintVar(variable, name, value, usage)
		return nil
	}
}

// WithUintVarP adds a uint flag to the command with the specified name,
// shorthand, value, and usage and uses the specified variable to store the
// value of the flag.
func WithUintVarP(variable *uint, name, shorthand string, value uint, usage string) FlagInitializer {
	return func(flags *pflag.FlagSet) error {
		flags.UintVarP(variable, name, shorthand, value, usage)
		return nil
	}
}

// WithUint8Var adds a uint8 flag to the command with the specified name, value,
// and usage and uses the specified variable to store the value of the flag.
func WithUint8Var(variable *uint8, name string, value uint8, usage string) FlagInitializer {
	return func(flags *pflag.FlagSet) error {
		flags.Uint8Var(variable, name, value, usage)
		return nil
	}
}

// WithUint8VarP adds a uint8 flag to the command with the specified name,
// shorthand, value, and usage and uses the specified variable to store the
// value of the flag.
func WithUint8VarP(variable *uint8, name, shorthand string, value uint8, usage string) FlagInitializer {
	return func(flags *pflag.FlagSet) error {
		flags.Uint8VarP(variable, name, shorthand, value, usage)
		return nil
	}
}

// WithUint16Var adds a uint16 flag to the command with the specified name,
// value, and usage and uses the specified variable to store the value of the
// flag.
func WithUint16Var(variable *uint16, name string, value uint16, usage string) FlagInitializer {
	return func(flags *pflag.FlagSet) error {
		flags.Uint16Var(variable, name, value, usage)
		return nil
	}
}

// WithUint16VarP adds a uint16 flag to the command with the specified name,
// shorthand, value, and usage and uses the specified variable to store the
// value of the flag.
func WithUint16VarP(variable *uint16, name, shorthand string, value uint16, usage string) FlagInitializer {
	return func(flags *pflag.FlagSet) error {
		flags.Uint16VarP(variable, name, shorthand, value, usage)
		return nil
	}
}

// WithUint32Var adds a uint32 flag to the command with the specified name,
// value, and usage and uses the specified variable to store the value of the
// flag.
func WithUint32Var(variable *uint32, name string, value uint32, usage string) FlagInitializer {
	return func(flags *pflag.FlagSet) error {
		flags.Uint32Var(variable, name, value, usage)
		return nil
	}
}

// WithUint32VarP adds a uint32 flag to the command with the specified name,
// shorthand, value, and usage and uses the specified variable to store the
// value of the flag.
func WithUint32VarP(variable *uint32, name, shorthand string, value uint32, usage string) FlagInitializer {
	return func(flags *pflag.FlagSet) error {
		flags.Uint32VarP(variable, name, shorthand, value, usage)
		return nil
	}
}

// WithUint64Var adds a uint64 flag to the command with the specified name,
// value, and usage and uses the specified variable to store the value of the
// flag.
func WithUint64Var(variable *uint64, name string, value uint64, usage string) FlagInitializer {
	return func(flags *pflag.FlagSet) error {
		flags.Uint64Var(variable, name, value, usage)
		return nil
	}
}

// WithUint64VarP adds a uint64 flag to the command with the specified name,
// shorthand, value, and usage and uses the specified variable to store the
// value of the flag.
func WithUint64VarP(variable *uint64, name, shorthand string, value uint64, usage string) FlagInitializer {
	return func(flags *pflag.FlagSet) error {
		flags.Uint64VarP(variable, name, shorthand, value, usage)
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
