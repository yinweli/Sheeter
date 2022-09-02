package builds

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/testdata"
)

func TestInitialize(t *testing.T) {
	suite.Run(t, new(SuiteInitialize))
}

type SuiteInitialize struct {
	suite.Suite
}

func (this *SuiteInitialize) target() *Content {
	target := &Content{
		LineOfField: 1,
		LineOfLayer: 2,
		LineOfNote:  3,
		Excel:       testdata.Path(testdata.ExcelNameReal),
		Sheet:       testdata.SheetName,
	}
	return target
}

func (this *SuiteInitialize) TestInitialize() {
	target := this.target()
	assert.Nil(this.T(), Initialize(target))
	assert.NotNil(this.T(), target.excel)
	target.Close()

	target = this.target()
	target.LineOfField = 10
	assert.NotNil(this.T(), Initialize(target))
	target.Close()

	target = this.target()
	target.LineOfLayer = 10
	assert.NotNil(this.T(), Initialize(target))
	target.Close()

	target = this.target()
	target.LineOfNote = 10
	assert.NotNil(this.T(), Initialize(target))
	target.Close()

	target = this.target()
	target.Excel = testdata.Path(testdata.UnknownStr)
	assert.NotNil(this.T(), Initialize(target))
	target.Close()

	target = this.target()
	target.Excel = testdata.Path(testdata.ExcelNameInvalidFile)
	assert.NotNil(this.T(), Initialize(target))
	target.Close()

	target = this.target()
	target.Excel = testdata.Path(testdata.ExcelNameCleanAll)
	assert.NotNil(this.T(), Initialize(target))
	target.Close()

	target = this.target()
	target.Excel = testdata.Path(testdata.ExcelNameCleanField)
	assert.NotNil(this.T(), Initialize(target))
	target.Close()

	target = this.target()
	target.Excel = testdata.Path(testdata.ExcelNameInvalidField)
	assert.NotNil(this.T(), Initialize(target))
	target.Close()

	target = this.target()
	target.Excel = testdata.Path(testdata.ExcelNameInvalidLayer)
	assert.NotNil(this.T(), Initialize(target))
	target.Close()

	target = this.target()
	target.Excel = testdata.Path(testdata.ExcelNameInvalidLayout)
	assert.NotNil(this.T(), Initialize(target))
	target.Close()

	target = this.target()
	target.Excel = testdata.Path(testdata.ExcelNameInvalidPkeyZero)
	assert.NotNil(this.T(), Initialize(target))
	target.Close()

	target = this.target()
	target.Excel = testdata.Path(testdata.ExcelNameInvalidPkeyDupl)
	assert.NotNil(this.T(), Initialize(target))
	target.Close()

	target = this.target()
	target.Sheet = testdata.Path(testdata.UnknownStr)
	assert.NotNil(this.T(), Initialize(target))
	target.Close()
}
