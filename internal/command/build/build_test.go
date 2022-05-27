package build

import (
	"bytes"
	"testing"

	"Sheeter/testdata"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestNewCommand(t *testing.T) {
	assert.NotNil(t, NewCommand())
}

func TestExecute(t *testing.T) {
	dir := testdata.ChangeWorkDir()
	defer testdata.RestoreWorkDir(dir)

	buffer := &bytes.Buffer{}
	command := &cobra.Command{}
	command.SetOut(buffer)

	assert.Nil(t, execute(command, []string{testdata.Path(testdata.RealConfig)}))
	assert.NotNil(t, execute(command, []string{testdata.Path(testdata.FakeConfig)}))
}
