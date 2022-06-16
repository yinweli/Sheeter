package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	assert.Equal(t, "", FirstUpper(""))
	assert.Equal(t, "TestString", FirstUpper("testString"))
	assert.Equal(t, "", FirstLower(""))
	assert.Equal(t, "testString", FirstLower("TestString"))

	value1, err := StrToBool("1")
	assert.Nil(t, err)
	assert.Equal(t, true, value1)
	value1, err = StrToBool("true")
	assert.Nil(t, err)
	assert.Equal(t, true, value1)
	value1, err = StrToBool("TRUE")
	assert.Nil(t, err)
	assert.Equal(t, true, value1)
	value1, err = StrToBool("0")
	assert.Nil(t, err)
	assert.Equal(t, false, value1)
	value1, err = StrToBool("false")
	assert.Nil(t, err)
	assert.Equal(t, false, value1)
	value1, err = StrToBool("FALSE")
	assert.Nil(t, err)
	assert.Equal(t, false, value1)
	value1, err = StrToBool("?????")
	assert.NotNil(t, err)

	value2, err := StrToBoolArray("1,0,1,0,1")
	assert.Nil(t, err)
	assert.Equal(t, []bool{true, false, true, false, true}, value2)
	value2, err = StrToBoolArray("true,false,true,false,true")
	assert.Nil(t, err)
	assert.Equal(t, []bool{true, false, true, false, true}, value2)
	value2, err = StrToBoolArray("TRUE,FALSE,TRUE,FALSE,TRUE")
	assert.Nil(t, err)
	assert.Equal(t, []bool{true, false, true, false, true}, value2)
	value2, err = StrToBoolArray("???,???,???,???,???")
	assert.NotNil(t, err)

	value3, err := StrToInt("123456789")
	assert.Nil(t, err)
	assert.Equal(t, int64(123456789), value3)
	value3, err = StrToInt("?????")
	assert.NotNil(t, err)

	value4, err := StrToIntArray("123,456,789")
	assert.Nil(t, err)
	assert.Equal(t, []int64{123, 456, 789}, value4)
	value4, err = StrToIntArray("?????")
	assert.NotNil(t, err)

	value5, err := StrToFloat("0.123456789")
	assert.Nil(t, err)
	assert.Equal(t, 0.123456789, value5)
	value5, err = StrToFloat("?????")
	assert.NotNil(t, err)

	value6, err := StrToFloatArray("0.123,0.456,0.789")
	assert.Nil(t, err)
	assert.Equal(t, []float64{0.123, 0.456, 0.789}, value6)
	value6, err = StrToFloatArray("?????")
	assert.NotNil(t, err)

	value7 := StrToStrArray("ball,book,pack")
	assert.Equal(t, []string{"ball", "book", "pack"}, value7)
	value7 = StrToStrArray("?????#?????#?????")
	assert.Equal(t, []string{"?????#?????#?????"}, value7)

	value8 := LuaArrayWrapper("testString")
	assert.Equal(t, "{testString}", value8)
}
