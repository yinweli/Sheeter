package util

import (
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yinweli/Sheeter/testdata"
)

func TestTmplWrite(t *testing.T) {
	dir := testdata.ChangeWorkDir()
	defer testdata.RestoreWorkDir(dir)

	filePath := "tmpl/test.txt"
	content := "{{$.Value}}"
	tmplTest := &TmplTest{Value: "Value"}

	err := TmplWrite(filePath, true, content, tmplTest)
	assert.Nil(t, err)

	err = TmplWrite(filePath, false, content, tmplTest)
	assert.Nil(t, err)

	bytes, err := os.ReadFile(filePath)
	assert.Nil(t, err)
	assert.Equal(t, []byte("Value"), bytes)

	err = TmplWrite(filePath, false, "{{{$.Value}}", nil)
	assert.NotNil(t, err)

	err = TmplWrite(filePath, false, content, "nothing!")
	assert.NotNil(t, err)

	err = TmplWrite("?????", false, content, tmplTest)
	assert.NotNil(t, err)

	err = os.RemoveAll(path.Dir(filePath))
	assert.Nil(t, err)
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
