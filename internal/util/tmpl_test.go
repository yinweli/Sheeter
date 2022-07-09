package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTmplExecute(t *testing.T) {

	buffer, err := TmplExecute("tmpl", "{{$.Value}}", TmplTest{Value: "Value"})
	assert.Nil(t, err)
	assert.Equal(t, "Value", buffer.String())
	buffer, err = TmplExecute("tmpl", "{{{$.Value}}", nil)
	assert.NotNil(t, err)
	buffer, err = TmplExecute("tmpl", "{{$.Value}}", "nothing!")
	assert.NotNil(t, err)
}

type TmplTest struct {
	Value string
}

func TestTmplLine(t *testing.T) {
	tmplLine := mockTmplLine()
	assert.Equal(t, "", tmplLine.SetLine(2))
	assert.Equal(t, "\n", tmplLine.NewLine())
	assert.Equal(t, "\n", tmplLine.NewLine())
	assert.Equal(t, "", tmplLine.NewLine())
	assert.Equal(t, "", tmplLine.NewLine())
}

func mockTmplLine() *TmplLine {
	return &TmplLine{}
}

func TestTmplFirstChar(t *testing.T) {
	tmplFirstChar := mockTmplFirstChar()
	assert.Equal(t, "TestColumn", tmplFirstChar.FirstUpper("testColumn"))
	assert.Equal(t, "testColumn", tmplFirstChar.FirstLower("TestColumn"))
}

func mockTmplFirstChar() *TmplFirstChar {
	return &TmplFirstChar{}
}
