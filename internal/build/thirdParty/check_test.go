package thirdParty

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestShell(t *testing.T) {
	suite.Run(t, new(SuiteShell))
}

type SuiteShell struct {
	suite.Suite
}

func (this *SuiteShell) TestCheck() {
	assert.Nil(this.T(), Check())
}
