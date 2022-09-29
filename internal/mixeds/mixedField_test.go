package mixeds

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

func TestField(t *testing.T) {
	suite.Run(t, new(SuiteField))
}

type SuiteField struct {
	suite.Suite
	workDir         string
	excel           string
	sheet           string
	name            string
	note            string
	alter           string
	field           fields.Field
	fieldValue      *layouts.Field
	fieldValueArray *layouts.Field
	fieldAlter      *layouts.Field
	fieldAlterArray *layouts.Field
}

func (this *SuiteField) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
	this.excel = "excelField"
	this.sheet = "sheetField"
	this.name = "name"
	this.note = "note"
	this.alter = "alter"
	this.field = &fields.Int{}
	this.fieldValue = &layouts.Field{
		Name:  this.name,
		Note:  this.note,
		Field: this.field,
		Alter: "",
		Array: false,
	}
	this.fieldValueArray = &layouts.Field{
		Name:  this.name,
		Note:  this.note,
		Field: this.field,
		Alter: "",
		Array: true,
	}
	this.fieldAlter = &layouts.Field{
		Name:  this.name,
		Note:  this.note,
		Field: nil,
		Alter: this.alter,
		Array: false,
	}
	this.fieldAlterArray = &layouts.Field{
		Name:  this.name,
		Note:  this.note,
		Field: nil,
		Alter: this.alter,
		Array: true,
	}
}

func (this *SuiteField) TearDownSuite() {
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteField) target() *Mixed {
	target := NewMixed(this.excel, this.sheet)
	return target
}
func (this *SuiteField) TestField() {
	name := utils.FirstUpper(this.name)
	note := this.note

	target := this.target()
	assert.Equal(this.T(), name, target.FieldName(this.fieldValue))
	assert.Equal(this.T(), note, target.FieldNote(this.fieldValue))
}

func (this *SuiteField) TestFieldTypeCs() {
	fieldValue := this.field.ToTypeCs()
	fieldValueArray := this.field.ToTypeCs() + internal.TokenArray
	fieldAlter := this.alter
	fieldAlterArray := this.alter + internal.TokenArray

	target := this.target()
	assert.Equal(this.T(), fieldValue, target.FieldTypeCs(this.fieldValue))
	assert.Equal(this.T(), fieldValueArray, target.FieldTypeCs(this.fieldValueArray))
	assert.Equal(this.T(), fieldAlter, target.FieldTypeCs(this.fieldAlter))
	assert.Equal(this.T(), fieldAlterArray, target.FieldTypeCs(this.fieldAlterArray))
}

func (this *SuiteField) TestFieldTypeGo() {
	fieldValue := this.field.ToTypeGo()
	fieldValueArray := internal.TokenArray + this.field.ToTypeGo()
	fieldAlter := this.alter
	fieldAlterArray := internal.TokenArray + this.alter

	target := this.target()
	assert.Equal(this.T(), fieldValue, target.FieldTypeGo(this.fieldValue))
	assert.Equal(this.T(), fieldValueArray, target.FieldTypeGo(this.fieldValueArray))
	assert.Equal(this.T(), fieldAlter, target.FieldTypeGo(this.fieldAlter))
	assert.Equal(this.T(), fieldAlterArray, target.FieldTypeGo(this.fieldAlterArray))
}

func (this *SuiteField) TestFieldTypeProto() {
	fieldValue := internal.TokenOptional + " " + this.field.ToTypeProto()
	fieldValueArray := internal.TokenRepeated + " " + this.field.ToTypeProto()
	fieldAlter := internal.TokenOptional + " " + this.alter
	fieldAlterArray := internal.TokenRepeated + " " + this.alter

	target := this.target()
	assert.Equal(this.T(), fieldValue, target.FieldTypeProto(this.fieldValue))
	assert.Equal(this.T(), fieldValueArray, target.FieldTypeProto(this.fieldValueArray))
	assert.Equal(this.T(), fieldAlter, target.FieldTypeProto(this.fieldAlter))
	assert.Equal(this.T(), fieldAlterArray, target.FieldTypeProto(this.fieldAlterArray))
}

func (this *SuiteField) TestPkeyType() {
	target := this.target()
	pkey := fields.Pkey{}
	assert.Equal(this.T(), pkey.ToTypeCs(), target.PkeyTypeCs())
	assert.Equal(this.T(), pkey.ToTypeGo(), target.PkeyTypeGo())
	assert.Equal(this.T(), pkey.ToTypeProto(), target.PkeyTypeProto())
}
