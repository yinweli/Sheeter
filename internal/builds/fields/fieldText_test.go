package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestText(t *testing.T) {
	suite.Run(t, new(SuiteText))
}

type SuiteText struct {
	suite.Suite
}

func (this *SuiteText) target() *Text {
	return &Text{}
}

func (this *SuiteText) TestType() {
	assert.Equal(this.T(), "text", this.target().Type())
}

func (this *SuiteText) TestIsShow() {
	assert.Equal(this.T(), true, this.target().IsShow())
}

func (this *SuiteText) TestIsPkey() {
	assert.Equal(this.T(), false, this.target().IsPkey())
}

func (this *SuiteText) TestToJsonDefault() {
	assert.Equal(this.T(), "", this.target().ToJsonDefault())
}

func (this *SuiteText) TestToJsonValue() {
	target := this.target()

	result, err := target.ToJsonValue("ball,book,pack")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), "ball,book,pack", result)
}
