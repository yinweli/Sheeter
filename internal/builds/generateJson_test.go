package builds

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/internal/fields"
	"github.com/yinweli/Sheeter/internal/layouts"
	"github.com/yinweli/Sheeter/internal/mixeds"
	"github.com/yinweli/Sheeter/testdata"
)

func TestGenerateJson(t *testing.T) {
	suite.Run(t, new(SuiteGenerateJson))
}

type SuiteGenerateJson struct {
	suite.Suite
	workDir string
}

func (this *SuiteGenerateJson) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
}

func (this *SuiteGenerateJson) TearDownSuite() {
	_ = os.RemoveAll(filepath.Join(internal.PathJson, internal.PathCs))
	_ = os.RemoveAll(filepath.Join(internal.PathJson, internal.PathGo))
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteGenerateJson) target() *RuntimeStruct {
	target := &RuntimeStruct{
		Mixed: mixeds.NewMixed("test", "data"),
		Type: &layouts.Type{
			Excel:  "test",
			Sheet:  "data",
			Reader: true,
			Fields: []*layouts.Field{
				{Name: "name1", Note: "note1", Field: &fields.Pkey{}, Alter: "", Array: false},
				{Name: "name2", Note: "note2", Field: &fields.Int{}, Alter: "", Array: false},
				{Name: "name3", Note: "note3", Field: &fields.IntArray{}, Alter: "", Array: false},
				{Name: "name4", Note: "note4", Field: nil, Alter: "Data", Array: false},
				{Name: "name5", Note: "note5", Field: nil, Alter: "Data", Array: true},
			},
		},
	}
	return target
}

func (this *SuiteGenerateJson) TestGenerateJsonCsStruct() {
	data := []byte(`// generated by sheeter, DO NOT EDIT.

using Newtonsoft.Json;

namespace sheeter {
    public partial class TestData {
        // note1
        [JsonProperty("Name1")]
        public long Name1 { get; set; }
        // note2
        [JsonProperty("Name2")]
        public long Name2 { get; set; }
        // note3
        [JsonProperty("Name3")]
        public long[] Name3 { get; set; }
        // note4
        [JsonProperty("Name4")]
        public Data Name4 { get; set; }
        // note5
        [JsonProperty("Name5")]
        public Data[] Name5 { get; set; }
    }
}
`)

	target := this.target()
	assert.Nil(this.T(), generateJsonCsStruct(target))
	testdata.CompareFile(this.T(), target.FileJsonCsStruct(), data)
}

func (this *SuiteGenerateJson) TestGenerateJsonCsReader() {
	data := []byte(`// generated by sheeter, DO NOT EDIT.

using Newtonsoft.Json;
using System;
using System.IO;
using System.Collections.Generic;

namespace sheeter {
    public partial class TestDataReader {
        public static readonly string Json = "testData.json";

        public static Dictionary<long, TestData> FromJsonFile(string path) {
            return FromJsonString(File.ReadAllText(path));
        }

        public static Dictionary<long, TestData> FromJsonString(string data) {
            var datas = JsonConvert.DeserializeObject<Dictionary<long, TestData>>(data);
            return datas;
        }
    }
}
`)

	target := this.target()
	assert.Nil(this.T(), generateJsonCsReader(target))
	testdata.CompareFile(this.T(), target.FileJsonCsReader(), data)
}

func (this *SuiteGenerateJson) TestGenerateJsonGoStruct() {
	data := []byte(`// generated by sheeter, DO NOT EDIT.

package sheeter

type TestData struct {
	// note1
	Name1 int64 ` + "`json:\"Name1\"`" + `
	// note2
	Name2 int64 ` + "`json:\"Name2\"`" + `
	// note3
	Name3 []int64 ` + "`json:\"Name3\"`" + `
	// note4
	Name4 Data ` + "`json:\"Name4\"`" + `
	// note5
	Name5 []Data ` + "`json:\"Name5\"`" + `
}
`)

	target := this.target()
	assert.Nil(this.T(), generateJsonGoStruct(target))
	testdata.CompareFile(this.T(), target.FileJsonGoStruct(), data)
}

func (this *SuiteGenerateJson) TestGenerateJsonGoReader() {
	data := []byte(`// generated by sheeter, DO NOT EDIT.

package sheeter

import (
	"encoding/json"
	"fmt"
	"os"
)

type TestDataReader struct {
	Datas map[int64]TestData
}

func (this *TestDataReader) Json() string {
	return "testData.json"
}

func (this *TestDataReader) FromJsonFile(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("TestDataReader: from json file failed: %w", err)
	}

	return this.FromJsonBytes(data)
}

func (this *TestDataReader) FromJsonBytes(data []byte) error {
	datas := map[int64]TestData{}

	if err := json.Unmarshal(data, &datas); err != nil {
		return err
	}

	this.Datas = datas
	return nil
}
`)

	target := this.target()
	assert.Nil(this.T(), generateJsonGoReader(target))
	testdata.CompareFile(this.T(), target.FileJsonGoReader(), data)
}
