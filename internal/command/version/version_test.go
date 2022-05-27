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
	assert.NotNil(t, NewCommand())
}

func TestExecute(t *testing.T) {
	buffer := &bytes.Buffer{}
	command := &cobra.Command{}
	command.SetOut(buffer)

	execute(command, []string{})
	assert.Equal(t, fmt.Sprintf("%s %s", internal.Title, internal.Version), buffer.String())
}
