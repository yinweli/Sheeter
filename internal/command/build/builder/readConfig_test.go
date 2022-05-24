package builder

import (
	"testing"

	"Sheeter/internal/command/build/config"
	"Sheeter/testdata"

	"github.com/stretchr/testify/assert"
)

func TestReadConfig(t *testing.T) {
	var result *config.Config
	var errs []error

	result, errs = ReadConfig(testdata.Path(testdata.RealYaml))
	assert.NotNil(t, result, "read real config failed")
	assert.Equal(t, 0, len(errs), "read real config failed")

	result, errs = ReadConfig(testdata.Path(testdata.FakeYaml))
	assert.Equal(t, 1, len(errs), "read fake config failed")

	result, errs = ReadConfig(testdata.Path(testdata.DefectYml))
	assert.Equal(t, 1, len(errs), "read defect config failed")

	result, errs = ReadConfig(testdata.Path(testdata.UnknownYml))
	assert.Equal(t, 1, len(errs), "read unknown config failed")
}
