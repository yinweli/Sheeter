package core

import (
	"fmt"

	"Sheeter/internal/util"
)

// WriteCpp 寫入c++
type WriteCpp struct {
}

// LongName 取得長名稱
func (this *WriteCpp) LongName() string {
	return "cpp"
}

// ShortName 取得短名稱
func (this *WriteCpp) ShortName() string {
	return "c"
}

// Note 取得註解
func (this *WriteCpp) Note() string {
	return "generate cpp file"
}

// Progress 取得進度值
func (this *WriteCpp) Progress(sheetSize int) int {
	return 2
}

// Execute 執行工作
func (this *WriteCpp) Execute(cargo *Cargo) (filePath string, err error) {
	cargo.Progress.Add(1)
	bytes, err := CodeGenerate(codeCpp, cargo)

	if err != nil {
		return "", fmt.Errorf("convert cpp failed: %s [%s]", cargo.LogName(), err)
	} // if

	cargo.Progress.Add(1)
	filePath, err = util.FileWrite(OutputPathCpp, cargo.CppFileName(), bytes)

	if err != nil {
		return "", fmt.Errorf("write to cpp failed: %s [%s]", cargo.LogName(), err)
	} // if

	return filePath, nil
}

// codeCpp c++程式碼模板
var codeCpp string = `// generation by sheeter ^o<
// use nlohmann json library
// github: https://github.com/nlohmann/json

#pragma once

#include <stdint.h>
#include <string>
#include <vector>

#include "{{.Global.CppLibraryPath}}"

namespace {{cppNamespace}} {
using nlohmann::json;

#ifndef PKEY
#define PKEY
using pkey = int32_t;
#endif // !PKEY

struct {{.StructName}} { {{setline .Columns}}
{{range .Columns}}{{if .Field.Show}}    {{.Field.TypeCpp}} {{.MemberName}}; // {{.Note}}{{newline}}{{end}}{{end}}

    static std::string get_filename() {
        return "{{.JsonFileName}}"
    }
};

inline json get_untyped(const json& j, const char* property) {
    return j.find(property) != j.end() ? j.at(property).get<json>() : json();
}
} // namespace {{cppNamespace}}

namespace nlohmann {
inline void from_json(const json& _j, struct {{cppNamespace}}::{{.StructName}}& _x) { {{setline .Columns}}
{{range .Columns}}{{if .Field.Show}}    _x.{{.MemberName}} = _j.at("{{.MemberName}}").get<{{.Field.TypeCpp}}>();{{newline}}{{end}}{{end}}
}

inline void to_json(json& _j, const struct {{cppNamespace}}::{{.StructName}}& _x) { {{setline .Columns}}
    _j = json::object();
{{range .Columns}}{{if .Field.Show}}    _j["{{.MemberName}}"] = _x.{{.MemberName}};{{newline}}{{end}}{{end}}
}
} // namespace nlohmann`
