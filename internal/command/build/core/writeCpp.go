package core

// templateCpp c++程式碼模板
var templateCpp string = `// generation by sheeter ^o<
// use nlohmann json library
// github: https://github.com/nlohmann/json

#pragma once

#include <stdint.h>
#include <string>
#include <vector>

#include "{{.Global.CppLibraryPath}}"

namespace {{.Global.CppNamespace}} {
using nlohmann::json;

#ifndef PKEY
#define PKEY
using pkey = int32_t;
#endif // !PKEY

struct {{.Element.Excel}}{{.Element.Sheet}} { {{range .Columns}}{{printf "\n"}}    {{.Field.TypeCpp}} {{.Name}}; // {{.Note}}{{end}}
};

inline json get_untyped(const json& j, const char* property) {
    return j.find(property) != j.end() ? j.at(property).get<json>() : json();
}
} // namespace Sheeter`
