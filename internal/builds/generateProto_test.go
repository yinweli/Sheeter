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
	testdata.CompareFile(this.T(), target.PathProtoName(), data1)

	target = this.target()
	target.Reader = false
	assert.Nil(this.T(), generateProtoSchema(target))
	testdata.CompareFile(this.T(), target.PathProtoName(), data2)
}

func (this *SuiteGenerateProto) TestGenerateProtoCsReader() {
	data := []byte(`// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

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
	testdata.CompareFile(this.T(), target.PathProtoCsReader(), data)
}

func (this *SuiteGenerateProto) TestGenerateProtoGoReader() {
	data := []byte(`// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type TestDataReader struct {
	TestDataStorer
}

func (this *TestDataReader) FileName() string {
	return "testData.pbd"
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
	if err := proto.Unmarshal(data, &this.TestDataStorer); err != nil {
		return err
	}

	return nil
}
`)

	target := this.target()
	assert.Nil(this.T(), generateProtoGoReader(target))
	testdata.CompareFile(this.T(), target.PathProtoGoReader(), data)
}
