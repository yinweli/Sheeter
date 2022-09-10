package builds

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/testdata"
)

func TestSectorJson(t *testing.T) {
	suite.Run(t, new(SuiteSectorJson))
}

type SuiteSectorJson struct {
	suite.Suite
	workDir string
	json    []byte
	schema  []byte
	empty   []byte
}

func (this *SuiteSectorJson) SetupSuite() {
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
	this.empty = []byte("{}")
}

func (this *SuiteSectorJson) TearDownSuite() {
	_ = os.RemoveAll(internal.PathJson)
	_ = os.RemoveAll(internal.PathJsonSchema)
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteSectorJson) target() *Sector {
	target := &Sector{
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

func (this *SuiteSectorJson) TestSectorJson() {
	target := this.target()
	assert.Nil(this.T(), SectorInit(target))
	assert.Nil(this.T(), SectorJson(target))
	testdata.CompareFile(this.T(), target.FileJson(), this.json)
	target.Close()

	target = this.target()
	assert.Nil(this.T(), SectorInit(target))
	target.LineOfData = -1
	assert.NotNil(this.T(), SectorJson(target))
	target.Close()

	target = this.target()
	target.Excel = testdata.ExcelNameEmpty
	assert.Nil(this.T(), SectorInit(target))
	assert.Nil(this.T(), SectorJson(target))
	testdata.CompareFile(this.T(), target.FileJson(), this.empty)
	target.Close()

	target = this.target()
	target.Excel = testdata.ExcelNameInvalidData
	assert.Nil(this.T(), SectorInit(target))
	assert.NotNil(this.T(), SectorJson(target))
	target.Close()

	target = this.target()
	assert.Nil(this.T(), SectorInit(target))
	target.Sheet = testdata.UnknownStr
	assert.NotNil(this.T(), SectorJson(target))
	target.Close()

	// 由於linux下檔案名稱幾乎沒有非法字元, 所以這項檢查只針對windows
	if testdata.IsWindows() {
		target = this.target()
		assert.Nil(this.T(), SectorInit(target))
		target.Excel = testdata.UnknownStr
		assert.NotNil(this.T(), SectorJson(target))
		target.Close()
	} // if
}

func (this *SuiteSectorJson) TestSectorJsonSchema() {
	target := this.target()
	assert.Nil(this.T(), SectorInit(target))
	assert.Nil(this.T(), SectorJsonSchema(target))
	testdata.CompareFile(this.T(), target.FileJsonSchema(), this.schema)
	target.Close()

	// 由於linux下檔案名稱幾乎沒有非法字元, 所以這項檢查只針對windows
	if testdata.IsWindows() {
		target = this.target()
		assert.Nil(this.T(), SectorInit(target))
		target.Excel = testdata.UnknownStr
		target.Sheet = testdata.UnknownStr
		assert.NotNil(this.T(), SectorJsonSchema(target))
		target.Close()
	} // if
}
