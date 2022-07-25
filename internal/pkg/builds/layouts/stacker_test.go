package layouts

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestStacker(t *testing.T) {
	suite.Run(t, new(SuiteStacker))
}

type SuiteStacker struct {
	suite.Suite
	name1  string
	name2  string
	value1 string
	value2 string
}

func (this *SuiteStacker) SetupSuite() {
	this.name1 = "name1"
	this.name2 = "name2"
	this.value1 = "value1"
	this.value2 = "value2"
}

func (this *SuiteStacker) target() *stacker {
	return newStacker()
}

func (this *SuiteStacker) TestPushArray() {
	target := this.target()

	assert.True(this.T(), target.pushArray(this.name1))
	assert.IsType(this.T(), &layoutArray{}, target.datas.Back().Value)
	assert.IsType(this.T(), layoutStruct{}, target.datas.Back().Prev().Value)
	assert.False(this.T(), target.pushArray(this.name2))
}

func (this *SuiteStacker) TestPushStructA() {
	target := this.target()

	assert.True(this.T(), target.pushArray(this.name1))
	assert.True(this.T(), target.pushStructA())
	assert.IsType(this.T(), layoutStruct{}, target.datas.Back().Value)
	assert.IsType(this.T(), &layoutArray{}, target.datas.Back().Prev().Value)
	assert.IsType(this.T(), layoutStruct{}, target.datas.Back().Prev().Prev().Value)
	assert.False(this.T(), target.pushStructA())
}

func (this *SuiteStacker) TestPushStructS() {
	target := this.target()

	assert.True(this.T(), target.pushStructS(this.name1))
	assert.True(this.T(), target.pushStructS(this.name2))
	assert.IsType(this.T(), layoutStruct{}, target.datas.Back().Value)
	assert.IsType(this.T(), layoutStruct{}, target.datas.Back().Prev().Value)
	assert.IsType(this.T(), layoutStruct{}, target.datas.Back().Prev().Prev().Value)
	assert.False(this.T(), target.pushStructA())
}

func (this *SuiteStacker) TestPushValue() {
	target := this.target()

	assert.True(this.T(), target.pushValue(this.name1, this.value1))
	last, ok := target.datas.Back().Value.(layoutStruct)
	assert.True(this.T(), ok)
	assert.Equal(this.T(), this.value1, last[this.name1])
	assert.True(this.T(), target.pushArray(this.name2))
	assert.False(this.T(), target.pushValue(this.name2, this.value1))
}

func (this *SuiteStacker) TestPop() {
	target := this.target()

	assert.True(this.T(), target.pushArray(this.name1))
	assert.True(this.T(), target.pushStructA())
	assert.IsType(this.T(), layoutStruct{}, target.datas.Back().Value)
	target.pop(1, false)
	assert.IsType(this.T(), &layoutArray{}, target.datas.Back().Value)
	target.pop(1, false)
	assert.IsType(this.T(), layoutStruct{}, target.datas.Back().Value)

	assert.True(this.T(), target.pushArray(this.name1))
	assert.True(this.T(), target.pushStructA())
	assert.IsType(this.T(), layoutStruct{}, target.datas.Back().Value)
	target.pop(1, true)
	assert.IsType(this.T(), layoutStruct{}, target.datas.Back().Value)
}

func (this *SuiteStacker) TestClosure() {
	target := this.target()

	assert.True(this.T(), target.pushArray(this.name1))
	assert.True(this.T(), target.pushStructA())
	assert.False(this.T(), target.closure())
	target.pop(1, true)
	assert.True(this.T(), target.closure())
}

func (this *SuiteStacker) TestResult() {
	target := this.target()

	assert.IsType(this.T(), layoutStruct{}, target.result())
}

func (this *SuiteStacker) TestNewStacker() {
	assert.NotNil(this.T(), newStacker())
}
