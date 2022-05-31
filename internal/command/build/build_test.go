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

	cmd := &cobra.Command{}
	cmd.SetOut(&bytes.Buffer{})

	err := execute(cmd, []string{testdata.Path(testdata.RealConfig)})
	assert.Nil(t, err)
	err = execute(cmd, []string{testdata.Path(testdata.FakeConfig)})
	assert.NotNil(t, err)
}
