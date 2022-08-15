package tasks

import (
	"os"
	"testing"

	"github.com/yinweli/Sheeter/testdata"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestTaskJsonCsReader(t *testing.T) {
	suite.Run(t, new(SuiteTaskJsonCsReader))
}

type SuiteTaskJsonCsReader struct {
	suite.Suite
	workDir   string
	dataBytes []byte
}

func (this *SuiteTaskJsonCsReader) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
	this.dataBytes = []byte(`// generated by sheeter, DO NOT EDIT.

namespace sheeter {
    using System;
    using System.Collections.Generic;

    using Newtonsoft.Json;

    public partial class RealDataReader {
		public static readonly string JsonFileName = "realData.json";

        public static Dictionary<string, RealData> FromJson(string data) {
            return JsonConvert.DeserializeObject<Dictionary<string, RealData>>(data);
        }
    }
}
`)
}

func (this *SuiteTaskJsonCsReader) TearDownSuite() {
	_ = os.RemoveAll(pathJsonCs)
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteTaskJsonCsReader) target() *Task {
	target := &Task{
		Excel: testdata.RealExcel,
		Sheet: testdata.SheetName,
	}
	return target
}

func (this *SuiteTaskJsonCsReader) TestJsonCsReader() {
	target := this.target()
	assert.Nil(this.T(), target.jsonCsReader())
	testdata.CompareFile(this.T(), target.jsonCsReaderFilePath(), this.dataBytes)
	target.close()

	target = this.target()
	target.Excel = testdata.UnknownStr
	assert.NotNil(this.T(), target.jsonCsReader())
	target.close()

	target = this.target()
	target.Sheet = testdata.UnknownStr
	assert.NotNil(this.T(), target.jsonCsReader())
	target.close()
}
