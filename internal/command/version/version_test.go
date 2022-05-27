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
	assert.NotNil(t, command)
}

func TestExecute(t *testing.T) {
	buffer := &bytes.Buffer{}
	command := &cobra.Command{}
	command.SetOut(buffer)

	err := execute(command, []string{})
	assert.Nil(t, err)
	assert.Equal(t, fmt.Sprintf("%s %s", internal.Title, internal.Version), buffer.String())
}
