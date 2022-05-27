package core

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"Sheeter/internal/util"
	"Sheeter/testdata"

	"github.com/stretchr/testify/assert"
)

func TestWriteJson(t *testing.T) {
	dir := testdata.ChangeWorkDir()
	defer testdata.RestoreWorkDir(dir)

	cargo := mockWriteJsonCargo()

	err := WriteJson(cargo)
	assert.Nil(t, err)

	cargo = mockWriteJsonCargo()
	cargo.Columns = []*Column{
		{Note: "note0", Name: "name0", Field: &FieldInt{}, Datas: []string{"1", "2", "3", "4", "5"}},
		{Note: "note1", Name: "name1", Field: &FieldInt{}, Datas: []string{"x", "2", "3", "4", "5"}},
		{Note: "note2", Name: "name2", Field: &FieldInt{}, Datas: []string{"1", "2", "3", "4", "5"}},
	}
	err = WriteJson(cargo)
	assert.NotNil(t, err)

	err = os.RemoveAll(OutputPathJson)
	assert.Nil(t, err)
}

func TestBuildJBoxes(t *testing.T) {
	cargo := mockWriteJsonCargo()

	expected := []jbox{
		{"name0": int64(1), "name1": int64(1), "name2": int64(1)},
		{"name0": int64(2), "name1": int64(2), "name2": int64(2)},
		{"name0": int64(3), "name1": int64(3), "name2": int64(3)},
		{"name0": int64(4), "name1": int64(4), "name2": int64(4)},
		{"name0": int64(5), "name1": int64(5), "name2": int64(5)},
	}
	result, err := buildJBoxes(cargo)
	assert.Nil(t, err)
	assert.Equal(t, expected, result)

	cargo.Columns = []*Column{
		{Note: "note0", Name: "name0", Field: &FieldInt{}, Datas: []string{"1", "2", "3", "4", "5"}},
		{Note: "note1", Name: "name1", Field: &FieldInt{}, Datas: []string{"x", "2", "3", "4", "5"}},
		{Note: "note2", Name: "name2", Field: &FieldInt{}, Datas: []string{"1", "2", "3", "4", "5"}},
	}

	result, err = buildJBoxes(cargo)
	assert.NotNil(t, err)
}

func TestBuildJson(t *testing.T) {
	cargo := mockWriteJsonCargo()
	jboxes := []jbox{
		{"name0": int64(1), "name1": int64(1), "name2": int64(1)},
		{"name0": int64(2), "name1": int64(2), "name2": int64(2)},
		{"name0": int64(3), "name1": int64(3), "name2": int64(3)},
		{"name0": int64(4), "name1": int64(4), "name2": int64(4)},
		{"name0": int64(5), "name1": int64(5), "name2": int64(5)},
	}

	expected, err := json.MarshalIndent(jboxes, "", "    ")
	assert.Nil(t, err)

	result, err := buildJson(cargo, jboxes)
	assert.Nil(t, err)
	assert.Equal(t, expected, result)

	// 要讓json吐出錯誤還真不容易...
	// 目前已知通道物件是不能轉換為json的, 所以就塞了個通道物件給json, 強制json轉換失敗
	jboxes = []jbox{
		{"name0": make(chan int), "name1": int64(1), "name2": int64(1)},
		{"name0": int64(2), "name1": int64(2), "name2": int64(2)},
		{"name0": int64(3), "name1": int64(3), "name2": int64(3)},
		{"name0": int64(4), "name1": int64(4), "name2": int64(4)},
		{"name0": int64(5), "name1": int64(5), "name2": int64(5)},
	}

	result, err = buildJson(cargo, jboxes)
	assert.NotNil(t, err)
}

func TestWriteFile(t *testing.T) {
	dir := testdata.ChangeWorkDir()
	defer testdata.RestoreWorkDir(dir)

	cargo := mockWriteJsonCargo()
	input := []byte(testdata.Text)

	err := writeFile(cargo, input)
	assert.Nil(t, err)

	output, _ := ioutil.ReadFile(filepath.Join(OutputPathJson, cargo.JsonFileName()))
	assert.Equal(t, input, output)

	err = os.RemoveAll(OutputPathJson)
	assert.Nil(t, err)
}

func mockWriteJsonCargo() *Cargo {
	return &Cargo{
		Progress: util.NewProgressBar("test", ioutil.Discard),
		Global: &Global{
			ExcelPath:   testdata.RootPath,
			LineOfNote:  1,
			LineOfField: 2,
			LineOfData:  3,
		},
		Element: &Element{
			Excel: testdata.RealExcel,
			Sheet: testdata.RealSheet,
		},
		Columns: []*Column{
			{Note: "note0", Name: "name0", Field: &FieldInt{}, Datas: []string{"1", "2", "3", "4", "5"}},
			{Note: "note1", Name: "name1", Field: &FieldInt{}, Datas: []string{"1", "2", "3", "4", "5"}},
			{Note: "note2", Name: "name2", Field: &FieldInt{}, Datas: []string{"1", "2", "3", "4", "5"}},
		},
	}
}
