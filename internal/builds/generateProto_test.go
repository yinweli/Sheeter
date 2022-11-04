package builds

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/internal/fields"
	"github.com/yinweli/Sheeter/internal/layouts"
	"github.com/yinweli/Sheeter/internal/nameds"
	"github.com/yinweli/Sheeter/testdata"
)

func TestGenerateProto(t *testing.T) {
	suite.Run(t, new(SuiteGenerateProto))
}

type SuiteGenerateProto struct {
	suite.Suite
	testdata.TestEnv
}

func (this *SuiteGenerateProto) SetupSuite() {
	this.Change("test-generateProto")
}

func (this *SuiteGenerateProto) TearDownSuite() {
	this.Restore()
}

func (this *SuiteGenerateProto) target() *generateProto {
	const excelName = "test"
	const sheetName = "data"

	target := &generateProto{
		Global: &Global{},
		Named:  &nameds.Named{ExcelName: excelName, SheetName: sheetName},
		Field:  &nameds.Field{},
		Proto:  &nameds.Proto{ExcelName: excelName, SheetName: sheetName},
		Reader: true,
		Fields: []*layouts.Field{
			{Name: "name1", Note: "note1", Field: &fields.Pkey{}, Alter: "", Array: false},
			{Name: "name2", Note: "note2", Field: &fields.Int{}, Alter: "", Array: false},
			{Name: "name3", Note: "note3", Field: &fields.IntArray{}, Alter: "", Array: false},
			{Name: "name4", Note: "note4", Field: nil, Alter: "Data", Array: false},
			{Name: "name5", Note: "note5", Field: nil, Alter: "Data", Array: true},
		},
		Depend: []string{"depend1", "depend2", "depend3"},
	}
	return target
}

func (this *SuiteGenerateProto) TestGenerateProtoSchema() {
	data1 := []byte(`// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

syntax = "proto3";
package sheeterProto;
option go_package = ".;sheeterProto";
import 'depend1.proto';
import 'depend2.proto';
import 'depend3.proto';

message TestData {
  optional int64 Name1 = 1; // note1
  optional int64 Name2 = 2; // note2
  repeated int64 Name3 = 3; // note3
  optional Data Name4 = 4; // note4
  repeated Data Name5 = 5; // note5
}

message TestDataStorer {
  map<int64, TestData> Datas = 1;
}
`)
	data2 := []byte(`// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

syntax = "proto3";
package sheeterProto;
option go_package = ".;sheeterProto";
import 'depend1.proto';
import 'depend2.proto';
import 'depend3.proto';

message TestData {
  optional int64 Name1 = 1; // note1
  optional int64 Name2 = 2; // note2
  repeated int64 Name3 = 3; // note3
  optional Data Name4 = 4; // note4
  repeated Data Name5 = 5; // note5
}
`)

	target := this.target()
	assert.Nil(this.T(), GenerateProtoSchema(target, nil))
	testdata.CompareFile(this.T(), target.ProtoPath(), data1)

	target = this.target()
	target.Reader = false
	assert.Nil(this.T(), GenerateProtoSchema(target, nil))
	testdata.CompareFile(this.T(), target.ProtoPath(), data2)

	assert.Nil(this.T(), GenerateProtoSchema(nil, nil))
}

func (this *SuiteGenerateProto) TestGenerateProtoReaderCs() {
	data := []byte(`// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

using System.Collections.Generic;

namespace SheeterProto {
    using Data_ = TestData;
    using PKey_ = System.Int64;
    using Storer_ = TestDataStorer;

    public partial class TestDataReader : Reader {
        public string DataName() {
            return "testData";
        }

        public string DataExt() {
            return ".bytes";
        }

        public string DataFile() {
            return "testData.bytes";
        }

        public string FromData(byte[] data) {
            Storer_ result;

            try {
                result = Storer_.Parser.ParseFrom(data);
            } catch {
                return "from data failed: deserialize failed";
            }

            if (result == null)
                return "from data failed: result null";

            storer = result;
            return string.Empty;
        }

        public string MergeData(byte[] data) {
            Storer_ result;

            try {
                result = Storer_.Parser.ParseFrom(data);
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
	assert.Nil(this.T(), GenerateProtoReaderCs(target, nil))
	testdata.CompareFile(this.T(), target.ProtoReaderCsPath(), data)

	target = this.target()
	target.Reader = false
	assert.Nil(this.T(), GenerateProtoReaderCs(target, nil))

	assert.Nil(this.T(), GenerateProtoReaderCs(nil, nil))
}

func (this *SuiteGenerateProto) TestGenerateProtoReaderGo() {
	data := []byte(`// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"

	"google.golang.org/protobuf/proto"
)

type TestDataReader struct {
	*TestDataStorer
}

func (this *TestDataReader) DataName() string {
	return "testData"
}

func (this *TestDataReader) DataExt() string {
	return ".bytes"
}

func (this *TestDataReader) DataFile() string {
	return "testData.bytes"
}

func (this *TestDataReader) FromData(data []byte) error {
	this.TestDataStorer = &TestDataStorer{
		Datas: map[int64]*TestData{},
	}

	if err := proto.Unmarshal(data, this.TestDataStorer); err != nil {
		return fmt.Errorf("from data failed: %w", err)
	}

	return nil
}

func (this *TestDataReader) MergeData(data []byte) error {
	tmpl := &TestDataStorer{
		Datas: map[int64]*TestData{},
	}

	if err := proto.Unmarshal(data, tmpl); err != nil {
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
	assert.Nil(this.T(), GenerateProtoReaderGo(target, nil))
	testdata.CompareFile(this.T(), target.ProtoReaderGoPath(), data)

	target = this.target()
	target.Reader = false
	assert.Nil(this.T(), GenerateProtoReaderGo(target, nil))

	assert.Nil(this.T(), GenerateProtoReaderGo(nil, nil))
}
