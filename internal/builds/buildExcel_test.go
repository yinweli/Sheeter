package builds

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/testdata"
)

func TestBuildExcel(t *testing.T) {
	suite.Run(t, new(SuiteBuildExcel))
}

type SuiteBuildExcel struct {
	suite.Suite
}

func (this *SuiteBuildExcel) target() *Content {
	target := &Content{
		Excel: filepath.Join(testdata.RootPath, testdata.ExcelNameReal),
		Sheet: testdata.SheetName,
	}
	return target
}

func (this *SuiteBuildExcel) TestReadExcel() {
	target := this.target()
	assert.Nil(this.T(), readExcel(target))
	assert.NotNil(this.T(), target.excel)
	target.close()

	target = this.target()
	target.Excel = testdata.ExcelNameInvalidFile
	assert.NotNil(this.T(), readExcel(target))
	target.close()

	target = this.target()
	target.Excel = testdata.UnknownStr
	assert.NotNil(this.T(), readExcel(target))
	target.close()

	target = this.target()
	target.Sheet = testdata.UnknownStr
	assert.NotNil(this.T(), readExcel(target))
	target.close()
}
