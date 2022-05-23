package config

import (
	"io/ioutil"
	"testing"

	"Sheeter/testdata"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestConfig_Check(t *testing.T) {
	var config Config
	var result bool
	var errs []string

	loadConfig(t, &config)
	result, errs = config.Check()
	assert.Equal(t, true, result, "check result failed")
	assert.Equal(t, 0, len(errs), "check errs failed")

	loadConfig(t, &config)
	config.Global.ExcelPath = ""
	result, errs = config.Check()
	assert.Equal(t, false, result, "check result failed(Global.ExcelPath)")
	assert.Equal(t, 1, len(errs), "check errs failed(Global.ExcelPath)")

	loadConfig(t, &config)
	config.Global.OutputPathJson = ""
	result, errs = config.Check()
	assert.Equal(t, false, result, "check result failed(Global.OutputPathJson)")
	assert.Equal(t, 1, len(errs), "check errs failed(Global.OutputPathJson)")

	loadConfig(t, &config)
	config.Global.OutputPathCpp = ""
	result, errs = config.Check()
	assert.Equal(t, false, result, "check result failed(Global.OutputPathCpp)")
	assert.Equal(t, 1, len(errs), "check errs failed(Global.OutputPathCpp)")

	loadConfig(t, &config)
	config.Global.OutputPathCs = ""
	result, errs = config.Check()
	assert.Equal(t, false, result, "check result failed(Global.OutputPathCs)")
	assert.Equal(t, 1, len(errs), "check errs failed(Global.OutputPathCs)")

	loadConfig(t, &config)
	config.Global.OutputPathGo = ""
	result, errs = config.Check()
	assert.Equal(t, false, result, "check result failed(Global.OutputPathGo)")
	assert.Equal(t, 1, len(errs), "check errs failed(Global.OutputPathGo)")

	loadConfig(t, &config)
	config.Global.GoPackage = ""
	result, errs = config.Check()
	assert.Equal(t, false, result, "check result failed(Global.GoPackage)")
	assert.Equal(t, 1, len(errs), "check errs failed(Global.GoPackage)")

	loadConfig(t, &config)
	config.Global.LineOfNote = 3
	result, errs = config.Check()
	assert.Equal(t, false, result, "check result failed(Global.LineOfNote)")
	assert.Equal(t, 1, len(errs), "check errs failed(Global.LineOfNote)")

	loadConfig(t, &config)
	config.Global.LineOfField = 3
	result, errs = config.Check()
	assert.Equal(t, false, result, "check result failed(Global.LineOfField)")
	assert.Equal(t, 1, len(errs), "check errs failed(Global.LineOfField)")

	loadConfig(t, &config)
	config.Elements = []Element{}
	result, errs = config.Check()
	assert.Equal(t, false, result, "check result failed(Elements empty)")
	assert.Equal(t, 1, len(errs), "check errs failed(Elements empty)")

	loadConfig(t, &config)
	config.Elements[0].Excel = ""
	result, errs = config.Check()
	assert.Equal(t, false, result, "check result failed(Elements.Excel)")
	assert.Equal(t, 1, len(errs), "check errs failed(Elements.Excel)")

	loadConfig(t, &config)
	config.Elements[0].Sheet = ""
	result, errs = config.Check()
	assert.Equal(t, false, result, "check result failed(Elements.Sheet)")
	assert.Equal(t, 1, len(errs), "check errs failed(Elements.Sheet)")
}

// loadConfig 讀取編譯設定
func loadConfig(t *testing.T, config *Config) {
	file, err := ioutil.ReadFile(testdata.RealYaml())

	assert.Nil(t, err, "load file failed")

	err = yaml.Unmarshal(file, config)

	assert.Nil(t, err, "load config failed")
}
