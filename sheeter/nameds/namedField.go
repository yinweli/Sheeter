package nameds

import (
	"strings"

	"github.com/yinweli/Sheeter/sheeter"
	"github.com/yinweli/Sheeter/sheeter/fields"
	"github.com/yinweli/Sheeter/sheeter/layouts"
	"github.com/yinweli/Sheeter/sheeter/utils"
)

// Field 欄位命名工具
type Field struct {
}

// FieldName 取得欄位名稱
func (this *Field) FieldName(field *layouts.Field) string {
	return utils.FirstUpper(field.Name)
}

// FieldNote 取得欄位註解
func (this *Field) FieldNote(field *layouts.Field) string {
	return field.Note
}

// FieldTypeCs 取得cs欄位類型
func (this *Field) FieldTypeCs(field *layouts.Field) string {
	fieldName := ""

	if field.Field != nil {
		fieldName = field.Field.ToTypeCs()
	} else {
		fieldName = field.Alter
	} // if

	fieldArray := ""

	if field.Array && strings.HasSuffix(fieldName, sheeter.TokenArray) == false {
		fieldArray = sheeter.TokenArray
	} // if

	return fieldName + fieldArray
}

// FieldTypeGo 取得go欄位類型
func (this *Field) FieldTypeGo(field *layouts.Field) string {
	fieldName := ""

	if field.Field != nil {
		fieldName = field.Field.ToTypeGo()
	} else {
		fieldName = field.Alter
	} // if

	fieldArray := ""

	if field.Array && strings.HasPrefix(fieldName, sheeter.TokenArray) == false {
		fieldArray = sheeter.TokenArray
	} // if

	return fieldArray + fieldName
}

// FieldTypeProto 取得proto欄位類型
func (this *Field) FieldTypeProto(field *layouts.Field) string {
	fieldName := ""

	if field.Field != nil {
		fieldName = field.Field.ToTypeProto()
	} else {
		fieldName = field.Alter
	} // if

	fieldPrefix := ""

	if field.Array == false && strings.HasPrefix(fieldName, sheeter.TokenRepeated) == false {
		fieldPrefix = sheeter.TokenOptional
	} else if strings.HasPrefix(fieldName, sheeter.TokenRepeated) == false {
		fieldPrefix = sheeter.TokenRepeated
	} // if

	if len(fieldPrefix) > 0 {
		return fieldPrefix + " " + fieldName
	} else {
		return fieldName
	} // if
}

// PkeyTypeCs 取得pkey的cs類型
func (this *Field) PkeyTypeCs() string {
	pkey := fields.Pkey{}
	return pkey.ToTypeCs()
}

// PkeyTypeGo 取得pkey的go類型
func (this *Field) PkeyTypeGo() string {
	pkey := fields.Pkey{}
	return pkey.ToTypeGo()
}

// PkeyTypeProto 取得pkey的proto類型
func (this *Field) PkeyTypeProto() string {
	pkey := fields.Pkey{}
	return pkey.ToTypeProto()
}
