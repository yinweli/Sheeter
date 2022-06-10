package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFieldPkey(t *testing.T) {
	field := mockFieldPkey()
	assert.Equal(t, "pkey", field.TypeExcel())
	assert.Equal(t, CppNamespace+"::pkey", field.TypeCpp())
	assert.Equal(t, "int", field.TypeCs())
	assert.Equal(t, "int32", field.TypeGo())
	assert.Equal(t, true, field.IsShow())
	assert.Equal(t, true, field.IsPkey())
	result, err := field.Transform("123456789")
	assert.Nil(t, err)
	assert.Equal(t, int64(123456789), result)
	result, err = field.Transform("?????")
	assert.NotNil(t, err)
}

func mockFieldPkey() *FieldPkey {
	return &FieldPkey{}
}
