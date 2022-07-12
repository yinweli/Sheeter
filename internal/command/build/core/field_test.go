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

	name, field, err = ParseField("fa-ke#fake")
	assert.NotNil(t, err)

	name, field, err = ParseField("fake#fake")
	assert.NotNil(t, err)

	name, field, err = ParseField("unknown")
	assert.NotNil(t, err)
}
