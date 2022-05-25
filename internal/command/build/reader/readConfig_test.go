package reader

import (
	"testing"

	"Sheeter/internal/command/build/config"
	"Sheeter/testdata"

	"github.com/stretchr/testify/assert"
)

func TestReadConfig(t *testing.T) {
	var result *config.Config
	var err error

	dir := testdata.ChangeWorkDir()
	defer testdata.RestoreWorkDir(dir)

	result, err = ReadConfig(testdata.Path(testdata.RealConfig))
	assert.NotNil(t, result, "read real config failed")
	assert.Nil(t, err, "read real config failed")

	result, err = ReadConfig(testdata.Path(testdata.FakeConfig))
	assert.NotNil(t, err, "read fake config failed")

	result, err = ReadConfig(testdata.Path(testdata.DefectConfig))
	assert.NotNil(t, err, "read defect config failed")

	result, err = ReadConfig(testdata.Path(testdata.UnknownConfig))
	assert.NotNil(t, err, "read unknown config failed")
}
