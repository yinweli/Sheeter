package core

import (
	"testing"

	"Sheeter/testdata"

	"github.com/stretchr/testify/assert"
)

func TestReadConfig(t *testing.T) {
	result, err := ReadConfig(testdata.Path(testdata.RealConfig))
	assert.Nil(t, err)
	assert.NotNil(t, result)

	result, err = ReadConfig(testdata.Path(testdata.Defect1Config))
	assert.NotNil(t, err)

	result, err = ReadConfig(testdata.Path(testdata.Defect2Config))
	assert.NotNil(t, err)

	result, err = ReadConfig(testdata.Path("?????"))
	assert.NotNil(t, err)
}
