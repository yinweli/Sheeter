package build

import (
	"bytes"
	"testing"

	"Sheeter/internal/command/build/config"
	"Sheeter/testdata"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestNewCommand(t *testing.T) {
	command := NewCommand()

	assert.NotNil(t, command, "command nil")
}

func TestExecute(t *testing.T) {
	buffer, command := fakeCommand()

	execute(command, []string{testdata.Path(testdata.RealYaml())})
	assert.Equal(t, "", buffer.String(), "read real config failed")

	execute(command, []string{testdata.Path(testdata.FakeYaml())})
	assert.NotEqual(t, "", buffer.String(), "read fake config failed")
}

func TestReadBuildConfig(t *testing.T) {
	var buildConfig config.Config

	assert.Nil(t, readBuildConfig(testdata.RealYaml(), &buildConfig), "read real config failed")
	assert.NotNil(t, readBuildConfig(testdata.FakeYaml(), &buildConfig), "read fake config failed")
	assert.NotNil(t, readBuildConfig(testdata.UnknownYaml(), &buildConfig), "read unknown config failed")
}

// fakeCommand 取得假的命令物件
func fakeCommand() (buffer *bytes.Buffer, command *cobra.Command) {
	buffer = &bytes.Buffer{}
	command = &cobra.Command{}
	command.SetOut(buffer)

	return buffer, command
}
