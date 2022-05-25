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

// NewFields 建立欄位列表
func NewFields() []Field {
	return []Field{
		&Bool{},
		&BoolArray{},
		&Double{},
		&DoubleArray{},
		&Empty{},
		&Float{},
		&FloatArray{},
		&Int{},
		&IntArray{},
		&Long{},
		&LongArray{},
		&Pkey{},
		&Text{},
		&TextArray{},
	}
}
