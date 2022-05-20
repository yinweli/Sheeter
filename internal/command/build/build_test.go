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

func TestReadBuildConfig(t *testing.T) {
	buildConfig := config.Config{}
	filename1 := testdata.Path("config.yaml")

	err := readBuildConfig(filename1, &buildConfig)
	assert.Nil(t, err, err)

	filename2 := testdata.Path("failed.yaml")

	err = readBuildConfig(filename2, &buildConfig)
	assert.NotNil(t, err, err)

	filename3 := testdata.Path("??????.yaml")

	err = readBuildConfig(filename3, &buildConfig)
	assert.NotNil(t, err, err)
}

// fakeCommand 取得假的命令物件
func fakeCommand() (buffer *bytes.Buffer, command *cobra.Command) {
	buffer = &bytes.Buffer{}
	command = &cobra.Command{}
	command.SetOut(buffer)

	return buffer, command
}
