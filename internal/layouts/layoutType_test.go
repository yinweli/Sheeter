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
	workDir string
}

func (this *SuiteLayoutType) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
}

func (this *SuiteLayoutType) TearDownSuite() {
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteLayoutType) target() *LayoutType {
	return NewLayoutType()
}

func (this *SuiteLayoutType) layer(input string) []layers.Layer {
	layer, _, _ := layers.Parser(input)
	return layer
}

func (this *SuiteLayoutType) TestNewLayoutType() {
	assert.NotNil(this.T(), NewLayoutType())
}

func (this *SuiteLayoutType) TestBegin() {
	target := this.target()
	assert.Nil(this.T(), target.Begin("type1", "", ""))
	assert.NotNil(this.T(), target.Begin("type2", "", ""))
}

func (this *SuiteLayoutType) TestEnd() {
	target := this.target()
	assert.Nil(this.T(), target.Begin("type", "", ""))
	assert.Nil(this.T(), target.End())
	assert.NotNil(this.T(), target.End())

	target = this.target()
	assert.Nil(this.T(), target.Begin("type", "", ""))
	assert.Nil(this.T(), target.Add("name", "", &fields.Int{}, this.layer("{[]layer1 {layer2"), 0))
	assert.NotNil(this.T(), target.End())
}

func (this *SuiteLayoutType) TestAdd() {
	target := this.target()
	assert.Nil(this.T(), target.Begin("type1", "", ""))
	assert.Nil(this.T(), target.Add("name", "", &fields.Int{}, this.layer("{[]layer1 {layer2"), 2))

	target = this.target()
	assert.Nil(this.T(), target.Begin("type1", "", ""))
	assert.NotNil(this.T(), target.Add("name", "", &fields.Int{}, this.layer("{[]type1"), 0))

	target = this.target()
	assert.NotNil(this.T(), target.Add("name", "", &fields.Int{}, this.layer("{[]type1"), 0))

	target = this.target()
	assert.NotNil(this.T(), target.Add("name", "", &fields.Int{}, nil, 0))

	target = this.target()
	assert.Nil(this.T(), target.Begin("type1", "", ""))
	assert.NotNil(this.T(), target.Add("name", "", &fields.Int{}, nil, 2))
}

func (this *SuiteLayoutType) TestMerge() {
	source1 := this.target()
	assert.Nil(this.T(), source1.Begin("type1", "", ""))
	assert.Nil(this.T(), source1.Add("name1", "", &fields.Int{}, this.layer("{typeA"), 1))
	assert.Nil(this.T(), source1.Add("name2", "", &fields.Int{}, nil, 0))
	assert.Nil(this.T(), source1.Add("name3", "", &fields.Int{}, this.layer("{typeB"), 1))
	assert.Nil(this.T(), source1.Add("name4", "", &fields.Int{}, nil, 0))
	assert.Nil(this.T(), source1.Add("name5", "", &fields.Int{}, this.layer("{typeC"), 1))
	assert.Nil(this.T(), source1.Add("name6", "", &fields.Int{}, nil, 0))
	assert.Nil(this.T(), source1.End())
	source2 := this.target()
	assert.Nil(this.T(), source2.Begin("type2", "", ""))
	assert.Nil(this.T(), source2.Add("name1", "", &fields.Int{}, nil, 0))
	assert.Nil(this.T(), source2.Add("name2", "", &fields.Int{}, this.layer("{typeA"), 1))
	assert.Nil(this.T(), source2.Add("name3", "", &fields.Int{}, nil, 0))
	assert.Nil(this.T(), source2.Add("name4", "", &fields.Int{}, this.layer("{typeB"), 1))
	assert.Nil(this.T(), source2.Add("name5", "", &fields.Int{}, nil, 0))
	assert.Nil(this.T(), source2.Add("name6", "", &fields.Int{}, this.layer("{typeD"), 1))
	assert.Nil(this.T(), source2.End())
	target := this.target()
	assert.Nil(this.T(), target.Merge(source1))
	assert.Nil(this.T(), target.Merge(source2))
	assert.Equal(this.T(), []string{"type1", "type2", "typeA", "typeB", "typeC", "typeD"}, target.TypeNames())
	assert.Equal(this.T(), []string{"name2", "name4", "name6", "typeA", "typeB", "typeC"}, target.FieldNames("type1"))
	assert.Equal(this.T(), []string{"name1", "name3", "name5", "typeA", "typeB", "typeD"}, target.FieldNames("type2"))
	assert.Equal(this.T(), []string{"name1", "name2"}, target.FieldNames("typeA"))
	assert.Equal(this.T(), []string{"name3", "name4"}, target.FieldNames("typeB"))
	assert.Equal(this.T(), []string{"name5"}, target.FieldNames("typeC"))
	assert.Equal(this.T(), []string{"name6"}, target.FieldNames("typeD"))

	failed := this.target()
	assert.Nil(this.T(), failed.Begin("type", "", ""))
	target = this.target()
	assert.NotNil(this.T(), target.Merge(failed))

	target = this.target()
	assert.Nil(this.T(), target.Begin("type", "", ""))
	assert.NotNil(this.T(), target.Merge(this.target()))
}

func (this *SuiteLayoutType) TestTypes() {
	field1 := &Field{Name: "name1", Note: "note1", Field: &fields.Pkey{}, Alter: "alter1", Array: true}
	field2 := &Field{Name: "name2", Note: "note2", Field: &fields.Text{}, Alter: "alter2", Array: false}
	field3 := &Field{Name: "name3", Note: "note3", Field: &fields.TextArray{}, Alter: "alter3", Array: true}

	target := this.target()
	assert.True(this.T(), target.pushType("type", "excel", "sheet", true))
	assert.True(this.T(), target.pushField(field1.Name, field1.Note, field1.Field, field1.Alter, field1.Array))
	assert.True(this.T(), target.pushField(field2.Name, field2.Note, field2.Field, field2.Alter, field2.Array))
	assert.True(this.T(), target.pushField(field3.Name, field3.Note, field3.Field, field3.Alter, field3.Array))
	type_ := target.Types("type")
	assert.NotNil(this.T(), type_)
	assert.Equal(this.T(), "excel", type_.Excel)
	assert.Equal(this.T(), "sheet", type_.Sheet)
	assert.True(this.T(), type_.Reader)
	assert.Equal(this.T(), []*Field{field1, field2, field3}, type_.Fields)

	target = this.target()
	assert.Nil(this.T(), target.Types("type"))
}

func (this *SuiteLayoutType) TestTypeNames() {
	target := this.target()
	assert.True(this.T(), target.pushType("type1", "", "", false))
	assert.True(this.T(), target.pushType("type2", "", "", false))
	assert.Equal(this.T(), []string{"type1", "type2"}, target.TypeNames())
}

func (this *SuiteLayoutType) TestFieldNames() {
	target := this.target()
	assert.True(this.T(), target.pushType("type", "", "", false))
	assert.True(this.T(), target.pushField("name1", "", &fields.Int{}, "", false))
	assert.True(this.T(), target.pushField("name2", "", &fields.Int{}, "", false))
	assert.Equal(this.T(), []string{"name1", "name2"}, target.FieldNames("type"))
}

func (this *SuiteLayoutType) TestClosure() {
	target := this.target()
	assert.True(this.T(), target.Closure())
	assert.True(this.T(), target.pushType("type", "", "", false))
	assert.False(this.T(), target.Closure())
}

func (this *SuiteLayoutType) TestPushType() {
	target := this.target()
	assert.True(this.T(), target.pushType("type", "", "", false))
	assert.False(this.T(), target.pushType("type", "", "", false))
}

func (this *SuiteLayoutType) TestPushField() {
	target := this.target()
	assert.True(this.T(), target.pushType("type", "", "", false))
	assert.True(this.T(), target.pushField("name1", "", &fields.Int{}, "", false))
	assert.True(this.T(), target.pushField("name2", "", &fields.Int{}, "", false))
	assert.True(this.T(), target.pushField("name3", "", &fields.Int{}, "", false))

	target = this.target()
	assert.False(this.T(), target.pushField("name", "", &fields.Int{}, "", false))
}

func (this *SuiteLayoutType) TestPop() {
	target := this.target()
	assert.True(this.T(), target.pushType("type1", "", "", false))
	assert.True(this.T(), target.pushType("type2", "", "", false))
	assert.True(this.T(), target.pop(2))
	assert.False(this.T(), target.pop(1))
}
