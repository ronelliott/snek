package snek_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ronelliott/snek"
)

type flagEnvTest[T any] struct {
	args        []string
	envValue    string
	setEnv      bool
	expected    T
	wantInitErr bool
}

func runFlagEnvTest[T any](
	t *testing.T,
	envVar string,
	defaultValue T,
	initializer func(*T, string, T) snek.FlagInitializer,
	tests map[string]flagEnvTest[T],
) {
	t.Helper()
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if test.setEnv {
				t.Setenv(envVar, test.envValue)
			}

			var testValue T
			cmd, err := snek.NewCommand(snek.WithFlag(initializer(&testValue, envVar, defaultValue)))
			if test.wantInitErr {
				require.ErrorIs(t, err, snek.ErrFlagEnvVarInvalid)
				return
			}

			require.NoError(t, err)
			require.NotNil(t, cmd)

			cmd.SetArgs(test.args)
			require.NoError(t, cmd.Execute())
			assert.Equal(t, test.expected, testValue)
		})
	}
}

// ---- bool ----

func TestWithBoolVarE(t *testing.T) {
	runFlagEnvTest(t,
		"TEST_BOOL",
		false,
		func(variable *bool, envVar string, value bool) snek.FlagInitializer {
			return snek.WithBoolVarE(variable, "test", envVar, value, "test bool")
		},
		map[string]flagEnvTest[bool]{
			"no env var, no flag": {
				args:     []string{},
				expected: false,
			},
			"env var true, no flag": {
				args:     []string{},
				envValue: "true",
				setEnv:   true,
				expected: true,
			},
			"env var false, no flag": {
				args:     []string{},
				envValue: "false",
				setEnv:   false,
				expected: false,
			},
			"cli flag overrides env var false": {
				args:     []string{"--test"},
				envValue: "false",
				setEnv:   true,
				expected: true,
			},
			"invalid env var": {
				envValue:    "notabool",
				setEnv:      true,
				wantInitErr: true,
			},
		})
}

func TestWithBoolVarPE(t *testing.T) {
	runFlagEnvTest(t,
		"TEST_BOOL",
		false,
		func(variable *bool, envVar string, value bool) snek.FlagInitializer {
			return snek.WithBoolVarPE(variable, "test", "t", envVar, value, "test bool")
		},
		map[string]flagEnvTest[bool]{
			"no env var, no flag": {
				args:     []string{},
				expected: false,
			},
			"env var true, no flag": {
				args:     []string{},
				envValue: "true",
				setEnv:   true,
				expected: true,
			},
			"long flag overrides env var false": {
				args:     []string{"--test"},
				envValue: "false",
				setEnv:   true,
				expected: true,
			},
			"short flag overrides env var false": {
				args:     []string{"-t"},
				envValue: "false",
				setEnv:   true,
				expected: true,
			},
			"invalid env var": {
				envValue:    "notabool",
				setEnv:      true,
				wantInitErr: true,
			},
		})
}

// ---- duration ----

func TestWithDurationVarE(t *testing.T) {
	runFlagEnvTest(t,
		"TEST_DURATION",
		time.Duration(0),
		func(variable *time.Duration, envVar string, value time.Duration) snek.FlagInitializer {
			return snek.WithDurationVarE(variable, "test", envVar, value, "test duration")
		},
		map[string]flagEnvTest[time.Duration]{
			"no env var, no flag": {
				args:     []string{},
				expected: 0,
			},
			"env var set, no flag": {
				args:     []string{},
				envValue: "5s",
				setEnv:   true,
				expected: 5 * time.Second,
			},
			"cli flag overrides env var": {
				args:     []string{"--test=30s"},
				envValue: "5s",
				setEnv:   true,
				expected: 30 * time.Second,
			},
			"cli flag set, no env var": {
				args:     []string{"--test", "1m"},
				expected: time.Minute,
			},
			"invalid env var": {
				envValue:    "notaduration",
				setEnv:      true,
				wantInitErr: true,
			},
		})
}

func TestWithDurationVarPE(t *testing.T) {
	runFlagEnvTest(t,
		"TEST_DURATION",
		time.Duration(0),
		func(variable *time.Duration, envVar string, value time.Duration) snek.FlagInitializer {
			return snek.WithDurationVarPE(variable, "test", "t", envVar, value, "test duration")
		},
		map[string]flagEnvTest[time.Duration]{
			"no env var, no flag": {
				args:     []string{},
				expected: 0,
			},
			"env var set, no flag": {
				args:     []string{},
				envValue: "5s",
				setEnv:   true,
				expected: 5 * time.Second,
			},
			"long flag overrides env var": {
				args:     []string{"--test=30s"},
				envValue: "5s",
				setEnv:   true,
				expected: 30 * time.Second,
			},
			"short flag overrides env var": {
				args:     []string{"-t", "10s"},
				envValue: "5s",
				setEnv:   true,
				expected: 10 * time.Second,
			},
			"invalid env var": {
				envValue:    "notaduration",
				setEnv:      true,
				wantInitErr: true,
			},
		})
}

// ---- float32 ----

func TestWithFloat32VarE(t *testing.T) {
	runFlagEnvTest(t,
		"TEST_FLOAT32",
		float32(0),
		func(variable *float32, envVar string, value float32) snek.FlagInitializer {
			return snek.WithFloat32VarE(variable, "test", envVar, value, "test float32")
		},
		map[string]flagEnvTest[float32]{
			"no env var, no flag": {
				args:     []string{},
				expected: 0,
			},
			"env var set, no flag": {
				args:     []string{},
				envValue: "1.5",
				setEnv:   true,
				expected: 1.5,
			},
			"cli flag overrides env var": {
				args:     []string{"--test=3.14"},
				envValue: "1.5",
				setEnv:   true,
				expected: 3.14,
			},
			"cli flag set, no env var": {
				args:     []string{"--test", "2.5"},
				expected: 2.5,
			},
			"invalid env var": {
				envValue:    "notafloat",
				setEnv:      true,
				wantInitErr: true,
			},
		})
}

func TestWithFloat32VarPE(t *testing.T) {
	runFlagEnvTest(t,
		"TEST_FLOAT32",
		float32(0),
		func(variable *float32, envVar string, value float32) snek.FlagInitializer {
			return snek.WithFloat32VarPE(variable, "test", "t", envVar, value, "test float32")
		},
		map[string]flagEnvTest[float32]{
			"no env var, no flag": {
				args:     []string{},
				expected: 0,
			},
			"env var set, no flag": {
				args:     []string{},
				envValue: "1.5",
				setEnv:   true,
				expected: 1.5,
			},
			"long flag overrides env var": {
				args:     []string{"--test=3.14"},
				envValue: "1.5",
				setEnv:   true,
				expected: 3.14,
			},
			"short flag overrides env var": {
				args:     []string{"-t", "2.5"},
				envValue: "1.5",
				setEnv:   true,
				expected: 2.5,
			},
			"invalid env var": {
				envValue:    "notafloat",
				setEnv:      true,
				wantInitErr: true,
			},
		})
}

// ---- float64 ----

func TestWithFloat64VarE(t *testing.T) {
	runFlagEnvTest(t,
		"TEST_FLOAT64",
		float64(0),
		func(variable *float64, envVar string, value float64) snek.FlagInitializer {
			return snek.WithFloat64VarE(variable, "test", envVar, value, "test float64")
		},
		map[string]flagEnvTest[float64]{
			"no env var, no flag": {
				args:     []string{},
				expected: 0,
			},
			"env var set, no flag": {
				args:     []string{},
				envValue: "1.5",
				setEnv:   true,
				expected: 1.5,
			},
			"cli flag overrides env var": {
				args:     []string{"--test=3.14"},
				envValue: "1.5",
				setEnv:   true,
				expected: 3.14,
			},
			"cli flag set, no env var": {
				args:     []string{"--test", "2.5"},
				expected: 2.5,
			},
			"invalid env var": {
				envValue:    "notafloat",
				setEnv:      true,
				wantInitErr: true,
			},
		})
}

func TestWithFloat64VarPE(t *testing.T) {
	runFlagEnvTest(t,
		"TEST_FLOAT64",
		float64(0),
		func(variable *float64, envVar string, value float64) snek.FlagInitializer {
			return snek.WithFloat64VarPE(variable, "test", "t", envVar, value, "test float64")
		},
		map[string]flagEnvTest[float64]{
			"no env var, no flag": {
				args:     []string{},
				expected: 0,
			},
			"env var set, no flag": {
				args:     []string{},
				envValue: "1.5",
				setEnv:   true,
				expected: 1.5,
			},
			"long flag overrides env var": {
				args:     []string{"--test=3.14"},
				envValue: "1.5",
				setEnv:   true,
				expected: 3.14,
			},
			"short flag overrides env var": {
				args:     []string{"-t", "2.5"},
				envValue: "1.5",
				setEnv:   true,
				expected: 2.5,
			},
			"invalid env var": {
				envValue:    "notafloat",
				setEnv:      true,
				wantInitErr: true,
			},
		})
}

// ---- int ----

func TestWithIntVarE(t *testing.T) {
	runFlagEnvTest(t,
		"TEST_INT",
		0,
		func(variable *int, envVar string, value int) snek.FlagInitializer {
			return snek.WithIntVarE(variable, "test", envVar, value, "test int")
		},
		map[string]flagEnvTest[int]{
			"no env var, no flag": {
				args:     []string{},
				expected: 0,
			},
			"env var set, no flag": {
				args:     []string{},
				envValue: "42",
				setEnv:   true,
				expected: 42,
			},
			"cli flag overrides env var": {
				args:     []string{"--test=99"},
				envValue: "42",
				setEnv:   true,
				expected: 99,
			},
			"cli flag set, no env var": {
				args:     []string{"--test", "7"},
				expected: 7,
			},
			"invalid env var": {
				envValue:    "notanint",
				setEnv:      true,
				wantInitErr: true,
			},
		})
}

func TestWithIntVarPE(t *testing.T) {
	runFlagEnvTest(t,
		"TEST_INT",
		0,
		func(variable *int, envVar string, value int) snek.FlagInitializer {
			return snek.WithIntVarPE(variable, "test", "t", envVar, value, "test int")
		},
		map[string]flagEnvTest[int]{
			"no env var, no flag": {
				args:     []string{},
				expected: 0,
			},
			"env var set, no flag": {
				args:     []string{},
				envValue: "42",
				setEnv:   true,
				expected: 42,
			},
			"long flag overrides env var": {
				args:     []string{"--test=99"},
				envValue: "42",
				setEnv:   true,
				expected: 99,
			},
			"short flag overrides env var": {
				args:     []string{"-t", "7"},
				envValue: "42",
				setEnv:   true,
				expected: 7,
			},
			"invalid env var": {
				envValue:    "notanint",
				setEnv:      true,
				wantInitErr: true,
			},
		})
}

// ---- int8 ----

func TestWithInt8VarE(t *testing.T) {
	runFlagEnvTest(t,
		"TEST_INT8",
		int8(0),
		func(variable *int8, envVar string, value int8) snek.FlagInitializer {
			return snek.WithInt8VarE(variable, "test", envVar, value, "test int8")
		},
		map[string]flagEnvTest[int8]{
			"no env var, no flag": {
				args:     []string{},
				expected: 0,
			},
			"env var set, no flag": {
				args:     []string{},
				envValue: "42",
				setEnv:   true,
				expected: 42,
			},
			"cli flag overrides env var": {
				args:     []string{"--test=99"},
				envValue: "42",
				setEnv:   true,
				expected: 99,
			},
			"cli flag set, no env var": {
				args:     []string{"--test", "7"},
				expected: 7,
			},
			"invalid env var": {
				envValue:    "notanint",
				setEnv:      true,
				wantInitErr: true,
			},
		})
}

func TestWithInt8VarPE(t *testing.T) {
	runFlagEnvTest(t,
		"TEST_INT8",
		int8(0),
		func(variable *int8, envVar string, value int8) snek.FlagInitializer {
			return snek.WithInt8VarPE(variable, "test", "t", envVar, value, "test int8")
		},
		map[string]flagEnvTest[int8]{
			"no env var, no flag": {
				args:     []string{},
				expected: 0,
			},
			"env var set, no flag": {
				args:     []string{},
				envValue: "42",
				setEnv:   true,
				expected: 42,
			},
			"long flag overrides env var": {
				args:     []string{"--test=99"},
				envValue: "42",
				setEnv:   true,
				expected: 99,
			},
			"short flag overrides env var": {
				args:     []string{"-t", "7"},
				envValue: "42",
				setEnv:   true,
				expected: 7,
			},
			"invalid env var": {
				envValue:    "notanint",
				setEnv:      true,
				wantInitErr: true,
			},
		})
}

// ---- int16 ----

func TestWithInt16VarE(t *testing.T) {
	runFlagEnvTest(t,
		"TEST_INT16",
		int16(0),
		func(variable *int16, envVar string, value int16) snek.FlagInitializer {
			return snek.WithInt16VarE(variable, "test", envVar, value, "test int16")
		},
		map[string]flagEnvTest[int16]{
			"no env var, no flag": {
				args:     []string{},
				expected: 0,
			},
			"env var set, no flag": {
				args:     []string{},
				envValue: "42",
				setEnv:   true,
				expected: 42,
			},
			"cli flag overrides env var": {
				args:     []string{"--test=99"},
				envValue: "42",
				setEnv:   true,
				expected: 99,
			},
			"cli flag set, no env var": {
				args:     []string{"--test", "7"},
				expected: 7,
			},
			"invalid env var": {
				envValue:    "notanint",
				setEnv:      true,
				wantInitErr: true,
			},
		})
}

func TestWithInt16VarPE(t *testing.T) {
	runFlagEnvTest(t,
		"TEST_INT16",
		int16(0),
		func(variable *int16, envVar string, value int16) snek.FlagInitializer {
			return snek.WithInt16VarPE(variable, "test", "t", envVar, value, "test int16")
		},
		map[string]flagEnvTest[int16]{
			"no env var, no flag": {
				args:     []string{},
				expected: 0,
			},
			"env var set, no flag": {
				args:     []string{},
				envValue: "42",
				setEnv:   true,
				expected: 42,
			},
			"long flag overrides env var": {
				args:     []string{"--test=99"},
				envValue: "42",
				setEnv:   true,
				expected: 99,
			},
			"short flag overrides env var": {
				args:     []string{"-t", "7"},
				envValue: "42",
				setEnv:   true,
				expected: 7,
			},
			"invalid env var": {
				envValue:    "notanint",
				setEnv:      true,
				wantInitErr: true,
			},
		})
}

// ---- int32 ----

func TestWithInt32VarE(t *testing.T) {
	runFlagEnvTest(t,
		"TEST_INT32",
		int32(0),
		func(variable *int32, envVar string, value int32) snek.FlagInitializer {
			return snek.WithInt32VarE(variable, "test", envVar, value, "test int32")
		},
		map[string]flagEnvTest[int32]{
			"no env var, no flag": {
				args:     []string{},
				expected: 0,
			},
			"env var set, no flag": {
				args:     []string{},
				envValue: "42",
				setEnv:   true,
				expected: 42,
			},
			"cli flag overrides env var": {
				args:     []string{"--test=99"},
				envValue: "42",
				setEnv:   true,
				expected: 99,
			},
			"cli flag set, no env var": {
				args:     []string{"--test", "7"},
				expected: 7,
			},
			"invalid env var": {
				envValue:    "notanint",
				setEnv:      true,
				wantInitErr: true,
			},
		})
}

func TestWithInt32VarPE(t *testing.T) {
	runFlagEnvTest(t,
		"TEST_INT32",
		int32(0),
		func(variable *int32, envVar string, value int32) snek.FlagInitializer {
			return snek.WithInt32VarPE(variable, "test", "t", envVar, value, "test int32")
		},
		map[string]flagEnvTest[int32]{
			"no env var, no flag": {
				args:     []string{},
				expected: 0,
			},
			"env var set, no flag": {
				args:     []string{},
				envValue: "42",
				setEnv:   true,
				expected: 42,
			},
			"long flag overrides env var": {
				args:     []string{"--test=99"},
				envValue: "42",
				setEnv:   true,
				expected: 99,
			},
			"short flag overrides env var": {
				args:     []string{"-t", "7"},
				envValue: "42",
				setEnv:   true,
				expected: 7,
			},
			"invalid env var": {
				envValue:    "notanint",
				setEnv:      true,
				wantInitErr: true,
			},
		})
}

// ---- int64 ----

func TestWithInt64VarE(t *testing.T) {
	runFlagEnvTest(t,
		"TEST_INT64",
		int64(0),
		func(variable *int64, envVar string, value int64) snek.FlagInitializer {
			return snek.WithInt64VarE(variable, "test", envVar, value, "test int64")
		},
		map[string]flagEnvTest[int64]{
			"no env var, no flag": {
				args:     []string{},
				expected: 0,
			},
			"env var set, no flag": {
				args:     []string{},
				envValue: "42",
				setEnv:   true,
				expected: 42,
			},
			"cli flag overrides env var": {
				args:     []string{"--test=99"},
				envValue: "42",
				setEnv:   true,
				expected: 99,
			},
			"cli flag set, no env var": {
				args:     []string{"--test", "7"},
				expected: 7,
			},
			"invalid env var": {
				envValue:    "notanint",
				setEnv:      true,
				wantInitErr: true,
			},
		})
}

func TestWithInt64VarPE(t *testing.T) {
	runFlagEnvTest(t,
		"TEST_INT64",
		int64(0),
		func(variable *int64, envVar string, value int64) snek.FlagInitializer {
			return snek.WithInt64VarPE(variable, "test", "t", envVar, value, "test int64")
		},
		map[string]flagEnvTest[int64]{
			"no env var, no flag": {
				args:     []string{},
				expected: 0,
			},
			"env var set, no flag": {
				args:     []string{},
				envValue: "42",
				setEnv:   true,
				expected: 42,
			},
			"long flag overrides env var": {
				args:     []string{"--test=99"},
				envValue: "42",
				setEnv:   true,
				expected: 99,
			},
			"short flag overrides env var": {
				args:     []string{"-t", "7"},
				envValue: "42",
				setEnv:   true,
				expected: 7,
			},
			"invalid env var": {
				envValue:    "notanint",
				setEnv:      true,
				wantInitErr: true,
			},
		})
}

// ---- uint ----

func TestWithUintVarE(t *testing.T) {
	runFlagEnvTest(t,
		"TEST_UINT",
		uint(0),
		func(variable *uint, envVar string, value uint) snek.FlagInitializer {
			return snek.WithUintVarE(variable, "test", envVar, value, "test uint")
		},
		map[string]flagEnvTest[uint]{
			"no env var, no flag": {
				args:     []string{},
				expected: 0,
			},
			"env var set, no flag": {
				args:     []string{},
				envValue: "42",
				setEnv:   true,
				expected: 42,
			},
			"cli flag overrides env var": {
				args:     []string{"--test=99"},
				envValue: "42",
				setEnv:   true,
				expected: 99,
			},
			"cli flag set, no env var": {
				args:     []string{"--test", "7"},
				expected: 7,
			},
			"invalid env var": {
				envValue:    "notauint",
				setEnv:      true,
				wantInitErr: true,
			},
		})
}

func TestWithUintVarPE(t *testing.T) {
	runFlagEnvTest(t,
		"TEST_UINT",
		uint(0),
		func(variable *uint, envVar string, value uint) snek.FlagInitializer {
			return snek.WithUintVarPE(variable, "test", "t", envVar, value, "test uint")
		},
		map[string]flagEnvTest[uint]{
			"no env var, no flag": {
				args:     []string{},
				expected: 0,
			},
			"env var set, no flag": {
				args:     []string{},
				envValue: "42",
				setEnv:   true,
				expected: 42,
			},
			"long flag overrides env var": {
				args:     []string{"--test=99"},
				envValue: "42",
				setEnv:   true,
				expected: 99,
			},
			"short flag overrides env var": {
				args:     []string{"-t", "7"},
				envValue: "42",
				setEnv:   true,
				expected: 7,
			},
			"invalid env var": {
				envValue:    "notauint",
				setEnv:      true,
				wantInitErr: true,
			},
		})
}

// ---- uint8 ----

func TestWithUint8VarE(t *testing.T) {
	runFlagEnvTest(t,
		"TEST_UINT8",
		uint8(0),
		func(variable *uint8, envVar string, value uint8) snek.FlagInitializer {
			return snek.WithUint8VarE(variable, "test", envVar, value, "test uint8")
		},
		map[string]flagEnvTest[uint8]{
			"no env var, no flag": {
				args:     []string{},
				expected: 0,
			},
			"env var set, no flag": {
				args:     []string{},
				envValue: "42",
				setEnv:   true,
				expected: 42,
			},
			"cli flag overrides env var": {
				args:     []string{"--test=99"},
				envValue: "42",
				setEnv:   true,
				expected: 99,
			},
			"cli flag set, no env var": {
				args:     []string{"--test", "7"},
				expected: 7,
			},
			"invalid env var": {
				envValue:    "notauint",
				setEnv:      true,
				wantInitErr: true,
			},
		})
}

func TestWithUint8VarPE(t *testing.T) {
	runFlagEnvTest(t,
		"TEST_UINT8",
		uint8(0),
		func(variable *uint8, envVar string, value uint8) snek.FlagInitializer {
			return snek.WithUint8VarPE(variable, "test", "t", envVar, value, "test uint8")
		},
		map[string]flagEnvTest[uint8]{
			"no env var, no flag": {
				args:     []string{},
				expected: 0,
			},
			"env var set, no flag": {
				args:     []string{},
				envValue: "42",
				setEnv:   true,
				expected: 42,
			},
			"long flag overrides env var": {
				args:     []string{"--test=99"},
				envValue: "42",
				setEnv:   true,
				expected: 99,
			},
			"short flag overrides env var": {
				args:     []string{"-t", "7"},
				envValue: "42",
				setEnv:   true,
				expected: 7,
			},
			"invalid env var": {
				envValue:    "notauint",
				setEnv:      true,
				wantInitErr: true,
			},
		})
}

// ---- uint16 ----

func TestWithUint16VarE(t *testing.T) {
	runFlagEnvTest(t,
		"TEST_UINT16",
		uint16(0),
		func(variable *uint16, envVar string, value uint16) snek.FlagInitializer {
			return snek.WithUint16VarE(variable, "test", envVar, value, "test uint16")
		},
		map[string]flagEnvTest[uint16]{
			"no env var, no flag": {
				args:     []string{},
				expected: 0,
			},
			"env var set, no flag": {
				args:     []string{},
				envValue: "42",
				setEnv:   true,
				expected: 42,
			},
			"cli flag overrides env var": {
				args:     []string{"--test=99"},
				envValue: "42",
				setEnv:   true,
				expected: 99,
			},
			"cli flag set, no env var": {
				args:     []string{"--test", "7"},
				expected: 7,
			},
			"invalid env var": {
				envValue:    "notauint",
				setEnv:      true,
				wantInitErr: true,
			},
		})
}

func TestWithUint16VarPE(t *testing.T) {
	runFlagEnvTest(t,
		"TEST_UINT16",
		uint16(0),
		func(variable *uint16, envVar string, value uint16) snek.FlagInitializer {
			return snek.WithUint16VarPE(variable, "test", "t", envVar, value, "test uint16")
		},
		map[string]flagEnvTest[uint16]{
			"no env var, no flag": {
				args:     []string{},
				expected: 0,
			},
			"env var set, no flag": {
				args:     []string{},
				envValue: "42",
				setEnv:   true,
				expected: 42,
			},
			"long flag overrides env var": {
				args:     []string{"--test=99"},
				envValue: "42",
				setEnv:   true,
				expected: 99,
			},
			"short flag overrides env var": {
				args:     []string{"-t", "7"},
				envValue: "42",
				setEnv:   true,
				expected: 7,
			},
			"invalid env var": {
				envValue:    "notauint",
				setEnv:      true,
				wantInitErr: true,
			},
		})
}

// ---- uint32 ----

func TestWithUint32VarE(t *testing.T) {
	runFlagEnvTest(t,
		"TEST_UINT32",
		uint32(0),
		func(variable *uint32, envVar string, value uint32) snek.FlagInitializer {
			return snek.WithUint32VarE(variable, "test", envVar, value, "test uint32")
		},
		map[string]flagEnvTest[uint32]{
			"no env var, no flag": {
				args:     []string{},
				expected: 0,
			},
			"env var set, no flag": {
				args:     []string{},
				envValue: "42",
				setEnv:   true,
				expected: 42,
			},
			"cli flag overrides env var": {
				args:     []string{"--test=99"},
				envValue: "42",
				setEnv:   true,
				expected: 99,
			},
			"cli flag set, no env var": {
				args:     []string{"--test", "7"},
				expected: 7,
			},
			"invalid env var": {
				envValue:    "notauint",
				setEnv:      true,
				wantInitErr: true,
			},
		})
}

func TestWithUint32VarPE(t *testing.T) {
	runFlagEnvTest(t,
		"TEST_UINT32",
		uint32(0),
		func(variable *uint32, envVar string, value uint32) snek.FlagInitializer {
			return snek.WithUint32VarPE(variable, "test", "t", envVar, value, "test uint32")
		},
		map[string]flagEnvTest[uint32]{
			"no env var, no flag": {
				args:     []string{},
				expected: 0,
			},
			"env var set, no flag": {
				args:     []string{},
				envValue: "42",
				setEnv:   true,
				expected: 42,
			},
			"long flag overrides env var": {
				args:     []string{"--test=99"},
				envValue: "42",
				setEnv:   true,
				expected: 99,
			},
			"short flag overrides env var": {
				args:     []string{"-t", "7"},
				envValue: "42",
				setEnv:   true,
				expected: 7,
			},
			"invalid env var": {
				envValue:    "notauint",
				setEnv:      true,
				wantInitErr: true,
			},
		})
}

// ---- uint64 ----

func TestWithUint64VarE(t *testing.T) {
	runFlagEnvTest(t,
		"TEST_UINT64",
		uint64(0),
		func(variable *uint64, envVar string, value uint64) snek.FlagInitializer {
			return snek.WithUint64VarE(variable, "test", envVar, value, "test uint64")
		},
		map[string]flagEnvTest[uint64]{
			"no env var, no flag": {
				args:     []string{},
				expected: 0,
			},
			"env var set, no flag": {
				args:     []string{},
				envValue: "42",
				setEnv:   true,
				expected: 42,
			},
			"cli flag overrides env var": {
				args:     []string{"--test=99"},
				envValue: "42",
				setEnv:   true,
				expected: 99,
			},
			"cli flag set, no env var": {
				args:     []string{"--test", "7"},
				expected: 7,
			},
			"invalid env var": {
				envValue:    "notauint",
				setEnv:      true,
				wantInitErr: true,
			},
		})
}

func TestWithUint64VarPE(t *testing.T) {
	runFlagEnvTest(t,
		"TEST_UINT64",
		uint64(0),
		func(variable *uint64, envVar string, value uint64) snek.FlagInitializer {
			return snek.WithUint64VarPE(variable, "test", "t", envVar, value, "test uint64")
		},
		map[string]flagEnvTest[uint64]{
			"no env var, no flag": {
				args:     []string{},
				expected: 0,
			},
			"env var set, no flag": {
				args:     []string{},
				envValue: "42",
				setEnv:   true,
				expected: 42,
			},
			"long flag overrides env var": {
				args:     []string{"--test=99"},
				envValue: "42",
				setEnv:   true,
				expected: 99,
			},
			"short flag overrides env var": {
				args:     []string{"-t", "7"},
				envValue: "42",
				setEnv:   true,
				expected: 7,
			},
			"invalid env var": {
				envValue:    "notauint",
				setEnv:      true,
				wantInitErr: true,
			},
		})
}

// ---- string ----

func TestWithStringVarE(t *testing.T) {
	runFlagEnvTest(t,
		"TEST_STRING",
		"default",
		func(variable *string, envVar string, value string) snek.FlagInitializer {
			return snek.WithStringVarE(variable, "test", envVar, value, "test string")
		},
		map[string]flagEnvTest[string]{
			"no env var, no flag": {
				args:     []string{},
				expected: "default",
			},
			"env var set, no flag": {
				args:     []string{},
				envValue: "from-env",
				setEnv:   true,
				expected: "from-env",
			},
			"cli flag overrides env var": {
				args:     []string{"--test=from-flag"},
				envValue: "from-env",
				setEnv:   true,
				expected: "from-flag",
			},
			"cli flag set, no env var": {
				args:     []string{"--test", "from-flag"},
				expected: "from-flag",
			},
		})
}

func TestWithStringVarPE(t *testing.T) {
	runFlagEnvTest(t,
		"TEST_STRING",
		"default",
		func(variable *string, envVar string, value string) snek.FlagInitializer {
			return snek.WithStringVarPE(variable, "test", "t", envVar, value, "test string")
		},
		map[string]flagEnvTest[string]{
			"no env var, no flag": {
				args:     []string{},
				expected: "default",
			},
			"env var set, no flag": {
				args:     []string{},
				envValue: "from-env",
				setEnv:   true,
				expected: "from-env",
			},
			"long flag overrides env var": {
				args:     []string{"--test=from-flag"},
				envValue: "from-env",
				setEnv:   true,
				expected: "from-flag",
			},
			"short flag overrides env var": {
				args:     []string{"-t", "from-flag"},
				envValue: "from-env",
				setEnv:   true,
				expected: "from-flag",
			},
			"cli flag set, no env var": {
				args:     []string{"--test", "from-flag"},
				expected: "from-flag",
			},
		})
}
