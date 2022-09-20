package builds

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/internal/mixeds"
	"github.com/yinweli/Sheeter/testdata"
)

func TestEncodingJson(t *testing.T) {
	suite.Run(t, new(SuiteEncodingJson))
}

type SuiteEncodingJson struct {
	suite.Suite
	workDir string
}

func (this *SuiteEncodingJson) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
}

func (this *SuiteEncodingJson) TearDownSuite() {
	_ = os.RemoveAll(internal.PathJsonData)
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteEncodingJson) target() *RuntimeSector {
	target := &RuntimeSector{
		Mixed: mixeds.NewMixed(testdata.ExcelNameReal, testdata.SheetName),
		Global: Global{
			LineOfField: 1,
			LineOfLayer: 2,
			LineOfNote:  3,
			LineOfData:  4,
		},
		Element: Element{
			Excel: testdata.ExcelNameReal,
			Sheet: testdata.SheetName,
		},
	}
	return target
}

func (this *SuiteEncodingJson) TestEncodingJson() {
	data := []byte(`{
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
	empty := []byte("{}")

	target := this.target()
	assert.Nil(this.T(), initializeSector(target))
	assert.Nil(this.T(), encodingJson(target))
	testdata.CompareFile(this.T(), target.FileJsonData(), data)
	target.Close()

	target = this.target()
	target.LineOfData = -1
	assert.Nil(this.T(), initializeSector(target))
	assert.NotNil(this.T(), encodingJson(target))
	target.Close()

	target = this.target()
	target.Excel = testdata.ExcelNameEmpty
	assert.Nil(this.T(), initializeSector(target))
	assert.Nil(this.T(), encodingJson(target))
	testdata.CompareFile(this.T(), target.FileJsonData(), empty)
	target.Close()

	target = this.target()
	target.Excel = testdata.ExcelNameInvalidData
	assert.Nil(this.T(), initializeSector(target))
	assert.NotNil(this.T(), encodingJson(target))
	target.Close()

	// 由於linux下檔案名稱幾乎沒有非法字元, 所以這項檢查只針對windows
	if testdata.IsWindows() {
		target = this.target()
		assert.Nil(this.T(), initializeSector(target))
		target.Mixed = mixeds.NewMixed(testdata.UnknownStr, target.Sheet)
		assert.NotNil(this.T(), encodingJson(target))
		target.Close()

		target = this.target()
		assert.Nil(this.T(), initializeSector(target))
		target.Mixed = mixeds.NewMixed(target.Excel, testdata.UnknownStr)
		assert.NotNil(this.T(), encodingJson(target))
		target.Close()
	} // if
}
