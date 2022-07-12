package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrings(t *testing.T) {
	assert.Equal(t, "", FirstUpper(""))
	assert.Equal(t, "TestString", FirstUpper("testString"))
	assert.Equal(t, "", FirstLower(""))
	assert.Equal(t, "testString", FirstLower("TestString"))

	assert.True(t, VariableCheck("value"))
	assert.True(t, VariableCheck("Value"))
	assert.True(t, VariableCheck("value1"))
	assert.True(t, VariableCheck("Value1"))
	assert.True(t, VariableCheck("value_a"))
	assert.True(t, VariableCheck("value_1"))
	assert.True(t, VariableCheck("_value"))
	assert.False(t, VariableCheck(""))
	assert.False(t, VariableCheck("0value"))
	assert.False(t, VariableCheck("-value"))
	assert.False(t, VariableCheck("value-"))
	assert.False(t, VariableCheck("#value"))
	assert.False(t, VariableCheck("@value"))
	assert.False(t, VariableCheck("value{}"))
}
