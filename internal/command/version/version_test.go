package version

import (
	"bytes"
	"fmt"
	"testing"

	"Sheeter/internal"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestNewCommand(t *testing.T) {
	command := NewCommand()

	assert.NotNil(t, command, "command nil")
}

func TestExecute(t *testing.T) {
	buffer := &bytes.Buffer{}
	command := &cobra.Command{}
	command.SetOut(buffer)
	expected := fmt.Sprintf("%s %s", internal.Title, internal.Version)

	execute(command, []string{})
	assert.Equal(t, expected, buffer.String(), "version failed")
}
