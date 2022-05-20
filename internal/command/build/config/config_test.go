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

	loadConfig(t, &config)
	assert.Equal(t, true, config.Check(), "check failed")

	loadConfig(t, &config)
	config.Global.ExcelPath = ""
	assert.Equal(t, false, config.Check(), "check failed(Global.ExcelPath)")

	loadConfig(t, &config)
	config.Global.OutputPathJson = ""
	assert.Equal(t, false, config.Check(), "check failed(Global.OutputPathJson)")

	loadConfig(t, &config)
	config.Global.OutputPathGo = ""
	assert.Equal(t, false, config.Check(), "check failed(Global.OutputPathGo)")

	loadConfig(t, &config)
	config.Global.OutputPathCs = ""
	assert.Equal(t, false, config.Check(), "check failed(Global.OutputPathCs)")

	loadConfig(t, &config)
	config.Global.OutputPathCpp = ""
	assert.Equal(t, false, config.Check(), "check failed(Global.OutputPathCpp)")

	loadConfig(t, &config)
	config.Global.LineOfNote = 3
	assert.Equal(t, false, config.Check(), "check failed(Global.LineOfNote)")

	loadConfig(t, &config)
	config.Global.LineOfField = 3
	assert.Equal(t, false, config.Check(), "check failed(Global.LineOfField)")

	loadConfig(t, &config)
	config.Elements = []Element{}
	assert.Equal(t, false, config.Check(), "check failed(Elements empty)")

	loadConfig(t, &config)
	config.Elements[0].ExcelName = ""
	assert.Equal(t, false, config.Check(), "check failed(Elements.ExcelName)")

	loadConfig(t, &config)
	config.Elements[0].SheetName = ""
	assert.Equal(t, false, config.Check(), "check failed(Elements.SheetName)")
}

// loadConfig 讀取編譯設定
func loadConfig(t *testing.T, config *Config) {
	path := testdata.Path("config.yaml")
	file, err := ioutil.ReadFile(path)

	assert.NotNil(t, file, "load config failed")
	assert.Nil(t, err, err)

	err = yaml.Unmarshal(file, config)

	assert.Nil(t, err, err)
}
