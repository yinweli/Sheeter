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
	return filepath.Join(internal.ProtoPath, internal.CsPath)
}

// ProtoGoPath 取得proto-go路徑
func (this *Proto) ProtoGoPath() string {
	return filepath.Join(internal.ProtoPath, internal.GoPath)
}

// ProtoSchemaPath 取得proto-schema路徑
func (this *Proto) ProtoSchemaPath() string {
	return filepath.Join(internal.ProtoPath, internal.SchemaPath)
}

// ProtoName 取得proto架構檔名
func (this *Proto) ProtoName() string {
	return this.mixed.combine(params{
		sheetUpper: true,
		ext:        internal.ProtoSchemaExt,
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
	return internal.ProtoDataExt
}

// ProtoDataFile 取得proto資料檔名
func (this *Proto) ProtoDataFile() string {
	return this.mixed.combine(params{
		sheetUpper: true,
		ext:        internal.ProtoDataExt,
	})
}

// ProtoDataPath 取得proto資料路徑
func (this *Proto) ProtoDataPath() string {
	return filepath.Join(internal.ProtoPath, internal.DataPath, this.ProtoDataFile())
}

// ProtoCsReaderPath 取得proto-cs讀取器程式碼路徑
func (this *Proto) ProtoCsReaderPath() string {
	return filepath.Join(internal.ProtoPath, internal.CsPath, this.mixed.combine(params{
		excelUpper: true, // 因為protoc產生出來的檔名為大寫開頭, 所以這裡也用大寫開頭
		sheetUpper: true,
		last:       internal.Reader,
		ext:        internal.CsExt,
	}))
}

// ProtoCsDepotPath 取得proto-cs倉庫程式碼路徑
func (this *Proto) ProtoCsDepotPath() string {
	return filepath.Join(internal.ProtoPath, internal.CsPath, utils.FirstUpper(internal.Depot)+"."+internal.CsExt) // 因為protoc產生出來的檔名為大寫開頭, 所以這裡也用大寫開頭
}

// ProtoGoReaderPath 取得proto-go讀取器程式碼路徑
func (this *Proto) ProtoGoReaderPath() string {
	return filepath.Join(internal.ProtoPath, internal.GoPath, this.mixed.combine(params{
		sheetUpper: true, // 因為protoc產生出來的檔名為大寫開頭, 所以這裡也用大寫開頭
		last:       internal.Reader,
		ext:        internal.GoExt,
	}))
}

// ProtoGoDepotPath 取得proto-go倉庫程式碼路徑
func (this *Proto) ProtoGoDepotPath() string {
	return filepath.Join(internal.ProtoPath, internal.GoPath, utils.FirstUpper(internal.Depot)+"."+internal.GoExt) // 因為protoc產生出來的檔名為大寫開頭, 所以這裡也用大寫開頭
}

// ProtoCsBatFile 取得proto-cs-bat檔名
func (this *Proto) ProtoCsBatFile() string {
	return internal.ProtoCsBatFile
}

// ProtoCsShFile 取得proto-cs-sh檔名
func (this *Proto) ProtoCsShFile() string {
	return internal.ProtoCsShFile
}

// ProtoGoBatFile 取得proto-go-bat檔名
func (this *Proto) ProtoGoBatFile() string {
	return internal.ProtoGoBatFile
}

// ProtoGoShFile 取得proto-go-sh檔名
func (this *Proto) ProtoGoShFile() string {
	return internal.ProtoGoShFile
}

// ProtoDepend 取得proto依賴檔名
func (this *Proto) ProtoDepend(name string) string {
	// proto依賴檔名必須跟已建立的proto檔名相符
	// 因為proto檔名是小寫駝峰, 所以這裡也必須是小寫駝峰

	return utils.FirstLower(name) + "." + internal.ProtoSchemaExt
}
