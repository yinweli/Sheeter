package core

import (
	"fmt"
	"strings"

	"github.com/yinweli/Sheeter/internal/util"
)

// ParseField 解析字串為欄位, 格式為 name#field
func ParseField(input string) (name string, field Field, err error) {
	before, after, ok := strings.Cut(input, "#") // 欄位字串以'#'符號分割為名稱與欄位

	if ok == false {
		return "", nil, fmt.Errorf("field format failed: %s", input)
	} // if

	if util.VariableCheck(before) == false {
		return "", nil, fmt.Errorf("name invalid: %s", input)
	} // if

	for _, itor := range fields {
		if itor.Type() == after {
			return before, itor, nil
		} // if
	} // for

	return "", nil, fmt.Errorf("field not found: %s", input)
}

// 實作新的欄位結構需要: 製作欄位結構, 實作欄位介面函式, 把結構加入到fields全域變數中
// 欄位資料在多執行緒環境下, 可能會造成讀寫衝突, 所以在製作新的欄位結構時, 要注意不可以在欄位結構中儲存任何資料

// Field 欄位介面
type Field interface {
	// Type 取得excel欄位類型
	Type() string

	// IsShow 是否顯示
	IsShow() bool

	// IsPkey 是否是主要索引
	IsPkey() bool

	// ToJsonDefault 轉換為json預設值
	ToJsonDefault() interface{}

	// ToJsonValue 轉換為json值
	ToJsonValue(input string) (result interface{}, err error)

	// ToLuaValue 轉換為lua值
	ToLuaValue(input string) (result string, err error)
}

// 欄位列表選擇用slice而非map, 是因為map要加入項目需要指定索引, 而Field的索引應該是Type函式
// 這會造成初始化的時候的麻煩, 加上欄位解析次數應該很少, 所以用slice對於效率的衝擊應該還好

// fields 欄位列表
var fields = []Field{
	&FieldEmpty{},
	&FieldPkey{},
	&FieldBool{},
	&FieldBoolArray{},
	&FieldInt{},
	&FieldIntArray{},
	&FieldFloat{},
	&FieldFloatArray{},
	&FieldText{},
	&FieldTextArray{},
}
