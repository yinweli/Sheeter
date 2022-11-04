package utils

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/testdata"
)

func TestJson(t *testing.T) {
	suite.Run(t, new(SuiteJson))
}

type SuiteJson struct {
	suite.Suite
	testdata.TestEnv
}

func (this *SuiteJson) SetupSuite() {
	this.Change("test-json")
}

func (this *SuiteJson) TearDownSuite() {
	this.Restore()
}

func (this *SuiteJson) TestJsonMarshal() {
	datas := map[string]string{"data": "value"}
	bytes, _ := json.MarshalIndent(datas, internal.JsonPrefix, internal.JsonIdent)

	result, err := JsonMarshal(datas)
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), bytes, result)
}
