package field

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFields(t *testing.T) {
	assert.NotNil(t, NewFields(), "new fields failed")
}

func TestAddFields(t *testing.T) {
	fields := make(Fields)
	field := &Bool{}

	addFields(fields, field)
	result, ok := fields[field.TypeExcel()]
	assert.Equal(t, field, result, "add field failed")
	assert.Equal(t, true, ok, "add field failed")

}
