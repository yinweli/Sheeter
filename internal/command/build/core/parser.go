package core

import (
	"fmt"
	"strings"
)

const fieldSeparator string = "#" // 欄位分隔字串

// Parser 欄位解析器
type Parser struct {
	fields map[string]Field // 欄位列表
}

// Add 新增欄位
func (this *Parser) Add(field Field) {
	this.fields[field.TypeExcel()] = field
}

// Parse 解析欄位
func (this *Parser) Parse(input string) (name string, field Field, err error) {
	tokens := strings.Split(input, fieldSeparator)

	if len(tokens) != 2 {
		return "", nil, fmt.Errorf("parse failed: %s", input)
	} // if

	field, ok := this.fields[tokens[1]]

	if field == nil || ok == false {
		return "", nil, fmt.Errorf("field not found: %s", input)
	} // if

	return tokens[0], field, nil
}

// NewParser 建立欄位解析器
func NewParser() *Parser {
	parser := &Parser{
		fields: make(map[string]Field),
	}
	parser.Add(&FieldBool{})
	parser.Add(&FieldBoolArray{})
	parser.Add(&FieldDouble{})
	parser.Add(&FieldDoubleArray{})
	parser.Add(&FieldEmpty{})
	parser.Add(&FieldFloat{})
	parser.Add(&FieldFloatArray{})
	parser.Add(&FieldInt{})
	parser.Add(&FieldIntArray{})
	parser.Add(&FieldLong{})
	parser.Add(&FieldLongArray{})
	parser.Add(&FieldPkey{})
	parser.Add(&FieldText{})
	parser.Add(&FieldTextArray{})

	return parser
}
