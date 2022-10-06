package mixeds

import (
	"path/filepath"

	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/internal/utils"
)

// Proto proto綜合工具
type Proto struct {
	mixed *Mixed // 綜合工具
}

// ProtoCsPath 取得proto-cs路徑
func (this *Proto) ProtoCsPath() string {
	return filepath.Join(internal.PathProto, internal.PathCs)
}

// ProtoGoPath 取得proto-go路徑
func (this *Proto) ProtoGoPath() string {
	return filepath.Join(internal.PathProto, internal.PathGo)
}

// ProtoSchemaPath 取得proto-schema路徑
func (this *Proto) ProtoSchemaPath() string {
	return filepath.Join(internal.PathProto, internal.PathSchema)
}

// ProtoName 取得proto架構檔名
func (this *Proto) ProtoName() string {
	return this.mixed.combine(params{
		sheetUpper: true,
		ext:        internal.ExtProtoSchema,
	})
}

// ProtoPath 取得proto架構路徑
func (this *Proto) ProtoPath() string {
	return filepath.Join(this.ProtoSchemaPath(), this.ProtoName())
}

// ProtoDataName 取得proto資料名稱
func (this *Proto) ProtoDataName() string {
	return this.mixed.combine(params{
		sheetUpper: true,
	})
}

// ProtoDataExt 取得proto資料副檔名
func (this *Proto) ProtoDataExt() string {
	return internal.ExtProtoData
}

// ProtoDataFile 取得proto資料檔名
func (this *Proto) ProtoDataFile() string {
	return this.mixed.combine(params{
		sheetUpper: true,
		ext:        internal.ExtProtoData,
	})
}

// ProtoDataPath 取得proto資料路徑
func (this *Proto) ProtoDataPath() string {
	return filepath.Join(internal.PathProto, internal.PathData, this.ProtoDataFile())
}

// ProtoCsReaderPath 取得proto-cs讀取器程式碼路徑
func (this *Proto) ProtoCsReaderPath() string {
	return filepath.Join(internal.PathProto, internal.PathCs, this.mixed.combine(params{
		excelUpper: true, // 因為protoc產生出來的cs結構程式碼檔名是大寫開頭, 所以cs讀取器名稱也用大寫開頭
		sheetUpper: true,
		last:       internal.Reader,
		ext:        internal.ExtCs,
	}))
}

// ProtoGoReaderPath 取得proto-go讀取器程式碼路徑
func (this *Proto) ProtoGoReaderPath() string {
	return filepath.Join(internal.PathProto, internal.PathGo, this.mixed.combine(params{
		sheetUpper: true, // 因為protoc產生出來的go結構程式碼檔名是小寫開頭, 所以go讀取器名稱也用小寫開頭
		last:       internal.Reader,
		ext:        internal.ExtGo,
	}))
}

// ProtoDepend 取得proto依賴檔名
func (this *Proto) ProtoDepend(name string) string {
	// proto依賴檔名必須跟已建立的proto檔名相符
	// 因為proto檔名是小寫駝峰, 所以這裡也必須是小寫駝峰

	return utils.FirstLower(name) + "." + internal.ExtProtoSchema
}
