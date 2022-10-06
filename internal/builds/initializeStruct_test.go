package builds

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/testdata"
)

func TestInitializeStruct(t *testing.T) {
	suite.Run(t, new(SuiteInitializeStruct))
}

type SuiteInitializeStruct struct {
	suite.Suite
	workDir string
}

func (this *SuiteInitializeStruct) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
}

func (this *SuiteInitializeStruct) TearDownSuite() {
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteInitializeStruct) target() *Runtime {
	target := &Runtime{
		Sector: []*RuntimeSector{
			{
				Global: Global{
					LineOfField: 1,
					LineOfLayer: 2,
					LineOfNote:  3,
				},
				Element: Element{
					Excel: testdata.ExcelNameReal,
					Sheet: testdata.SheetName,
				},
			},
			{
				Global: Global{
					LineOfField: 1,
					LineOfLayer: 2,
					LineOfNote:  3,
				},
				Element: Element{
					Excel: testdata.ExcelNameReal,
					Sheet: testdata.SheetName,
				},
			},
		},
	}
	return target
}

func (this *SuiteInitializeStruct) TestInitializeStruct() {
	target := this.target()
	assert.Nil(this.T(), initializeSector(target.Sector[0]))
	assert.Nil(this.T(), initializeSector(target.Sector[1]))
	assert.Nil(this.T(), initializeStruct(target, &Config{}))

	structNames := []string{}

	for _, itor := range target.Struct {
		structNames = append(structNames, itor.StructName())
	} // for

	assert.ElementsMatch(this.T(), []string{"RealData", "S", "A"}, structNames)
}
