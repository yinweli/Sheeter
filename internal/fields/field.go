package fields

import (
	"fmt"
	"strings"

	"github.com/yinweli/Sheeter/internal"
)

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

	// ToTypeCs 取得cs類型字串
	ToTypeCs() string

	// ToTypeGo 取得go類型字串
	ToTypeGo() string

	// ToTypeProto 取得proto類型字串
	ToTypeProto() string

	// ToJsonValue 轉換為json值
	ToJsonValue(input string) (result interface{}, err error)
}

// 欄位列表選擇用slice而非map, 是因為map要加入項目需要指定索引, 而Field的索引應該是Type函式
// 這會造成初始化的時候的麻煩, 加上欄位解析次數應該很少, 所以用slice對於效率的衝擊應該還好

// fields 欄位列表
var fields = []Field{
	&Empty{},
	&Pkey{},
	&Bool{},
	&BoolArray{},
	&Int{},
	&IntArray{},
	&Float{},
	&FloatArray{},
	&Text{},
	&TextArray{},
}

// Parser 欄位解析, 格式為 field#tag
func Parser(input string) (field Field, tag string, err error) {
	type_, tag, _ := strings.Cut(input, internal.SeparateField)

	for _, itor := range fields {
		if itor.Type() == type_ {
			return itor, tag, nil
		} // if
	} // for

	return nil, "", fmt.Errorf("%s: parser field failed: field not found", input)
}
