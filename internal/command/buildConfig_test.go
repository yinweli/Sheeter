package command

import (
	"io/ioutil"
	"testing"

	"Sheeter/testdata"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestBuildConfig_Check(t *testing.T) {
	var buildConfig BuildConfig

	loadBuildConfig(t, &buildConfig)
	assert.Equal(t, true, buildConfig.Check(), "check failed")

	loadBuildConfig(t, &buildConfig)
	buildConfig.Global.ExcelPath = ""
	assert.Equal(t, false, buildConfig.Check(), "check failed(Global.ExcelPath)")

	loadBuildConfig(t, &buildConfig)
	buildConfig.Global.OutputPathJson = ""
	assert.Equal(t, false, buildConfig.Check(), "check failed(Global.OutputPathJson)")

	loadBuildConfig(t, &buildConfig)
	buildConfig.Global.OutputPathGo = ""
	assert.Equal(t, false, buildConfig.Check(), "check failed(Global.OutputPathGo)")

	loadBuildConfig(t, &buildConfig)
	buildConfig.Global.OutputPathCs = ""
	assert.Equal(t, false, buildConfig.Check(), "check failed(Global.OutputPathCs)")

	loadBuildConfig(t, &buildConfig)
	buildConfig.Global.OutputPathCpp = ""
	assert.Equal(t, false, buildConfig.Check(), "check failed(Global.OutputPathCpp)")

	loadBuildConfig(t, &buildConfig)
	buildConfig.Global.LineOfNote = 3
	assert.Equal(t, false, buildConfig.Check(), "check failed(Global.LineOfNote)")

	loadBuildConfig(t, &buildConfig)
	buildConfig.Global.LineOfField = 3
	assert.Equal(t, false, buildConfig.Check(), "check failed(Global.LineOfField)")

	loadBuildConfig(t, &buildConfig)
	buildConfig.Elements = []Element{}
	assert.Equal(t, false, buildConfig.Check(), "check failed(Elements empty)")

	loadBuildConfig(t, &buildConfig)
	buildConfig.Elements[0].ExcelName = ""
	assert.Equal(t, false, buildConfig.Check(), "check failed(Elements.ExcelName)")

	loadBuildConfig(t, &buildConfig)
	buildConfig.Elements[0].SheetName = ""
	assert.Equal(t, false, buildConfig.Check(), "check failed(Elements.SheetName)")

}

// loadBuildConfig 讀取建立表格設定資料
func loadBuildConfig(t *testing.T, buildConfig *BuildConfig) {
	yamlPath := testdata.Path("test.yaml")
	yamlFile, err := ioutil.ReadFile(yamlPath)

	assert.NotNil(t, yamlFile, "load buildConfig failed")
	assert.Nil(t, err, err)

	err = yaml.Unmarshal(yamlFile, buildConfig)

	assert.Nil(t, err, err)
}
