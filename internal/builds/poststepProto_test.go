package builds

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/internal/layouts"
	"github.com/yinweli/Sheeter/internal/mixeds"
	"github.com/yinweli/Sheeter/testdata"
)

func TestPoststepProto(t *testing.T) {
	suite.Run(t, new(SuitePoststepProto))
}

type SuitePoststepProto struct {
	suite.Suite
	workDir string
}

func (this *SuitePoststepProto) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
}

func (this *SuitePoststepProto) TearDownSuite() {
	target := this.target()
	_ = os.Remove(target.ProtoCsBatFile())
	_ = os.Remove(target.ProtoCsShFile())
	_ = os.Remove(target.ProtoGoBatFile())
	_ = os.Remove(target.ProtoGoShFile())
	_ = os.RemoveAll(internal.ProtoPath)
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuitePoststepProto) target() *poststepData {
	target := &poststepData{
		Global: &Global{},
		Mixed:  mixeds.NewMixed("", ""),
		Struct: []poststepStruct{
			{
				Mixed: mixeds.NewMixed("test", "data"),
				Type: &layouts.Type{
					Excel:  "test",
					Sheet:  "data",
					Reader: true,
				},
			},
		},
	}
	return target
}

func (this *SuitePoststepJson) TestPoststepProtoCsDepot() {
	data := []byte(`// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

using System.Collections.Generic;

namespace SheeterProto {
    public partial class Depot {
        public Loader Loader { get; set; }
        public readonly TestDataReader TestData = new TestDataReader();
        private readonly List<Reader> Readers = new List<Reader>();

        public Depot() {
            Readers.Add(TestData);
        }

        public bool FromData() {
            if (Loader == null)
                return false;

            var result = true;

            foreach (var itor in Readers) {
                var data = Loader.Load(itor.DataName(), itor.DataExt(), itor.DataFile());

                if (data == null || data.Length == 0)
                    continue;

                var message = itor.FromData(data);

                if (message.Length != 0) {
                    result = false;
                    Loader.Error(itor.DataName(), message);
                }
            }

            return result;
        }

        public bool MergeData() {
            if (Loader == null)
                return false;

            var result = true;

            foreach (var itor in Readers) {
                var data = Loader.Load(itor.DataName(), itor.DataExt(), itor.DataFile());

                if (data == null || data.Length == 0)
                    continue;

                var message = itor.MergeData(data);

                if (message.Length != 0) {
                    result = false;
                    Loader.Error(itor.DataName(), message);
                }
            }

            return result;
        }
    }

    public interface Loader {
        public void Error(string name, string message);
        public byte[] Load(string name, string ext, string fullname);
    }

    public interface Reader {
        public string DataName();
        public string DataExt();
        public string DataFile();
        public string FromData(byte[] data);
        public string MergeData(byte[] data);
    }
}
`)

	target := this.target()
	assert.Nil(this.T(), poststepProtoCsDepot(target))
	testdata.CompareFile(this.T(), target.ProtoCsDepotPath(), data)
}

func (this *SuitePoststepProto) TestPoststepProtoGoDepot() {
	data := []byte(`// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

type Depot struct {
	TestData TestDataReader
	loader   Loader
	readers  []Reader
}

func NewDepot(loader Loader) *Depot {
	depot := &Depot{}
	depot.loader = loader
	depot.readers = append(
		depot.readers,
		&depot.TestData,
	)
	return depot
}

func (this *Depot) FromData() bool {
	if this.loader == nil {
		return false
	}

	result := true

	for _, itor := range this.readers {
		data := this.loader.Load(itor.DataName(), itor.DataExt(), itor.DataFile())

		if data == nil || len(data) == 0 {
			continue
		}

		if err := itor.FromData(data); err != nil {
			result = false
			this.loader.Error(itor.DataName(), err)
		}
	}

	return result
}

func (this *Depot) MergeData() bool {
	if this.loader == nil {
		return false
	}

	result := true

	for _, itor := range this.readers {
		data := this.loader.Load(itor.DataName(), itor.DataExt(), itor.DataFile())

		if data == nil || len(data) == 0 {
			continue
		}

		if err := itor.MergeData(data); err != nil {
			result = false
			this.loader.Error(itor.DataName(), err)
		}
	}

	return result
}

type Loader interface {
	Error(name string, err error)
	Load(name, ext, fullname string) []byte
}

type Reader interface {
	DataName() string
	DataExt() string
	DataFile() string
	FromData(data []byte) error
	MergeData(data []byte) error
}
`)

	target := this.target()
	assert.Nil(this.T(), poststepProtoGoDepot(target))
	testdata.CompareFile(this.T(), target.ProtoGoDepotPath(), data)
}

func (this *SuitePoststepProto) TestPoststepProtoCsBat() {
	proto := filepath.Join(internal.ProtoPath, internal.SchemaPath)
	code := filepath.Join(internal.ProtoPath, internal.CsPath)
	file := filepath.Join(internal.ProtoPath, internal.SchemaPath, "testData.proto")
	data := []byte(fmt.Sprintf(`REM Code generated by sheeter. DO NOT EDIT.
REM Sheeter: https://github.com/yinweli/Sheeter

mkdir %s
protoc --experimental_allow_proto3_optional --proto_path=./%s --csharp_out=./%s ./%s
`, code, proto, code, file))

	target := this.target()
	assert.Nil(this.T(), poststepProtoCsBat(target))
	testdata.CompareFile(this.T(), target.ProtoCsBatFile(), data)
}

func (this *SuitePoststepProto) TestPoststepProtoCsSh() {
	proto := filepath.Join(internal.ProtoPath, internal.SchemaPath)
	code := filepath.Join(internal.ProtoPath, internal.CsPath)
	file := filepath.Join(internal.ProtoPath, internal.SchemaPath, "testData.proto")
	data := []byte(fmt.Sprintf(`# Code generated by sheeter. DO NOT EDIT.
# Sheeter: https://github.com/yinweli/Sheeter

mkdir %s
protoc --experimental_allow_proto3_optional --proto_path=./%s --csharp_out=./%s ./%s
`, code, proto, code, file))

	target := this.target()
	assert.Nil(this.T(), poststepProtoCsSh(target))
	testdata.CompareFile(this.T(), target.ProtoCsShFile(), data)
}

func (this *SuitePoststepProto) TestPoststepProtoGoBat() {
	proto := filepath.Join(internal.ProtoPath, internal.SchemaPath)
	code := filepath.Join(internal.ProtoPath, internal.GoPath)
	file := filepath.Join(internal.ProtoPath, internal.SchemaPath, "testData.proto")
	data := []byte(fmt.Sprintf(`REM Code generated by sheeter. DO NOT EDIT.
REM Sheeter: https://github.com/yinweli/Sheeter

mkdir %s
protoc --experimental_allow_proto3_optional --proto_path=./%s --go_out=./%s ./%s
`, code, proto, code, file))

	target := this.target()
	assert.Nil(this.T(), poststepProtoGoBat(target))
	testdata.CompareFile(this.T(), target.ProtoGoBatFile(), data)
}

func (this *SuitePoststepProto) TestPoststepProtoGoSh() {
	proto := filepath.Join(internal.ProtoPath, internal.SchemaPath)
	code := filepath.Join(internal.ProtoPath, internal.GoPath)
	file := filepath.Join(internal.ProtoPath, internal.SchemaPath, "testData.proto")
	data := []byte(fmt.Sprintf(`# Code generated by sheeter. DO NOT EDIT.
# Sheeter: https://github.com/yinweli/Sheeter

mkdir %s
protoc --experimental_allow_proto3_optional --proto_path=./%s --go_out=./%s ./%s
`, code, proto, code, file))

	target := this.target()
	assert.Nil(this.T(), poststepProtoGoSh(target))
	testdata.CompareFile(this.T(), target.ProtoGoShFile(), data)
}
