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
	_ = os.RemoveAll(internal.PathProto)
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
	data := []byte(`// generated by sheeter, DO NOT EDIT.

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

	target := this.target()
	assert.Nil(this.T(), generateProtoSchema(target))
	testdata.CompareFileByte(this.T(), target.PathProtoName(), data)
}

func (this *SuiteGenerateProto) TestGenerateProtoCsReader() {
	data := []byte(`// generated by sheeter, DO NOT EDIT.

using System.IO;
using System.Collections.Generic;

namespace SheeterProto {
    public partial class TestDataReader {
        public static string FileName() {
            return "testData.pbd";
        }

        public bool FromPathFull(string path) {
            return FromData(File.ReadAllBytes(path));
        }

        public bool FromPathHalf(string path) {
            return FromData(File.ReadAllBytes(Path.Combine(path, FileName())));
        }

        public bool FromData(byte[] data) {
            Datas = TestDataStorer.Parser.ParseFrom(data);
            return Datas != null;
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
	testdata.CompareFileByte(this.T(), target.PathProtoCsReader(), data)
}
