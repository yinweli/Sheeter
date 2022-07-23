package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestLua(t *testing.T) {
	suite.Run(t, new(SuiteLua))
}

type SuiteLua struct {
	suite.Suite
}

func (this *SuiteLua) TestLuaBool() {
	value, err := LuaBool("true")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), "true", value)

	value, err = LuaBool("false")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), "false", value)

	value, err = LuaBool("TRUE")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), "true", value)

	value, err = LuaBool("FALSE")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), "false", value)

	value, err = LuaBool("1")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), "true", value)

	value, err = LuaBool("0")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), "false", value)

	_, err = LuaBool("?????")
	assert.NotNil(this.T(), err)
}

func (this *SuiteLua) TestLuaBoolArray() {
	value, err := LuaBoolArray("true,false,true,false,true")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), "true,false,true,false,true", value)

	value, err = LuaBoolArray("TRUE,FALSE,TRUE,FALSE,TRUE")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), "true,false,true,false,true", value)

	value, err = LuaBoolArray("1,0,1,0,1")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), "true,false,true,false,true", value)

	_, err = LuaBoolArray("???,???,???,???,???")
	assert.NotNil(this.T(), err)
}

func (this *SuiteLua) TestLuaWrapperArray() {
	assert.Equal(this.T(), "{testString}", LuaWrapperArray("testString"))
}

func (this *SuiteLua) TestLuaWrapperString() {
	assert.Equal(this.T(), "\"testString\"", LuaWrapperString("testString"))
}
