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

func TestWrite(t *testing.T) {
	suite.Run(t, new(SuiteWrite))
}

type SuiteWrite struct {
	suite.Suite
	workDir      string
	name         string
	path         string
	filePathReal string
	filePathFake string
	fileBytes    []byte
	fileBytesBom []byte
	jsonPathReal string
	jsonPathFake string
	jsonDatas    map[string]string
	jsonBytes    []byte
	jsonBytesBom []byte
	tmplPathReal string
	tmplPathFake string
	tmplContent  string
	tmplDatas    map[string]string
	tmplBytes    []byte
	tmplBytesBom []byte
}

func (this *SuiteWrite) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
	this.name = "name"
	this.path = filepath.Join("dir1", "dir2", "dir3", this.name+"."+"txt")
	this.filePathReal = "file/test.file"
	this.filePathFake = "?file/test.file"
	this.fileBytes = []byte("this is a string")
	this.fileBytesBom = bomPrefix
	this.fileBytesBom = append(this.fileBytesBom, this.fileBytes...)
	this.jsonPathReal = "json/test.json"
	this.jsonPathFake = "?json/test.json"
	this.jsonDatas = map[string]string{"data": "value"}
	this.jsonBytes, _ = json.MarshalIndent(this.jsonDatas, jsonPrefix, jsonIdent)
	this.jsonBytesBom = bomPrefix
	this.jsonBytesBom = append(this.jsonBytesBom, this.jsonBytes...)
	this.tmplPathReal = "tmpl/test.tmpl"
	this.tmplPathFake = "?tmpl/test.tmpl"
	this.tmplContent = "{{$.Value}}"
	this.tmplDatas = map[string]string{"Value": "Value"}
	this.tmplBytes = []byte("Value")
	this.tmplBytesBom = bomPrefix
	this.tmplBytesBom = append(this.tmplBytesBom, this.tmplBytes...)
}

func (this *SuiteWrite) TearDownSuite() {
	_ = os.RemoveAll(filepath.Dir(this.filePathReal))
	_ = os.RemoveAll(filepath.Dir(this.jsonPathReal))
	_ = os.RemoveAll(filepath.Dir(this.tmplPathReal))
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteWrite) TestFileName() {
	assert.Equal(this.T(), this.name, FileName(this.path))
}

func (this *SuiteWrite) TestExistFile() {
	assert.True(this.T(), ExistFile(testdata.ConfigNameReal))
	assert.False(this.T(), ExistFile(testdata.UnknownStr))
}

func (this *SuiteWrite) TestWriteFile() {
	assert.Nil(this.T(), WriteFile(this.filePathReal, this.fileBytes, false))
	testdata.CompareFile(this.T(), this.filePathReal, this.fileBytes)

	assert.Nil(this.T(), WriteFile(this.filePathReal, this.fileBytes, true))
	testdata.CompareFile(this.T(), this.filePathReal, this.fileBytesBom)

	// 由於linux下檔案名稱幾乎沒有非法字元, 所以這項檢查只針對windows
	if testdata.IsWindows() {
		assert.NotNil(this.T(), WriteFile(this.filePathFake, this.fileBytes, false))
	} // if
}

func (this *SuiteWrite) TestWriteJson() {
	assert.Nil(this.T(), WriteJson(this.jsonPathReal, this.jsonDatas, false))
	testdata.CompareFile(this.T(), this.jsonPathReal, this.jsonBytes)

	assert.Nil(this.T(), WriteJson(this.jsonPathReal, this.jsonDatas, true))
	testdata.CompareFile(this.T(), this.jsonPathReal, this.jsonBytesBom)

	// 由於linux下檔案名稱幾乎沒有非法字元, 所以這項檢查只針對windows
	if testdata.IsWindows() {
		assert.NotNil(this.T(), WriteJson(this.filePathFake, this.jsonDatas, false))
	} // if
}

func (this *SuiteWrite) TestWriteTmpl() {
	assert.Nil(this.T(), WriteTmpl(this.tmplPathReal, this.tmplContent, this.tmplDatas, false))
	testdata.CompareFile(this.T(), this.tmplPathReal, this.tmplBytes)

	assert.Nil(this.T(), WriteTmpl(this.tmplPathReal, this.tmplContent, this.tmplDatas, true))
	testdata.CompareFile(this.T(), this.tmplPathReal, this.tmplBytesBom)

	assert.NotNil(this.T(), WriteTmpl(this.tmplPathReal, "{{{$.Value}}", nil, false))
	assert.NotNil(this.T(), WriteTmpl(this.tmplPathReal, this.tmplContent, "nothing!", false))

	// 由於linux下檔案名稱幾乎沒有非法字元, 所以這項檢查只針對windows
	if testdata.IsWindows() {
		assert.NotNil(this.T(), WriteTmpl(this.tmplPathFake, this.tmplContent, this.tmplDatas, false))
	} // if
}
