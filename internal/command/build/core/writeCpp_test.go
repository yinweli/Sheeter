package core

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"Sheeter/testdata"

	"github.com/stretchr/testify/assert"
)

func TestWriteCpp(t *testing.T) {
	dir := testdata.ChangeWorkDir()
	defer testdata.RestoreWorkDir(dir)

	writeCpp := &WriteCpp{}
	assert.Equal(t, "cpp", writeCpp.LongName())
	assert.Equal(t, "c", writeCpp.ShortName())
	assert.Equal(t, "generate cpp file", writeCpp.Note())
	assert.Equal(t, 2, writeCpp.Progress(0))

	cargo := mockWriteCppCargo()
	path, err := writeCpp.Execute(cargo)
	assert.Nil(t, err)
	assert.Equal(t, filepath.Join(OutputPathCpp, "realData.hpp"), path)
	assert.FileExists(t, path)

	bytes, err := ioutil.ReadFile(path)
	assert.Nil(t, err)
	assert.Equal(t, mockWriteCppString(), string(bytes[:]))

	err = os.RemoveAll(OutputPathCpp)
	assert.Nil(t, err)
}

func TestWriteCppFailed(t *testing.T) {
	dir := testdata.ChangeWorkDir()
	defer testdata.RestoreWorkDir(dir)

	writeCpp := &WriteCpp{}
	cargo := mockWriteCppCargo()
	cargo.Global = nil
	_, err := writeCpp.Execute(cargo)
	assert.NotNil(t, err)

	err = os.RemoveAll(OutputPathCpp)
	assert.Nil(t, err)
}

func mockWriteCppCargo() *Cargo {
	return &Cargo{
		Progress: NewProgress(0, "test", ioutil.Discard),
		Global: &Global{
			CppLibraryPath: "nlohmann/json.hpp",
		},
		Element: &Element{
			Excel: "real.xlsx",
			Sheet: "data",
		},
		Columns: []*Column{
			{Note: "note0", Name: "name0", Field: &FieldInt{}},
			{Note: "note1", Name: "name1", Field: &FieldBool{}},
			{Note: "note2", Name: "name2", Field: &FieldText{}},
		},
	}
}

func mockWriteCppString() string {
	return `// generation by sheeter ^o<
// use nlohmann json library
// github: https://github.com/nlohmann/json

#pragma once

#include <stdint.h>
#include <string>
#include <vector>

#include "nlohmann/json.hpp"

namespace Sheeter {
using nlohmann::json;

#ifndef PKEY
#define PKEY
using pkey = int32_t;
#endif // !PKEY

struct RealData { 
    int32_t Name0; // note0
    bool Name1; // note1
    std::string Name2; // note2

    static std::string get_filename() {
        return "realData.json"
    }
};

inline json get_untyped(const json& j, const char* property) {
    return j.find(property) != j.end() ? j.at(property).get<json>() : json();
}
} // namespace Sheeter

namespace nlohmann {
inline void from_json(const json& _j, struct Sheeter::RealData& _x) { 
    _x.Name0 = _j.at("Name0").get<int32_t>();
    _x.Name1 = _j.at("Name1").get<bool>();
    _x.Name2 = _j.at("Name2").get<std::string>();
}

inline void to_json(json& _j, const struct Sheeter::RealData& _x) { 
    _j = json::object();
    _j["Name0"] = _x.Name0;
    _j["Name1"] = _x.Name1;
    _j["Name2"] = _x.Name2;
}
} // namespace nlohmann`
}
