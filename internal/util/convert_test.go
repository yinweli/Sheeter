package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringToBoolArray(t *testing.T) {
	var result []bool
	var err error

	result, err = StringToBoolArray("true,false,t,f,1,0")
	assert.Equal(t, []bool{true, false, true, false, true, false}, result, "convert real failed")
	assert.Nil(t, err, "convert real failed")

	fakeString := "fake"
	result, err = StringToBoolArray(fakeString)
	assert.Nil(t, result, "convert fake failed")
	assert.NotNil(t, err, "convert fake failed")
}
