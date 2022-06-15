package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yinweli/Sheeter/testdata"
)

func TestTaskFields(t *testing.T) {
	task := mockTaskFields()
	task.excel = testdata.GetTestExcel(testdata.RealExcel)
	err := task.executeFields()
	assert.Nil(t, err)
	assert.Equal(t, 4, len(task.columns))
	assert.Equal(t, "name0", task.columns[0].Name)
	assert.Equal(t, (&FieldPkey{}).TypeExcel(), task.columns[0].Field.TypeExcel())
	assert.Equal(t, "name1", task.columns[1].Name)
	assert.Equal(t, (&FieldBool{}).TypeExcel(), task.columns[1].Field.TypeExcel())
	assert.Equal(t, "name2", task.columns[2].Name)
	assert.Equal(t, (&FieldInt{}).TypeExcel(), task.columns[2].Field.TypeExcel())
	assert.Equal(t, "name3", task.columns[3].Name)
	assert.Equal(t, (&FieldText{}).TypeExcel(), task.columns[3].Field.TypeExcel())
	task.close()

	task = mockTaskFields()
	task.global.LineOfField = 10
	task.excel = testdata.GetTestExcel(testdata.RealExcel)
	err = task.executeFields()
	assert.NotNil(t, err)
	task.close()

	task = mockTaskFields()
	task.excel = testdata.GetTestExcel(testdata.Defect3Excel)
	err = task.executeFields()
	assert.Nil(t, err)
	assert.Equal(t, 4, len(task.columns))
	task.close()

	task = mockTaskFields()
	task.excel = testdata.GetTestExcel(testdata.Defect4Excel)
	err = task.executeFields()
	assert.NotNil(t, err)
	task.close()

	task = mockTaskFields()
	task.excel = testdata.GetTestExcel(testdata.Defect5Excel)
	err = task.executeFields()
	assert.NotNil(t, err)
	task.close()

	task = mockTaskFields()
	task.excel = testdata.GetTestExcel(testdata.Defect6Excel)
	err = task.executeFields()
	assert.NotNil(t, err)
	task.close()

	task = mockTaskFields()
	task.excel = testdata.GetTestExcel(testdata.Defect7Excel)
	err = task.executeFields()
	assert.NotNil(t, err)
	task.close()

	task = mockTaskFields()
	task.excel = testdata.GetTestExcel(testdata.Defect8Excel)
	err = task.executeFields()
	assert.NotNil(t, err)
	task.close()

	task = mockTaskFields()
	task.excel = testdata.GetTestExcel(testdata.Defect9Excel)
	err = task.executeFields()
	assert.NotNil(t, err)
	task.close()
}

func mockTaskFields() *Task {
	return &Task{
		global: &Global{
			LineOfField: 1,
		},
		element: &Element{
			Excel: testdata.RealExcel,
			Sheet: testdata.SheetName,
		},
	}
}
