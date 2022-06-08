package util

import (
	"io/ioutil"
	"os"
	"path"
	"testing"

	"Sheeter/testdata"

	"github.com/stretchr/testify/assert"
)

func TestFileWrite(t *testing.T) {
	dir := testdata.ChangeWorkDir()
	defer testdata.RestoreWorkDir(dir)

	filePath := "test/test.txt"
	input := []byte(testdata.Text)

	err := FileWrite(filePath, input)
	assert.Nil(t, err)

	bytes, err := ioutil.ReadFile(filePath)
	assert.Nil(t, err)
	assert.Equal(t, input, bytes)

	err = FileWrite("????/????.txt", input)
	assert.NotNil(t, err)

	err = FileWrite("????.txt", input)
	assert.NotNil(t, err)

	err = os.RemoveAll(path.Dir(filePath))
	assert.Nil(t, err)
}
