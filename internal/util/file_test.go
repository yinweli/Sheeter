package util

import (
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/yinweli/Sheeter/testdata"

	"github.com/stretchr/testify/assert"
)

func TestFileWrite(t *testing.T) {
	dir := testdata.ChangeWorkDir()
	defer testdata.RestoreWorkDir(dir)

	filePath := "test/test.txt"
	input := []byte(testdata.Text)

	err := FileWrite(filePath, input, true)
	assert.Nil(t, err)

	err = FileWrite(filePath, input, false)
	assert.Nil(t, err)

	bytes, err := ioutil.ReadFile(filePath)
	assert.Nil(t, err)
	assert.Equal(t, input, bytes)

	err = FileWrite("????/????.txt", input, false)
	assert.NotNil(t, err)

	err = FileWrite("????.txt", input, false)
	assert.NotNil(t, err)

	err = os.RemoveAll(path.Dir(filePath))
	assert.Nil(t, err)
}
