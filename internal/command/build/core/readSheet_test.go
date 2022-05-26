package core

import (
	"testing"

	"Sheeter/internal"
	"Sheeter/testdata"

	"github.com/stretchr/testify/assert"

	"github.com/schollz/progressbar/v3"
)

func TestReadSheet(t *testing.T) {

}

func TestBuildColumns(t *testing.T) {
	task := mockTask()
	fields := []string{"field0#pkey", "field1#bool", "field2#int", "", "field3#text"}

	pkey, err := buildColumns(task, [][]string{{}, fields})
	assert.Nil(t, err, "build columns failed")
	assert.NotNil(t, pkey, "build columns failed")
}

func TestBuildNotes(t *testing.T) {
	task := mockTask()
	notes := []string{"note0", "note1", "note2"}

	err := buildNotes(task, [][]string{notes})
	assert.Nil(t, err, "build notes failed")
	assert.Equal(t, "note0", task.Columns[0].Note, "build notes failed")
	assert.Equal(t, "note1", task.Columns[1].Note, "build notes failed")
	assert.Equal(t, "note2", task.Columns[2].Note, "build notes failed")
}

func TestBuildDatas(t *testing.T) {
	task := mockTask()
	data0 := []string{"data0", "data1", "data2"}
	data1 := []string{"data4", "data5", "data6"}
	data2 := []string{"data7", "data8", "data9"}

	err := buildDatas(task, [][]string{{}, {}, data0, data1, data2})
	assert.Nil(t, err, "build datas failed")
	assert.Equal(t, []string{"data0", "data4", "data7"}, task.Columns[0].Datas, "build datas failed")
	assert.Equal(t, []string{"data1", "data5", "data8"}, task.Columns[1].Datas, "build datas failed")
	assert.Equal(t, []string{"data2", "data6", "data9"}, task.Columns[2].Datas, "build datas failed")
}

func TestPkeyCheck(t *testing.T) {
	task := mockTask()
	pkey := &Column{Datas: []string{"1", "2", "3", "4", "5"}}

	err := pkeyCheck(task, pkey)
	assert.Nil(t, err, "pkey check failed")

	pkey.Datas = append(pkey.Datas, "5")
	err = pkeyCheck(task, pkey)
	assert.NotNil(t, err, "pkey check failed")
}

func mockTask() *Task {
	return &Task{
		Progress: progressbar.Default(internal.ProgressDefault),
		Global: &Global{
			ExcelPath:   testdata.RootPath,
			LineOfNote:  1,
			LineOfField: 2,
			LineOfData:  3,
		},
		Element: &Element{
			Excel: testdata.TestExcel,
			Sheet: testdata.TestSheet,
		},
		Columns: []*Column{
			{Note: "note0", Name: "name0", Field: &FieldInt{}, Datas: []string{}},
			{Note: "note1", Name: "name1", Field: &FieldInt{}, Datas: []string{}},
			{Note: "note2", Name: "name2", Field: &FieldInt{}, Datas: []string{}},
		},
	}
}
