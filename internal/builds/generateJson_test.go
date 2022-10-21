package builds

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/internal/fields"
	"github.com/yinweli/Sheeter/internal/layouts"
	"github.com/yinweli/Sheeter/internal/nameds"
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
	_ = os.RemoveAll(internal.JsonPath)
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteGenerateJson) target() *generateJson {
	const excelName = "test"
	const sheetName = "data"

	target := &generateJson{
		Global: &Global{},
		Named:  &nameds.Named{ExcelName: excelName, SheetName: sheetName},
		Field:  &nameds.Field{},
		Json:   &nameds.Json{ExcelName: excelName, SheetName: sheetName},
		Reader: true,
		Fields: []*layouts.Field{
			{Name: "name1", Note: "note1", Field: &fields.Pkey{}, Alter: "", Array: false},
			{Name: "name2", Note: "note2", Field: &fields.Int{}, Alter: "", Array: false},
			{Name: "name3", Note: "note3", Field: &fields.IntArray{}, Alter: "", Array: false},
			{Name: "name4", Note: "note4", Field: nil, Alter: "Data", Array: false},
			{Name: "name5", Note: "note5", Field: nil, Alter: "Data", Array: true},
		},
	}
	return target
}

func (this *SuiteGenerateJson) TestGenerateJsonStructCs() {
	data1 := []byte(`// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

using Newtonsoft.Json;
using System.Collections.Generic;

namespace SheeterJson {
    public partial class TestData {
        // note1
        [JsonProperty("Name1")]
        public System.Int64 Name1 { get; set; }
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
        public Dictionary<System.Int64, TestData> Datas = new Dictionary<System.Int64, TestData>(); 
    }
}
`)
	data2 := []byte(`// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

using Newtonsoft.Json;
using System.Collections.Generic;

namespace SheeterJson {
    public partial class TestData {
        // note1
        [JsonProperty("Name1")]
        public System.Int64 Name1 { get; set; }
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
	assert.Nil(this.T(), GenerateJsonStructCs(target))
	testdata.CompareFile(this.T(), target.JsonStructCsPath(), data1)

	target = this.target()
	target.Reader = false
	assert.Nil(this.T(), GenerateJsonStructCs(target))
	testdata.CompareFile(this.T(), target.JsonStructCsPath(), data2)

	assert.Nil(this.T(), GenerateJsonStructCs(nil))
}

func (this *SuiteGenerateJson) TestGenerateJsonReaderCs() {
	data := []byte(`// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

using Newtonsoft.Json;
using System.Collections.Generic;

namespace SheeterJson {
    using Data_ = TestData;
    using PKey_ = System.Int64;
    using Storer_ = TestDataStorer;

    public partial class TestDataReader : Reader {
        public string DataName() {
            return "testData";
        }

        public string DataExt() {
            return "json";
        }

        public string DataFile() {
            return "testData.json";
        }

        public string FromData(string data) {
            Storer_ result;

            try {
                result = JsonConvert.DeserializeObject<Storer_>(data);
            } catch {
                return "from data failed: deserialize failed";
            }

            if (result == null)
                return "from data failed: result null";

            storer = result;
            return string.Empty;
        }

        public string MergeData(string data) {
            Storer_ result;

            try {
                result = JsonConvert.DeserializeObject<Storer_>(data);
            } catch {
                return "merge data failed: deserialize failed";
            }

            if (result == null)
                return "merge data failed: result null";

            foreach (var itor in result.Datas) {
                if (storer.Datas.ContainsKey(itor.Key))
                    return "merge data failed: key repeat";

                storer.Datas[itor.Key] = itor.Value;
            }

            return string.Empty;
        }

        public void Clear() {
            storer.Datas.Clear();
        }

        public bool TryGetValue(PKey_ key, out Data_ value) {
            return storer.Datas.TryGetValue(key, out value);
        }

        public bool ContainsKey(PKey_ key) {
            return storer.Datas.ContainsKey(key);
        }

        public IEnumerator<KeyValuePair<PKey_, Data_>> GetEnumerator() {
            return storer.Datas.GetEnumerator();
        }

        public Data_ this[PKey_ key] {
            get {
                return storer.Datas[key];
            }
        }

        public ICollection<PKey_> Keys {
            get {
                return storer.Datas.Keys;
            }
        }

        public ICollection<Data_> Values {
            get {
                return storer.Datas.Values;
            }
        }

        public int Count {
            get {
                return storer.Datas.Count;
            }
        }

        private Storer_ storer = new Storer_();
    }
}
`)

	target := this.target()
	assert.Nil(this.T(), GenerateJsonReaderCs(target))
	testdata.CompareFile(this.T(), target.JsonReaderCsPath(), data)

	target = this.target()
	target.Reader = false
	assert.Nil(this.T(), GenerateJsonReaderCs(target))

	assert.Nil(this.T(), GenerateJsonReaderCs(nil))
}

func (this *SuiteGenerateJson) TestGenerateJsonStructGo() {
	data1 := []byte(`// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

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
	Datas map[int64]*TestData
}
`)
	data2 := []byte(`// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

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
	assert.Nil(this.T(), GenerateJsonStructGo(target))
	testdata.CompareFile(this.T(), target.JsonStructGoPath(), data1)

	target = this.target()
	target.Reader = false
	assert.Nil(this.T(), GenerateJsonStructGo(target))
	testdata.CompareFile(this.T(), target.JsonStructGoPath(), data2)

	assert.Nil(this.T(), GenerateJsonStructGo(nil))
}

func (this *SuiteGenerateJson) TestGenerateJsonReaderGo() {
	data := []byte(`// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
)

type TestDataReader struct {
	*TestDataStorer
}

func (this *TestDataReader) DataName() string {
	return "testData"
}

func (this *TestDataReader) DataExt() string {
	return "json"
}

func (this *TestDataReader) DataFile() string {
	return "testData.json"
}

func (this *TestDataReader) FromData(data []byte) error {
	this.TestDataStorer = &TestDataStorer{
		Datas: map[int64]*TestData{},
	}

	if err := json.Unmarshal(data, this.TestDataStorer); err != nil {
		return fmt.Errorf("from data failed: %w", err)
	}

	return nil
}

func (this *TestDataReader) MergeData(data []byte) error {
	tmpl := &TestDataStorer{
		Datas: map[int64]*TestData{},
	}

	if err := json.Unmarshal(data, tmpl); err != nil {
		return fmt.Errorf("merge data failed: %w", err)
	}

	if this.TestDataStorer == nil {
		this.TestDataStorer = &TestDataStorer{
			Datas: map[int64]*TestData{},
		}
	}

	for k, v := range tmpl.Datas {
		if _, ok := this.TestDataStorer.Datas[k]; ok {
			return fmt.Errorf("merge data failed: key repeat")
		}

		this.TestDataStorer.Datas[k] = v
	}

	return nil
}

func (this *TestDataReader) Clear() {
	this.TestDataStorer = nil
}

func (this *TestDataReader) Get(key int64) (result *TestData, ok bool) {
	result, ok = this.Datas[key]
	return result, ok
}

func (this *TestDataReader) Keys() (result []int64) {
	for itor := range this.Datas {
		result = append(result, itor)
	}

	return result
}

func (this *TestDataReader) Values() (result []*TestData) {
	for _, itor := range this.Datas {
		result = append(result, itor)
	}

	return result
}

func (this *TestDataReader) Count() int {
	return len(this.Datas)
}
`)

	target := this.target()
	assert.Nil(this.T(), GenerateJsonReaderGo(target))
	testdata.CompareFile(this.T(), target.JsonReaderGoPath(), data)

	target = this.target()
	target.Reader = false
	assert.Nil(this.T(), GenerateJsonReaderGo(target))

	assert.Nil(this.T(), GenerateJsonReaderGo(nil))
}
