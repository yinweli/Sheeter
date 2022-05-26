package core

import (
	"testing"

	"Sheeter/testdata"

	"github.com/stretchr/testify/assert"
)

func TestReadConfig(t *testing.T) {
	result, err := ReadConfig(testdata.Path(testdata.RealConfig))
	assert.Nil(t, err, "read real config failed")
	assert.NotNil(t, result, "read real config failed")

	result, err = ReadConfig(testdata.Path(testdata.FakeConfig))
	assert.NotNil(t, err, "read fake config failed")

	result, err = ReadConfig(testdata.Path(testdata.DefectConfig))
	assert.NotNil(t, err, "read defect config failed")

	result, err = ReadConfig(testdata.Path(testdata.UnknownConfig))
	assert.NotNil(t, err, "read unknown config failed")
}
