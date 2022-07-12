package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShellRun(t *testing.T) {
	err, detail := ShellRun("echo", "value")
	assert.Nil(t, err)
	assert.Equal(t, "", detail)
}

func TestShellExist(t *testing.T) {
	assert.Nil(t, ShellExist("go"))
	assert.NotNil(t, ShellExist("x"))
}
