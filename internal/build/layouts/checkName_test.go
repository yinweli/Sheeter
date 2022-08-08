package layouts

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestCheckName(t *testing.T) {
	suite.Run(t, new(SuiteCheckName))
}

type SuiteCheckName struct {
	suite.Suite
}

func (this *SuiteCheckName) target() *checkName {
	return &checkName{}
}

func (this *SuiteCheckName) TestCheck() {
	target := this.target()

	assert.True(this.T(), target.check("name1"))
	assert.True(this.T(), target.check("name2"))
	assert.True(this.T(), target.check("names"))
	assert.False(this.T(), target.check("names"))
}
