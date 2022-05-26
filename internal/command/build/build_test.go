package build

import (
	"testing"

	"Sheeter/testdata"

	"github.com/stretchr/testify/assert"
)

func TestNewCommand(t *testing.T) {
	assert.NotNil(t, NewCommand())
}

func TestExecute(t *testing.T) {
	buffer, command := testdata.MockCommand()
	dir := testdata.ChangeWorkDir()
	defer testdata.RestoreWorkDir(dir)

	execute(command, []string{testdata.Path(testdata.RealConfig)})
	assert.Equal(t, "", buffer.String())

	execute(command, []string{testdata.Path(testdata.FakeConfig)})
	assert.NotEqual(t, "", buffer.String())
}
