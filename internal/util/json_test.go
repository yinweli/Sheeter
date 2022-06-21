package util

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yinweli/Sheeter/testdata"
)

func TestJsonWrite(t *testing.T) {
	dir := testdata.ChangeWorkDir()
	defer testdata.RestoreWorkDir(dir)

	filePath := "json/test.txt"
	value := map[string]string{"data": "value"}
	jsons, _ := json.MarshalIndent(value, "", "    ")

	err := JsonWrite(value, filePath, true)
	assert.Nil(t, err)

	err = JsonWrite(value, filePath, false)
	assert.Nil(t, err)

	bytes, err := ioutil.ReadFile(filePath)
	assert.Nil(t, err)
	assert.Equal(t, jsons, bytes)

	err = JsonWrite(value, "????/????.txt", false)
	assert.NotNil(t, err)

	err = JsonWrite(value, "????.txt", false)
	assert.NotNil(t, err)

	err = os.RemoveAll(path.Dir(filePath))
	assert.Nil(t, err)
}
