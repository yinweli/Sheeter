package builds

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/testdata"
)

func TestPreJson(t *testing.T) {
	suite.Run(t, new(SuitePreJson))
}

type SuitePreJson struct {
	suite.Suite
	workDir string
	json    []byte
	empty   []byte
}

func (this *SuitePreJson) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
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

func (this *SuitePreJson) TearDownSuite() {
	_ = os.RemoveAll(pathJson)
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuitePreJson) target() *Content {
	target := &Content{
		LineOfField: 1,
		LineOfLayer: 2,
		LineOfNote:  3,
		LineOfData:  4,
		Excel:       testdata.Path(testdata.ExcelNameReal),
		Sheet:       testdata.SheetName,
	}
	return target
}

func (this *SuitePreJson) TestOutputJson() {
	target := this.target()
	assert.Nil(this.T(), Initialize(target))
	assert.Nil(this.T(), OutputJson(target))
	testdata.CompareFile(this.T(), target.JsonPath(), this.json)
	target.Close()

	target = this.target()
	assert.Nil(this.T(), Initialize(target))
	target.LineOfData = -1
	assert.NotNil(this.T(), OutputJson(target))
	target.Close()

	target = this.target()
	target.Excel = testdata.Path(testdata.ExcelNameEmpty)
	assert.Nil(this.T(), Initialize(target))
	assert.Nil(this.T(), OutputJson(target))
	testdata.CompareFile(this.T(), target.JsonPath(), this.empty)
	target.Close()

	target = this.target()
	target.Excel = testdata.Path(testdata.ExcelNameInvalidData)
	assert.Nil(this.T(), Initialize(target))
	assert.NotNil(this.T(), OutputJson(target))
	target.Close()

	// 由於linux下檔案名稱幾乎沒有非法字元, 所以這項檢查只針對windows
	if testdata.IsWindows() {
		target = this.target()
		assert.Nil(this.T(), Initialize(target))
		target.Excel = testdata.UnknownStr
		assert.NotNil(this.T(), OutputJson(target))
		target.Close()
	} // if

	target = this.target()
	assert.Nil(this.T(), Initialize(target))
	target.Sheet = testdata.UnknownStr
	assert.NotNil(this.T(), OutputJson(target))
	target.Close()
}
