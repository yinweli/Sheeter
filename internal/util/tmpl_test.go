package util

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/testdata"
)

func TestTmpl(t *testing.T) {
	suite.Run(t, new(SuiteTmpl))
}

type SuiteTmpl struct {
	suite.Suite
	workDir      string
	filePathReal string
	filePathFake string
	tmplContent  string
	tmplDatas    map[string]string
	tmplBytes    []byte
	tmplBytesBom []byte
}

func (this *SuiteTmpl) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
	this.filePathReal = "tmpl/test.tmpl"
	this.filePathFake = "?tmpl/test.tmpl"
	this.tmplContent = "{{$.Value}}"
	this.tmplDatas = map[string]string{"Value": "Value"}
	this.tmplBytes = []byte("Value")
	this.tmplBytesBom = bomPrefix
	this.tmplBytesBom = append(this.tmplBytesBom, this.tmplBytes...)
}

func (this *SuiteTmpl) TearDownSuite() {
	_ = os.RemoveAll(filepath.Dir(this.filePathReal))
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteTmpl) TestTmplWrite() {
	assert.Nil(this.T(), TmplWrite(this.filePathReal, this.tmplContent, this.tmplDatas, false))
	testdata.CompareFile(this.T(), this.filePathReal, this.tmplBytes)

	assert.Nil(this.T(), TmplWrite(this.filePathReal, this.tmplContent, this.tmplDatas, true))
	testdata.CompareFile(this.T(), this.filePathReal, this.tmplBytesBom)

	assert.NotNil(this.T(), TmplWrite(this.filePathReal, "{{{$.Value}}", nil, false))
	assert.NotNil(this.T(), TmplWrite(this.filePathReal, this.tmplContent, "nothing!", false))
	assert.NotNil(this.T(), TmplWrite(this.filePathFake, this.tmplContent, this.tmplDatas, false))
}
