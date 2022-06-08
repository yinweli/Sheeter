package core

import (
	"fmt"

	"Sheeter/internal/util"
)

// jsonCppCode json/c++程式碼模板
const jsonCppCode = `// generation by sheeter ^o<
// use nlohmann json library
// github: https://github.com/nlohmann/json

#pragma once

#include <stdint.h>
#include <string>
#include <vector>

#include "{{.Global.CppLibraryPath}}"

namespace {{.CppNamespace}} {
using nlohmann::json;

#ifndef PKEY
#define PKEY
using pkey = int32_t;
#endif // !PKEY

struct {{.StructName}} { {{setline .Columns}}
{{range .Columns}}{{if .Field.IsShow}}    {{.Field.TypeCpp}} {{.ColumnName}}; // {{.Note}}{{newline}}{{end}}{{end}}

    static std::string get_filename() {
        return "{{.JsonFileName}}"
    }
};

inline json get_untyped(const json& j, const char* property) {
    return j.find(property) != j.end() ? j.at(property).get<json>() : json();
}
} // namespace {{.CppNamespace}}

namespace nlohmann {
inline void from_json(const json& _j, struct {{.CppNamespace}}::{{.StructName}}& _x) { {{setline .Columns}}
{{range .Columns}}{{if .Field.IsShow}}    _x.{{.ColumnName}} = _j.at("{{.ColumnName}}").get<{{.Field.TypeCpp}}>();{{newline}}{{end}}{{end}}
}

inline void to_json(json& _j, const struct {{.CppNamespace}}::{{.StructName}}& _x) { {{setline .Columns}}
    _j = json::object();
{{range .Columns}}{{if .Field.IsShow}}    _j["{{.ColumnName}}"] = _x.{{.ColumnName}};{{newline}}{{end}}{{end}}
}
} // namespace nlohmann
`

// TaskJsonCpp 輸出json/c++
func TaskJsonCpp(ctx *Context) error {
	bytes, err := Coder(jsonCppCode, ctx)

	if err != nil {
		return fmt.Errorf("generate cpp failed: %s [%s]", ctx.LogName(), err)
	} // if

	err = util.FileWrite(ctx.JsonCppFilePath(), bytes)

	if err != nil {
		return fmt.Errorf("write to cpp failed: %s [%s]", ctx.LogName(), err)
	} // if

	_ = ctx.Progress.Add(1)
	return nil
}
