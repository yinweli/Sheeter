package utils

import (
	"encoding/json"

	"github.com/yinweli/Sheeter/sheeter"
)

// JsonMarshal 把物件轉換為json字串
func JsonMarshal(value any) (result []byte, err error) {
	return json.MarshalIndent(value, sheeter.JsonPrefix, sheeter.JsonIdent)
}
