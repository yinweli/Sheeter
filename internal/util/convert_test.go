package util

import (
	"testing"

	"Sheeter/testdata"

	"github.com/stretchr/testify/assert"
)

func TestStringToBoolArray(t *testing.T) {
	result, err := StringToBoolArray(testdata.BoolString())
	assert.Nil(t, err)
	assert.Equal(t, testdata.BoolArray(), result)

	result, err = StringToBoolArray("fake")
	assert.NotNil(t, err)
}

func TestBoolArrayToString(t *testing.T) {
	assert.Equal(t, testdata.BoolString(), BoolArrayToString(testdata.BoolArray()))
}

func TestStringToInt32Array(t *testing.T) {
	result, err := StringToInt32Array(testdata.Int32String())
	assert.Nil(t, err)
	assert.Equal(t, testdata.Int32Array(), result)

	result, err = StringToInt32Array("fake")
	assert.NotNil(t, err)
}

func TestInt32ArrayToString(t *testing.T) {
	assert.Equal(t, testdata.Int32String(), Int32ArrayToString(testdata.Int32Array()))
}
func TestStringToInt64Array(t *testing.T) {
	result, err := StringToInt64Array(testdata.Int64String())
	assert.Nil(t, err)
	assert.Equal(t, testdata.Int64Array(), result)

	result, err = StringToInt64Array("fake")
	assert.NotNil(t, err)
}

func TestInt64ArrayToString(t *testing.T) {
	assert.Equal(t, testdata.Int64String(), Int64ArrayToString(testdata.Int64Array()))
}

func TestStringToFloat32Array(t *testing.T) {
	result, err := StringToFloat32Array(testdata.Float32String())
	assert.Nil(t, err)
	assert.InDeltaSlice(t, testdata.Float32Array(), result, 0.000001)

	result, err = StringToFloat32Array("fake")
	assert.NotNil(t, err)
}

func TestFloat32ArrayToString(t *testing.T) {
	assert.Equal(t, testdata.Float32String(), Float32ArrayToString(testdata.Float32Array()))
}

func TestStringToFloat64Array(t *testing.T) {
	result, err := StringToFloat64Array(testdata.Float64String())
	assert.Nil(t, err)
	assert.InDeltaSlice(t, testdata.Float64Array(), result, 0.000001)

	result, err = StringToFloat64Array("fake")
	assert.NotNil(t, err)
}

func TestFloat64ArrayToString(t *testing.T) {
	assert.Equal(t, testdata.Float64String(), Float64ArrayToString(testdata.Float64Array()))
}

func TestStringToStringArray(t *testing.T) {
	assert.Equal(t, testdata.StringArray(), StringToStringArray(testdata.StringString()))

}

func TestStringArrayToString(t *testing.T) {
	assert.Equal(t, testdata.StringString(), StringArrayToString(testdata.StringArray()))
}

func TestTrimFloatString(t *testing.T) {
	assert.Equal(t, "trim", trimFloatString("trim"))
	assert.Equal(t, "trim", trimFloatString("trim.00"))
	assert.Equal(t, "trim.1", trimFloatString("trim.10"))
	assert.Equal(t, "trim.01", trimFloatString("trim.010"))
}
