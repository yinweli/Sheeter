package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLua(t *testing.T) {
	value1, err := LuaBool("true")
	assert.Nil(t, err)
	assert.Equal(t, "true", value1)
	value1, err = LuaBool("false")
	assert.Nil(t, err)
	assert.Equal(t, "false", value1)
	value1, err = LuaBool("TRUE")
	assert.Nil(t, err)
	assert.Equal(t, "true", value1)
	value1, err = LuaBool("FALSE")
	assert.Nil(t, err)
	assert.Equal(t, "false", value1)
	value1, err = LuaBool("1")
	assert.Nil(t, err)
	assert.Equal(t, "true", value1)
	value1, err = LuaBool("0")
	assert.Nil(t, err)
	assert.Equal(t, "false", value1)
	_, err = LuaBool("?????")
	assert.NotNil(t, err)

	value2, err := LuaBoolArray("true,false,true,false,true")
	assert.Nil(t, err)
	assert.Equal(t, "true,false,true,false,true", value2)
	value2, err = LuaBoolArray("TRUE,FALSE,TRUE,FALSE,TRUE")
	assert.Nil(t, err)
	assert.Equal(t, "true,false,true,false,true", value2)
	value2, err = LuaBoolArray("1,0,1,0,1")
	assert.Nil(t, err)
	assert.Equal(t, "true,false,true,false,true", value2)
	_, err = LuaBoolArray("???,???,???,???,???")
	assert.NotNil(t, err)

	value10 := LuaWrapperArray("testString")
	assert.Equal(t, "{testString}", value10)

	value11 := LuaWrapperString("testString")
	assert.Equal(t, "\"testString\"", value11)
}
