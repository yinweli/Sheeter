package utils

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/v3/sheeter"
	"github.com/yinweli/Sheeter/v3/testdata"
)

func TestJson(t *testing.T) {
	suite.Run(t, new(SuiteJson))
}

type SuiteJson struct {
	suite.Suite
	testdata.Env
}

func (this *SuiteJson) SetupSuite() {
	this.Env = testdata.EnvSetup("test-utils-json")
}

func (this *SuiteJson) TearDownSuite() {
	testdata.EnvRestore(this.Env)
}

func (this *SuiteJson) TestJsonMarshal() {
	datas := map[string]string{"data": "value"}
	bytes, _ := json.MarshalIndent(datas, sheeter.JsonPrefix, sheeter.JsonIdent)

	result, err := JsonMarshal(datas)
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), bytes, result)
}
