package build

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestFieldTextArray(t *testing.T) {
	suite.Run(t, new(SuiteFieldTextArray))
}

type SuiteFieldTextArray struct {
	suite.Suite
}

func (this *SuiteFieldTextArray) target() *FieldTextArray {
	return &FieldTextArray{}
}

func (this *SuiteFieldTextArray) TestType() {
	assert.Equal(this.T(), "textArray", this.target().Type())
}

func (this *SuiteFieldTextArray) TestIsShow() {
	assert.Equal(this.T(), true, this.target().IsShow())
}

func (this *SuiteFieldTextArray) TestIsPkey() {
	assert.Equal(this.T(), false, this.target().IsPkey())
}

func (this *SuiteFieldTextArray) TestToJsonDefault() {
	assert.Equal(this.T(), []string{}, this.target().ToJsonDefault())
}

func (this *SuiteFieldTextArray) TestToJsonValue() {
	target := this.target()

	result, err := target.ToJsonValue("ball,book,pack")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), []string{"ball", "book", "pack"}, result)
}

func (this *SuiteFieldTextArray) TestToLuaValue() {
	target := this.target()

	result, err := target.ToLuaValue("ball,book,pack")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), "{\"ball\",\"book\",\"pack\"}", result)
}
