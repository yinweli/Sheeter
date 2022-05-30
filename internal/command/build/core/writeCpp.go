package core

import (
	"fmt"

	"Sheeter/internal/util"
)

// WriteCpp 寫入c++
func WriteCpp(cargo *Cargo) (filePath string, err error) {
	bytes, err := CodeGenerate(codeCpp, cargo)

	if err != nil {
		return "", fmt.Errorf("convert cpp failed: %s [%s]", cargo.Element.GetFullName(), err)
	} // if

	_ = cargo.Progress.Add(1)
	filePath, err = util.FileWrite(OutputPathCpp, cargo.CppFileName(), bytes)

	if err != nil {
		return "", fmt.Errorf("write to cpp failed: %s [%s]", cargo.Element.GetFullName(), err)
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

struct {{structName .}} { {{len .Columns|setline}}
{{range .Columns}}    {{.Field.TypeCpp}} {{memberName .Name}}; // {{.Note}}{{newline}}{{end}}
};

inline json get_untyped(const json& j, const char* property) {
    return j.find(property) != j.end() ? j.at(property).get<json>() : json();
}
} // namespace Sheeter

namespace nlohmann {
inline void from_json(const json& _j, struct {{cppNamespace}}::{{structName .}}& _x) { {{len .Columns|setline}}
{{range .Columns}}    _x.{{memberName .Name}} = _j.at("{{memberName .Name}}").get<{{.Field.TypeCpp}}>();{{newline}}{{end}}
}

inline void to_json(json& _j, const struct {{cppNamespace}}::{{structName .}}& _x) { {{len .Columns|setline}}
    _j = json::object();
{{range .Columns}}    _j[{{memberName .Name}}] = _x.{{memberName .Name}};{{newline}}{{end}}
}
} // namespace nlohmann`
