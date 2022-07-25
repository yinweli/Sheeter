package util

import (
	"os"
	"path"
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
	_ = os.RemoveAll(path.Dir(this.filePathReal))
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

func TestTmplLine(t *testing.T) {
	suite.Run(t, new(SuiteTmplLine))
}

type SuiteTmplLine struct {
	suite.Suite
	lineEmpty string
	lineNew   string
}

func (this *SuiteTmplLine) SetupSuite() {
	this.lineEmpty = ""
	this.lineNew = "\n"
}

func (this *SuiteTmplLine) target() *TmplLine {
	return &TmplLine{}
}

func (this *SuiteTmplLine) TestTmplLine() {
	target := this.target()

	assert.Equal(this.T(), this.lineEmpty, target.SetLine(2))
	assert.Equal(this.T(), this.lineNew, target.NewLine())
	assert.Equal(this.T(), this.lineNew, target.NewLine())
	assert.Equal(this.T(), this.lineEmpty, target.NewLine())
	assert.Equal(this.T(), this.lineEmpty, target.NewLine())
}

func TestTmplFirstChar(t *testing.T) {
	suite.Run(t, new(SuiteTmplFirstChar))
}

type SuiteTmplFirstChar struct {
	suite.Suite
	firstUpper string
	firstLower string
}

func (this *SuiteTmplFirstChar) SetupSuite() {
	this.firstUpper = "TestData"
	this.firstLower = "testData"
}

func (this *SuiteTmplFirstChar) target() *TmplFirstChar {
	return &TmplFirstChar{}
}

func (this *SuiteTmplFirstChar) TestFirstUpper() {
	assert.Equal(this.T(), this.firstUpper, this.target().FirstUpper(this.firstLower))
}

func (this *SuiteTmplFirstChar) TestFirstLower() {
	assert.Equal(this.T(), this.firstLower, this.target().FirstLower(this.firstUpper))
}
