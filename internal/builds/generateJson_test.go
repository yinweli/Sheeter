package builds

import (
	"os"
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
	_ = os.RemoveAll(internal.PathJson)
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
using System.Collections.Generic;

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

    public partial class TestDataStorer {
        public Dictionary<long, TestData> Datas = new Dictionary<long, TestData>(); 
    }
}
`)

	target := this.target()
	assert.Nil(this.T(), generateJsonCsStruct(target))
	testdata.CompareFileByte(this.T(), target.PathJsonCsStruct(), data)
}

func (this *SuiteGenerateJson) TestGenerateJsonCsReader() {
	data := []byte(`// generated by sheeter, DO NOT EDIT.

using Newtonsoft.Json;
using System;
using System.IO;
using System.Collections.Generic;

namespace sheeter {
    public partial class TestDataReader {
        public static string FileName() {
            return "testData.json";
        }

        public bool FromPathFull(string path) {
            return FromData(File.ReadAllText(path));
        }

        public bool FromPathHalf(string path) {
            return FromData(File.ReadAllText(Path.Combine(path, FileName())));
        }

        public bool FromData(string data) {
            Datas = JsonConvert.DeserializeObject<TestDataStorer>(data);
            return Datas != null;
        }

        public TestDataStorer Datas = null;
    }
}
`)

	target := this.target()
	assert.Nil(this.T(), generateJsonCsReader(target))
	testdata.CompareFileByte(this.T(), target.PathJsonCsReader(), data)
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

type TestDataStorer struct {
	Datas map[int64]TestData
}
`)

	target := this.target()
	assert.Nil(this.T(), generateJsonGoStruct(target))
	testdata.CompareFileByte(this.T(), target.PathJsonGoStruct(), data)
}

func (this *SuiteGenerateJson) TestGenerateJsonGoReader() {
	data := []byte(`// generated by sheeter, DO NOT EDIT.

package sheeter

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type TestDataReader struct {
	TestDataStorer
}

func (this *TestDataReader) FileName() string {
	return "testData.json"
}

func (this *TestDataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("TestDataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *TestDataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("TestDataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *TestDataReader) FromData(data []byte) error {
	this.TestDataStorer = TestDataStorer{
		Datas: map[int64]TestData{},
	}

	if err := json.Unmarshal(data, &this.TestDataStorer); err != nil {
		return err
	}

	return nil
}
`)

	target := this.target()
	assert.Nil(this.T(), generateJsonGoReader(target))
	testdata.CompareFileByte(this.T(), target.PathJsonGoReader(), data)
}
