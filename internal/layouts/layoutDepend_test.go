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
	assert.Nil(this.T(), target.Begin("name1"))
	assert.NotNil(this.T(), target.Begin("name2"))
}

func (this *SuiteLayoutDepend) TestEnd() {
	target := this.target()
	assert.Nil(this.T(), target.End())
}

func (this *SuiteLayoutDepend) TestAdd() {
	target := this.target()
	assert.Nil(this.T(), target.Begin("type1"))
	assert.Nil(this.T(), target.Add(this.layer("{[]layer1 {layer2")))

	target = this.target()
	assert.NotNil(this.T(), target.Add(this.layer("{[]layer1 {layer2")))
}

func (this *SuiteLayoutDepend) TestMerge() {
	source1 := this.target()
	assert.Nil(this.T(), source1.Begin("type1"))
	assert.Nil(this.T(), source1.Add(this.layer("{typeA {typeB {typeC")))
	assert.Nil(this.T(), source1.End())
	source2 := this.target()
	assert.Nil(this.T(), source2.Begin("type1"))
	assert.Nil(this.T(), source2.Add(this.layer("{typeA {typeB {typeD")))
	assert.Nil(this.T(), source2.End())
	source3 := this.target()
	assert.Nil(this.T(), source2.Begin("type2"))
	assert.Nil(this.T(), source2.Add(this.layer("{typeC {typeD")))
	assert.Nil(this.T(), source2.End())
	target := this.target()
	assert.Nil(this.T(), target.Merge(source1))
	assert.Nil(this.T(), target.Merge(source2))
	assert.Nil(this.T(), target.Merge(source3))
	assert.Equal(this.T(), []string{"typeA", "typeB", "typeC", "typeD"}, target.Depends("type1"))
	assert.Equal(this.T(), []string{"typeC", "typeD"}, target.Depends("type2"))
}

func (this *SuiteLayoutDepend) TestDepends() {
	target := this.target()
	assert.Nil(this.T(), target.Begin("type"))
	assert.Nil(this.T(), target.Add(this.layer("{typeA {typeB {typeC")))
	assert.Nil(this.T(), target.Add(this.layer("{typeD {typeE")))
	assert.Nil(this.T(), target.Add(this.layer("{typeF")))
	assert.Nil(this.T(), target.End())
	assert.Equal(this.T(), []string{"typeA", "typeB", "typeC", "typeD", "typeE", "typeF"}, target.Depends("type"))
}
