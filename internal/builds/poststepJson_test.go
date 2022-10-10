package builds

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/internal/layouts"
	"github.com/yinweli/Sheeter/internal/mixeds"
	"github.com/yinweli/Sheeter/testdata"
)

func TestPoststepJson(t *testing.T) {
	suite.Run(t, new(SuitePoststepJson))
}

type SuitePoststepJson struct {
	suite.Suite
	workDir string
}

func (this *SuitePoststepJson) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
}

func (this *SuitePoststepJson) TearDownSuite() {
	_ = os.RemoveAll(internal.JsonPath)
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuitePoststepJson) target() *poststepData {
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

func (this *SuitePoststepJson) TestPoststepJsonCsDepot() {
	data := []byte(`// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

using System.Collections.Generic;

namespace SheeterJson {
    public partial class Depot {
        public readonly TestDataReader TestData = new TestDataReader();
        private readonly List<ReaderInterface> Readers = new List<ReaderInterface>();
        
        public Depot() {
            Readers.Add(TestData);
        }

        public bool FromData(DelegateLoad load, DelegateError error) {
            var result = true;

            foreach (var itor in Readers) {
                var data = load(itor.DataName(), itor.DataExt());

                if (data == null || data.Length == 0)
                    continue;

                var message = itor.FromData(data);

                if (message.Length != 0) {
                    result = false;
                    error(itor.DataName(), message);
                }
            }

            return result;
        }

        public bool MergeData(DelegateLoad load, DelegateError error) {
            var result = true;

            foreach (var itor in Readers) {
                var data = load(itor.DataName(), itor.DataExt());

                if (data == null || data.Length == 0)
                    continue;

                var message = itor.MergeData(data);

                if (message.Length != 0) {
                    result = false;
                    error(itor.DataName(), message);
                }
            }

            return result;
        }

        public delegate void DelegateError(string name, string message);
        public delegate string DelegateLoad(string name, string ext);
    }

    public interface ReaderInterface {
        public string DataName();
        public string DataExt();
        public string DataFile();
        public string FromData(string data);
        public string MergeData(string data);
    }
}
`)

	target := this.target()
	assert.Nil(this.T(), poststepJsonCsDepot(target))
	testdata.CompareFile(this.T(), target.JsonCsDepotPath(), data)
}

func (this *SuitePoststepJson) TestPoststepJsonGoDepot() {
	data := []byte(`// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

type Depot struct {
	TestData TestDataReader
	readers  []ReaderInterface
}

func (this *Depot) FromData(load DepotLoad, error DepotError) bool {
	this.build()
	result := true

	for _, itor := range this.readers {
		data := load(itor.DataName(), itor.DataExt())

		if data == nil || len(data) == 0 {
			continue
		}

		if err := itor.FromData(data); err != nil {
			result = false
			error(itor.DataName(), err)
		}
	}

	return result
}

func (this *Depot) MergeData(load DepotLoad, error DepotError) bool {
	this.build()
	result := true

	for _, itor := range this.readers {
		data := load(itor.DataName(), itor.DataExt())

		if data == nil || len(data) == 0 {
			continue
		}

		if err := itor.MergeData(data); err != nil {
			result = false
			error(itor.DataName(), err)
		}
	}

	return result
}

func (this *Depot) build() {
	if len(this.readers) == 0 {
		this.readers = append(
			this.readers,
			&this.TestData,
		)
	}
}

type DepotError func(name string, err error)
type DepotLoad func(name, ext string) []byte

type ReaderInterface interface {
	DataName() string
	DataExt() string
	DataFile() string
	FromData(data []byte) error
	MergeData(data []byte) error
}
`)

	target := this.target()
	assert.Nil(this.T(), poststepJsonGoDepot(target))
	testdata.CompareFile(this.T(), target.JsonGoDepotPath(), data)
}
