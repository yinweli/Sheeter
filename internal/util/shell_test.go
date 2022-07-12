package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShellRun(t *testing.T) {
	assert.Nil(t, ShellRun("go", "version"))
	assert.NotNil(t, ShellRun("unknown"))
}

func TestShellExist(t *testing.T) {
	assert.True(t, ShellExist("go"))
	assert.False(t, ShellExist("unknown"))
}
