package layouts

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/internal/fields"
	"github.com/yinweli/Sheeter/internal/layers"
	"github.com/yinweli/Sheeter/testdata"
)

func TestLayoutType(t *testing.T) {
	suite.Run(t, new(SuiteLayoutType))
}

type SuiteLayoutType struct {
	suite.Suite
	workDir    string
	type1      string
	type2      string
	typeA      string
	typeB      string
	typeC      string
	typeD      string
	field1     *Field
	field2     *Field
	fieldName  string
	fieldNote  string
	fieldField fields.Field
}

func (this *SuiteLayoutType) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
	this.type1 = "type1"
	this.type2 = "type2"
	this.typeA = "typeA"
	this.typeB = "typeB"
	this.typeC = "typeC"
	this.typeD = "typeD"
	this.field1 = &Field{
		Name:  "name1",
		Note:  "note1",
		Field: &fields.Text{},
		Alter: "alter1",
		Array: true,
	}
	this.field2 = &Field{
		Name:  "name2",
		Note:  "note2",
		Field: &fields.TextArray{},
		Alter: "alter2",
		Array: false,
	}
	this.fieldName = "name"
	this.fieldNote = "note"
	this.fieldField = &fields.Int{}
}

func (this *SuiteLayoutType) TearDownSuite() {
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteLayoutType) target() *LayoutType {
	return NewLayoutType()
}

func (this *SuiteLayoutType) TestNewLayoutType() {
	assert.NotNil(this.T(), NewLayoutType())
}

func (this *SuiteLayoutType) TestBegin() {
	target := this.target()
	assert.Nil(this.T(), target.Begin(this.type1, nil))
	assert.NotNil(this.T(), target.Begin(this.type2, nil))
}

func (this *SuiteLayoutType) TestEnd() {
	target := this.target()
	assert.Nil(this.T(), target.Begin(this.type1, nil))
	assert.Nil(this.T(), target.End())
	assert.NotNil(this.T(), target.End())
}

func (this *SuiteLayoutType) TestAdd() {
	target := this.target()
	assert.Nil(this.T(), target.Begin(this.type1, nil))
	assert.Nil(this.T(), target.Add(this.fieldName, this.fieldNote, this.fieldField, []layers.Layer{{Name: this.typeA, Type: layers.LayerArray}, {Name: this.typeB, Type: layers.LayerStruct}}, 2))

	target = this.target()
	assert.Nil(this.T(), target.Begin(this.type1, nil))
	assert.NotNil(this.T(), target.Add(this.fieldName, this.fieldNote, this.fieldField, []layers.Layer{{Name: this.type1, Type: layers.LayerArray}}, 0))

	target = this.target()
	assert.NotNil(this.T(), target.Add(this.fieldName, this.fieldNote, this.fieldField, []layers.Layer{{Name: this.type1, Type: layers.LayerArray}}, 0))

	target = this.target()
	assert.NotNil(this.T(), target.Add(this.fieldName, this.fieldNote, this.fieldField, nil, 0))

	target = this.target()
	assert.Nil(this.T(), target.Begin(this.type1, nil))
	assert.NotNil(this.T(), target.Add(this.fieldName, this.fieldNote, this.fieldField, nil, 2))
}

func (this *SuiteLayoutType) TestMerge() {
	source1 := this.target()
	assert.Nil(this.T(), source1.Begin(this.type1, nil))
	assert.Nil(this.T(), source1.Add("name1", this.fieldNote, this.fieldField, []layers.Layer{{Name: this.typeA, Type: layers.LayerStruct}}, 1))
	assert.Nil(this.T(), source1.Add("name2", this.fieldNote, this.fieldField, nil, 0))
	assert.Nil(this.T(), source1.Add("name3", this.fieldNote, this.fieldField, []layers.Layer{{Name: this.typeB, Type: layers.LayerStruct}}, 1))
	assert.Nil(this.T(), source1.Add("name4", this.fieldNote, this.fieldField, nil, 0))
	assert.Nil(this.T(), source1.Add("name5", this.fieldNote, this.fieldField, []layers.Layer{{Name: this.typeC, Type: layers.LayerStruct}}, 1))
	assert.Nil(this.T(), source1.Add("name6", this.fieldNote, this.fieldField, nil, 0))
	assert.Nil(this.T(), source1.End())
	source2 := this.target()
	assert.Nil(this.T(), source2.Begin(this.type2, nil))
	assert.Nil(this.T(), source2.Add("name1", this.fieldNote, this.fieldField, nil, 0))
	assert.Nil(this.T(), source2.Add("name2", this.fieldNote, this.fieldField, []layers.Layer{{Name: this.typeA, Type: layers.LayerStruct}}, 1))
	assert.Nil(this.T(), source2.Add("name3", this.fieldNote, this.fieldField, nil, 0))
	assert.Nil(this.T(), source2.Add("name4", this.fieldNote, this.fieldField, []layers.Layer{{Name: this.typeB, Type: layers.LayerStruct}}, 1))
	assert.Nil(this.T(), source2.Add("name5", this.fieldNote, this.fieldField, nil, 0))
	assert.Nil(this.T(), source2.Add("name6", this.fieldNote, this.fieldField, []layers.Layer{{Name: this.typeD, Type: layers.LayerStruct}}, 1))
	assert.Nil(this.T(), source2.End())
	target := this.target()
	assert.Nil(this.T(), target.Merge(source1))
	assert.Nil(this.T(), target.Merge(source2))
	assert.Equal(this.T(), []string{this.type1, this.type2, this.typeA, this.typeB, this.typeC, this.typeD}, target.TypeNames())
	assert.Equal(this.T(), []string{"name2", "name4", "name6", this.typeA, this.typeB, this.typeC}, target.FieldNames(this.type1))
	assert.Equal(this.T(), []string{"name1", "name3", "name5", this.typeA, this.typeB, this.typeD}, target.FieldNames(this.type2))
	assert.Equal(this.T(), []string{"name1", "name2"}, target.FieldNames(this.typeA))
	assert.Equal(this.T(), []string{"name3", "name4"}, target.FieldNames(this.typeB))
	assert.Equal(this.T(), []string{"name5"}, target.FieldNames(this.typeC))
	assert.Equal(this.T(), []string{"name6"}, target.FieldNames(this.typeD))

	failed := this.target()
	assert.Nil(this.T(), failed.Begin(this.type1, nil))
	target = this.target()
	assert.NotNil(this.T(), target.Merge(failed))

	target = this.target()
	assert.Nil(this.T(), target.Begin(this.type1, nil))
	assert.NotNil(this.T(), target.Merge(this.target()))
}

func (this *SuiteLayoutType) TestTypes() {
	target := this.target()
	assert.True(this.T(), target.pushType(this.type1, nil))
	assert.True(this.T(), target.pushField(this.field1.Name, this.field1.Note, this.field1.Field, this.field1.Alter, this.field1.Array))
	assert.True(this.T(), target.pushField(this.field2.Name, this.field2.Note, this.field2.Field, this.field2.Alter, this.field2.Array))
	type_ := target.Types(this.type1)
	assert.NotNil(this.T(), type_)
	assert.Nil(this.T(), type_.Named)
	assert.Equal(this.T(), []*Field{this.field1, this.field2}, type_.Field)

	target = this.target()
	assert.Nil(this.T(), target.Types(this.type2))
}

func (this *SuiteLayoutType) TestTypeNames() {
	target := this.target()
	assert.True(this.T(), target.pushType(this.type1, nil))
	assert.True(this.T(), target.pushType(this.type2, nil))
	assert.Equal(this.T(), []string{this.type1, this.type2}, target.TypeNames())
}

func (this *SuiteLayoutType) TestFieldNames() {
	target := this.target()
	assert.True(this.T(), target.pushType(this.type1, nil))
	assert.True(this.T(), target.pushField(this.field1.Name, this.field1.Note, this.field1.Field, this.field1.Alter, this.field1.Array))
	assert.True(this.T(), target.pushField(this.field2.Name, this.field2.Note, this.field2.Field, this.field2.Alter, this.field2.Array))
	assert.Equal(this.T(), []string{this.field1.Name, this.field2.Name}, target.FieldNames(this.type1))
}

func (this *SuiteLayoutType) TestClosure() {
	target := this.target()
	assert.True(this.T(), target.Closure())
	assert.True(this.T(), target.pushType(this.type1, nil))
	assert.False(this.T(), target.Closure())
}

func (this *SuiteLayoutType) TestPushType() {
	target := this.target()
	assert.True(this.T(), target.pushType(this.type1, nil))
	assert.False(this.T(), target.pushType(this.type1, nil))
}

func (this *SuiteLayoutType) TestPushField() {
	target := this.target()
	assert.True(this.T(), target.pushType(this.type1, nil))
	assert.True(this.T(), target.pushField(this.field1.Name, this.field1.Note, this.field1.Field, this.field1.Alter, this.field1.Array))

	target = this.target()
	assert.False(this.T(), target.pushField(this.field1.Name, this.field1.Note, this.field1.Field, this.field1.Alter, this.field1.Array))

	target = this.target()
	target.level.Push(this.type1)
	assert.False(this.T(), target.pushField(this.field1.Name, this.field1.Note, this.field1.Field, this.field1.Alter, this.field1.Array))
}

func (this *SuiteLayoutType) TestPop() {
	target := this.target()
	assert.True(this.T(), target.pushType(this.type1, nil))
	assert.True(this.T(), target.pushType(this.type2, nil))
	assert.True(this.T(), target.pop(2))
	assert.False(this.T(), target.pop(1))
}
