package util

import (
	"testing"

	"Sheeter/testdata"

	"github.com/stretchr/testify/assert"
)

func TestStringToBoolArray(t *testing.T) {
	var result []bool
	var err error

	result, err = StringToBoolArray(testdata.BoolString())
	assert.Equal(t, testdata.BoolArray(), result, "convert real failed")
	assert.Nil(t, err, "convert real failed")

	result, err = StringToBoolArray("fake")
	assert.Nil(t, result, "convert fake failed")
	assert.NotNil(t, err, "convert fake failed")
}

func TestBoolArrayToString(t *testing.T) {
	assert.Equal(t, testdata.BoolString(), BoolArrayToString(testdata.BoolArray()), "convert failed")
}

func TestStringToFloat32Array(t *testing.T) {
	var result []float32
	var err error

	result, err = StringToFloat32Array(testdata.Float32String())
	assert.InDeltaSlice(t, testdata.Float32Array(), result, 0.000001, "convert real failed")
	assert.Nil(t, err, "convert real failed")

	result, err = StringToFloat32Array("fake")
	assert.Nil(t, result, "convert fake failed")
	assert.NotNil(t, err, "convert fake failed")
}

func TestFloat32ArrayToString(t *testing.T) {
	assert.Equal(t, testdata.Float32String(), Float32ArrayToString(testdata.Float32Array()), "convert failed")
}

func TestStringToFloat64Array(t *testing.T) {
	var result []float64
	var err error

	result, err = StringToFloat64Array(testdata.Float64String())
	assert.InDeltaSlice(t, testdata.Float64Array(), result, 0.000001, "convert real failed")
	assert.Nil(t, err, "convert real failed")

	result, err = StringToFloat64Array("fake")
	assert.Nil(t, result, "convert fake failed")
	assert.NotNil(t, err, "convert fake failed")
}

func TestFloat64ArrayToString(t *testing.T) {
	assert.Equal(t, testdata.Float64String(), Float64ArrayToString(testdata.Float64Array()), "convert failed")
}

func TestTrimFloatString(t *testing.T) {
	assert.Equal(t, "trim", trimFloatString("trim"), "trim float string failed")
	assert.Equal(t, "trim", trimFloatString("trim.00"), "trim float string failed")
	assert.Equal(t, "trim.1", trimFloatString("trim.10"), "trim float string failed")
	assert.Equal(t, "trim.01", trimFloatString("trim.010"), "trim float string failed")
}
