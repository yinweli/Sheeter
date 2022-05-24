package version

import (
	"fmt"
	"testing"

	"Sheeter/internal"
	"Sheeter/testdata"

	"github.com/stretchr/testify/assert"
)

func TestNewCommand(t *testing.T) {
	command := NewCommand()

	assert.NotNil(t, command, "command nil")
}

func TestExecute(t *testing.T) {
	buffer, command := testdata.MockCommand()

	execute(command, []string{})
	assert.Equal(t, fmt.Sprintf("%s %s", internal.Title, internal.Version), buffer.String(), "version failed")
}
