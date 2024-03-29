package builds

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/v2/sheeter"
	"github.com/yinweli/Sheeter/v2/sheeter/excels"
	"github.com/yinweli/Sheeter/v2/sheeter/nameds"
	"github.com/yinweli/Sheeter/v2/testdata"
)

func TestPoststep(t *testing.T) {
	suite.Run(t, new(SuitePoststep))
}

type SuitePoststep struct {
	suite.Suite
	testdata.Env
	folder      string
	excel       string
	sheet       string
	lineOfTag   int
	lineOfName  int
	lineOfNote  int
	lineOfField int
	lintOfData  int
	tag         string
}

func (this *SuitePoststep) SetupSuite() {
	this.Env = testdata.EnvSetup("test-builds-poststep", "poststep")
	this.folder = "poststep"
	this.excel = "poststep.xlsx"
	this.sheet = "Sheet"
	this.lineOfTag = 1
	this.lineOfName = 2
	this.lineOfNote = 3
	this.lineOfField = 4
	this.lintOfData = 5
	this.tag = "1"
}

func (this *SuitePoststep) TearDownSuite() {
	excels.CloseAll()
	testdata.EnvRestore(this.Env)
}

func (this *SuitePoststep) TestPoststep() {
	config := this.prepareConfig(false)
	context, _ := Initialize(config)
	time.Sleep(testdata.Timeout)
	file, err := Poststep(config, context)
	time.Sleep(testdata.Timeout)
	assert.Len(this.T(), err, 0)
	assert.Len(this.T(), file, 2)

	for _, itor := range file {
		assert.FileExists(this.T(), itor.(string))
	} // for
}

func (this *SuitePoststep) TestGenerateSheeterCs() {
	result := make(chan any, sheeter.MaxExcel)
	poststepData := this.prepareData(this.excel, this.sheet)
	assert.Nil(this.T(), generateSheeterCs(poststepData, result))
	assert.FileExists(this.T(), poststepData.SheeterPathCs())
}

func (this *SuitePoststep) TestGenerateSheeterGo() {
	result := make(chan any, sheeter.MaxExcel)
	poststepData := this.prepareData(this.excel, this.sheet)
	assert.Nil(this.T(), generateSheeterGo(poststepData, result))
	assert.FileExists(this.T(), poststepData.SheeterPathGo())
}

func (this *SuitePoststep) prepareConfig(empty bool) *Config {
	config := &Config{
		Tag:         this.tag,
		LineOfTag:   this.lineOfTag,
		LineOfName:  this.lineOfName,
		LineOfNote:  this.lineOfNote,
		LineOfField: this.lineOfField,
		LineOfData:  this.lintOfData,
	}

	if empty == false {
		config.Source = []string{this.folder}
		config.Merge = []string{"merge$poststep#Sheet"}
	} // if

	return config
}

func (this *SuitePoststep) prepareData(excelName, sheetName string) *PoststepData {
	return &PoststepData{
		Config: this.prepareConfig(true),
		Named:  &nameds.Named{},
		Alone: []*nameds.Named{
			{ExcelName: excelName, SheetName: sheetName},
		},
		Merge: []*nameds.Merge{
			{Name: "merge0", Member: []*nameds.Named{{ExcelName: excelName, SheetName: sheetName}}},
		},
	}
}
