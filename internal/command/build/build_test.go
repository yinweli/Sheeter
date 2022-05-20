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

	//
}

func TestReadBuildConfig(t *testing.T) {
	filename := testdata.Path("config.yaml")
	buildConfig := config.Config{}

	err := readBuildConfig(filename, &buildConfig)
	assert.Nil(t, err, err)
}

// fakeCommand 取得假的命令物件
func fakeCommand() (buffer *bytes.Buffer, command *cobra.Command) {
	buffer = &bytes.Buffer{}
	command = &cobra.Command{}
	command.SetOut(buffer)

	return buffer, command
}
