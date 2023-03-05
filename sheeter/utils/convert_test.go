package utils

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
	testdata.TestEnv
}

func (this *SuiteConvert) SetupSuite() {
	this.Change("test-convert")
}

func (this *SuiteConvert) TearDownSuite() {
	this.Restore()
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

func (this *SuiteConvert) TestStrToInt32() {
	value, err := StrToInt32("123456789")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), int32(123456789), value)

	_, err = StrToInt32(testdata.UnknownStr)
	assert.NotNil(this.T(), err)
}

func (this *SuiteConvert) TestStrToInt32Array() {
	value, err := StrToInt32Array("123,456,789")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), []int32{123, 456, 789}, value)

	_, err = StrToInt32Array(testdata.UnknownStr)
	assert.NotNil(this.T(), err)
}

func (this *SuiteConvert) TestStrToInt64() {
	value, err := StrToInt64("123456789")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), int64(123456789), value)

	_, err = StrToInt32(testdata.UnknownStr)
	assert.NotNil(this.T(), err)
}

func (this *SuiteConvert) TestStrToInt64Array() {
	value, err := StrToInt64Array("123,456,789")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), []int64{123, 456, 789}, value)

	_, err = StrToInt64Array(testdata.UnknownStr)
	assert.NotNil(this.T(), err)
}

func (this *SuiteConvert) TestStrToFloat32() {
	value, err := StrToFloat32("0.12345")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), float32(0.12345), value)

	_, err = StrToFloat64(testdata.UnknownStr)
	assert.NotNil(this.T(), err)
}

func (this *SuiteConvert) TestStrToFloat32Array() {
	value, err := StrToFloat32Array("0.123,0.456,0.789")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), []float32{0.123, 0.456, 0.789}, value)

	_, err = StrToFloat32Array(testdata.UnknownStr)
	assert.NotNil(this.T(), err)
}

func (this *SuiteConvert) TestStrToFloat64() {
	value, err := StrToFloat64("0.12345")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), 0.12345, value)

	_, err = StrToFloat64(testdata.UnknownStr)
	assert.NotNil(this.T(), err)
}

func (this *SuiteConvert) TestStrToFloat64Array() {
	value, err := StrToFloat64Array("0.123,0.456,0.789")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), []float64{0.123, 0.456, 0.789}, value)

	_, err = StrToFloat64Array(testdata.UnknownStr)
	assert.NotNil(this.T(), err)
}

func (this *SuiteConvert) TestStrToStrArray() {
	assert.Equal(this.T(), []string{"ball", "book", "pack"}, StrToStrArray("ball,book,pack"))
	assert.Equal(this.T(), []string{"ball#book#pack"}, StrToStrArray("ball#book#pack"))
}
