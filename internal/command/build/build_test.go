package build

import (
	"testing"

	"Sheeter/testdata"

	"github.com/stretchr/testify/assert"
)

func TestExecute(t *testing.T) {
	buffer, command := testdata.FakeCommand()
	dir := testdata.ChangeWorkDir()
	defer testdata.RestoreWorkDir(dir)

	execute(command, []string{testdata.Path(testdata.RealConfig)})
	assert.Equal(t, "", buffer.String(), "execute failed")

	execute(command, []string{testdata.Path(testdata.FakeConfig)})
	assert.NotEqual(t, "", buffer.String(), "execute fake config failed")
}
