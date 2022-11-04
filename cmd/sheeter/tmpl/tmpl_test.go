package tmpl

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/testdata"
)

func TestTmpl(t *testing.T) {
	suite.Run(t, new(SuiteTmpl))
}

type SuiteTmpl struct {
	suite.Suite
	testdata.TestEnv
}

func (this *SuiteTmpl) SetupSuite() {
	this.Change("test-cmd-tmpl")
}

func (this *SuiteTmpl) TearDownSuite() {
	this.Restore()
}

func (this *SuiteTmpl) TestExecute() {
	cmd := NewCommand()
	assert.Nil(this.T(), cmd.Execute())
}
