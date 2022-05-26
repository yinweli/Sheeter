package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseField(t *testing.T) {
	var name string
	var field Field
	var err error

	name, field, err = ParseField("real#bool")
	assert.Equal(t, "real", name, "parse real failed")
	assert.Equal(t, (&FieldBool{}).TypeExcel(), field.TypeExcel(), "parse real failed")
	assert.Nil(t, err, "parse real failed")

	name, field, err = ParseField("fake#fake")
	assert.NotNil(t, err, "parse fake failed")

	name, field, err = ParseField("unknown")
	assert.NotNil(t, err, "parse unknown failed")
}
