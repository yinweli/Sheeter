package builder

import (
	"testing"

	"Sheeter/testdata"

	"github.com/stretchr/testify/assert"
)

func TestReadConfig(t *testing.T) {
	result, err := ReadConfig(testdata.Path(testdata.RealYaml))
	assert.NotNil(t, result, "read real config failed")
	assert.Nil(t, err, "read real config failed")

	result, err = ReadConfig(testdata.Path(testdata.FakeYaml))
	assert.NotNil(t, err, "read fake config failed")

	result, err = ReadConfig(testdata.Path(testdata.DefectYml))
	assert.NotNil(t, err, "read defect config failed")

	result, err = ReadConfig(testdata.Path(testdata.UnknownYml))
	assert.NotNil(t, err, "read unknown config failed")
}
