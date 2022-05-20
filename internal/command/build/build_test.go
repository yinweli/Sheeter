package build

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCommand(t *testing.T) {
	command := NewCommand()

	assert.NotNil(t, command, "command nil")
}
