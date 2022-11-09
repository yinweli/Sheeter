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
	testdata.TestEnv
}

func (this *SuiteStructor) SetupSuite() {
	this.Change("test-structor")
}

func (this *SuiteStructor) TearDownSuite() {
	this.Restore()
}

func (this *SuiteStructor) target() *structor {
	return newStructor()
}

func (this *SuiteStructor) TestNewStructor() {
	assert.NotNil(this.T(), newStructor())
}

func (this *SuiteStructor) TestPushArray() {
	target := this.target()
	assert.True(this.T(), target.pushArray("array1"))
	assert.IsType(this.T(), &array_{}, target.back().Value)
	assert.IsType(this.T(), struct_{}, target.back().Prev().Value)
	assert.False(this.T(), target.pushArray("array2"))
}

func (this *SuiteStructor) TestPushStructA() {
	target := this.target()
	assert.True(this.T(), target.pushArray("array"))
	assert.True(this.T(), target.pushStructA())
	assert.IsType(this.T(), struct_{}, target.back().Value)
	assert.IsType(this.T(), &array_{}, target.back().Prev().Value)
	assert.IsType(this.T(), struct_{}, target.back().Prev().Prev().Value)
	assert.False(this.T(), target.pushStructA())
}

func (this *SuiteStructor) TestPushStructS() {
	target := this.target()
	assert.True(this.T(), target.pushStructS("struct1"))
	assert.True(this.T(), target.pushStructS("struct2"))
	assert.IsType(this.T(), struct_{}, target.back().Value)
	assert.IsType(this.T(), struct_{}, target.back().Prev().Value)
	assert.IsType(this.T(), struct_{}, target.back().Prev().Prev().Value)
	assert.False(this.T(), target.pushStructA())
}

func (this *SuiteStructor) TestPushValue() {
	field := "field"
	value := "value"

	target := this.target()
	assert.True(this.T(), target.pushValue(field, value))
	last, ok := target.back().Value.(struct_)
	assert.True(this.T(), ok)
	assert.Equal(this.T(), value, last[field])
	assert.True(this.T(), target.pushArray("array"))
	assert.False(this.T(), target.pushValue(field, value))
}

func (this *SuiteStructor) TestPop() {
	target := this.target()
	assert.True(this.T(), target.pushArray("array"))
	assert.True(this.T(), target.pushStructA())
	assert.IsType(this.T(), struct_{}, target.back().Value)
	target.pop(1, false)
	assert.IsType(this.T(), &array_{}, target.back().Value)
	target.pop(1, false)
	assert.IsType(this.T(), struct_{}, target.back().Value)

	target = this.target()
	assert.True(this.T(), target.pushArray("array"))
	assert.True(this.T(), target.pushStructA())
	assert.IsType(this.T(), struct_{}, target.back().Value)
	target.pop(1, true)
	assert.IsType(this.T(), struct_{}, target.back().Value)
}

func (this *SuiteStructor) TestBack() {
	target := this.target()
	assert.True(this.T(), target.pushArray("array"))
	assert.IsType(this.T(), &array_{}, target.back().Value)
}

func (this *SuiteStructor) TestClosure() {
	target := this.target()
	assert.True(this.T(), target.pushArray("array"))
	assert.True(this.T(), target.pushStructA())
	assert.False(this.T(), target.closure())
	target.pop(1, true)
	assert.True(this.T(), target.closure())
}

func (this *SuiteStructor) TestResult() {
	target := this.target()
	assert.IsType(this.T(), struct_{}, target.result())
}
