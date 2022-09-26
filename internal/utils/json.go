package utils

import (
	"encoding/json"
)

const jsonPrefix = ""    // json前綴字串
const jsonIdent = "    " // json縮排字串

// JsonMarshal 把物件轉換為json字串
func JsonMarshal(value any) (results []byte, err error) {
	return json.MarshalIndent(value, jsonPrefix, jsonIdent)
}
