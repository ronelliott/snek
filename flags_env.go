package snek

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/spf13/pflag"
)

// WithBoolVarE adds a bool flag to the command with the specified name, value,
// and usage. If the environment variable envVar is set, its value is used as
// the default instead of value. The variable stores the final flag value.
func WithBoolVarE(variable *bool, name, envVar string, value bool, usage string) FlagInitializer {
	if v, ok := os.LookupEnv(envVar); ok {
		parsed, err := strconv.ParseBool(v)
		if err != nil {
			return func(*pflag.FlagSet) error {
				return fmt.Errorf("%w: %s=%q: %v", ErrFlagEnvVarInvalid, envVar, v, err)
			}
		}
		value = parsed
	}
	return func(flags *pflag.FlagSet) error {
		flags.BoolVar(variable, name, value, usage)
		return nil
	}
}

// WithBoolVarPE adds a bool flag to the command with the specified name,
// shorthand, value, and usage. If the environment variable envVar is set, its
// value is used as the default instead of value. The variable stores the final
// flag value.
func WithBoolVarPE(variable *bool, name, shorthand, envVar string, value bool, usage string) FlagInitializer {
	if v, ok := os.LookupEnv(envVar); ok {
		parsed, err := strconv.ParseBool(v)
		if err != nil {
			return func(*pflag.FlagSet) error {
				return fmt.Errorf("%w: %s=%q: %v", ErrFlagEnvVarInvalid, envVar, v, err)
			}
		}
		value = parsed
	}
	return func(flags *pflag.FlagSet) error {
		flags.BoolVarP(variable, name, shorthand, value, usage)
		return nil
	}
}

// WithDurationVarE adds a duration flag to the command with the specified name,
// value, and usage. If the environment variable envVar is set, its value is
// used as the default instead of value. The variable stores the final flag value.
func WithDurationVarE(variable *time.Duration, name, envVar string, value time.Duration, usage string) FlagInitializer {
	if v, ok := os.LookupEnv(envVar); ok {
		parsed, err := time.ParseDuration(v)
		if err != nil {
			return func(*pflag.FlagSet) error {
				return fmt.Errorf("%w: %s=%q: %v", ErrFlagEnvVarInvalid, envVar, v, err)
			}
		}
		value = parsed
	}
	return func(flags *pflag.FlagSet) error {
		flags.DurationVar(variable, name, value, usage)
		return nil
	}
}

// WithDurationVarPE adds a duration flag to the command with the specified
// name, shorthand, value, and usage. If the environment variable envVar is
// set, its value is used as the default instead of value. The variable stores
// the final flag value.
func WithDurationVarPE(variable *time.Duration, name, shorthand, envVar string, value time.Duration, usage string) FlagInitializer {
	if v, ok := os.LookupEnv(envVar); ok {
		parsed, err := time.ParseDuration(v)
		if err != nil {
			return func(*pflag.FlagSet) error {
				return fmt.Errorf("%w: %s=%q: %v", ErrFlagEnvVarInvalid, envVar, v, err)
			}
		}
		value = parsed
	}
	return func(flags *pflag.FlagSet) error {
		flags.DurationVarP(variable, name, shorthand, value, usage)
		return nil
	}
}

// WithFloat32VarE adds a float32 flag to the command with the specified name,
// value, and usage. If the environment variable envVar is set, its value is
// used as the default instead of value. The variable stores the final flag value.
func WithFloat32VarE(variable *float32, name, envVar string, value float32, usage string) FlagInitializer {
	if v, ok := os.LookupEnv(envVar); ok {
		parsed, err := strconv.ParseFloat(v, 32)
		if err != nil {
			return func(*pflag.FlagSet) error {
				return fmt.Errorf("%w: %s=%q: %v", ErrFlagEnvVarInvalid, envVar, v, err)
			}
		}
		value = float32(parsed)
	}
	return func(flags *pflag.FlagSet) error {
		flags.Float32Var(variable, name, value, usage)
		return nil
	}
}

// WithFloat32VarPE adds a float32 flag to the command with the specified name,
// shorthand, value, and usage. If the environment variable envVar is set, its
// value is used as the default instead of value. The variable stores the final
// flag value.
func WithFloat32VarPE(variable *float32, name, shorthand, envVar string, value float32, usage string) FlagInitializer {
	if v, ok := os.LookupEnv(envVar); ok {
		parsed, err := strconv.ParseFloat(v, 32)
		if err != nil {
			return func(*pflag.FlagSet) error {
				return fmt.Errorf("%w: %s=%q: %v", ErrFlagEnvVarInvalid, envVar, v, err)
			}
		}
		value = float32(parsed)
	}
	return func(flags *pflag.FlagSet) error {
		flags.Float32VarP(variable, name, shorthand, value, usage)
		return nil
	}
}

// WithFloat64VarE adds a float64 flag to the command with the specified name,
// value, and usage. If the environment variable envVar is set, its value is
// used as the default instead of value. The variable stores the final flag value.
func WithFloat64VarE(variable *float64, name, envVar string, value float64, usage string) FlagInitializer {
	if v, ok := os.LookupEnv(envVar); ok {
		parsed, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return func(*pflag.FlagSet) error {
				return fmt.Errorf("%w: %s=%q: %v", ErrFlagEnvVarInvalid, envVar, v, err)
			}
		}
		value = parsed
	}
	return func(flags *pflag.FlagSet) error {
		flags.Float64Var(variable, name, value, usage)
		return nil
	}
}

// WithFloat64VarPE adds a float64 flag to the command with the specified name,
// shorthand, value, and usage. If the environment variable envVar is set, its
// value is used as the default instead of value. The variable stores the final
// flag value.
func WithFloat64VarPE(variable *float64, name, shorthand, envVar string, value float64, usage string) FlagInitializer {
	if v, ok := os.LookupEnv(envVar); ok {
		parsed, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return func(*pflag.FlagSet) error {
				return fmt.Errorf("%w: %s=%q: %v", ErrFlagEnvVarInvalid, envVar, v, err)
			}
		}
		value = parsed
	}
	return func(flags *pflag.FlagSet) error {
		flags.Float64VarP(variable, name, shorthand, value, usage)
		return nil
	}
}

// WithIntVarE adds an int flag to the command with the specified name, value,
// and usage. If the environment variable envVar is set, its value is used as
// the default instead of value. The variable stores the final flag value.
func WithIntVarE(variable *int, name, envVar string, value int, usage string) FlagInitializer {
	if v, ok := os.LookupEnv(envVar); ok {
		parsed, err := strconv.Atoi(v)
		if err != nil {
			return func(*pflag.FlagSet) error {
				return fmt.Errorf("%w: %s=%q: %v", ErrFlagEnvVarInvalid, envVar, v, err)
			}
		}
		value = parsed
	}
	return func(flags *pflag.FlagSet) error {
		flags.IntVar(variable, name, value, usage)
		return nil
	}
}

// WithIntVarPE adds an int flag to the command with the specified name,
// shorthand, value, and usage. If the environment variable envVar is set, its
// value is used as the default instead of value. The variable stores the final
// flag value.
func WithIntVarPE(variable *int, name, shorthand, envVar string, value int, usage string) FlagInitializer {
	if v, ok := os.LookupEnv(envVar); ok {
		parsed, err := strconv.Atoi(v)
		if err != nil {
			return func(*pflag.FlagSet) error {
				return fmt.Errorf("%w: %s=%q: %v", ErrFlagEnvVarInvalid, envVar, v, err)
			}
		}
		value = parsed
	}
	return func(flags *pflag.FlagSet) error {
		flags.IntVarP(variable, name, shorthand, value, usage)
		return nil
	}
}

// WithInt8VarE adds an int8 flag to the command with the specified name, value,
// and usage. If the environment variable envVar is set, its value is used as
// the default instead of value. The variable stores the final flag value.
func WithInt8VarE(variable *int8, name, envVar string, value int8, usage string) FlagInitializer {
	if v, ok := os.LookupEnv(envVar); ok {
		parsed, err := strconv.ParseInt(v, 10, 8)
		if err != nil {
			return func(*pflag.FlagSet) error {
				return fmt.Errorf("%w: %s=%q: %v", ErrFlagEnvVarInvalid, envVar, v, err)
			}
		}
		value = int8(parsed)
	}
	return func(flags *pflag.FlagSet) error {
		flags.Int8Var(variable, name, value, usage)
		return nil
	}
}

// WithInt8VarPE adds an int8 flag to the command with the specified name,
// shorthand, value, and usage. If the environment variable envVar is set, its
// value is used as the default instead of value. The variable stores the final
// flag value.
func WithInt8VarPE(variable *int8, name, shorthand, envVar string, value int8, usage string) FlagInitializer {
	if v, ok := os.LookupEnv(envVar); ok {
		parsed, err := strconv.ParseInt(v, 10, 8)
		if err != nil {
			return func(*pflag.FlagSet) error {
				return fmt.Errorf("%w: %s=%q: %v", ErrFlagEnvVarInvalid, envVar, v, err)
			}
		}
		value = int8(parsed)
	}
	return func(flags *pflag.FlagSet) error {
		flags.Int8VarP(variable, name, shorthand, value, usage)
		return nil
	}
}

// WithInt16VarE adds an int16 flag to the command with the specified name,
// value, and usage. If the environment variable envVar is set, its value is
// used as the default instead of value. The variable stores the final flag value.
func WithInt16VarE(variable *int16, name, envVar string, value int16, usage string) FlagInitializer {
	if v, ok := os.LookupEnv(envVar); ok {
		parsed, err := strconv.ParseInt(v, 10, 16)
		if err != nil {
			return func(*pflag.FlagSet) error {
				return fmt.Errorf("%w: %s=%q: %v", ErrFlagEnvVarInvalid, envVar, v, err)
			}
		}
		value = int16(parsed)
	}
	return func(flags *pflag.FlagSet) error {
		flags.Int16Var(variable, name, value, usage)
		return nil
	}
}

// WithInt16VarPE adds an int16 flag to the command with the specified name,
// shorthand, value, and usage. If the environment variable envVar is set, its
// value is used as the default instead of value. The variable stores the final
// flag value.
func WithInt16VarPE(variable *int16, name, shorthand, envVar string, value int16, usage string) FlagInitializer {
	if v, ok := os.LookupEnv(envVar); ok {
		parsed, err := strconv.ParseInt(v, 10, 16)
		if err != nil {
			return func(*pflag.FlagSet) error {
				return fmt.Errorf("%w: %s=%q: %v", ErrFlagEnvVarInvalid, envVar, v, err)
			}
		}
		value = int16(parsed)
	}
	return func(flags *pflag.FlagSet) error {
		flags.Int16VarP(variable, name, shorthand, value, usage)
		return nil
	}
}

// WithInt32VarE adds an int32 flag to the command with the specified name,
// value, and usage. If the environment variable envVar is set, its value is
// used as the default instead of value. The variable stores the final flag value.
func WithInt32VarE(variable *int32, name, envVar string, value int32, usage string) FlagInitializer {
	if v, ok := os.LookupEnv(envVar); ok {
		parsed, err := strconv.ParseInt(v, 10, 32)
		if err != nil {
			return func(*pflag.FlagSet) error {
				return fmt.Errorf("%w: %s=%q: %v", ErrFlagEnvVarInvalid, envVar, v, err)
			}
		}
		value = int32(parsed)
	}
	return func(flags *pflag.FlagSet) error {
		flags.Int32Var(variable, name, value, usage)
		return nil
	}
}

// WithInt32VarPE adds an int32 flag to the command with the specified name,
// shorthand, value, and usage. If the environment variable envVar is set, its
// value is used as the default instead of value. The variable stores the final
// flag value.
func WithInt32VarPE(variable *int32, name, shorthand, envVar string, value int32, usage string) FlagInitializer {
	if v, ok := os.LookupEnv(envVar); ok {
		parsed, err := strconv.ParseInt(v, 10, 32)
		if err != nil {
			return func(*pflag.FlagSet) error {
				return fmt.Errorf("%w: %s=%q: %v", ErrFlagEnvVarInvalid, envVar, v, err)
			}
		}
		value = int32(parsed)
	}
	return func(flags *pflag.FlagSet) error {
		flags.Int32VarP(variable, name, shorthand, value, usage)
		return nil
	}
}

// WithInt64VarE adds an int64 flag to the command with the specified name,
// value, and usage. If the environment variable envVar is set, its value is
// used as the default instead of value. The variable stores the final flag value.
func WithInt64VarE(variable *int64, name, envVar string, value int64, usage string) FlagInitializer {
	if v, ok := os.LookupEnv(envVar); ok {
		parsed, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return func(*pflag.FlagSet) error {
				return fmt.Errorf("%w: %s=%q: %v", ErrFlagEnvVarInvalid, envVar, v, err)
			}
		}
		value = parsed
	}
	return func(flags *pflag.FlagSet) error {
		flags.Int64Var(variable, name, value, usage)
		return nil
	}
}

// WithInt64VarPE adds an int64 flag to the command with the specified name,
// shorthand, value, and usage. If the environment variable envVar is set, its
// value is used as the default instead of value. The variable stores the final
// flag value.
func WithInt64VarPE(variable *int64, name, shorthand, envVar string, value int64, usage string) FlagInitializer {
	if v, ok := os.LookupEnv(envVar); ok {
		parsed, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return func(*pflag.FlagSet) error {
				return fmt.Errorf("%w: %s=%q: %v", ErrFlagEnvVarInvalid, envVar, v, err)
			}
		}
		value = parsed
	}
	return func(flags *pflag.FlagSet) error {
		flags.Int64VarP(variable, name, shorthand, value, usage)
		return nil
	}
}

// WithUintVarE adds a uint flag to the command with the specified name, value,
// and usage. If the environment variable envVar is set, its value is used as
// the default instead of value. The variable stores the final flag value.
func WithUintVarE(variable *uint, name, envVar string, value uint, usage string) FlagInitializer {
	if v, ok := os.LookupEnv(envVar); ok {
		parsed, err := strconv.ParseUint(v, 10, strconv.IntSize)
		if err != nil {
			return func(*pflag.FlagSet) error {
				return fmt.Errorf("%w: %s=%q: %v", ErrFlagEnvVarInvalid, envVar, v, err)
			}
		}
		value = uint(parsed)
	}
	return func(flags *pflag.FlagSet) error {
		flags.UintVar(variable, name, value, usage)
		return nil
	}
}

// WithUintVarPE adds a uint flag to the command with the specified name,
// shorthand, value, and usage. If the environment variable envVar is set, its
// value is used as the default instead of value. The variable stores the final
// flag value.
func WithUintVarPE(variable *uint, name, shorthand, envVar string, value uint, usage string) FlagInitializer {
	if v, ok := os.LookupEnv(envVar); ok {
		parsed, err := strconv.ParseUint(v, 10, strconv.IntSize)
		if err != nil {
			return func(*pflag.FlagSet) error {
				return fmt.Errorf("%w: %s=%q: %v", ErrFlagEnvVarInvalid, envVar, v, err)
			}
		}
		value = uint(parsed)
	}
	return func(flags *pflag.FlagSet) error {
		flags.UintVarP(variable, name, shorthand, value, usage)
		return nil
	}
}

// WithUint8VarE adds a uint8 flag to the command with the specified name,
// value, and usage. If the environment variable envVar is set, its value is
// used as the default instead of value. The variable stores the final flag value.
func WithUint8VarE(variable *uint8, name, envVar string, value uint8, usage string) FlagInitializer {
	if v, ok := os.LookupEnv(envVar); ok {
		parsed, err := strconv.ParseUint(v, 10, 8)
		if err != nil {
			return func(*pflag.FlagSet) error {
				return fmt.Errorf("%w: %s=%q: %v", ErrFlagEnvVarInvalid, envVar, v, err)
			}
		}
		value = uint8(parsed)
	}
	return func(flags *pflag.FlagSet) error {
		flags.Uint8Var(variable, name, value, usage)
		return nil
	}
}

// WithUint8VarPE adds a uint8 flag to the command with the specified name,
// shorthand, value, and usage. If the environment variable envVar is set, its
// value is used as the default instead of value. The variable stores the final
// flag value.
func WithUint8VarPE(variable *uint8, name, shorthand, envVar string, value uint8, usage string) FlagInitializer {
	if v, ok := os.LookupEnv(envVar); ok {
		parsed, err := strconv.ParseUint(v, 10, 8)
		if err != nil {
			return func(*pflag.FlagSet) error {
				return fmt.Errorf("%w: %s=%q: %v", ErrFlagEnvVarInvalid, envVar, v, err)
			}
		}
		value = uint8(parsed)
	}
	return func(flags *pflag.FlagSet) error {
		flags.Uint8VarP(variable, name, shorthand, value, usage)
		return nil
	}
}

// WithUint16VarE adds a uint16 flag to the command with the specified name,
// value, and usage. If the environment variable envVar is set, its value is
// used as the default instead of value. The variable stores the final flag value.
func WithUint16VarE(variable *uint16, name, envVar string, value uint16, usage string) FlagInitializer {
	if v, ok := os.LookupEnv(envVar); ok {
		parsed, err := strconv.ParseUint(v, 10, 16)
		if err != nil {
			return func(*pflag.FlagSet) error {
				return fmt.Errorf("%w: %s=%q: %v", ErrFlagEnvVarInvalid, envVar, v, err)
			}
		}
		value = uint16(parsed)
	}
	return func(flags *pflag.FlagSet) error {
		flags.Uint16Var(variable, name, value, usage)
		return nil
	}
}

// WithUint16VarPE adds a uint16 flag to the command with the specified name,
// shorthand, value, and usage. If the environment variable envVar is set, its
// value is used as the default instead of value. The variable stores the final
// flag value.
func WithUint16VarPE(variable *uint16, name, shorthand, envVar string, value uint16, usage string) FlagInitializer {
	if v, ok := os.LookupEnv(envVar); ok {
		parsed, err := strconv.ParseUint(v, 10, 16)
		if err != nil {
			return func(*pflag.FlagSet) error {
				return fmt.Errorf("%w: %s=%q: %v", ErrFlagEnvVarInvalid, envVar, v, err)
			}
		}
		value = uint16(parsed)
	}
	return func(flags *pflag.FlagSet) error {
		flags.Uint16VarP(variable, name, shorthand, value, usage)
		return nil
	}
}

// WithUint32VarE adds a uint32 flag to the command with the specified name,
// value, and usage. If the environment variable envVar is set, its value is
// used as the default instead of value. The variable stores the final flag value.
func WithUint32VarE(variable *uint32, name, envVar string, value uint32, usage string) FlagInitializer {
	if v, ok := os.LookupEnv(envVar); ok {
		parsed, err := strconv.ParseUint(v, 10, 32)
		if err != nil {
			return func(*pflag.FlagSet) error {
				return fmt.Errorf("%w: %s=%q: %v", ErrFlagEnvVarInvalid, envVar, v, err)
			}
		}
		value = uint32(parsed)
	}
	return func(flags *pflag.FlagSet) error {
		flags.Uint32Var(variable, name, value, usage)
		return nil
	}
}

// WithUint32VarPE adds a uint32 flag to the command with the specified name,
// shorthand, value, and usage. If the environment variable envVar is set, its
// value is used as the default instead of value. The variable stores the final
// flag value.
func WithUint32VarPE(variable *uint32, name, shorthand, envVar string, value uint32, usage string) FlagInitializer {
	if v, ok := os.LookupEnv(envVar); ok {
		parsed, err := strconv.ParseUint(v, 10, 32)
		if err != nil {
			return func(*pflag.FlagSet) error {
				return fmt.Errorf("%w: %s=%q: %v", ErrFlagEnvVarInvalid, envVar, v, err)
			}
		}
		value = uint32(parsed)
	}
	return func(flags *pflag.FlagSet) error {
		flags.Uint32VarP(variable, name, shorthand, value, usage)
		return nil
	}
}

// WithUint64VarE adds a uint64 flag to the command with the specified name,
// value, and usage. If the environment variable envVar is set, its value is
// used as the default instead of value. The variable stores the final flag value.
func WithUint64VarE(variable *uint64, name, envVar string, value uint64, usage string) FlagInitializer {
	if v, ok := os.LookupEnv(envVar); ok {
		parsed, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			return func(*pflag.FlagSet) error {
				return fmt.Errorf("%w: %s=%q: %v", ErrFlagEnvVarInvalid, envVar, v, err)
			}
		}
		value = parsed
	}
	return func(flags *pflag.FlagSet) error {
		flags.Uint64Var(variable, name, value, usage)
		return nil
	}
}

// WithUint64VarPE adds a uint64 flag to the command with the specified name,
// shorthand, value, and usage. If the environment variable envVar is set, its
// value is used as the default instead of value. The variable stores the final
// flag value.
func WithUint64VarPE(variable *uint64, name, shorthand, envVar string, value uint64, usage string) FlagInitializer {
	if v, ok := os.LookupEnv(envVar); ok {
		parsed, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			return func(*pflag.FlagSet) error {
				return fmt.Errorf("%w: %s=%q: %v", ErrFlagEnvVarInvalid, envVar, v, err)
			}
		}
		value = parsed
	}
	return func(flags *pflag.FlagSet) error {
		flags.Uint64VarP(variable, name, shorthand, value, usage)
		return nil
	}
}

// WithStringVarE adds a string flag to the command with the specified name,
// value, and usage. If the environment variable envVar is set, its value is
// used as the default instead of value. The variable stores the final flag value.
func WithStringVarE(variable *string, name, envVar, value, usage string) FlagInitializer {
	if v, ok := os.LookupEnv(envVar); ok {
		value = v
	}
	return func(flags *pflag.FlagSet) error {
		flags.StringVar(variable, name, value, usage)
		return nil
	}
}

// WithStringVarPE adds a string flag to the command with the specified name,
// shorthand, value, and usage. If the environment variable envVar is set, its
// value is used as the default instead of value. The variable stores the final
// flag value.
func WithStringVarPE(variable *string, name, shorthand, envVar, value, usage string) FlagInitializer {
	if v, ok := os.LookupEnv(envVar); ok {
		value = v
	}
	return func(flags *pflag.FlagSet) error {
		flags.StringVarP(variable, name, shorthand, value, usage)
		return nil
	}
}
