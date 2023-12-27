package snek_test

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ronelliott/snek"
)

func TestNewCommand_WithInitializers(t *testing.T) {
	called := false
	cmd, err := snek.NewCommand(func(cmd *cobra.Command) error {
		cmd.Use = "foo"
		called = true
		return nil
	})
	assert.NoError(t, err, "NewCommand should not return an error")
	assert.NotNil(t, cmd, "NewCommand should return a command")
	assert.Equal(t, "foo", cmd.Use, "Name should be set")
	assert.True(t, called, "Initializer should be called")
}

func TestNewCommand_WithInitializers_Error(t *testing.T) {
	cmd, err := snek.NewCommand(func(cmd *cobra.Command) error {
		return assert.AnError
	})
	assert.ErrorIs(t, err, assert.AnError, "NewCommand should return the error produced by the initializer")
	assert.Nil(t, cmd, "NewCommand should not return a command")
}

func TestWithAliases(t *testing.T) {
	cmd, err := snek.NewCommand(snek.WithAliases("foo", "bar"))
	require.NoError(t, err, "NewCommand should not return an error")
	assert.Equal(t, []string{"foo", "bar"}, cmd.Aliases, "Aliases should be set")
}

func TestWithDeprecated(t *testing.T) {
	cmd, err := snek.NewCommand(snek.WithDeprecated("foo"))
	require.NoError(t, err, "NewCommand should not return an error")
	assert.Equal(t, "foo", cmd.Deprecated, "Deprecated should be set")
}

func TestWithExample(t *testing.T) {
	cmd, err := snek.NewCommand(snek.WithExample("foo"))
	require.NoError(t, err, "NewCommand should not return an error")
	assert.Equal(t, "foo", cmd.Example, "Example should be set")
}

func TestWithLong(t *testing.T) {
	cmd, err := snek.NewCommand(snek.WithLong("foo"))
	require.NoError(t, err, "NewCommand should not return an error")
	assert.Equal(t, "foo", cmd.Long, "Long description should be set")
}

func TestWithRun(t *testing.T) {
	cmd, err := snek.NewCommand(snek.WithRun(func(*cobra.Command, []string) {}))
	require.NoError(t, err, "NewCommand should not return an error")
	assert.NotNil(t, cmd.Run, "Run should be set")
}

func TestWithRunE(t *testing.T) {
	cmd, err := snek.NewCommand(snek.WithRunE(func(*cobra.Command, []string) error { return nil }))
	require.NoError(t, err, "NewCommand should not return an error")
	assert.NotNil(t, cmd.RunE, "RunE should be set")
}

func TestWithShort(t *testing.T) {
	cmd, err := snek.NewCommand(snek.WithShort("foo"))
	require.NoError(t, err, "NewCommand should not return an error")
	assert.Equal(t, "foo", cmd.Short, "Short description should be set")
}

func TestWithSimpleRun(t *testing.T) {
	called := false
	cmd, err := snek.NewCommand(snek.WithSimpleRun(func(args []string) {
		called = true
		assert.Equal(t, []string{"foo", "bar"}, args, "Args should be passed values")
	}))
	require.NoError(t, err, "NewCommand should not return an error")
	assert.NotNil(t, cmd.Run, "Run should be set")

	cmd.Run(cmd, []string{"foo", "bar"})
	assert.True(t, called, "The function should be called")
}

func TestWithSimpleRunE(t *testing.T) {
	called := false
	cmd, err := snek.NewCommand(snek.WithSimpleRunE(func(args []string) error {
		called = true
		assert.Equal(t, []string{"foo", "bar"}, args, "Args should be passed values")
		return nil
	}))
	require.NoError(t, err, "NewCommand should not return an error")
	assert.NotNil(t, cmd.RunE, "RunE should be set")

	err = cmd.RunE(cmd, []string{"foo", "bar"})
	assert.NoError(t, err, "The function should not return an error")
	assert.True(t, called, "The function should be called")
}

func TestWithSubCommand(t *testing.T) {
	cmd, err := snek.NewCommand(snek.WithSubCommand(
		&cobra.Command{
			Use: "foo",
		},
		&cobra.Command{
			Use: "bar",
		},
	))
	require.NoError(t, err, "NewCommand should not return an error")

	cmds := cmd.Commands()
	require.Len(t, cmds, 2, "There should be two subcommands")
}

func TestWithSubCommand_Error(t *testing.T) {
	cmd, err := snek.NewCommand(snek.WithSubCommandGenerator(
		func() (*cobra.Command, error) {
			return nil, assert.AnError
		},
	))
	assert.ErrorIs(t, err, assert.AnError, "NewCommand should return the error produced by the initializer")
	assert.Nil(t, cmd, "NewCommand should not return a command")
}

func TestWithSubCommandGenerator_NoError(t *testing.T) {
	cmd, err := snek.NewCommand(snek.WithSubCommandGenerator(
		func() (*cobra.Command, error) {
			return &cobra.Command{
				Use: "foo",
			}, nil
		},
		func() (*cobra.Command, error) {
			return &cobra.Command{
				Use: "bar",
			}, nil
		},
	))
	require.NoError(t, err, "NewCommand should not return an error")

	cmds := cmd.Commands()
	require.Len(t, cmds, 2, "There should be two subcommands")
}

func TestWithUse(t *testing.T) {
	cmd, err := snek.NewCommand(snek.WithUse("foo"))
	require.NoError(t, err, "NewCommand should not return an error")
	assert.Equal(t, "foo", cmd.Use, "Name should be set")
}

func TestWithValidArgs(t *testing.T) {
	cmd, err := snek.NewCommand(snek.WithValidArgs("foo"))
	require.NoError(t, err, "NewCommand should not return an error")
	assert.Equal(t, []string{"foo"}, cmd.ValidArgs, "ValidArgs should be set")
}

func TestWithVersion(t *testing.T) {
	cmd, err := snek.NewCommand(snek.WithVersion("foo"))
	require.NoError(t, err, "NewCommand should not return an error")
	assert.Equal(t, "foo", cmd.Version, "Version should be set")
}
