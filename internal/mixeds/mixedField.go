package mixeds

import (
	"strings"

	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/internal/fields"
	"github.com/yinweli/Sheeter/internal/layouts"
	"github.com/yinweli/Sheeter/internal/utils"
)

// Field 欄位綜合工具
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

	if field.Array && strings.HasSuffix(fieldName, internal.TokenArray) == false {
		fieldArray = internal.TokenArray
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

	if field.Array && strings.HasPrefix(fieldName, internal.TokenArray) == false {
		fieldArray = internal.TokenArray
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

	if field.Array == false && strings.HasPrefix(fieldName, internal.TokenRepeated) == false {
		fieldPrefix = internal.TokenOptional
	} else if strings.HasPrefix(fieldName, internal.TokenRepeated) == false {
		fieldPrefix = internal.TokenRepeated
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
