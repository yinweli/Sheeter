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
	buffer := &bytes.Buffer{}
	command := &cobra.Command{}
	command.SetOut(buffer)

	dir := testdata.ChangeWorkDir()
	defer testdata.RestoreWorkDir(dir)

	execute(command, []string{testdata.Path(testdata.RealConfig)})
	assert.Equal(t, "", buffer.String())

	execute(command, []string{testdata.Path(testdata.FakeConfig)})
	assert.NotEqual(t, "", buffer.String())
}
