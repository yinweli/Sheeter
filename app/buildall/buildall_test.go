package buildall

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

// TODO: 要記得做buildall的測試

func TestBuildAll(t *testing.T) {
	suite.Run(t, new(SuiteBuildAll))
}

type SuiteBuildAll struct {
	suite.Suite
}

/*
func (this *SuiteConfig) TestReadConfig() {
	config, err := readConfig(testdata.Path(testdata.RealConfig))
	assert.Nil(this.T(), err)
	assert.NotNil(this.T(), config)

	_, err = readConfig(testdata.Path(testdata.Defect1Config))
	assert.NotNil(this.T(), err)

	_, err = readConfig(testdata.Path(testdata.Defect2Config))
	assert.NotNil(this.T(), err)

	_, err = readConfig(testdata.UnknownStr)
	assert.NotNil(this.T(), err)
}
*/
