package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/testdata"
)

func TestDuplicate(t *testing.T) {
	suite.Run(t, new(SuiteDuplicate))
}

type SuiteDuplicate struct {
	suite.Suite
	workDir string
}

func (this *SuiteDuplicate) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
}

func (this *SuiteDuplicate) TearDownSuite() {
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteDuplicate) target() *Duplicate {
	return NewDuplicate()
}

func (this *SuiteDuplicate) TestNewDuplicate() {
	assert.NotNil(this.T(), NewDuplicate())
}

func (this *SuiteDuplicate) TestCheck() {
	target := this.target()
	assert.True(this.T(), target.Check("1"))
	assert.True(this.T(), target.Check("1", "2"))
	assert.True(this.T(), target.Check("1", "2", "3"))
	assert.False(this.T(), target.Check("1", "2", "3"))
}
