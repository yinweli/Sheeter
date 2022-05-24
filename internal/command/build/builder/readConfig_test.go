package builder

import (
	"testing"

	"Sheeter/internal/command/build/config"
	"Sheeter/testdata"

	"github.com/stretchr/testify/assert"
)

func TestReadConfig(t *testing.T) {
	var buildConfig config.Config

	assert.Nil(t, ReadConfig(testdata.RealYaml(), &buildConfig), "read real config failed")
	assert.NotNil(t, ReadConfig(testdata.FakeYaml(), &buildConfig), "read fake config failed")
	assert.NotNil(t, ReadConfig(testdata.UnknownYaml(), &buildConfig), "read unknown config failed")
}
