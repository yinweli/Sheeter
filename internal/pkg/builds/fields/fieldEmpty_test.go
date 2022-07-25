package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestEmpty(t *testing.T) {
	suite.Run(t, new(SuiteEmpty))
}

type SuiteEmpty struct {
	suite.Suite
}

func (this *SuiteEmpty) target() *Empty {
	return &Empty{}
}

func (this *SuiteEmpty) TestType() {
	assert.Equal(this.T(), "empty", this.target().Type())
}

func (this *SuiteEmpty) TestIsShow() {
	assert.Equal(this.T(), false, this.target().IsShow())
}

func (this *SuiteEmpty) TestIsPkey() {
	assert.Equal(this.T(), false, this.target().IsPkey())
}

func (this *SuiteEmpty) TestToJsonValue() {
	target := this.target()

	result, err := target.ToJsonValue("", true)
	assert.Nil(this.T(), err)
	assert.Nil(this.T(), result)

	result, err = target.ToJsonValue("test", false)
	assert.Nil(this.T(), err)
	assert.Nil(this.T(), result)
}
