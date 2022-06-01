package util

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"Sheeter/testdata"

	"github.com/stretchr/testify/assert"
)

func TestFileWrite(t *testing.T) {
	dir := testdata.ChangeWorkDir()
	defer testdata.RestoreWorkDir(dir)

	path := "test"
	name := "test.txt"
	input := []byte(testdata.Text)

	filePath, err := FileWrite(path, name, input)
	assert.Nil(t, err)
	assert.Equal(t, filepath.Join(path, name), filePath)

	bytes, err := ioutil.ReadFile(filePath)
	assert.Nil(t, err)
	assert.Equal(t, input, bytes)

	filePath, err = FileWrite(path, "", input)
	assert.NotNil(t, err)

	err = os.RemoveAll(path)
	assert.Nil(t, err)
}
