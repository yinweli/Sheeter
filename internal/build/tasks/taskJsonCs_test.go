package tasks

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/yinweli/Sheeter/internal/build/fields"
	"github.com/yinweli/Sheeter/testdata"
)

func TestTaskJsonCs(t *testing.T) {
	suite.Run(t, new(SuiteTaskJsonCs))
}

type SuiteTaskJsonCs struct {
	suite.Suite
	workDir   string
	dataBytes []byte
}

func (this *SuiteTaskJsonCs) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
	this.dataBytes = []byte(`namespace sheeter
{
    using System;
    using System.Collections.Generic;

    using System.Globalization;
    using Newtonsoft.Json;
    using Newtonsoft.Json.Converters;

    public partial class RealData
    {
        [JsonProperty("name0")]
        public long Name0 { get; set; }

        [JsonProperty("name1")]
        public bool Name1 { get; set; }

        [JsonProperty("name2")]
        public long Name2 { get; set; }

        [JsonProperty("name3")]
        public string Name3 { get; set; }
    }
}
`)
}

func (this *SuiteTaskJsonCs) TearDownSuite() {
	_ = os.RemoveAll(pathSchema)
	_ = os.RemoveAll(pathJsonCs)
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteTaskJsonCs) target() *Task {
	target := NewTask(nil, nil)
	target.global = &Global{}
	target.element = &Element{
		Excel: testdata.RealExcel,
		Sheet: testdata.SheetName,
	}
	target.columns = []*Column{
		{Name: "name0", Note: "note0", Field: &fields.FieldPkey{}},
		{Name: "name1", Note: "note1", Field: &fields.FieldBool{}},
		{Name: "name2", Note: "note2", Field: &fields.FieldInt{}},
		{Name: "name3", Note: "note3", Field: &fields.FieldText{}},
	}
	return target
}

func (this *SuiteTaskJsonCs) TestTaskJsonCs() {
	target := this.target()
	assert.Nil(this.T(), target.runJsonSchema())
	assert.Nil(this.T(), target.runJsonCs())
	testdata.CompareFile(this.T(), target.jsonCsFilePath(), this.dataBytes)
	target.close()

	target = this.target()
	target.element.Excel = testdata.UnknownStr
	assert.NotNil(this.T(), target.runJsonCs())
	target.close()

	target = this.target()
	target.element.Sheet = testdata.UnknownStr
	assert.NotNil(this.T(), target.runJsonCs())
	target.close()
}
