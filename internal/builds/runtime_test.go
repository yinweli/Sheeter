package builds

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/internal/fields"
	"github.com/yinweli/Sheeter/internal/layouts"
	"github.com/yinweli/Sheeter/internal/utils"
	"github.com/yinweli/Sheeter/testdata"
)

func TestRuntime(t *testing.T) {
	suite.Run(t, new(SuiteRuntime))
}

type SuiteRuntime struct {
	suite.Suite
	workDir string
	name    string
	note    string
	alter   string
	field   fields.Field
	field1  *layouts.Field
	field2  *layouts.Field
	field3  *layouts.Field
}

func (this *SuiteRuntime) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
	this.name = "name"
	this.note = "note"
	this.alter = "alter"
	this.field = &fields.Pkey{}
	this.field1 = &layouts.Field{
		Name:  this.name,
		Note:  this.note,
		Field: this.field,
		Alter: "",
		Array: false,
	}
	this.field2 = &layouts.Field{
		Name:  this.name,
		Note:  this.note,
		Field: nil,
		Alter: this.alter,
		Array: false,
	}
	this.field3 = &layouts.Field{
		Name:  this.name,
		Note:  this.note,
		Field: nil,
		Alter: this.alter,
		Array: true,
	}
}

func (this *SuiteRuntime) TearDownSuite() {
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteRuntime) runtimeSector() *RuntimeSector {
	sector := &RuntimeSector{
		Global: Global{
			LineOfField: 1,
			LineOfLayer: 2,
			LineOfNote:  3,
			LineOfData:  4,
		},
		Element: Element{
			Excel: testdata.ExcelNameReal,
			Sheet: testdata.SheetName,
		},
	}
	return sector
}

func (this *SuiteRuntime) runtimeStruct() *RuntimeStruct {
	return &RuntimeStruct{}
}

func (this *SuiteRuntime) TestGetRows() {
	sector := this.runtimeSector()
	sector.excel = testdata.GetTestExcel(testdata.ExcelNameReal)

	rows, err := sector.GetRows(1)
	assert.Nil(this.T(), err)
	assert.NotNil(this.T(), rows)
	_ = rows.Close()

	rows, err = sector.GetRows(10)
	assert.Nil(this.T(), err)
	assert.NotNil(this.T(), rows)
	_ = rows.Close()

	_, err = sector.GetRows(0)
	assert.NotNil(this.T(), err)

	sector.Sheet = testdata.UnknownStr
	_, err = sector.GetRows(1)
	assert.NotNil(this.T(), err)

	sector.Close()
}

func (this *SuiteRuntime) TestGetColumns() {
	sector := this.runtimeSector()
	sector.excel = testdata.GetTestExcel(testdata.ExcelNameReal)

	cols, err := sector.GetColumns(1)
	assert.Nil(this.T(), err)
	assert.Equal(this.T(),
		[]string{
			"name0#pkey",
			"empty#empty",
			"name1#bool",
			"name2#int",
			"name3#text",
			"name2#int",
			"name3#text",
			"name2#int",
			"name3#text",
		}, cols)

	_, err = sector.GetColumns(10)
	assert.NotNil(this.T(), err)

	_, err = sector.GetColumns(0)
	assert.NotNil(this.T(), err)

	sector.Sheet = testdata.UnknownStr
	_, err = sector.GetColumns(1)
	assert.NotNil(this.T(), err)

	sector.Close()
}

func (this *SuiteRuntime) TestRuntimeStruct() {
	target := this.runtimeStruct()
	assert.Equal(this.T(), utils.FirstUpper(this.name), target.FieldName(this.field1))
	assert.Equal(this.T(), this.note, target.FieldNote(this.field1))
	assert.Equal(this.T(), this.field.ToTypeCs(), target.FieldTypeCs(this.field1))
	assert.Equal(this.T(), this.field.ToTypeGo(), target.FieldTypeGo(this.field1))
	assert.Equal(this.T(), this.alter, target.FieldTypeCs(this.field2))
	assert.Equal(this.T(), this.alter, target.FieldTypeGo(this.field2))
	assert.Equal(this.T(), this.alter+internal.TokenArray, target.FieldTypeCs(this.field3))
	assert.Equal(this.T(), internal.TokenArray+this.alter, target.FieldTypeGo(this.field3))
}
