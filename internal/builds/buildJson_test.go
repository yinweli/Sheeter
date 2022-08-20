package builds

import (
	"os"
	"path/filepath"
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
    "S": {
        "A": [
            {
                "name2": 0,
                "name3": ""
            },
            {
                "name2": 0,
                "name3": ""
            },
            {
                "name2": 0,
                "name3": ""
            }
        ],
        "name1": false
    },
    "name0": 0
}`)
	this.json = []byte(`{
    "1": {
        "S": {
            "A": [
                {
                    "name2": 1,
                    "name3": "a"
                },
                {
                    "name2": 1,
                    "name3": "a"
                },
                {
                    "name2": 1,
                    "name3": "a"
                }
            ],
            "name1": true
        },
        "name0": 1
    },
    "2": {
        "S": {
            "A": [
                {
                    "name2": 2,
                    "name3": "b"
                },
                {
                    "name2": 2,
                    "name3": "b"
                },
                {
                    "name2": 2,
                    "name3": "b"
                }
            ],
            "name1": false
        },
        "name0": 2
    },
    "3": {
        "S": {
            "A": [
                {
                    "name2": 3,
                    "name3": "c"
                },
                {
                    "name2": 3,
                    "name3": "c"
                },
                {
                    "name2": 3,
                    "name3": "c"
                }
            ],
            "name1": true
        },
        "name0": 3
    }
}`)
	this.empty = []byte("{}")
}

func (this *SuiteBuildJson) TearDownSuite() {
	_ = os.RemoveAll(pathSchema)
	_ = os.RemoveAll(pathJson)
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteBuildJson) target() *Content {
	target := &Content{
		LineOfField: 1,
		LineOfLayer: 2,
		LineOfNote:  3,
		LineOfData:  4,
		Excel:       filepath.Join(testdata.RootPath, testdata.ExcelNameReal),
		Sheet:       testdata.SheetName,
	}
	return target
}

func (this *SuiteBuildJson) TestWriteSchema() {
	target := this.target()
	target.excel = testdata.GetTestExcel(testdata.ExcelNameReal)
	assert.Nil(this.T(), buildLayout(target))
	assert.Nil(this.T(), writeSchema(target))
	testdata.CompareFile(this.T(), target.SchemaPath(), this.schema)
	target.close()

	target = this.target()
	target.excel = testdata.GetTestExcel(testdata.ExcelNameReal)
	assert.Nil(this.T(), buildLayout(target))
	target.Excel = testdata.UnknownStr
	assert.NotNil(this.T(), writeSchema(target))
	target.close()

	target = this.target()
	target.excel = testdata.GetTestExcel(testdata.ExcelNameReal)
	assert.Nil(this.T(), buildLayout(target))
	target.Sheet = testdata.UnknownStr
	assert.NotNil(this.T(), writeSchema(target))
	target.close()
}

func (this *SuiteBuildJson) TestWriteJson() {
	target := this.target()
	target.excel = testdata.GetTestExcel(testdata.ExcelNameReal)
	assert.Nil(this.T(), buildLayout(target))
	assert.Nil(this.T(), writeJson(target))
	testdata.CompareFile(this.T(), target.JsonPath(), this.json)
	target.close()

	target = this.target()
	target.excel = testdata.GetTestExcel(testdata.ExcelNameEmpty)
	assert.Nil(this.T(), buildLayout(target))
	assert.Nil(this.T(), writeJson(target))
	testdata.CompareFile(this.T(), target.JsonPath(), this.empty)
	target.close()

	target = this.target()
	target.excel = testdata.GetTestExcel(testdata.ExcelNameInvalidData)
	assert.Nil(this.T(), buildLayout(target))
	assert.NotNil(this.T(), writeJson(target))
	target.close()

	target = this.target()
	target.excel = testdata.GetTestExcel(testdata.ExcelNameReal)
	assert.Nil(this.T(), buildLayout(target))
	target.LineOfData = -1
	assert.NotNil(this.T(), writeJson(target))
	target.close()

	target = this.target()
	target.excel = testdata.GetTestExcel(testdata.ExcelNameReal)
	assert.Nil(this.T(), buildLayout(target))
	target.Excel = testdata.UnknownStr
	assert.NotNil(this.T(), writeJson(target))
	target.close()

	target = this.target()
	target.excel = testdata.GetTestExcel(testdata.ExcelNameReal)
	assert.Nil(this.T(), buildLayout(target))
	target.Sheet = testdata.UnknownStr
	assert.NotNil(this.T(), writeJson(target))
	target.close()
}
