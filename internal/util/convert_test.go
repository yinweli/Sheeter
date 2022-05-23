package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringToBoolArray(t *testing.T) {
	var result []bool
	var err error

	realString := "true,false,true,false,0,0,1,1"
	realArray := []bool{true, false, true, false, false, false, true, true}
	result, err = StringToBoolArray(realString)
	assert.Equal(t, realArray, result, "")
	assert.Nil(t, err, "")

	fakeString := "fake"
	result, err = StringToBoolArray(fakeString)
	assert.NotNil(t, err, "")
}
