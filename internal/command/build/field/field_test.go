package field

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	var name string
	var field Field
	var err error

	name, field, err = Parse("real#bool")
	assert.Equal(t, "real", name, "parse real failed")
	assert.Equal(t, (&Bool{}).TypeExcel(), field.TypeExcel(), "parse real failed")
	assert.Nil(t, err, "parse real failed")

	name, field, err = Parse("fake#fake")
	assert.NotNil(t, err, "parse fake failed")

	name, field, err = Parse("unknown")
	assert.NotNil(t, err, "parse unknown failed")
}

func TestFields(t *testing.T) {
	field := &Bool{}

	result, ok := fields[field.TypeExcel()]
	assert.Equal(t, field, result, "fields failed")
	assert.Equal(t, true, ok, "fields failed")
}
