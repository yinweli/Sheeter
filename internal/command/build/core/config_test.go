package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yinweli/Sheeter/testdata"
)

func TestConfig(t *testing.T) {
	config := mockConfig()
	err := config.Check()
	assert.Nil(t, err)

	config = mockConfig()
	config.Global.CppLibraryPath = ""
	err = config.Check()
	assert.NotNil(t, err)

	config = mockConfig()
	config.Global.LineOfField = 0
	err = config.Check()
	assert.NotNil(t, err)

	config = mockConfig()
	config.Global.LineOfNote = 0
	err = config.Check()
	assert.NotNil(t, err)

	config = mockConfig()
	config.Global.LineOfData = 0
	err = config.Check()
	assert.NotNil(t, err)

	config = mockConfig()
	config.Global.LineOfField = 3
	err = config.Check()
	assert.NotNil(t, err)

	config = mockConfig()
	config.Global.LineOfNote = 3
	err = config.Check()
	assert.NotNil(t, err)

	config = mockConfig()
	config.Elements[0].Excel = ""
	err = config.Check()
	assert.NotNil(t, err)

	config = mockConfig()
	config.Elements[0].Sheet = ""
	err = config.Check()
	assert.NotNil(t, err)
}

func TestReadConfig(t *testing.T) {
	result, err := ReadConfig(testdata.Path(testdata.RealConfig))
	assert.Nil(t, err)
	assert.NotNil(t, result)

	result, err = ReadConfig(testdata.Path(testdata.Defect1Config))
	assert.NotNil(t, err)

	result, err = ReadConfig(testdata.Path(testdata.Defect2Config))
	assert.NotNil(t, err)

	result, err = ReadConfig(testdata.Path("?????"))
	assert.NotNil(t, err)
}

func mockConfig() *Config {
	return &Config{
		Global: Global{
			ExcelPath:      "excel",
			CppLibraryPath: "cpp",
			Bom:            true,
			LineOfField:    1,
			LineOfNote:     2,
			LineOfData:     3,
		},
		Elements: []Element{{
			Excel: "excel.xlsx",
			Sheet: "sheet",
		}},
	}
}
