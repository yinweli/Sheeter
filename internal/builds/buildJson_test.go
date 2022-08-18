package builds

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/testdata"
)

func TestBuildJson(t *testing.T) {
	suite.Run(t, new(SuiteBuildJson))
}

type SuiteBuildJson struct {
	suite.Suite
	workDir string
	schema  []byte
	json    []byte
	empty   []byte
}

func (this *SuiteBuildJson) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
	this.schema = []byte(`{
    "name0": 0,
    "S" : {
        "name1": false,
        "A": [
            {
                "name2": 0,
                "name3": "",
            },
        ]
    }
}`)
	this.json = []byte(`{
    "1": {
        "name0": 1,
        "S" : {
            "name1": true,
            "A": [
                {
                    "name2": 1,
                    "name3": "a",
                },
                {
                    "name2": 1,
                    "name3": "a",
                },
                {
                    "name2": 1,
                    "name3": "a",
                }
            ]
        }
    },
    "2": {
        "name0": 2,
        "S" : {
            "name1": false,
            "A": [
                {
                    "name2": 2,
                    "name3": "b",
                },
                {
                    "name2": 2,
                    "name3": "b",
                },
                {
                    "name2": 2,
                    "name3": "b",
                }
            ]
        }
    },
    "3": {
        "name0": 3,
        "S" : {
            "name1": true,
            "A": [
                {
                    "name2": 3,
                    "name3": "c",
                },
                {
                    "name2": 3,
                    "name3": "c",
                },
                {
                    "name2": 3,
                    "name3": "c",
                }
            ]
        }
    }
}`)
	this.empty = []byte("{}")
}

func (this *SuiteBuildJson) TearDownSuite() {
	_ = os.RemoveAll(pathSchema)
	_ = os.RemoveAll(pathJson)
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteBuildJson) content() *Content {
	content := &Content{
		Path:        testdata.RootPath,
		LineOfField: 1,
		LineOfLayer: 2,
		LineOfNote:  3,
		LineOfData:  4,
		Excel:       testdata.ExcelNameReal,
		Sheet:       testdata.SheetName,
	}
	return content
}

func (this *SuiteBuildJson) TestWriteSchema() {
	content := this.content()
	content.excel = testdata.GetTestExcel(testdata.ExcelNameReal)
	assert.Nil(this.T(), buildLayout(content))
	assert.Nil(this.T(), writeSchema(content))
	testdata.CompareFile(this.T(), content.SchemaFilePath(), this.schema)
	content.close()

	content = this.content()
	content.excel = testdata.GetTestExcel(testdata.ExcelNameInvalidData)
	assert.Nil(this.T(), buildLayout(content))
	assert.NotNil(this.T(), writeSchema(content))
	content.close()

	content = this.content()
	content.Excel = testdata.UnknownStr
	content.excel = testdata.GetTestExcel(testdata.ExcelNameReal)
	assert.Nil(this.T(), buildLayout(content))
	assert.NotNil(this.T(), writeSchema(content))
	content.close()

	content = this.content()
	content.Sheet = testdata.UnknownStr
	content.excel = testdata.GetTestExcel(testdata.ExcelNameReal)
	assert.Nil(this.T(), buildLayout(content))
	assert.NotNil(this.T(), writeSchema(content))
	content.close()
}

func (this *SuiteBuildJson) TestWriteJson() {
	content := this.content()
	content.excel = testdata.GetTestExcel(testdata.ExcelNameReal)
	assert.Nil(this.T(), buildLayout(content))
	assert.Nil(this.T(), writeJson(content))
	testdata.CompareFile(this.T(), content.JsonFilePath(), this.json)
	content.close()

	content = this.content()
	content.excel = testdata.GetTestExcel(testdata.ExcelNameEmpty)
	assert.Nil(this.T(), buildLayout(content))
	assert.Nil(this.T(), writeJson(content))
	testdata.CompareFile(this.T(), content.JsonFilePath(), this.empty)
	content.close()

	content = this.content()
	content.excel = testdata.GetTestExcel(testdata.ExcelNameInvalidData)
	assert.Nil(this.T(), buildLayout(content))
	assert.NotNil(this.T(), writeJson(content))
	content.close()

	content = this.content()
	content.LineOfData = -1
	content.excel = testdata.GetTestExcel(testdata.ExcelNameReal)
	assert.Nil(this.T(), buildLayout(content))
	assert.NotNil(this.T(), writeJson(content))
	content.close()

	content = this.content()
	content.Excel = testdata.UnknownStr
	content.excel = testdata.GetTestExcel(testdata.ExcelNameReal)
	assert.Nil(this.T(), buildLayout(content))
	assert.NotNil(this.T(), writeJson(content))
	content.close()

	content = this.content()
	content.Sheet = testdata.UnknownStr
	content.excel = testdata.GetTestExcel(testdata.ExcelNameReal)
	assert.Nil(this.T(), buildLayout(content))
	assert.NotNil(this.T(), writeJson(content))
	content.close()
}
