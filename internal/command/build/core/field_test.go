package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseField(t *testing.T) {
	name, field, err := ParseField("real#bool")
	assert.Nil(t, err)
	assert.Equal(t, "real", name)
	assert.Equal(t, (&FieldBool{}).Type(), field.Type())
	_, _, err = ParseField("fa-ke#fake")
	assert.NotNil(t, err)
	_, _, err = ParseField("fake#fake")
	assert.NotNil(t, err)
	_, _, err = ParseField("unknown")
	assert.NotNil(t, err)
}
