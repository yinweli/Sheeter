package layouts

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/testdata"
)

func TestStructor(t *testing.T) {
	suite.Run(t, new(SuiteStructor))
}

type SuiteStructor struct {
	suite.Suite
	workDir string
	name1   string
	name2   string
	value1  string
	value2  string
}

func (this *SuiteStructor) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
	this.name1 = "name1"
	this.name2 = "name2"
	this.value1 = "value1"
	this.value2 = "value2"
}

func (this *SuiteStructor) TearDownSuite() {
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteStructor) target() *structor {
	return newStructor()
}

func (this *SuiteStructor) TestNewStructor() {
	assert.NotNil(this.T(), newStructor())
}

func (this *SuiteStructor) TestPushArray() {
	target := this.target()

	assert.True(this.T(), target.pushArray(this.name1))
	assert.IsType(this.T(), &array_{}, target.datas.Back().Value)
	assert.IsType(this.T(), struct_{}, target.datas.Back().Prev().Value)
	assert.False(this.T(), target.pushArray(this.name2))
}

func (this *SuiteStructor) TestPushStructA() {
	target := this.target()

	assert.True(this.T(), target.pushArray(this.name1))
	assert.True(this.T(), target.pushStructA())
	assert.IsType(this.T(), struct_{}, target.datas.Back().Value)
	assert.IsType(this.T(), &array_{}, target.datas.Back().Prev().Value)
	assert.IsType(this.T(), struct_{}, target.datas.Back().Prev().Prev().Value)
	assert.False(this.T(), target.pushStructA())
}

func (this *SuiteStructor) TestPushStructS() {
	target := this.target()

	assert.True(this.T(), target.pushStructS(this.name1))
	assert.True(this.T(), target.pushStructS(this.name2))
	assert.IsType(this.T(), struct_{}, target.datas.Back().Value)
	assert.IsType(this.T(), struct_{}, target.datas.Back().Prev().Value)
	assert.IsType(this.T(), struct_{}, target.datas.Back().Prev().Prev().Value)
	assert.False(this.T(), target.pushStructA())
}

func (this *SuiteStructor) TestPushValue() {
	target := this.target()

	assert.True(this.T(), target.pushValue(this.name1, this.value1))
	last, ok := target.datas.Back().Value.(struct_)
	assert.True(this.T(), ok)
	assert.Equal(this.T(), this.value1, last[this.name1])
	assert.True(this.T(), target.pushArray(this.name2))
	assert.False(this.T(), target.pushValue(this.name2, this.value1))
}

func (this *SuiteStructor) TestPop() {
	target := this.target()

	assert.True(this.T(), target.pushArray(this.name1))
	assert.True(this.T(), target.pushStructA())
	assert.IsType(this.T(), struct_{}, target.datas.Back().Value)
	target.pop(1, false)
	assert.IsType(this.T(), &array_{}, target.datas.Back().Value)
	target.pop(1, false)
	assert.IsType(this.T(), struct_{}, target.datas.Back().Value)

	assert.True(this.T(), target.pushArray(this.name1))
	assert.True(this.T(), target.pushStructA())
	assert.IsType(this.T(), struct_{}, target.datas.Back().Value)
	target.pop(1, true)
	assert.IsType(this.T(), struct_{}, target.datas.Back().Value)
}

func (this *SuiteStructor) TestClosure() {
	target := this.target()

	assert.True(this.T(), target.pushArray(this.name1))
	assert.True(this.T(), target.pushStructA())
	assert.False(this.T(), target.closure())
	target.pop(1, true)
	assert.True(this.T(), target.closure())
}

func (this *SuiteStructor) TestResult() {
	target := this.target()

	assert.IsType(this.T(), struct_{}, target.result())
}
