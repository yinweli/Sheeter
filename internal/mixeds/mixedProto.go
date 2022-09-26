package mixeds

import (
	"path/filepath"

	"github.com/yinweli/Sheeter/internal"
)

// Proto proto綜合工具
type Proto struct {
	mixed *Mixed // 綜合工具
}

// PathProtoCs 取得proto-cs路徑
func (this *Proto) PathProtoCs() string {
	return filepath.Join(internal.PathProto, internal.PathCs)
}

// PathProtoGo 取得proto-go路徑
func (this *Proto) PathProtoGo() string {
	return filepath.Join(internal.PathProto, internal.PathGo)
}

// PathProtoSchema 取得proto-schema路徑
func (this *Proto) PathProtoSchema() string {
	return filepath.Join(internal.PathProto, internal.PathSchema)
}

// FileProtoName 取得proto架構檔名
func (this *Proto) FileProtoName() string {
	return this.mixed.combine(params{
		sheetUpper: true,
		ext:        internal.ExtProtoSchema,
	})
}

// PathProtoName 取得proto架構路徑
func (this *Proto) PathProtoName() string {
	return filepath.Join(this.PathProtoSchema(), this.FileProtoName())
}

// FileProtoData 取得proto資料檔名
func (this *Proto) FileProtoData() string {
	return this.mixed.combine(params{
		sheetUpper: true,
		ext:        internal.ExtProtoData,
	})
}

// PathProtoData 取得proto資料路徑
func (this *Proto) PathProtoData() string {
	return filepath.Join(internal.PathProto, internal.PathData, this.FileProtoData())
}

// PathProtoCsReader 取得proto-cs讀取器程式碼路徑
func (this *Proto) PathProtoCsReader() string {
	return filepath.Join(internal.PathProto, internal.PathCs, this.mixed.combine(params{
		sheetUpper: true,
		last:       internal.Reader,
		ext:        internal.ExtCs,
	}))
}

// PathProtoGoReader 取得proto-go讀取器程式碼路徑
func (this *Proto) PathProtoGoReader() string {
	return filepath.Join(internal.PathProto, internal.PathGo, this.mixed.combine(params{
		sheetUpper: true,
		last:       internal.Reader,
		ext:        internal.ExtGo,
	}))
}

// ProtoDepend 取得proto依賴檔案名稱
func (this *Proto) ProtoDepend(name string) string {
	return name + "." + internal.ExtProtoSchema
}
