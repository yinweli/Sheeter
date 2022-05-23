package config

import (
	"io/ioutil"
	"testing"

	"Sheeter/testdata"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestConfig_Check(t *testing.T) {
	var config *Config
	var result bool
	var errs []string

	config = loadConfig(t)
	result, errs = config.Check()
	assert.Equal(t, true, result, "check result failed")
	assert.Equal(t, 0, len(errs), "check errs failed")

	config = loadConfig(t)
	config.Global.ExcelPath = ""
	result, errs = config.Check()
	assert.Equal(t, false, result, "check result failed(Global.ExcelPath)")
	assert.Equal(t, 1, len(errs), "check errs failed(Global.ExcelPath)")

	config = loadConfig(t)
	config.Global.OutputPathJson = ""
	result, errs = config.Check()
	assert.Equal(t, false, result, "check result failed(Global.OutputPathJson)")
	assert.Equal(t, 1, len(errs), "check errs failed(Global.OutputPathJson)")

	config = loadConfig(t)
	config.Global.OutputPathCpp = ""
	result, errs = config.Check()
	assert.Equal(t, false, result, "check result failed(Global.OutputPathCpp)")
	assert.Equal(t, 1, len(errs), "check errs failed(Global.OutputPathCpp)")

	config = loadConfig(t)
	config.Global.OutputPathCs = ""
	result, errs = config.Check()
	assert.Equal(t, false, result, "check result failed(Global.OutputPathCs)")
	assert.Equal(t, 1, len(errs), "check errs failed(Global.OutputPathCs)")

	config = loadConfig(t)
	config.Global.OutputPathGo = ""
	result, errs = config.Check()
	assert.Equal(t, false, result, "check result failed(Global.OutputPathGo)")
	assert.Equal(t, 1, len(errs), "check errs failed(Global.OutputPathGo)")

	config = loadConfig(t)
	config.Global.GoPackage = ""
	result, errs = config.Check()
	assert.Equal(t, false, result, "check result failed(Global.GoPackage)")
	assert.Equal(t, 1, len(errs), "check errs failed(Global.GoPackage)")

	config = loadConfig(t)
	config.Global.LineOfNote = 3
	result, errs = config.Check()
	assert.Equal(t, false, result, "check result failed(Global.LineOfNote)")
	assert.Equal(t, 1, len(errs), "check errs failed(Global.LineOfNote)")

	config = loadConfig(t)
	config.Global.LineOfField = 3
	result, errs = config.Check()
	assert.Equal(t, false, result, "check result failed(Global.LineOfField)")
	assert.Equal(t, 1, len(errs), "check errs failed(Global.LineOfField)")

	config = loadConfig(t)
	config.Elements = []Element{}
	result, errs = config.Check()
	assert.Equal(t, false, result, "check result failed(Elements empty)")
	assert.Equal(t, 1, len(errs), "check errs failed(Elements empty)")

	config = loadConfig(t)
	config.Elements[0].Excel = ""
	result, errs = config.Check()
	assert.Equal(t, false, result, "check result failed(Elements.Excel)")
	assert.Equal(t, 1, len(errs), "check errs failed(Elements.Excel)")

	config = loadConfig(t)
	config.Elements[0].Sheet = ""
	result, errs = config.Check()
	assert.Equal(t, false, result, "check result failed(Elements.Sheet)")
	assert.Equal(t, 1, len(errs), "check errs failed(Elements.Sheet)")
}

// loadConfig 讀取編譯設定
func loadConfig(t *testing.T) *Config {
	file, err := ioutil.ReadFile(testdata.RealYaml())

	assert.Nil(t, err, "load file failed")

	config := &Config{}
	err = yaml.Unmarshal(file, config)

	assert.Nil(t, err, "load config failed")

	return config
}
