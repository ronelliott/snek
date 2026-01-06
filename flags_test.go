package snek_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ronelliott/snek"
)

func TestWithBoolVar(t *testing.T) {
	runFlagTest(t,
		func(variable *bool, value bool) snek.FlagInitializer {
			return snek.WithBoolVar(variable, "test", value, "test flag")
		},
		map[string]flagTest[bool]{
			"long flag": {
				args:     []string{"--test"},
				expected: true,
			},
		})
}

func TestWithBoolVarP(t *testing.T) {
	runFlagTest(t,
		func(variable *bool, value bool) snek.FlagInitializer {
			return snek.WithBoolVarP(variable, "test", "t", value, "test flag")
		},
		map[string]flagTest[bool]{
			"long flag": {
				args:     []string{"--test"},
				expected: true,
			},
			"short flag": {
				args:     []string{"-t"},
				expected: true,
			},
		})
}

func TestWithDurationVar(t *testing.T) {
	runFlagTest(t,
		func(variable *time.Duration, value time.Duration) snek.FlagInitializer {
			return snek.WithDurationVar(variable, "test", value, "test duration")
		},
		map[string]flagTest[time.Duration]{
			"long flag with =": {
				args:     []string{"--test=5s"},
				expected: time.Second * 5,
			},
			"long flag with space": {
				args:     []string{"--test", "30s"},
				expected: time.Second * 30,
			},
		})
}

func TestWithDurationVarP(t *testing.T) {
	runFlagTest(t,
		func(variable *time.Duration, value time.Duration) snek.FlagInitializer {
			return snek.WithDurationVarP(variable, "test", "t", value, "test duration")
		},
		map[string]flagTest[time.Duration]{
			"short flag with =": {
				args:     []string{"-t=1s"},
				expected: time.Second,
			},
			"short flag with space": {
				args:     []string{"-t", "1m"},
				expected: time.Minute,
			},
			"long flag with =": {
				args:     []string{"--test=5s"},
				expected: time.Second * 5,
			},
			"long flag with space": {
				args:     []string{"--test", "30s"},
				expected: time.Second * 30,
			},
		})
}

func TestWithFloat32Var(t *testing.T) {
	runFlagTest(t,
		func(variable *float32, value float32) snek.FlagInitializer {
			return snek.WithFloat32Var(variable, "test", value, "test float32")
		},
		map[string]flagTest[float32]{
			"long flag with =": {
				args:     []string{"--test=1.5"},
				expected: 1.5,
			},
			"long flag with space": {
				args:     []string{"--test", "3.14"},
				expected: 3.14,
			},
		})
}

func TestWithFloat32VarP(t *testing.T) {
	runFlagTest(t,
		func(variable *float32, value float32) snek.FlagInitializer {
			return snek.WithFloat32VarP(variable, "test", "t", value, "test float32")
		},
		map[string]flagTest[float32]{
			"short flag with =": {
				args:     []string{"-t=1.5"},
				expected: 1.5,
			},
			"short flag with space": {
				args:     []string{"-t", "3.14"},
				expected: 3.14,
			},
			"long flag with =": {
				args:     []string{"--test=1.5"},
				expected: 1.5,
			},
			"long flag with space": {
				args:     []string{"--test", "3.14"},
				expected: 3.14,
			},
		})
}

func TestWithFloat64Var(t *testing.T) {
	runFlagTest(t,
		func(variable *float64, value float64) snek.FlagInitializer {
			return snek.WithFloat64Var(variable, "test", value, "test float64")
		},
		map[string]flagTest[float64]{
			"long flag with =": {
				args:     []string{"--test=1.5"},
				expected: 1.5,
			},
			"long flag with space": {
				args:     []string{"--test", "3.14"},
				expected: 3.14,
			},
		})
}

func TestWithFloat64VarP(t *testing.T) {
	runFlagTest(t,
		func(variable *float64, value float64) snek.FlagInitializer {
			return snek.WithFloat64VarP(variable, "test", "t", value, "test float64")
		},
		map[string]flagTest[float64]{
			"short flag with =": {
				args:     []string{"-t=1.5"},
				expected: 1.5,
			},
			"short flag with space": {
				args:     []string{"-t", "3.14"},
				expected: 3.14,
			},
			"long flag with =": {
				args:     []string{"--test=1.5"},
				expected: 1.5,
			},
			"long flag with space": {
				args:     []string{"--test", "3.14"},
				expected: 3.14,
			},
		})
}

func TestWithIntVar(t *testing.T) {
	runFlagTest(t,
		func(variable *int, value int) snek.FlagInitializer {
			return snek.WithIntVar(variable, "test", value, "test int")
		},
		map[string]flagTest[int]{
			"long flag with =": {
				args:     []string{"--test=1"},
				expected: 1,
			},
			"long flag with space": {
				args:     []string{"--test", "2"},
				expected: 2,
			},
		})
}

func TestWithIntVarP(t *testing.T) {
	runFlagTest(t,
		func(variable *int, value int) snek.FlagInitializer {
			return snek.WithIntVarP(variable, "test", "t", value, "test int")
		},
		map[string]flagTest[int]{
			"short flag with =": {
				args:     []string{"-t=1"},
				expected: 1,
			},
			"short flag with space": {
				args:     []string{"-t", "2"},
				expected: 2,
			},
			"long flag with =": {
				args:     []string{"--test=1"},
				expected: 1,
			},
			"long flag with space": {
				args:     []string{"--test", "2"},
				expected: 2,
			},
		})
}

func TestWithInt64Var(t *testing.T) {
	runFlagTest(t,
		func(variable *int64, value int64) snek.FlagInitializer {
			return snek.WithInt64Var(variable, "test", value, "test int64")
		},
		map[string]flagTest[int64]{
			"long flag with =": {
				args:     []string{"--test=1"},
				expected: 1,
			},
			"long flag with space": {
				args:     []string{"--test", "2"},
				expected: 2,
			},
		})
}

func TestWithInt64VarP(t *testing.T) {
	runFlagTest(t,
		func(variable *int64, value int64) snek.FlagInitializer {
			return snek.WithInt64VarP(variable, "test", "t", value, "test int64")
		},
		map[string]flagTest[int64]{
			"short flag with =": {
				args:     []string{"-t=1"},
				expected: 1,
			},
			"short flag with space": {
				args:     []string{"-t", "2"},
				expected: 2,
			},
			"long flag with =": {
				args:     []string{"--test=1"},
				expected: 1,
			},
			"long flag with space": {
				args:     []string{"--test", "2"},
				expected: 2,
			},
		})
}

func TestWithInt8Var(t *testing.T) {
	runFlagTest(t,
		func(variable *int8, value int8) snek.FlagInitializer {
			return snek.WithInt8Var(variable, "test", value, "test int8")
		},
		map[string]flagTest[int8]{
			"long flag with =": {
				args:     []string{"--test=1"},
				expected: 1,
			},
			"long flag with space": {
				args:     []string{"--test", "2"},
				expected: 2,
			},
		})
}

func TestWithInt8VarP(t *testing.T) {
	runFlagTest(t,
		func(variable *int8, value int8) snek.FlagInitializer {
			return snek.WithInt8VarP(variable, "test", "t", value, "test int8")
		},
		map[string]flagTest[int8]{
			"short flag with =": {
				args:     []string{"-t=1"},
				expected: 1,
			},
			"short flag with space": {
				args:     []string{"-t", "2"},
				expected: 2,
			},
			"long flag with =": {
				args:     []string{"--test=1"},
				expected: 1,
			},
			"long flag with space": {
				args:     []string{"--test", "2"},
				expected: 2,
			},
		})
}

func TestWithInt16Var(t *testing.T) {
	runFlagTest(t,
		func(variable *int16, value int16) snek.FlagInitializer {
			return snek.WithInt16Var(variable, "test", value, "test int16")
		},
		map[string]flagTest[int16]{
			"long flag with =": {
				args:     []string{"--test=1"},
				expected: 1,
			},
			"long flag with space": {
				args:     []string{"--test", "2"},
				expected: 2,
			},
		})
}

func TestWithInt16VarP(t *testing.T) {
	runFlagTest(t,
		func(variable *int16, value int16) snek.FlagInitializer {
			return snek.WithInt16VarP(variable, "test", "t", value, "test int16")
		},
		map[string]flagTest[int16]{
			"short flag with =": {
				args:     []string{"-t=1"},
				expected: 1,
			},
			"short flag with space": {
				args:     []string{"-t", "2"},
				expected: 2,
			},
			"long flag with =": {
				args:     []string{"--test=1"},
				expected: 1,
			},
			"long flag with space": {
				args:     []string{"--test", "2"},
				expected: 2,
			},
		})
}

func TestWithInt32Var(t *testing.T) {
	runFlagTest(t,
		func(variable *int32, value int32) snek.FlagInitializer {
			return snek.WithInt32Var(variable, "test", value, "test int32")
		},
		map[string]flagTest[int32]{
			"long flag with =": {
				args:     []string{"--test=1"},
				expected: 1,
			},
			"long flag with space": {
				args:     []string{"--test", "2"},
				expected: 2,
			},
		})
}

func TestWithInt32VarP(t *testing.T) {
	runFlagTest(t,
		func(variable *int32, value int32) snek.FlagInitializer {
			return snek.WithInt32VarP(variable, "test", "t", value, "test int32")
		},
		map[string]flagTest[int32]{
			"short flag with =": {
				args:     []string{"-t=1"},
				expected: 1,
			},
			"short flag with space": {
				args:     []string{"-t", "2"},
				expected: 2,
			},
			"long flag with =": {
				args:     []string{"--test=1"},
				expected: 1,
			},
			"long flag with space": {
				args:     []string{"--test", "2"},
				expected: 2,
			},
		})
}

func TestWithUintVar(t *testing.T) {
	runFlagTest(t,
		func(variable *uint, value uint) snek.FlagInitializer {
			return snek.WithUintVar(variable, "test", value, "test uint")
		},
		map[string]flagTest[uint]{
			"long flag with =": {
				args:     []string{"--test=1"},
				expected: 1,
			},
			"long flag with space": {
				args:     []string{"--test", "2"},
				expected: 2,
			},
		})
}

func TestWithUintVarP(t *testing.T) {
	runFlagTest(t,
		func(variable *uint, value uint) snek.FlagInitializer {
			return snek.WithUintVarP(variable, "test", "t", value, "test uint")
		},
		map[string]flagTest[uint]{
			"short flag with =": {
				args:     []string{"-t=1"},
				expected: 1,
			},
			"short flag with space": {
				args:     []string{"-t", "2"},
				expected: 2,
			},
			"long flag with =": {
				args:     []string{"--test=1"},
				expected: 1,
			},
			"long flag with space": {
				args:     []string{"--test", "2"},
				expected: 2,
			},
		})
}

func TestWithUint8Var(t *testing.T) {
	runFlagTest(t,
		func(variable *uint8, value uint8) snek.FlagInitializer {
			return snek.WithUint8Var(variable, "test", value, "test uint8")
		},
		map[string]flagTest[uint8]{
			"long flag with =": {
				args:     []string{"--test=1"},
				expected: 1,
			},
			"long flag with space": {
				args:     []string{"--test", "2"},
				expected: 2,
			},
		})
}

func TestWithUint8VarP(t *testing.T) {
	runFlagTest(t,
		func(variable *uint8, value uint8) snek.FlagInitializer {
			return snek.WithUint8VarP(variable, "test", "t", value, "test uint8")
		},
		map[string]flagTest[uint8]{
			"short flag with =": {
				args:     []string{"-t=1"},
				expected: 1,
			},
			"short flag with space": {
				args:     []string{"-t", "2"},
				expected: 2,
			},
			"long flag with =": {
				args:     []string{"--test=1"},
				expected: 1,
			},
			"long flag with space": {
				args:     []string{"--test", "2"},
				expected: 2,
			},
		})
}

func TestWithUint16Var(t *testing.T) {
	runFlagTest(t,
		func(variable *uint16, value uint16) snek.FlagInitializer {
			return snek.WithUint16Var(variable, "test", value, "test uint16")
		},
		map[string]flagTest[uint16]{
			"long flag with =": {
				args:     []string{"--test=1"},
				expected: 1,
			},
			"long flag with space": {
				args:     []string{"--test", "2"},
				expected: 2,
			},
		})
}

func TestWithUint16VarP(t *testing.T) {
	runFlagTest(t,
		func(variable *uint16, value uint16) snek.FlagInitializer {
			return snek.WithUint16VarP(variable, "test", "t", value, "test uint16")
		},
		map[string]flagTest[uint16]{
			"short flag with =": {
				args:     []string{"-t=1"},
				expected: 1,
			},
			"short flag with space": {
				args:     []string{"-t", "2"},
				expected: 2,
			},
			"long flag with =": {
				args:     []string{"--test=1"},
				expected: 1,
			},
			"long flag with space": {
				args:     []string{"--test", "2"},
				expected: 2,
			},
		})
}

func TestWithUint32Var(t *testing.T) {
	runFlagTest(t,
		func(variable *uint32, value uint32) snek.FlagInitializer {
			return snek.WithUint32Var(variable, "test", value, "test uint32")
		},
		map[string]flagTest[uint32]{
			"long flag with =": {
				args:     []string{"--test=1"},
				expected: 1,
			},
			"long flag with space": {
				args:     []string{"--test", "2"},
				expected: 2,
			},
		})
}

func TestWithUint32VarP(t *testing.T) {
	runFlagTest(t,
		func(variable *uint32, value uint32) snek.FlagInitializer {
			return snek.WithUint32VarP(variable, "test", "t", value, "test uint32")
		},
		map[string]flagTest[uint32]{
			"short flag with =": {
				args:     []string{"-t=1"},
				expected: 1,
			},
			"short flag with space": {
				args:     []string{"-t", "2"},
				expected: 2,
			},
			"long flag with =": {
				args:     []string{"--test=1"},
				expected: 1,
			},
			"long flag with space": {
				args:     []string{"--test", "2"},
				expected: 2,
			},
		})
}

func TestWithUint64Var(t *testing.T) {
	runFlagTest(t,
		func(variable *uint64, value uint64) snek.FlagInitializer {
			return snek.WithUint64Var(variable, "test", value, "test uint64")
		},
		map[string]flagTest[uint64]{
			"long flag with =": {
				args:     []string{"--test=1"},
				expected: 1,
			},
			"long flag with space": {
				args:     []string{"--test", "2"},
				expected: 2,
			},
		})
}

func TestWithUint64VarP(t *testing.T) {
	runFlagTest(t,
		func(variable *uint64, value uint64) snek.FlagInitializer {
			return snek.WithUint64VarP(variable, "test", "t", value, "test uint64")
		},
		map[string]flagTest[uint64]{
			"short flag with =": {
				args:     []string{"-t=1"},
				expected: 1,
			},
			"short flag with space": {
				args:     []string{"-t", "2"},
				expected: 2,
			},
			"long flag with =": {
				args:     []string{"--test=1"},
				expected: 1,
			},
			"long flag with space": {
				args:     []string{"--test", "2"},
				expected: 2,
			},
		})
}

func TestWithStringVar(t *testing.T) {
	runFlagTest(t,
		func(variable *string, value string) snek.FlagInitializer {
			return snek.WithStringVar(variable, "test", value, "test string")
		},
		map[string]flagTest[string]{
			"long flag with =": {
				args:     []string{"--test=foo"},
				expected: "foo",
			},
			"long flag with space": {
				args:     []string{"--test", "bar"},
				expected: "bar",
			},
		})
}

func TestWithStringVarP(t *testing.T) {
	runFlagTest(t,
		func(variable *string, value string) snek.FlagInitializer {
			return snek.WithStringVarP(variable, "test", "t", value, "test string")
		},
		map[string]flagTest[string]{
			"short flag with =": {
				args:     []string{"-t=foo"},
				expected: "foo",
			},
			"short flag with space": {
				args:     []string{"-t", "bar"},
				expected: "bar",
			},
			"long flag with =": {
				args:     []string{"--test=foo"},
				expected: "foo",
			},
			"long flag with space": {
				args:     []string{"--test", "bar"},
				expected: "bar",
			},
		})
}

type flagTest[T any] struct {
	args     []string
	expected T
}

func runFlagTest[T any](
	t *testing.T,
	initializer func(*T, T) snek.FlagInitializer,
	tests map[string]flagTest[T],
) {
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			testValue := *new(T)
			require.Empty(t, testValue, "The test variable should be empty")

			cmd, err := snek.NewCommand(snek.WithFlag(initializer(&testValue, test.expected)))
			require.NoError(t, err, "NewCommand should not return an error")
			require.NotNil(t, cmd, "NewCommand should return a command")

			cmd.SetArgs(test.args)
			require.NoError(t, cmd.Execute(), "Execute should not return an error")
			assert.Equal(t, test.expected, testValue, "The parsed value should be the expected value")
		})
	}
}
