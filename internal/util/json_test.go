package util

import (
	"encoding/json"
	"os"
	"path/filepath"
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
	workDir      string
	filePathReal string
	filePathFake string
	jsonDatas    map[string]string
	jsonBytes    []byte
	jsonBytesBom []byte
}

func (this *SuiteJsonWrite) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
	this.filePathReal = "json/test.json"
	this.filePathFake = "?json/test.json"
	this.jsonDatas = map[string]string{"data": "value"}
	this.jsonBytes, _ = json.MarshalIndent(this.jsonDatas, jsonPrefix, jsonIdent)
	this.jsonBytesBom = bomPrefix
	this.jsonBytesBom = append(this.jsonBytesBom, this.jsonBytes...)
}

func (this *SuiteJsonWrite) TearDownSuite() {
	_ = os.RemoveAll(filepath.Dir(this.filePathReal))
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteJsonWrite) TestJsonWrite() {
	assert.Nil(this.T(), JsonWrite(this.filePathReal, this.jsonDatas, false))
	testdata.CompareFile(this.T(), this.filePathReal, this.jsonBytes)

	assert.Nil(this.T(), JsonWrite(this.filePathReal, this.jsonDatas, true))
	testdata.CompareFile(this.T(), this.filePathReal, this.jsonBytesBom)

	// 由於linux下檔案名稱幾乎沒有非法字元, 所以這項檢查只針對windows
	if testdata.IsWindows() {
		assert.NotNil(this.T(), JsonWrite(this.filePathFake, this.jsonDatas, false))
	} // if
}
