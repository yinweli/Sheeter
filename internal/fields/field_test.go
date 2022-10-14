package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/testdata"
)

func TestParser(t *testing.T) {
	suite.Run(t, new(SuiteParser))
}

type SuiteParser struct {
	suite.Suite
	workDir string
}

func (this *SuiteParser) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
}

func (this *SuiteParser) TearDownSuite() {
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteParser) TestParser() {
	field, tag, err := Parser("bool#tag")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), (&Bool{}).Type(), field.Type())
	assert.Equal(this.T(), "tag", tag)

	field, tag, err = Parser("bool")
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), (&Bool{}).Type(), field.Type())
	assert.Equal(this.T(), "", tag)

	_, _, err = Parser(testdata.UnknownStr)
	assert.NotNil(this.T(), err)

	_, _, err = Parser("fake#fake")
	assert.NotNil(this.T(), err)

	_, _, err = Parser("fake:fake")
	assert.NotNil(this.T(), err)
}
