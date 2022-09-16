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

func TestEntire(t *testing.T) {
	suite.Run(t, new(SuiteEntire))
}

type SuiteEntire struct {
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

func (this *SuiteEntire) SetupSuite() {
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

func (this *SuiteEntire) TearDownSuite() {
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteEntire) target() *Entire {
	target := &Entire{}
	return target
}

func (this *SuiteEntire) TestEntire() {
	target := this.target()
	assert.Equal(this.T(), utils.FirstUpper(this.name), target.FieldName(this.field1))
	assert.Equal(this.T(), this.note, target.FieldNote(this.field1))
	assert.Equal(this.T(), this.field.ToTypeCs(), target.FieldTypeCs(this.field1))
	assert.Equal(this.T(), this.field.ToTypeGo(), target.FieldTypeGo(this.field1))
	assert.Equal(this.T(), this.alter, target.FieldTypeCs(this.field2))
	assert.Equal(this.T(), this.alter, target.FieldTypeGo(this.field2))
	assert.Equal(this.T(), this.alter+internal.TokenArray, target.FieldTypeCs(this.field3))
	assert.Equal(this.T(), internal.TokenArray+this.alter, target.FieldTypeGo(this.field3))
}
