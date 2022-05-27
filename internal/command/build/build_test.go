package build

import (
	"bytes"
	"testing"

	"Sheeter/testdata"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestNewCommand(t *testing.T) {
	command := NewCommand()
	assert.NotNil(t, command)
}

func TestExecute(t *testing.T) {
	dir := testdata.ChangeWorkDir()
	defer testdata.RestoreWorkDir(dir)

	buffer := &bytes.Buffer{}
	command := &cobra.Command{}
	command.SetOut(buffer)

	err := execute(command, []string{testdata.Path(testdata.RealConfig)})
	assert.Nil(t, err)
	err = execute(command, []string{testdata.Path(testdata.FakeConfig)})
	assert.NotNil(t, err)
}
