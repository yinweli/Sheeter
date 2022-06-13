package core

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/yinweli/Sheeter/testdata"

	"github.com/stretchr/testify/assert"
)

func TestTaskJsonCpp(t *testing.T) {
	dir := testdata.ChangeWorkDir()
	defer testdata.RestoreWorkDir(dir)

	task := mockTaskJsonCpp()
	err := task.executeJsonCpp()
	assert.Nil(t, err)
	bytes, err := ioutil.ReadFile(task.jsonCppFilePath())
	assert.Nil(t, err)
	assert.Equal(t, mockTaskJsonCppString(), string(bytes[:]))
	task.close()

	task = mockTaskJsonCpp()
	task.element.Excel = "?????.xlsx"
	err = task.executeJsonCpp()
	assert.NotNil(t, err)
	task.close()

	err = os.RemoveAll(pathJsonCpp)
	assert.Nil(t, err)
}

func mockTaskJsonCpp() *Task {
	return &Task{
		global: &Global{
			CppLibraryPath: "nlohmann/json.hpp",
		},
		element: &Element{
			Excel: testdata.RealExcel,
			Sheet: testdata.SheetName,
		},
		columns: []*Column{
			{Name: "name0", Note: "note0", Field: &FieldPkey{}},
			{Name: "name1", Note: "note1", Field: &FieldBool{}},
			{Name: "name2", Note: "note2", Field: &FieldInt{}},
			{Name: "name3", Note: "note3", Field: &FieldText{}},
		},
	}
}

func mockTaskJsonCppString() string {
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
    Sheeter::pkey Name0; // note0
    bool Name1; // note1
    int32_t Name2; // note2
    std::string Name3; // note3

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
    _x.Name0 = _j.at("Name0").get<Sheeter::pkey>();
    _x.Name1 = _j.at("Name1").get<bool>();
    _x.Name2 = _j.at("Name2").get<int32_t>();
    _x.Name3 = _j.at("Name3").get<std::string>();
}

inline void to_json(json& _j, const struct Sheeter::RealData& _x) { 
    _j = json::object();
    _j["Name0"] = _x.Name0;
    _j["Name1"] = _x.Name1;
    _j["Name2"] = _x.Name2;
    _j["Name3"] = _x.Name3;
}
} // namespace nlohmann
`
}
