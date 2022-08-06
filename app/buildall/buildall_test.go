package buildall

import (
	"testing"

	"github.com/yinweli/Sheeter/internal/build/tasks"
	"github.com/yinweli/Sheeter/testdata"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestConfig(t *testing.T) {
	suite.Run(t, new(SuiteConfig))
}

type SuiteConfig struct {
	suite.Suite
}

func (this *SuiteConfig) target() *config {
	return &config{
		Global: tasks.Global{
			ExcelPath:   "excel",
			Bom:         true,
			LineOfField: 1,
			LineOfLayer: 2,
			LineOfNote:  3,
			LineOfData:  4,
		},
		Elements: []tasks.Element{{
			Excel: "excel.xlsx",
			Sheet: "sheet",
		}},
	}
}

func (this *SuiteConfig) TestReadConfig() {
	config, err := readConfig(testdata.Path(testdata.RealConfig))
	assert.Nil(this.T(), err)
	assert.NotNil(this.T(), config)

	_, err = readConfig(testdata.Path(testdata.Defect1Config))
	assert.NotNil(this.T(), err)

	_, err = readConfig(testdata.Path(testdata.Defect2Config))
	assert.NotNil(this.T(), err)

	_, err = readConfig(testdata.UnknownStr)
	assert.NotNil(this.T(), err)
}

func (this *SuiteConfig) TestCheck() {
	target := this.target()
	assert.Nil(this.T(), target.check())

	target = this.target()
	target.Global.LineOfField = 0
	assert.NotNil(this.T(), target.check())

	target = this.target()
	target.Global.LineOfLayer = 0
	assert.NotNil(this.T(), target.check())

	target = this.target()
	target.Global.LineOfNote = 0
	assert.NotNil(this.T(), target.check())

	target = this.target()
	target.Global.LineOfData = 0
	assert.NotNil(this.T(), target.check())

	target = this.target()
	target.Global.LineOfField = 4
	assert.NotNil(this.T(), target.check())

	target = this.target()
	target.Global.LineOfLayer = 4
	assert.NotNil(this.T(), target.check())

	target = this.target()
	target.Global.LineOfNote = 4
	assert.NotNil(this.T(), target.check())

	target = this.target()
	target.Elements[0].Excel = ""
	assert.NotNil(this.T(), target.check())

	target = this.target()
	target.Elements[0].Sheet = ""
	assert.NotNil(this.T(), target.check())
}
