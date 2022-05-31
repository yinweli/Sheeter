package core

import (
	"fmt"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestExecutor_Run(t *testing.T) {
	err1 := fmt.Errorf("err1")
	err2 := fmt.Errorf("err2")
	executor := NewExecutor(&cobra.Command{}, []ExecData{
		{LongName: "cmd1", ExecFunc: func(cargo *Cargo) (filePath string, err error) { return "", err1 }},
		{LongName: "cmd2", ExecFunc: func(cargo *Cargo) (filePath string, err error) { return "", err2 }},
		{LongName: "cmd3", ExecFunc: func(cargo *Cargo) (filePath string, err error) { return "", nil }},
	})

	err := executor.Cmd.Flags().Set("cmd1", "true")
	assert.Nil(t, err)
	assert.Equal(t, err1, executor.Run(nil))

	err = executor.Cmd.Flags().Set("cmd1", "false")
	assert.Nil(t, err)
	err = executor.Cmd.Flags().Set("cmd2", "true")
	assert.Nil(t, err)
	assert.Equal(t, err2, executor.Run(nil))

	err = executor.Cmd.Flags().Set("cmd2", "false")
	assert.Nil(t, err)
	err = executor.Cmd.Flags().Set("cmd3", "true")
	assert.Nil(t, err)
	assert.Nil(t, executor.Run(nil))
}

func TestExecutor_Count(t *testing.T) {
	executor := NewExecutor(&cobra.Command{}, []ExecData{
		{LongName: "cmd1"},
		{LongName: "cmd2"},
		{LongName: "cmd3"},
	})

	assert.Equal(t, 0, executor.Count())

	err := executor.Cmd.Flags().Set("cmd1", "true")
	assert.Nil(t, err)
	assert.Equal(t, 1, executor.Count())

	err = executor.Cmd.Flags().Set("cmd2", "true")
	assert.Nil(t, err)
	assert.Equal(t, 2, executor.Count())

	err = executor.Cmd.Flags().Set("cmd3", "true")
	assert.Nil(t, err)
	assert.Equal(t, 3, executor.Count())
}

func TestExecutor_State(t *testing.T) {
	executor := NewExecutor(&cobra.Command{}, []ExecData{
		{LongName: "cmd"},
	})

	err := executor.Cmd.Flags().Set("cmd", "true")
	assert.Nil(t, err)
	assert.True(t, executor.state("cmd"))
	assert.False(t, executor.state("????"))
}

func TestNewExecutor(t *testing.T) {
	executor := NewExecutor(&cobra.Command{}, []ExecData{})
	assert.NotNil(t, executor)
}
