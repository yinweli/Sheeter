package core

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yinweli/Sheeter/testdata"
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

using RealDataMaps = std::map<pkey, struct RealData>;

struct RealData { 
    Sheeter::pkey name0; // note0
    bool name1; // note1
    int32_t name2; // note2
    std::string name3; // note3

    static std::string get_filename() {
        return "realData.json";
    }
};

inline json get_untyped(const json& j, const char* property) {
    return j.find(property) != j.end() ? j.at(property).get<json>() : json();
}
} // namespace Sheeter

namespace nlohmann {
inline void from_json(const json& _j, Sheeter::RealDataMaps& _x) {
    _j.get_to(_x);
}

inline void from_json(const json& _j, struct Sheeter::RealData& _x) { 
    _x.name0 = _j.at("name0").get<Sheeter::pkey>();
    _x.name1 = _j.at("name1").get<bool>();
    _x.name2 = _j.at("name2").get<int32_t>();
    _x.name3 = _j.at("name3").get<std::string>();
}

inline void to_json(json& _j, const struct Sheeter::RealData& _x) { 
    _j = json::object();
    _j["name0"] = _x.name0;
    _j["name1"] = _x.name1;
    _j["name2"] = _x.name2;
    _j["name3"] = _x.name3;
}
} // namespace nlohmann
`
}
