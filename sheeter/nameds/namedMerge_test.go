package nameds

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/v2/testdata"
)

func TestMerge(t *testing.T) {
	suite.Run(t, new(SuiteMerge))
}

type SuiteMerge struct {
	suite.Suite
	testdata.Env
}

func (this *SuiteMerge) SetupSuite() {
	this.Env = testdata.EnvSetup("test-nameds-merge")
}

func (this *SuiteMerge) TearDownSuite() {
	testdata.EnvRestore(this.Env)
}

func (this *SuiteMerge) TestName() {
	target := &Merge{
		Name: "merge",
		Member: []*Named{
			{ExcelName: "excel1", SheetName: "sheet"},
			{ExcelName: "excel2", SheetName: "sheet"},
		},
	}
	assert.Equal(this.T(), "Merge", target.StructName())
	assert.Equal(this.T(), "merge by excel1#sheet, excel2#sheet", target.StructNote())
	assert.Equal(this.T(), "Excel1SheetReader", target.ReaderName())
	assert.Equal(this.T(), []string{"Excel1Sheet", "Excel2Sheet"}, target.MemberName())
}
