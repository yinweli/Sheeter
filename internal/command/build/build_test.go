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

	execute(command, []string{testdata.Path("real.yaml")})
	assert.Equal(t, "", buffer.String(), "read real config failed")

	execute(command, []string{testdata.Path("fake.yaml")})
	assert.NotEqual(t, "", buffer.String(), "read fake config failed")
}

func TestReadBuildConfig(t *testing.T) {
	var err error
	var buildConfig config.Config

	err = readBuildConfig(testdata.Path("real.yaml"), &buildConfig)
	assert.Nil(t, err, "read real config failed")

	err = readBuildConfig(testdata.Path("fake.yaml"), &buildConfig)
	assert.NotNil(t, err, "read fake config failed")

	err = readBuildConfig(testdata.Path("??????.yaml"), &buildConfig)
	assert.NotNil(t, err, "read unknown config failed")
}

// fakeCommand 取得假的命令物件
func fakeCommand() (buffer *bytes.Buffer, command *cobra.Command) {
	buffer = &bytes.Buffer{}
	command = &cobra.Command{}
	command.SetOut(buffer)

	return buffer, command
}
