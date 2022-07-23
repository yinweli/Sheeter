package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/yinweli/Sheeter/testdata"
)

func TestConvert(t *testing.T) {
	suite.Run(t, new(SuiteConvert))
}

type SuiteConvert struct {
	suite.Suite
}

func (this *SuiteConvert) TestStrToBool() {
	value, err := StrToBool("true")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), true, value)

	value, err = StrToBool("false")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), false, value)

	value, err = StrToBool("TRUE")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), true, value)

	value, err = StrToBool("FALSE")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), false, value)

	value, err = StrToBool("1")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), true, value)

	value, err = StrToBool("0")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), false, value)

	_, err = StrToBool(testdata.UnknownStr)
	assert.NotNil(this.T(), err)
}

func (this *SuiteConvert) TestStrToBoolArray() {
	value, err := StrToBoolArray("true,false,true,false,true")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), []bool{true, false, true, false, true}, value)

	value, err = StrToBoolArray("TRUE,FALSE,TRUE,FALSE,TRUE")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), []bool{true, false, true, false, true}, value)

	value, err = StrToBoolArray("1,0,1,0,1")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), []bool{true, false, true, false, true}, value)

	_, err = StrToBoolArray("???,???,???,???,???")
	assert.NotNil(this.T(), err)
}

func (this *SuiteConvert) TestStrToInt() {
	value, err := StrToInt("123456789")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), int64(123456789), value)

	_, err = StrToInt(testdata.UnknownStr)
	assert.NotNil(this.T(), err)
}

func (this *SuiteConvert) TestStrToIntArray() {
	value, err := StrToIntArray("123,456,789")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), []int64{123, 456, 789}, value)

	_, err = StrToIntArray(testdata.UnknownStr)
	assert.NotNil(this.T(), err)
}

func (this *SuiteConvert) TestStrToFloat() {
	value, err := StrToFloat("0.123456789")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), 0.123456789, value)

	_, err = StrToFloat(testdata.UnknownStr)
	assert.NotNil(this.T(), err)
}

func (this *SuiteConvert) TestStrToFloatArray() {
	value, err := StrToFloatArray("0.123,0.456,0.789")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), []float64{0.123, 0.456, 0.789}, value)

	_, err = StrToFloatArray(testdata.UnknownStr)
	assert.NotNil(this.T(), err)
}

func (this *SuiteConvert) TestStrToStrArray() {
	assert.Equal(this.T(), []string{"ball", "book", "pack"}, StrToStrArray("ball,book,pack"))
	assert.Equal(this.T(), []string{"ball#book#pack"}, StrToStrArray("ball#book#pack"))
}
