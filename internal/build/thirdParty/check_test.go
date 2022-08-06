package thirdParty

import (
	"fmt"
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

func (this *SuiteShell) println(i ...interface{}) {
	fmt.Println(i...)
}

func (this *SuiteShell) TestCheck() {
	assert.True(this.T(), Check(this.println))
	assert.True(this.T(), check(this.println, "go"))
	assert.False(this.T(), check(this.println, "unknown"))
}
