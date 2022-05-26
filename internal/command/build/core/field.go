package core

import (
	"fmt"
	"strings"

	"Sheeter/internal"
)

/***** 欄位相關 *****/

// Field 欄位介面
type Field interface {
	// TypeExcel 取得excel欄位類型
	TypeExcel() string

	// TypeCpp 取得c++欄位類型
	TypeCpp() string

	// TypeCs 取得c#欄位類型
	TypeCs() string

	// TypeGo 取得go欄位類型
	TypeGo() string

	// Hide 是否隱藏
	Hide() bool

	// PrimaryKey 是否是主要索引
	PrimaryKey() bool

	// Transform 字串轉換
	Transform(input string) (result interface{}, err error)
}

// ParseField 解析欄位
func ParseField(input string) (name string, field Field, err error) {
	tokens := strings.Split(input, internal.FieldSeparator)

	if len(tokens) != 2 {
		return "", nil, fmt.Errorf("parse failed: %s", input)
	} // if

	field, ok := fields[tokens[1]]

	if field == nil || ok == false {
		return "", nil, fmt.Errorf("field not found: %s", input)
	} // if

	return tokens[0], field, nil
}

/***** 欄位列表 *****/

var fields = make(map[string]Field) // 欄位列表

// addField 新增欄位到欄位列表
func addField(field Field) {
	fields[field.TypeExcel()] = field
}

/***** 初始執行 *****/

func init() {
	// TODO: 要記得到這裡新增欄位到欄位列表

	addField(&FieldBool{})
	addField(&FieldBoolArray{})
	addField(&FieldDouble{})
	addField(&FieldDoubleArray{})
	addField(&FieldEmpty{})
	addField(&FieldFloat{})
	addField(&FieldFloatArray{})
	addField(&FieldInt{})
	addField(&FieldIntArray{})
	addField(&FieldLong{})
	addField(&FieldLongArray{})
	addField(&FieldPkey{})
	addField(&FieldText{})
	addField(&FieldTextArray{})
}
