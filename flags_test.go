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
