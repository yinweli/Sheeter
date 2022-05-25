package field

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

// Fields 欄位列表型態
type Fields map[string]Field

// NewFields 建立欄位列表
func NewFields() Fields {
	fields := make(Fields)

	addFields(fields, &Bool{})
	addFields(fields, &BoolArray{})
	addFields(fields, &Double{})
	addFields(fields, &DoubleArray{})
	addFields(fields, &Empty{})
	addFields(fields, &Float{})
	addFields(fields, &FloatArray{})
	addFields(fields, &Int{})
	addFields(fields, &IntArray{})
	addFields(fields, &Long{})
	addFields(fields, &LongArray{})
	addFields(fields, &Pkey{})
	addFields(fields, &Text{})
	addFields(fields, &TextArray{})

	return fields
}

// addFields 新增欄位
func addFields(fields Fields, field Field) {
	fields[field.TypeExcel()] = field
}
