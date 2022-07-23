package util

import (
	"encoding/json"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/yinweli/Sheeter/testdata"
)

func TestJsonWrite(t *testing.T) {
	suite.Run(t, new(SuiteJsonWrite))
}

type SuiteJsonWrite struct {
	suite.Suite
	workDir       string
	filePathReal  string
	filePathFake1 string
	filePathFake2 string
	jsonDatas     map[string]string
	jsonBytes     []byte
	jsonBytesBom  []byte
}

func (this *SuiteJsonWrite) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
	this.filePathReal = "json/test.txt"
	this.filePathFake1 = "????.txt"
	this.filePathFake2 = "????/????.txt"
	this.jsonDatas = map[string]string{"data": "value"}
	this.jsonBytes, _ = json.MarshalIndent(this.jsonDatas, jsonPrefix, jsonIdent)
	this.jsonBytesBom = append(bomPrefix, this.jsonBytes...)
}

func (this *SuiteJsonWrite) TearDownSuite() {
	_ = os.RemoveAll(path.Dir(this.filePathReal))
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteJsonWrite) check(filepath string, expected []byte) {
	actual, err := os.ReadFile(filepath)
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), expected, actual)
}

func (this *SuiteJsonWrite) TestJsonWrite() {
	assert.Nil(this.T(), JsonWrite(this.filePathReal, this.jsonDatas, false))
	this.check(this.filePathReal, this.jsonBytes)
}

func (this *SuiteJsonWrite) TestJsonWriteBom() {
	assert.Nil(this.T(), JsonWrite(this.filePathReal, this.jsonDatas, true))
	this.check(this.filePathReal, this.jsonBytesBom)
}

func (this *SuiteJsonWrite) TestJsonWriteFailed() {
	assert.NotNil(this.T(), JsonWrite(this.filePathFake1, this.jsonDatas, false))
	assert.NotNil(this.T(), JsonWrite(this.filePathFake2, this.jsonDatas, false))
}
