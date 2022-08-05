package build

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestLayoutStacker(t *testing.T) {
	suite.Run(t, new(SuiteLayoutStacker))
}

type SuiteLayoutStacker struct {
	suite.Suite
	name1  string
	name2  string
	value1 string
	value2 string
}

func (this *SuiteLayoutStacker) SetupSuite() {
	this.name1 = "name1"
	this.name2 = "name2"
	this.value1 = "value1"
	this.value2 = "value2"
}

func (this *SuiteLayoutStacker) target() *layoutStacker {
	return NewLayoutStacker()
}

func (this *SuiteLayoutStacker) TestPushArray() {
	target := this.target()

	assert.True(this.T(), target.PushArray(this.name1))
	assert.IsType(this.T(), layoutArray{}, target.datas.Back().Value)
	assert.IsType(this.T(), layoutStruct{}, target.datas.Back().Prev().Value)
	assert.False(this.T(), target.PushArray(this.name2))
}

func (this *SuiteLayoutStacker) TestPushStructA() {
	target := this.target()

	assert.True(this.T(), target.PushArray(this.name1))
	assert.True(this.T(), target.PushStructA())
	assert.IsType(this.T(), layoutStruct{}, target.datas.Back().Value)
	assert.IsType(this.T(), layoutArray{}, target.datas.Back().Prev().Value)
	assert.IsType(this.T(), layoutStruct{}, target.datas.Back().Prev().Prev().Value)
	assert.False(this.T(), target.PushStructA())
}

func (this *SuiteLayoutStacker) TestPushStructS() {
	target := this.target()

	assert.True(this.T(), target.PushStructS(this.name1))
	assert.True(this.T(), target.PushStructS(this.name2))
	assert.IsType(this.T(), layoutStruct{}, target.datas.Back().Value)
	assert.IsType(this.T(), layoutStruct{}, target.datas.Back().Prev().Value)
	assert.IsType(this.T(), layoutStruct{}, target.datas.Back().Prev().Prev().Value)
	assert.False(this.T(), target.PushStructA())
}

func (this *SuiteLayoutStacker) TestPushValue() {
	target := this.target()

	assert.True(this.T(), target.PushValue(this.name1, this.value1))
	last, ok := target.datas.Back().Value.(layoutStruct)
	assert.True(this.T(), ok)
	assert.Equal(this.T(), this.value1, last[this.name1])
	assert.True(this.T(), target.PushArray(this.name2))
	assert.False(this.T(), target.PushValue(this.name2, this.value1))
}

func (this *SuiteLayoutStacker) TestPop() {
	target := this.target()

	assert.True(this.T(), target.PushArray(this.name1))
	assert.True(this.T(), target.PushStructA())
	assert.IsType(this.T(), layoutStruct{}, target.datas.Back().Value)
	target.Pop(1, false)
	assert.IsType(this.T(), layoutArray{}, target.datas.Back().Value)
	target.Pop(1, false)
	assert.IsType(this.T(), layoutStruct{}, target.datas.Back().Value)

	assert.True(this.T(), target.PushArray(this.name1))
	assert.True(this.T(), target.PushStructA())
	assert.IsType(this.T(), layoutStruct{}, target.datas.Back().Value)
	target.Pop(1, true)
	assert.IsType(this.T(), layoutStruct{}, target.datas.Back().Value)
}

func (this *SuiteLayoutStacker) TestClosure() {
	target := this.target()

	assert.True(this.T(), target.PushArray(this.name1))
	assert.True(this.T(), target.PushStructA())
	assert.False(this.T(), target.Closure())
	target.Pop(1, true)
	assert.True(this.T(), target.Closure())
}

func (this *SuiteLayoutStacker) TestResult() {
	target := this.target()

	assert.IsType(this.T(), layoutStruct{}, target.Result())
}

func (this *SuiteLayoutStacker) TestNewLayoutStacker() {
	assert.NotNil(this.T(), NewLayoutStacker())
}
