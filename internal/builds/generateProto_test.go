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

func TestGenerateProto(t *testing.T) {
	suite.Run(t, new(SuiteGenerateProto))
}

type SuiteGenerateProto struct {
	suite.Suite
	workDir string
}

func (this *SuiteGenerateProto) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
}

func (this *SuiteGenerateProto) TearDownSuite() {
	_ = os.RemoveAll(internal.ProtoPath)
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteGenerateProto) target() *RuntimeStruct {
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
	assert.Nil(this.T(), generateProtoSchema(target))
	testdata.CompareFile(this.T(), target.ProtoPath(), data1)

	target = this.target()
	target.Reader = false
	assert.Nil(this.T(), generateProtoSchema(target))
	testdata.CompareFile(this.T(), target.ProtoPath(), data2)
}

func (this *SuiteGenerateProto) TestGenerateProtoCsReader() {
	data := []byte(`// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

using System.Collections.Generic;

namespace SheeterProto {
    public partial class TestDataReader {
        public string DataName() {
            return "testData";
        }

        public string DataExt() {
            return "bytes";
        }

        public string DataFile() {
            return "testData.bytes";
        }

        public bool FromData(byte[] data) {
            Datas = TestDataStorer.Parser.ParseFrom(data);
            return Datas != null;
        }

        public long[] MergeData(byte[] data) {
            var repeats = new List<long>();
            var tmpl = TestDataStorer.Parser.ParseFrom(data);

            if (tmpl == null)
                return repeats.ToArray();

            if (Datas == null)
                Datas = new TestDataStorer();

            foreach (var itor in tmpl.Datas) {
                if (Data.ContainsKey(itor.Key) == false)
                    Data[itor.Key] = itor.Value;
                else
                    repeats.Add(itor.Key);
            }

            return repeats.ToArray();
        }

        public IDictionary<long, TestData> Data {
            get {
                return Datas.Datas;
            }
        }

        private TestDataStorer Datas = null;
    }
}
`)

	target := this.target()
	assert.Nil(this.T(), generateProtoCsReader(target))
	testdata.CompareFile(this.T(), target.ProtoCsReaderPath(), data)
}

func (this *SuiteGenerateProto) TestGenerateProtoGoReader() {
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
	return "bytes"
}

func (this *TestDataReader) DataFile() string {
	return "testData.bytes"
}

func (this *TestDataReader) FromData(data []byte) error {
	this.TestDataStorer = &TestDataStorer{
		Datas: map[int64]*TestData{},
	}

	if err := proto.Unmarshal(data, this.TestDataStorer); err != nil {
		return fmt.Errorf("TestDataReader: from data failed: %w", err)
	}

	return nil
}

func (this *TestDataReader) MergeData(data []byte) (repeats []int64) {
	tmpl := &TestDataStorer{
		Datas: map[int64]*TestData{},
	}

	if err := proto.Unmarshal(data, tmpl); err != nil {
		return repeats
	}

	if this.TestDataStorer == nil {
		this.TestDataStorer = &TestDataStorer{
			Datas: map[int64]*TestData{},
		}
	}

	for k, v := range tmpl.Datas {
		if _, ok := this.TestDataStorer.Datas[k]; ok == false {
			this.TestDataStorer.Datas[k] = v
		} else {
			repeats = append(repeats, k)
		}
	}

	return repeats
}
`)

	target := this.target()
	assert.Nil(this.T(), generateProtoGoReader(target))
	testdata.CompareFile(this.T(), target.ProtoGoReaderPath(), data)
}
