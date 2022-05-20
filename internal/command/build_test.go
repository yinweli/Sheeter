package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCommandBuild(t *testing.T) {
	build := NewCommandBuild()

	assert.NotNil(t, build, "build nil")
}
