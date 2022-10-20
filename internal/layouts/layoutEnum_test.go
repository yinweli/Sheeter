package layouts

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/testdata"
)

func TestLayoutEnum(t *testing.T) {
	suite.Run(t, new(SuiteLayoutEnum))
}

type SuiteLayoutEnum struct {
	suite.Suite
	workDir string
}

func (this *SuiteLayoutEnum) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
}

func (this *SuiteLayoutEnum) TearDownSuite() {
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteLayoutEnum) target() *LayoutEnum {
	return NewLayoutEnum()
}

func (this *SuiteLayoutEnum) TestNewLayoutEnum() {
	assert.NotNil(this.T(), NewLayoutEnum())
}

func (this *SuiteLayoutEnum) TestAdd() {
	target := this.target()
	assert.Nil(this.T(), target.Add([]string{"name0", "0", "commit0"}))
	assert.Nil(this.T(), target.Add([]string{"name1", "1", "commit1"}))
	assert.Nil(this.T(), target.Add([]string{"name2", "2"}))
	assert.Nil(this.T(), target.Add([]string{"name3", "3", "commit3", testdata.UnknownStr}))
	assert.NotNil(this.T(), target.Add([]string{"name"}))
	assert.NotNil(this.T(), target.Add([]string{"name", "?"}))
	assert.NotNil(this.T(), target.Add([]string{"?name", "0"}))
	assert.NotNil(this.T(), target.Add([]string{"name", "0"}))
	assert.NotNil(this.T(), target.Add([]string{"name0", "0"}))
}

func (this *SuiteLayoutEnum) TestEnums() {
	target := this.target()
	assert.Nil(this.T(), target.Add([]string{"name0", "0", "commit0"}))
	assert.Nil(this.T(), target.Add([]string{"name1", "1", "commit1"}))
	assert.Equal(this.T(), []*Enum{
		{Name: "name0", Index: 0, Comment: "commit0"},
		{Name: "name1", Index: 1, Comment: "commit1"},
	}, target.Enums())
}
