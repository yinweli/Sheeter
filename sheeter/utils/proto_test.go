package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/testdata"
)

func TestProto(t *testing.T) {
	suite.Run(t, new(SuiteProto))
}

type SuiteProto struct {
	suite.Suite
	testdata.TestEnv
}

func (this *SuiteProto) SetupSuite() {
	this.Change("test-proto")
}

func (this *SuiteProto) TearDownSuite() {
	this.Restore()
}

func (this *SuiteProto) TestJsonToProto() {
	// proto轉json時, 只要是int64的欄位, 轉為json時都會被轉為字串
	// 所以測試的欄位要注意不要改成int64, 會造成測試錯誤

	raw := []byte(`{"Value":123456,"Data":{"Value":654321}}`)
	data, err := JsonToProto(testdata.ProtoTest, "test.Test1", []string{}, raw)
	assert.Nil(this.T(), err)
	assert.NotNil(this.T(), data)
	json, err := ProtoToJson(testdata.ProtoTest, "test.Test1", []string{}, data)
	assert.Nil(this.T(), err)
	assert.NotNil(this.T(), json)
	assert.Equal(this.T(), string(raw), string(json))
}

func (this *SuiteProto) TestParseProto() {
	file, err := parseProto(testdata.ProtoTest, []string{})
	assert.Nil(this.T(), err)
	assert.NotNil(this.T(), file)
	assert.NotNil(this.T(), file.FindMessage("test.Test1"))
	assert.NotNil(this.T(), file.FindMessage("test.TestX"))

	_, err = parseProto(testdata.UnknownStr, []string{})
	assert.NotNil(this.T(), err)
}
