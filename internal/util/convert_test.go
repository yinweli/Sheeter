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
