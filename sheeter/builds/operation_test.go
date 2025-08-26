package builds

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/v3/sheeter/excels"
	"github.com/yinweli/Sheeter/v3/sheeter/fields"
	"github.com/yinweli/Sheeter/v3/sheeter/nameds"
	"github.com/yinweli/Sheeter/v3/sheeter/utils"
	"github.com/yinweli/Sheeter/v3/testdata"
)

func TestOperation(t *testing.T) {
	suite.Run(t, new(SuiteOperation))
}

type SuiteOperation struct {
	suite.Suite
	testdata.Env
	folder               string
	excel                string
	sheetSuccess         string
	sheetFieldNotExist   string
	sheetFieldEmpty      string
	sheetPrimaryNotExist string
	lineOfTag            int
	lineOfName           int
	lineOfNote           int
	lineOfField          int
	lintOfData           int
	tag                  string
}

func (this *SuiteOperation) SetupSuite() {
	this.Env = testdata.EnvSetup("test-builds-operation", "operation")
	this.folder = "operation"
	this.excel = "operation.xlsx"
	this.sheetSuccess = "Success"
	this.sheetFieldNotExist = "Failed1"
	this.sheetFieldEmpty = "Failed2"
	this.sheetPrimaryNotExist = "Failed3"
	this.lineOfTag = 1
	this.lineOfName = 2
	this.lineOfNote = 3
	this.lineOfField = 4
	this.lintOfData = 5
	this.tag = "1"
}

func (this *SuiteOperation) TearDownSuite() {
	excels.CloseAll()
	testdata.EnvRestore(this.Env)
}

func (this *SuiteOperation) TestOperation() {
	config := this.prepareConfig([]string{this.folder})
	initializeData, _ := Initialize(config)
	result, err := Operation(config, initializeData)
	assert.Len(this.T(), err, 0)
	assert.Len(this.T(), result, 3)

	for _, itor := range result {
		assert.FileExists(this.T(), itor.(string))
	} // for
}

func (this *SuiteOperation) TestParseLayout() {
	material := this.prepareData(this.excel, this.sheetSuccess)
	result := parseLayout(material)
	assert.Nil(this.T(), result.Error)
	assert.NotNil(this.T(), material.Named.Primary)
	assert.Equal(this.T(), &fields.Int{}, material.Named.Primary.Field)
	assert.NotNil(this.T(), material.Field)
	assert.Len(this.T(), material.Field, 5)
	assert.NotNil(this.T(), material.Layout)

	material.SheetName = testdata.Unknown
	result = parseLayout(material)
	assert.NotNil(this.T(), result.Error)

	material.SheetName = this.sheetFieldNotExist
	result = parseLayout(material)
	assert.NotNil(this.T(), result.Error)

	material.SheetName = this.sheetFieldEmpty
	result = parseLayout(material)
	assert.NotNil(this.T(), result.Error)

	material.SheetName = this.sheetPrimaryNotExist
	result = parseLayout(material)
	assert.NotNil(this.T(), result.Error)
}

func (this *SuiteOperation) TestGenerateData() {
	expected, _ := utils.JsonMarshal(map[string]interface{}{
		"1": map[string]interface{}{
			"name1": int32(1),
			"name2": int32(10),
			"name3": int32(11),
			"name4": int32(12),
			"name5": int32(13),
		},
		"2": map[string]interface{}{
			"name1": int32(2),
			"name2": int32(20),
			"name3": int32(21),
			"name4": int32(22),
			"name5": int32(23),
		},
		"4": map[string]interface{}{
			"name1": int32(4),
			"name2": int32(40),
			"name3": int32(41),
			"name4": int32(42),
			"name5": int32(43),
		},
		"5": map[string]interface{}{
			"name1": int32(5),
			"name2": int32(50),
			"name3": int32(51),
			"name4": int32(52),
			"name5": int32(53),
		},
	})
	material := this.prepareData(this.excel, this.sheetSuccess)
	result := parseLayout(material)
	assert.Nil(this.T(), result.Error)
	result = generateData(material)
	assert.Nil(this.T(), result.Error)
	assert.True(this.T(), testdata.CompareFile(material.DataPath(), expected))
}

func (this *SuiteOperation) TestGenerateReaderCs() {
	material := this.prepareData(this.excel, this.sheetSuccess)
	result := parseLayout(material)
	assert.Nil(this.T(), result.Error)
	result = generateReaderCs(material)
	assert.Nil(this.T(), result.Error)
	assert.FileExists(this.T(), material.ReaderPathCs())
}

func (this *SuiteOperation) TestGenerateReaderGo() {
	material := this.prepareData(this.excel, this.sheetSuccess)
	result := parseLayout(material)
	assert.Nil(this.T(), result.Error)
	result = generateReaderGo(material)
	assert.Nil(this.T(), result.Error)
	assert.FileExists(this.T(), material.ReaderPathGo())
}

func (this *SuiteOperation) prepareConfig(source []string) *Config {
	return &Config{
		Source:      source,
		Tag:         this.tag,
		LineOfTag:   this.lineOfTag,
		LineOfName:  this.lineOfName,
		LineOfNote:  this.lineOfNote,
		LineOfField: this.lineOfField,
		LineOfData:  this.lintOfData,
	}
}

func (this *SuiteOperation) prepareData(excelName, sheetName string) *OperationData {
	excel := &excels.Excel{}
	assert.Nil(this.T(), excel.Open(excelName))
	sheet, err := excel.Get(sheetName)
	assert.Nil(this.T(), err)
	return &OperationData{
		Config: this.prepareConfig(nil),
		Excel:  excel,
		Sheet:  sheet,
		Named: &nameds.Named{
			ExcelName: excelName,
			SheetName: sheetName,
		},
	}
}
