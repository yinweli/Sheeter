package utils

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/testdata"
)

func TestJson(t *testing.T) {
	suite.Run(t, new(SuiteJson))
}

type SuiteJson struct {
	suite.Suite
	workDir string
}

func (this *SuiteJson) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
}

func (this *SuiteJson) TearDownSuite() {
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteJson) TestJsonMarshal() {
	datas := map[string]string{"data": "value"}
	bytes, _ := json.MarshalIndent(datas, jsonPrefix, jsonIdent)

	result, err := JsonMarshal(datas)
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), bytes, result)
}
