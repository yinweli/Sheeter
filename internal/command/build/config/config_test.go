package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	var config *Config
	var err error

	config = mockConfig()
	err = config.Check()
	assert.Nil(t, err, "check failed")

	config = mockConfig()
	config.Global.ExcelPath = ""
	err = config.Check()
	assert.NotNil(t, err, "check failed(Global.ExcelPath)")

	config = mockConfig()
	config.Global.OutputPathJson = ""
	err = config.Check()
	assert.NotNil(t, err, "check failed(Global.OutputPathJson)")

	config = mockConfig()
	config.Global.OutputPathCpp = ""
	err = config.Check()
	assert.NotNil(t, err, "check failed(Global.OutputPathCpp)")

	config = mockConfig()
	config.Global.OutputPathCs = ""
	err = config.Check()
	assert.NotNil(t, err, "check failed(Global.OutputPathCs)")

	config = mockConfig()
	config.Global.OutputPathGo = ""
	err = config.Check()
	assert.NotNil(t, err, "check failed(Global.OutputPathGo)")

	config = mockConfig()
	config.Global.CppLibraryPath = ""
	err = config.Check()
	assert.NotNil(t, err, "check failed(Global.CppLibraryPath)")

	config = mockConfig()
	config.Global.GoPackage = ""
	err = config.Check()
	assert.NotNil(t, err, "check failed(Global.GoPackage)")

	config = mockConfig()
	config.Global.LineOfNote = 3
	err = config.Check()
	assert.NotNil(t, err, "check failed(Global.LineOfNote)")

	config = mockConfig()
	config.Global.LineOfField = 3
	err = config.Check()
	assert.NotNil(t, err, "check failed(Global.LineOfField)")

	config = mockConfig()
	config.Elements = []Element{}
	err = config.Check()
	assert.NotNil(t, err, "check failed(Elements empty)")

	config = mockConfig()
	config.Elements[0].Excel = ""
	err = config.Check()
	assert.NotNil(t, err, "check failed(Elements.Excel)")

	config = mockConfig()
	config.Elements[0].Sheet = ""
	err = config.Check()
	assert.NotNil(t, err, "check failed(Elements.Sheet)")
}

func mockConfig() *Config {
	return &Config{
		Global: Global{
			ExcelPath:      "test",
			OutputPathJson: "json",
			OutputPathCpp:  "cpp",
			OutputPathCs:   "cs",
			OutputPathGo:   "go",
			CppLibraryPath: "nlohmann",
			GoPackage:      "sheeter",
			Bom:            true,
			LineOfNote:     1,
			LineOfField:    2,
			LineOfData:     3,
		},
		Elements: []Element{{
			Excel: "Test.xlsx",
			Sheet: "Data",
		}},
	}
}
