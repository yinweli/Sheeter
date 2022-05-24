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

	assert.NotNil(t, command, "command nil")
}

func TestExecute(t *testing.T) {
	buffer, command := fakeCommand()

	execute(command, []string{testdata.RealYaml()})
	assert.Equal(t, "", buffer.String(), "read real config failed")

	execute(command, []string{testdata.FakeYaml()})
	assert.NotEqual(t, "", buffer.String(), "read fake config failed")

	execute(command, []string{testdata.EmptyYaml()})
	assert.NotEqual(t, "", buffer.String(), "read fake config failed")
}

// fakeCommand 取得假的命令物件
func fakeCommand() (buffer *bytes.Buffer, command *cobra.Command) {
	buffer = &bytes.Buffer{}
	command = &cobra.Command{}
	command.SetOut(buffer)

	return buffer, command
}
