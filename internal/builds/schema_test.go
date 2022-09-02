package builds

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/testdata"
)

func TestSchema(t *testing.T) {
	suite.Run(t, new(SuiteSchema))
}

type SuiteSchema struct {
	suite.Suite
	workDir string
	schema  []byte
}

func (this *SuiteSchema) SetupSuite() {
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
}

func (this *SuiteSchema) TearDownSuite() {
	_ = os.RemoveAll(pathJsonSchema)
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteSchema) target() *Content {
	target := &Content{
		LineOfField: 1,
		LineOfLayer: 2,
		LineOfNote:  3,
		LineOfData:  4,
		Excel:       testdata.ExcelNameReal,
		Sheet:       testdata.SheetName,
	}
	return target
}

func (this *SuiteSchema) TestOutputJsonSchema() {
	target := this.target()
	assert.Nil(this.T(), Initialize(target))
	assert.Nil(this.T(), OutputJsonSchema(target))
	testdata.CompareFile(this.T(), target.FileJsonSchema(), this.schema)
	target.Close()

	// 由於linux下檔案名稱幾乎沒有非法字元, 所以這項檢查只針對windows
	if testdata.IsWindows() {
		target = this.target()
		assert.Nil(this.T(), Initialize(target))
		target.Excel = testdata.UnknownStr
		assert.NotNil(this.T(), OutputJsonSchema(target))
		target.Close()
	} // if

	// 由於linux下檔案名稱幾乎沒有非法字元, 所以這項檢查只針對windows
	if testdata.IsWindows() {
		target = this.target()
		assert.Nil(this.T(), Initialize(target))
		target.Sheet = testdata.UnknownStr
		assert.NotNil(this.T(), OutputJsonSchema(target))
		target.Close()
	} // if
}
