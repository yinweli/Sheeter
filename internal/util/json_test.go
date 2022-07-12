package util

import (
	"encoding/json"
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

	err := JsonWrite(filePath, true, value)
	assert.Nil(t, err)

	err = JsonWrite(filePath, false, value)
	assert.Nil(t, err)

	bytes, err := os.ReadFile(filePath)
	assert.Nil(t, err)
	assert.Equal(t, jsons, bytes)

	err = JsonWrite("????/????.txt", false, value)
	assert.NotNil(t, err)

	err = JsonWrite("????.txt", false, value)
	assert.NotNil(t, err)

	err = os.RemoveAll(path.Dir(filePath))
	assert.Nil(t, err)
}
