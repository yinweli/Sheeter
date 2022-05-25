package build

import (
	"testing"

	"Sheeter/testdata"

	"github.com/stretchr/testify/assert"
)

func TestExecute(t *testing.T) {
	buffer, command := testdata.FakeCommand()

	execute(command, []string{""})
	assert.NotNil(t, buffer.String(), "execute failed")
}

func TestBuild_ReadConfig(t *testing.T) {
	_, command := testdata.FakeCommand()
	build := build{cmd: command}

	build.args = []string{testdata.Path(testdata.RealConfig)}
	build.err = nil
	build.readConfig()
	assert.Nil(t, build.err, "read real config failed")
	assert.NotNil(t, build.config, "read real config failed")

	build.args = []string{testdata.Path(testdata.FakeConfig)}
	build.err = nil
	build.readConfig()
	assert.NotNil(t, build.err, "read fake config failed")

	build.args = []string{testdata.Path(testdata.DefectConfig)}
	build.err = nil
	build.readConfig()
	assert.NotNil(t, build.err, "read defect config failed")

	build.args = []string{testdata.Path(testdata.UnknownConfig)}
	build.err = nil
	build.readConfig()
	assert.NotNil(t, build.err, "read unknown config failed")
}
