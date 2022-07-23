package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestConvert(t *testing.T) {
	value1, err := StrToBool("true")
	assert.Nil(t, err)
	assert.Equal(t, true, value1)
	value1, err = StrToBool("false")
	assert.Nil(t, err)
	assert.Equal(t, false, value1)
	value1, err = StrToBool("TRUE")
	assert.Nil(t, err)
	assert.Equal(t, true, value1)
	value1, err = StrToBool("FALSE")
	assert.Nil(t, err)
	assert.Equal(t, false, value1)
	value1, err = StrToBool("1")
	assert.Nil(t, err)
	assert.Equal(t, true, value1)
	value1, err = StrToBool("0")
	assert.Nil(t, err)
	assert.Equal(t, false, value1)
	_, err = StrToBool("?????")
	assert.NotNil(t, err)

	value2, err := StrToBoolArray("true,false,true,false,true")
	assert.Nil(t, err)
	assert.Equal(t, []bool{true, false, true, false, true}, value2)
	value2, err = StrToBoolArray("TRUE,FALSE,TRUE,FALSE,TRUE")
	assert.Nil(t, err)
	assert.Equal(t, []bool{true, false, true, false, true}, value2)
	value2, err = StrToBoolArray("1,0,1,0,1")
	assert.Nil(t, err)
	assert.Equal(t, []bool{true, false, true, false, true}, value2)
	_, err = StrToBoolArray("???,???,???,???,???")
	assert.NotNil(t, err)

	value3, err := StrToInt("123456789")
	assert.Nil(t, err)
	assert.Equal(t, int64(123456789), value3)
	_, err = StrToInt("?????")
	assert.NotNil(t, err)

	value4, err := StrToIntArray("123,456,789")
	assert.Nil(t, err)
	assert.Equal(t, []int64{123, 456, 789}, value4)
	_, err = StrToIntArray("?????")
	assert.NotNil(t, err)

	value5, err := StrToFloat("0.123456789")
	assert.Nil(t, err)
	assert.Equal(t, 0.123456789, value5)
	_, err = StrToFloat("?????")
	assert.NotNil(t, err)

	value6, err := StrToFloatArray("0.123,0.456,0.789")
	assert.Nil(t, err)
	assert.Equal(t, []float64{0.123, 0.456, 0.789}, value6)
	_, err = StrToFloatArray("?????")
	assert.NotNil(t, err)

	suite.Run(t, new(SuiteConvert))
}

type SuiteConvert struct {
	suite.Suite
	strArrayRightInput  string
	strArrayRightOutput []string
	strArrayErrorInput  string
	strArrayErrorOutput []string
}

func (this *SuiteConvert) SetupSuite() {
	this.strArrayRightInput = "ball,book,pack"
	this.strArrayRightOutput = []string{"ball", "book", "pack"}
	this.strArrayErrorInput = "ball#book#pack"
	this.strArrayErrorOutput = []string{"ball#book#pack"}
}

func (this *SuiteConvert) TestStrToStrArray() {
	value := StrToStrArray(this.strArrayRightInput)
	assert.Equal(this.T(), this.strArrayRightOutput, value)
	value = StrToStrArray(this.strArrayErrorInput)
	assert.Equal(this.T(), this.strArrayErrorOutput, value)
}
