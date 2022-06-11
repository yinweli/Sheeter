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

#include "{{$.CppLibraryPath}}"

namespace {{$.CppNamespace}} {
using nlohmann::json;

#ifndef PKEY
#define PKEY
using pkey = int32_t;
#endif // !PKEY

struct {{.StructName}} { {{$.SetLine}}
{{range .Columns}}{{if .Field.IsShow}}    {{.Field.TypeCpp}} {{$.ColumnName .Name}}; // {{.Note}}{{$.NewLine}}{{end}}{{end}}

    static std::string get_filename() {
        return "{{$.JsonFileName}}"
    }
};

inline json get_untyped(const json& j, const char* property) {
    return j.find(property) != j.end() ? j.at(property).get<json>() : json();
}
} // namespace {{$.CppNamespace}}

namespace nlohmann {
inline void from_json(const json& _j, struct {{$.CppNamespace}}::{{$.StructName}}& _x) { {{$.SetLine}}
{{range .Columns}}{{if .Field.IsShow}}    _x.{{$.ColumnName .Name}} = _j.at("{{$.ColumnName .Name}}").get<{{.Field.TypeCpp}}>();{{$.NewLine}}{{end}}{{end}}
}

inline void to_json(json& _j, const struct {{$.CppNamespace}}::{{$.StructName}}& _x) { {{$.SetLine}}
    _j = json::object();
{{range .Columns}}{{if .Field.IsShow}}    _j["{{$.ColumnName .Name}}"] = _x.{{$.ColumnName .Name}};{{$.NewLine}}{{end}}{{end}}
}
} // namespace nlohmann
`

// executeJsonCpp 輸出json/c++
func (this *Task) executeJsonCpp() error {
	bytes, err := NewCoder(this.columns, this.global.CppLibraryPath, this.jsonFileName(), this.structName()).Generate(jsonCppCode)

	if err != nil {
		return fmt.Errorf("generate cpp failed: %s [%s]", this.logName(), err)
	} // if

	err = util.FileWrite(this.jsonCppFilePath(), bytes, this.global.Bom)

	if err != nil {
		return fmt.Errorf("write to cpp failed: %s [%s]", this.logName(), err)
	} // if

	if this.bar != nil {
		this.bar.IncrBy(taskProgressM)
	} // if

	return nil
}
