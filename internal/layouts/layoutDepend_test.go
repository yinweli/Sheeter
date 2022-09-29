package layouts

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/internal/layers"
	"github.com/yinweli/Sheeter/testdata"
)

func TestLayoutDepend(t *testing.T) {
	suite.Run(t, new(SuiteLayoutDepend))
}

type SuiteLayoutDepend struct {
	suite.Suite
	workDir string
}

func (this *SuiteLayoutDepend) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
}

func (this *SuiteLayoutDepend) TearDownSuite() {
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteLayoutDepend) target() *LayoutDepend {
	return NewLayoutDepend()
}

func (this *SuiteLayoutDepend) layer(input string) []layers.Layer {
	layer, _, _ := layers.Parser(input)
	return layer
}

func (this *SuiteLayoutDepend) TestNewLayoutDepend() {
	assert.NotNil(this.T(), NewLayoutDepend())
}

func (this *SuiteLayoutDepend) TestBegin() {
	target := this.target()
	assert.Nil(this.T(), target.Begin("type1"))
	assert.NotNil(this.T(), target.Begin("type2"))
}

func (this *SuiteLayoutDepend) TestEnd() {
	target := this.target()
	assert.Nil(this.T(), target.Begin("type"))
	assert.Nil(this.T(), target.End())
	assert.NotNil(this.T(), target.End())

	target = this.target()
	assert.Nil(this.T(), target.Begin("type"))
	assert.Nil(this.T(), target.Add(this.layer("{[]layer1 {layer2"), 0))
	assert.NotNil(this.T(), target.End())
}

func (this *SuiteLayoutDepend) TestAdd() {
	target := this.target()
	assert.Nil(this.T(), target.Begin("type"))
	assert.Nil(this.T(), target.Add(this.layer("{[]layer1 {layer2"), 2))

	target = this.target()
	assert.NotNil(this.T(), target.Add(this.layer("{[]layer1 {layer2"), 0))
}

func (this *SuiteLayoutDepend) TestMerge() {
	source1 := this.target()
	assert.Nil(this.T(), source1.Begin("type1"))
	assert.Nil(this.T(), source1.Add(this.layer("{typeA {typeB {typeC"), 3))
	assert.Nil(this.T(), source1.End())
	source2 := this.target()
	assert.Nil(this.T(), source2.Begin("type1"))
	assert.Nil(this.T(), source2.Add(this.layer("{typeA {typeB {typeD"), 3))
	assert.Nil(this.T(), source2.End())
	source3 := this.target()
	assert.Nil(this.T(), source2.Begin("type2"))
	assert.Nil(this.T(), source2.Add(this.layer("{typeC {typeD"), 2))
	assert.Nil(this.T(), source2.End())
	target := this.target()
	assert.Nil(this.T(), target.Merge(source1))
	assert.Nil(this.T(), target.Merge(source2))
	assert.Nil(this.T(), target.Merge(source3))
	assert.Equal(this.T(), []string{"typeA"}, target.Depends("type1"))
	assert.Equal(this.T(), []string{"typeC"}, target.Depends("type2"))
	assert.Equal(this.T(), []string{"typeB"}, target.Depends("typeA"))
	assert.Equal(this.T(), []string{"typeC", "typeD"}, target.Depends("typeB"))
	assert.Equal(this.T(), []string{"typeD"}, target.Depends("typeC"))
	assert.Nil(this.T(), target.Depends("typeD"))
}

func (this *SuiteLayoutDepend) TestDepends() {
	target := this.target()
	assert.Nil(this.T(), target.Begin("type"))
	assert.Nil(this.T(), target.Add(this.layer("{typeA {typeB {typeC"), 0))
	assert.Nil(this.T(), target.Add(this.layer("{typeD {typeE"), 0))
	assert.Nil(this.T(), target.Add(this.layer("{typeF"), 6))
	assert.Nil(this.T(), target.End())
	assert.Equal(this.T(), []string{"typeA"}, target.Depends("type"))
	assert.Equal(this.T(), []string{"typeB"}, target.Depends("typeA"))
	assert.Equal(this.T(), []string{"typeC"}, target.Depends("typeB"))
	assert.Equal(this.T(), []string{"typeD"}, target.Depends("typeC"))
	assert.Equal(this.T(), []string{"typeE"}, target.Depends("typeD"))
	assert.Equal(this.T(), []string{"typeF"}, target.Depends("typeE"))
	assert.Nil(this.T(), target.Depends("typeF"))
}

func (this *SuiteLayoutDepend) TestPushField() {
	target := this.target()
	assert.Nil(this.T(), target.Begin("type1"))
	assert.True(this.T(), target.push("type2"))
	assert.True(this.T(), target.push("type3"))

	target = this.target()
	assert.False(this.T(), target.push("type1"))
}

func (this *SuiteLayoutDepend) TestPop() {
	target := this.target()
	assert.Nil(this.T(), target.Begin("type1"))
	assert.True(this.T(), target.push("type2"))
	assert.True(this.T(), target.push("type3"))
	assert.True(this.T(), target.pop(3))
	assert.False(this.T(), target.pop(1))
}
