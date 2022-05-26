package core

import (
	"testing"

	"Sheeter/internal"
	"Sheeter/testdata"

	"github.com/schollz/progressbar/v3"

	"github.com/stretchr/testify/assert"
)

func TestReadSheet(t *testing.T) {
	task := mockTask()

	err := ReadSheet(task)
	assert.Nil(t, err)
}

func TestBuildSheet(t *testing.T) {
	task := mockTask()

	sheet, err := buildSheet(task)
	assert.Nil(t, err)
	assert.Equal(t, 12, len(sheet))
	assert.Equal(t, 16, len(sheet[0]))
	assert.Equal(t, "checkpoint", sheet[0][15])
	assert.Equal(t, "checkpoint", sheet[11][15])

	task.Element.Excel = testdata.FakeExcel
	task.Element.Sheet = testdata.FakeSheet
	sheet, err = buildSheet(task)
	assert.NotNil(t, err)

	task.Element.Excel = testdata.FakeExcel
	task.Element.Sheet = testdata.UnknownSheet
	sheet, err = buildSheet(task)
	assert.NotNil(t, err)

	task.Element.Excel = testdata.UnknownExcel
	sheet, err = buildSheet(task)
	assert.NotNil(t, err)
}

func TestBuildColumns(t *testing.T) {
	task := mockTask()
	fields := []string{"field0#pkey", "field1#bool", "field2#int", "", "field3#text"}

	pkey, err := buildColumns(task, [][]string{{}, fields})
	assert.Nil(t, err)
	assert.NotNil(t, pkey)
	assert.Equal(t, "field0", pkey.Name)
	assert.Equal(t, (&FieldPkey{}).TypeExcel(), pkey.Field.TypeExcel())
	assert.Equal(t, 3, len(task.Columns))
	assert.Equal(t, "field0", task.Columns[0].Name)
	assert.Equal(t, (&FieldPkey{}).TypeExcel(), task.Columns[0].Field.TypeExcel())
	assert.Equal(t, "field1", task.Columns[1].Name)
	assert.Equal(t, (&FieldBool{}).TypeExcel(), task.Columns[1].Field.TypeExcel())
	assert.Equal(t, "field2", task.Columns[2].Name)
	assert.Equal(t, (&FieldInt{}).TypeExcel(), task.Columns[2].Field.TypeExcel())

	fields = []string{"field0#????", "field1#bool", "field2#int"}
	pkey, err = buildColumns(task, [][]string{{}, fields})
	assert.NotNil(t, err)

	fields = []string{"field0#pkey", "field1#pkey", "field2#int"}
	pkey, err = buildColumns(task, [][]string{{}, fields})
	assert.NotNil(t, err)

	fields = []string{"field0#int", "field1#int", "field2#int"}
	pkey, err = buildColumns(task, [][]string{{}, fields})
	assert.NotNil(t, err)

	fields = []string{}
	pkey, err = buildColumns(task, [][]string{{}, fields})
	assert.NotNil(t, err)
}

func TestBuildNotes(t *testing.T) {
	task := mockTask()
	notes := []string{"note0", "note1", "note2"}

	err := buildNotes(task, [][]string{notes})
	assert.Nil(t, err)
	assert.Equal(t, "note0", task.Columns[0].Note)
	assert.Equal(t, "note1", task.Columns[1].Note)
	assert.Equal(t, "note2", task.Columns[2].Note)
}

func TestBuildDatas(t *testing.T) {
	task := mockTask()
	data0 := []string{"data0", "data1", "data2"}
	data1 := []string{"data4", "data5", "data6"}
	data2 := []string{"data7", "data8", "data9"}

	err := buildDatas(task, [][]string{{}, {}, data0, data1, data2})
	assert.Nil(t, err)
	assert.Equal(t, []string{"data0", "data4", "data7"}, task.Columns[0].Datas)
	assert.Equal(t, []string{"data1", "data5", "data8"}, task.Columns[1].Datas)
	assert.Equal(t, []string{"data2", "data6", "data9"}, task.Columns[2].Datas)
}

func TestPkeyCheck(t *testing.T) {
	task := mockTask()
	pkey := &Column{Datas: []string{"1", "2", "3", "4", "5"}}

	err := pkeyCheck(task, pkey)
	assert.Nil(t, err)

	pkey.Datas = append(pkey.Datas, "5")
	err = pkeyCheck(task, pkey)
	assert.NotNil(t, err)
}

func mockTask() *Task {
	return &Task{
		Progress: progressbar.DefaultSilent(internal.ProgressDefault),
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
