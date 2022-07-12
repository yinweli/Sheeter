package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDuplicate(t *testing.T) {
	duplicate := Duplicate{}
	assert.True(t, duplicate.Check("001"))
	assert.True(t, duplicate.Check("001/002"))
	assert.True(t, duplicate.Check("001/002/003"))
	assert.False(t, duplicate.Check("001"))
	assert.False(t, duplicate.Check("001/002"))
	assert.False(t, duplicate.Check("001/002/003"))
}
