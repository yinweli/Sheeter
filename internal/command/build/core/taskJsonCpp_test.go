package core

import (
	"io/ioutil"
	"os"
	"testing"

	"Sheeter/internal/util"
	"Sheeter/testdata"

	"github.com/stretchr/testify/assert"
)

func TestTaskJsonCpp(t *testing.T) {
	dir := testdata.ChangeWorkDir()
	defer testdata.RestoreWorkDir(dir)

	ctx := mockTaskJsonCppContext()
	err := TaskJsonCpp(ctx)
	assert.Nil(t, err)
	assert.FileExists(t, ctx.JsonCppFilePath())

	bytes, err := ioutil.ReadFile(ctx.JsonCppFilePath())
	assert.Nil(t, err)
	assert.Equal(t, mockTaskJsonCppString(), string(bytes[:]))

	ctx = mockTaskJsonCppContext()
	ctx.Element.Excel = "?????.xlsx"
	err = TaskJsonCpp(ctx)
	assert.NotNil(t, err)

	err = os.RemoveAll(PathJsonCpp)
	assert.Nil(t, err)
}

func mockTaskJsonCppContext() *Context {
	return &Context{
		Progress: util.NewProgress(0, "", ioutil.Discard),
		Global: &Global{
			CppLibraryPath: "nlohmann/json.hpp",
		},
		Element: &Element{
			Excel: "excel.xlsx",
			Sheet: "sheet",
		},
		Columns: []*Column{
			{Note: "note0", Name: "name0", Field: &FieldPkey{}},
			{Note: "note1", Name: "name1", Field: &FieldBool{}},
			{Note: "note2", Name: "name2", Field: &FieldText{}},
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

struct ExcelSheet { 
    Sheeter::pkey Name0; // note0
    bool Name1; // note1
    std::string Name2; // note2

    static std::string get_filename() {
        return "excelSheet.json"
    }
};

inline json get_untyped(const json& j, const char* property) {
    return j.find(property) != j.end() ? j.at(property).get<json>() : json();
}
} // namespace Sheeter

namespace nlohmann {
inline void from_json(const json& _j, struct Sheeter::ExcelSheet& _x) { 
    _x.Name0 = _j.at("Name0").get<Sheeter::pkey>();
    _x.Name1 = _j.at("Name1").get<bool>();
    _x.Name2 = _j.at("Name2").get<std::string>();
}

inline void to_json(json& _j, const struct Sheeter::ExcelSheet& _x) { 
    _j = json::object();
    _j["Name0"] = _x.Name0;
    _j["Name1"] = _x.Name1;
    _j["Name2"] = _x.Name2;
}
} // namespace nlohmann
`
}
