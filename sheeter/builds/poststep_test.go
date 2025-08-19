package builds

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/v3/sheeter/excels"
	"github.com/yinweli/Sheeter/v3/sheeter/nameds"
	"github.com/yinweli/Sheeter/v3/testdata"
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
	initializeData, _ := Initialize(config)
	result, err := Poststep(config, initializeData)
	assert.Len(this.T(), err, 0)
	assert.Len(this.T(), result, 4)

	for _, itor := range result {
		assert.FileExists(this.T(), itor.(string))
	} // for
}

func (this *SuitePoststep) TestGenerateSheeterCs() {
	material := this.prepareData(this.excel, this.sheet)
	result := generateSheeterCs(material)
	assert.Nil(this.T(), result.Error)
	assert.FileExists(this.T(), material.SheeterPathCs())
}

func (this *SuitePoststep) TestGenerateHelperCs() {
	material := this.prepareData(this.excel, this.sheet)
	result := generateHelperCs(material)
	assert.Nil(this.T(), result.Error)
	assert.FileExists(this.T(), material.HelperPathCs())
}

func (this *SuitePoststep) TestGenerateSheeterGo() {
	material := this.prepareData(this.excel, this.sheet)
	result := generateSheeterGo(material)
	assert.Nil(this.T(), result.Error)
	assert.FileExists(this.T(), material.SheeterPathGo())
}

func (this *SuitePoststep) TestGenerateHelperGo() {
	material := this.prepareData(this.excel, this.sheet)
	result := generateHelperGo(material)
	assert.Nil(this.T(), result.Error)
	assert.FileExists(this.T(), material.HelperPathGo())
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
