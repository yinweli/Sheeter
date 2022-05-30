package util

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"Sheeter/testdata"

	"github.com/stretchr/testify/assert"
)

func TestWriteFile(t *testing.T) {
	path := "test"
	name := "test.txt"
	input := []byte(testdata.Text)

	err := WriteFile(path, name, input)
	assert.Nil(t, err)

	output, _ := ioutil.ReadFile(filepath.Join(path, name))
	assert.Equal(t, input, output)

	err = os.RemoveAll(path)
	assert.Nil(t, err)
}
