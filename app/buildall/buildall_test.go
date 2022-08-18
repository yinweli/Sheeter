package buildall

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/testdata"
)

func TestBuildAll(t *testing.T) {
	suite.Run(t, new(SuiteBuildAll))
}

type SuiteBuildAll struct {
	suite.Suite
}

func (this *SuiteBuildAll) TestExecute() {
	// 由於沒辦法刪除產生的檔案, 所以這裡只測試錯誤

	_, errs := execute(testdata.ConfigInvalid)
	assert.NotEmpty(this.T(), errs)
}
