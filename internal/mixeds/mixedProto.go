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

// ProtoDepend 取得proto依賴檔名
func (this *Proto) ProtoDepend(name string) string {
	// proto依賴檔名必須跟已建立的proto檔名相符
	// 因為proto檔名是小寫駝峰, 所以這裡也必須是小寫駝峰

	return utils.FirstLower(name) + "." + internal.ExtProtoSchema
}
