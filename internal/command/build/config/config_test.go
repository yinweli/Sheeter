package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	var config *Config
	var result bool
	var errs []error

	config = mockConfig()
	result, errs = config.Check()
	assert.Equal(t, true, result, "check failed")
	assert.Equal(t, 0, len(errs), "check failed")

	config = mockConfig()
	config.Global.ExcelPath = ""
	result, errs = config.Check()
	assert.Equal(t, false, result, "check failed(Global.ExcelPath)")
	assert.Equal(t, 1, len(errs), "check failed(Global.ExcelPath)")

	config = mockConfig()
	config.Global.OutputPathJson = ""
	result, errs = config.Check()
	assert.Equal(t, false, result, "check failed(Global.OutputPathJson)")
	assert.Equal(t, 1, len(errs), "check failed(Global.OutputPathJson)")

	config = mockConfig()
	config.Global.OutputPathCpp = ""
	result, errs = config.Check()
	assert.Equal(t, false, result, "check failed(Global.OutputPathCpp)")
	assert.Equal(t, 1, len(errs), "check failed(Global.OutputPathCpp)")

	config = mockConfig()
	config.Global.OutputPathCs = ""
	result, errs = config.Check()
	assert.Equal(t, false, result, "check failed(Global.OutputPathCs)")
	assert.Equal(t, 1, len(errs), "check failed(Global.OutputPathCs)")

	config = mockConfig()
	config.Global.OutputPathGo = ""
	result, errs = config.Check()
	assert.Equal(t, false, result, "check failed(Global.OutputPathGo)")
	assert.Equal(t, 1, len(errs), "check failed(Global.OutputPathGo)")

	config = mockConfig()
	config.Global.CppLibraryPath = ""
	result, errs = config.Check()
	assert.Equal(t, false, result, "check failed(Global.CppLibraryPath)")
	assert.Equal(t, 1, len(errs), "check failed(Global.CppLibraryPath)")

	config = mockConfig()
	config.Global.GoPackage = ""
	result, errs = config.Check()
	assert.Equal(t, false, result, "check failed(Global.GoPackage)")
	assert.Equal(t, 1, len(errs), "check failed(Global.GoPackage)")

	config = mockConfig()
	config.Global.LineOfNote = 3
	result, errs = config.Check()
	assert.Equal(t, false, result, "check failed(Global.LineOfNote)")
	assert.Equal(t, 1, len(errs), "check failed(Global.LineOfNote)")

	config = mockConfig()
	config.Global.LineOfField = 3
	result, errs = config.Check()
	assert.Equal(t, false, result, "check failed(Global.LineOfField)")
	assert.Equal(t, 1, len(errs), "check failed(Global.LineOfField)")

	config = mockConfig()
	config.Elements = []Element{}
	result, errs = config.Check()
	assert.Equal(t, false, result, "check failed(Elements empty)")
	assert.Equal(t, 1, len(errs), "check failed(Elements empty)")

	config = mockConfig()
	config.Elements[0].Excel = ""
	result, errs = config.Check()
	assert.Equal(t, false, result, "check failed(Elements.Excel)")
	assert.Equal(t, 1, len(errs), "check failed(Elements.Excel)")

	config = mockConfig()
	config.Elements[0].Sheet = ""
	result, errs = config.Check()
	assert.Equal(t, false, result, "check failed(Elements.Sheet)")
	assert.Equal(t, 1, len(errs), "check failed(Elements.Sheet)")
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
