package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/yinweli/Sheeter/testdata"
)

func TestConfig(t *testing.T) {
	suite.Run(t, new(SuiteConfig))
}

type SuiteConfig struct {
	suite.Suite
}

func (this *SuiteConfig) target() *Config {
	return &Config{
		Global: Global{
			ExcelPath:   "excel",
			Bom:         true,
			LineOfField: 1,
			LineOfNote:  2,
			LineOfData:  3,
		},
		Elements: []Element{{
			Excel: "excel.xlsx",
			Sheet: "sheet",
		}},
	}
}

func (this *SuiteConfig) TestCheck() {
	target := this.target()
	assert.Nil(this.T(), target.Check())

	target = this.target()
	target.Global.LineOfField = 0
	assert.NotNil(this.T(), target.Check())

	target = this.target()
	target.Global.LineOfNote = 0
	assert.NotNil(this.T(), target.Check())

	target = this.target()
	target.Global.LineOfData = 0
	assert.NotNil(this.T(), target.Check())

	target = this.target()
	target.Global.LineOfField = 3
	assert.NotNil(this.T(), target.Check())

	target = this.target()
	target.Global.LineOfNote = 3
	assert.NotNil(this.T(), target.Check())

	target = this.target()
	target.Elements[0].Excel = ""
	assert.NotNil(this.T(), target.Check())

	target = this.target()
	target.Elements[0].Sheet = ""
	assert.NotNil(this.T(), target.Check())
}

func (this *SuiteConfig) TestReadConfig() {
	config, err := ReadConfig(testdata.Path(testdata.RealConfig))
	assert.Nil(this.T(), err)
	assert.NotNil(this.T(), config)

	_, err = ReadConfig(testdata.Path(testdata.Defect1Config))
	assert.NotNil(this.T(), err)

	_, err = ReadConfig(testdata.Path(testdata.Defect2Config))
	assert.NotNil(this.T(), err)

	_, err = ReadConfig("?????")
	assert.NotNil(this.T(), err)
}
