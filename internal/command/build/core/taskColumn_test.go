package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yinweli/Sheeter/testdata"
)

func TestTaskColumn(t *testing.T) {
	task := mockTaskColumn()
	task.excel = testdata.GetTestExcel(testdata.RealExcel)
	err := task.executeColumn()
	assert.Nil(t, err)
	assert.Equal(t, 4, len(task.columns))
	assert.Equal(t, "name0", task.columns[0].Name)
	assert.Equal(t, "note0", task.columns[0].Note)
	assert.Equal(t, (&FieldPkey{}).Type(), task.columns[0].Field.Type())
	assert.Equal(t, "name1", task.columns[1].Name)
	assert.Equal(t, "note1", task.columns[1].Note)
	assert.Equal(t, (&FieldBool{}).Type(), task.columns[1].Field.Type())
	assert.Equal(t, "name2", task.columns[2].Name)
	assert.Equal(t, "note2", task.columns[2].Note)
	assert.Equal(t, (&FieldInt{}).Type(), task.columns[2].Field.Type())
	assert.Equal(t, "name3", task.columns[3].Name)
	assert.Equal(t, "note3", task.columns[3].Note)
	assert.Equal(t, (&FieldText{}).Type(), task.columns[3].Field.Type())
	task.close()

	task = mockTaskColumn()
	task.global.LineOfField = 10
	task.excel = testdata.GetTestExcel(testdata.RealExcel)
	err = task.executeColumn()
	assert.NotNil(t, err)
	task.close()

	task = mockTaskColumn()
	task.excel = testdata.GetTestExcel(testdata.RealExcel)
	task.global.LineOfNote = 10
	err = task.executeColumn()
	assert.NotNil(t, err)
	task.close()

	task = mockTaskColumn()
	task.excel = testdata.GetTestExcel(testdata.Defect3Excel)
	err = task.executeColumn()
	assert.Nil(t, err)
	assert.Equal(t, 4, len(task.columns))
	task.close()

	task = mockTaskColumn()
	task.excel = testdata.GetTestExcel(testdata.Defect4Excel)
	err = task.executeColumn()
	assert.NotNil(t, err)
	task.close()

	task = mockTaskColumn()
	task.excel = testdata.GetTestExcel(testdata.Defect5Excel)
	err = task.executeColumn()
	assert.NotNil(t, err)
	task.close()

	task = mockTaskColumn()
	task.excel = testdata.GetTestExcel(testdata.Defect6Excel)
	err = task.executeColumn()
	assert.NotNil(t, err)
	task.close()

	task = mockTaskColumn()
	task.excel = testdata.GetTestExcel(testdata.Defect7Excel)
	err = task.executeColumn()
	assert.NotNil(t, err)
	task.close()

	task = mockTaskColumn()
	task.excel = testdata.GetTestExcel(testdata.Defect8Excel)
	err = task.executeColumn()
	assert.NotNil(t, err)
	task.close()

	task = mockTaskColumn()
	task.excel = testdata.GetTestExcel(testdata.Defect9Excel)
	err = task.executeColumn()
	assert.NotNil(t, err)
	task.close()
}

func mockTaskColumn() *Task {
	return &Task{
		global: &Global{
			LineOfField: 1,
			LineOfNote:  2,
		},
		element: &Element{
			Excel: testdata.RealExcel,
			Sheet: testdata.SheetName,
		},
	}
}
