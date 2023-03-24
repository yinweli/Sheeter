package builds

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/v2/sheeter"
	"github.com/yinweli/Sheeter/v2/sheeter/excels"
	"github.com/yinweli/Sheeter/v2/sheeter/fields"
	"github.com/yinweli/Sheeter/v2/sheeter/nameds"
	"github.com/yinweli/Sheeter/v2/sheeter/utils"
	"github.com/yinweli/Sheeter/v2/testdata"
)

func TestOperation(t *testing.T) {
	suite.Run(t, new(SuiteOperation))
}

type SuiteOperation struct {
	suite.Suite
	testdata.Env
	folder             string
	excel              string
	sheetSuccess       string
	sheetFieldNotExist string
	sheetFieldEmpty    string
	sheetPkeyNotExist  string
	lineOfTag          int
	lineOfName         int
	lineOfNote         int
	lineOfField        int
	lintOfData         int
	tag                string
}

func (this *SuiteOperation) SetupSuite() {
	testdata.EnvSetup(&this.Env, "test-builds-operation", "operation")
	this.folder = "operation"
	this.excel = "operation.xlsx"
	this.sheetSuccess = "Success"
	this.sheetFieldNotExist = "Failed1"
	this.sheetFieldEmpty = "Failed2"
	this.sheetPkeyNotExist = "Failed3"
	this.lineOfTag = 1
	this.lineOfName = 2
	this.lineOfNote = 3
	this.lineOfField = 4
	this.lintOfData = 5
	this.tag = "1"
}

func (this *SuiteOperation) TearDownSuite() {
	excels.CloseAll()
	testdata.EnvRestore(&this.Env)
}

func (this *SuiteOperation) TestOperation() {
	config := this.prepareConfig([]string{this.folder})
	context, _ := Initialize(config)
	time.Sleep(testdata.Timeout)
	file, err := Operation(config, context)
	time.Sleep(testdata.Timeout)
	assert.Len(this.T(), err, 0)
	assert.Len(this.T(), file, 3)

	for _, itor := range file {
		assert.FileExists(this.T(), itor.(string))
	} // for
}

func (this *SuiteOperation) TestParseLayout() {
	operationData := this.prepareData(this.excel, this.sheetSuccess)
	assert.Nil(this.T(), parseLayout(operationData, nil))
	assert.NotNil(this.T(), operationData.Pkey)
	assert.Equal(this.T(), &fields.Pkey{}, operationData.Pkey.Pkey)
	assert.NotNil(this.T(), operationData.Field)
	assert.Len(this.T(), operationData.Field, 5)
	assert.NotNil(this.T(), operationData.Layout)

	operationData.SheetName = testdata.Unknown
	assert.NotNil(this.T(), parseLayout(operationData, nil))

	operationData.SheetName = this.sheetFieldNotExist
	assert.NotNil(this.T(), parseLayout(operationData, nil))

	operationData.SheetName = this.sheetFieldEmpty
	assert.NotNil(this.T(), parseLayout(operationData, nil))

	operationData.SheetName = this.sheetPkeyNotExist
	assert.NotNil(this.T(), parseLayout(operationData, nil))
}

func (this *SuiteOperation) TestGenerateData() {
	expected, _ := utils.JsonMarshal(map[string]interface{}{
		"1": map[string]interface{}{
			"pkey":  int32(1),
			"name1": int32(10),
			"name2": int32(11),
			"name3": int32(12),
			"name4": int32(13),
		},
		"2": map[string]interface{}{
			"pkey":  int32(2),
			"name1": int32(20),
			"name2": int32(21),
			"name3": int32(22),
			"name4": int32(23),
		},
		"4": map[string]interface{}{
			"pkey":  int32(4),
			"name1": int32(40),
			"name2": int32(41),
			"name3": int32(42),
			"name4": int32(43),
		},
		"5": map[string]interface{}{
			"pkey":  int32(5),
			"name1": int32(50),
			"name2": int32(51),
			"name3": int32(52),
			"name4": int32(53),
		},
	})

	result := make(chan any, sheeter.MaxExcel)
	operationData := this.prepareData(this.excel, this.sheetSuccess)
	assert.Nil(this.T(), parseLayout(operationData, result))
	assert.Nil(this.T(), generateData(operationData, result))
	assert.True(this.T(), testdata.CompareFile(operationData.DataPath(), expected))
}

func (this *SuiteOperation) TestGenerateReaderCs() {
	result := make(chan any, sheeter.MaxExcel)
	operationData := this.prepareData(this.excel, this.sheetSuccess)
	assert.Nil(this.T(), parseLayout(operationData, result))
	assert.Nil(this.T(), generateReaderCs(operationData, result))
	assert.FileExists(this.T(), operationData.ReaderPathCs())
}

func (this *SuiteOperation) TestGenerateReaderGo() {
	result := make(chan any, sheeter.MaxExcel)
	operationData := this.prepareData(this.excel, this.sheetSuccess)
	assert.Nil(this.T(), parseLayout(operationData, result))
	assert.Nil(this.T(), generateReaderGo(operationData, result))
	assert.FileExists(this.T(), operationData.ReaderPathGo())
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
