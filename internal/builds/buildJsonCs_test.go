package builds

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/testdata"
)

func TestBuildJsonCs(t *testing.T) {
	suite.Run(t, new(SuiteBuildJsonCs))
}

type SuiteBuildJsonCs struct {
	suite.Suite
	workDir string
	code    []byte
	reader  []byte
}

func (this *SuiteBuildJsonCs) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
	this.code = []byte(`namespace realdata
{
    using System;
    using System.Collections.Generic;

    using System.Globalization;
    using Newtonsoft.Json;
    using Newtonsoft.Json.Converters;

    public partial class RealData
    {
        [JsonProperty("S")]
        public S S { get; set; }

        [JsonProperty("name0")]
        public long Name0 { get; set; }
    }

    public partial class S
    {
        [JsonProperty("A")]
        public A[] A { get; set; }

        [JsonProperty("name1")]
        public bool Name1 { get; set; }
    }

    public partial class A
    {
        [JsonProperty("name2")]
        public long Name2 { get; set; }

        [JsonProperty("name3")]
        public string Name3 { get; set; }
    }
}
`)
	this.reader = []byte(`// generated by sheeter, DO NOT EDIT.

namespace realdata {
    using System;
    using System.Collections.Generic;

    using Newtonsoft.Json;

    public partial class RealDataReader {
        public static readonly string JsonPath = "json\realData.json";

        public static Dictionary<string, RealData> FromJson(string data) {
            return JsonConvert.DeserializeObject<Dictionary<string, RealData>>(data);
        }
    }
}
`)
}

func (this *SuiteBuildJsonCs) TearDownSuite() {
	_ = os.RemoveAll(pathSchema)
	_ = os.RemoveAll(pathJsonCs)
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteBuildJsonCs) target() *Content {
	target := &Content{
		LineOfField: 1,
		LineOfLayer: 2,
		LineOfNote:  3,
		Excel:       testdata.Path(testdata.ExcelNameReal),
		Sheet:       testdata.SheetName,
	}
	return target
}

func (this *SuiteBuildJsonCs) TestWriteJsonCs() {
	target := this.target()
	target.excel = testdata.GetTestExcel(testdata.ExcelNameReal)
	assert.Nil(this.T(), buildLayout(target))
	assert.Nil(this.T(), writeSchema(target))
	assert.Nil(this.T(), writeJsonCs(target))
	testdata.CompareFile(this.T(), target.JsonCsPath(), this.code)
	target.close()

	target = this.target()
	target.excel = testdata.GetTestExcel(testdata.ExcelNameReal)
	assert.Nil(this.T(), buildLayout(target))
	assert.Nil(this.T(), writeSchema(target))
	target.Excel = testdata.UnknownStr
	assert.NotNil(this.T(), writeJsonCs(target))
	target.close()

	target = this.target()
	target.excel = testdata.GetTestExcel(testdata.ExcelNameReal)
	assert.Nil(this.T(), buildLayout(target))
	assert.Nil(this.T(), writeSchema(target))
	target.Sheet = testdata.UnknownStr
	assert.NotNil(this.T(), writeJsonCs(target))
	target.close()
}

func (this *SuiteBuildJsonCs) TestWriteJsonCsReader() {
	target := this.target()
	target.excel = testdata.GetTestExcel(testdata.ExcelNameReal)
	assert.Nil(this.T(), buildLayout(target))
	assert.Nil(this.T(), writeSchema(target))
	assert.Nil(this.T(), writeJsonCsReader(target))
	testdata.CompareFile(this.T(), target.JsonCsReaderPath(), this.reader)
	target.close()

	target = this.target()
	target.excel = testdata.GetTestExcel(testdata.ExcelNameReal)
	assert.Nil(this.T(), buildLayout(target))
	assert.Nil(this.T(), writeSchema(target))
	target.Excel = testdata.UnknownStr
	assert.NotNil(this.T(), writeJsonCsReader(target))
	target.close()

	target = this.target()
	target.excel = testdata.GetTestExcel(testdata.ExcelNameReal)
	assert.Nil(this.T(), buildLayout(target))
	assert.Nil(this.T(), writeSchema(target))
	target.Sheet = testdata.UnknownStr
	assert.NotNil(this.T(), writeJsonCsReader(target))
	target.close()
}