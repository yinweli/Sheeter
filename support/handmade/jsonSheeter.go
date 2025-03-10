// 以下是模板驗證程式碼

package sheeter

// NewSheeter 建立表格資料
func NewSheeter(loader Loader) *Sheeter {
	sheeter := &Sheeter{}
	sheeter.loader = loader
	return sheeter
}

// Sheeter 表格資料
type Sheeter struct {
	loader Loader         // 裝載器物件
	Alone0 HandmadeReader // 獨立表格說明
	Alone1 HandmadeReader // 獨立表格說明
	Merge0 HandmadeReader // 合併表格說明
	Merge1 HandmadeReader // 合併表格說明
}

// FromData 讀取資料處理
func (this *Sheeter) FromData() bool {
	if this.loader == nil {
		return false
	} // if

	result := true

	for _, itor := range []Reader{
		&this.Alone0,
		&this.Alone1,
	} {
		filename := itor.FileName()
		data := this.loader.Load(filename)

		if len(data) == 0 {
			continue
		} // if

		if err := itor.FromData(data, true); err != nil {
			result = false
			this.loader.Error(filename.File(), err)
		} // if
	} // for

	for i, itor := range []Reader{
		&this.Alone0,
		&this.Alone1,
	} {
		filename := itor.FileName()
		data := this.loader.Load(filename)

		if len(data) == 0 {
			continue
		} // if

		if err := this.Merge0.FromData(data, i == 0); err != nil {
			result = false
			this.loader.Error("Merge0", err)
		} // if
	} // for

	for i, itor := range []Reader{
		&this.Alone0,
		&this.Alone1,
	} {
		filename := itor.FileName()
		data := this.loader.Load(filename)

		if len(data) == 0 {
			continue
		} // if

		if err := this.Merge1.FromData(data, i == 0); err != nil {
			result = false
			this.loader.Error("Merge1", err)
		} // if
	} // for

	return result
}

// Clear 清除資料
func (this *Sheeter) Clear() {
	this.Alone0.Clear()
	this.Alone1.Clear()
	this.Merge0.Clear()
	this.Merge1.Clear()
}

// Loader 裝載器介面
type Loader interface {
	// Load 讀取檔案
	Load(filename FileName) []byte

	// Error 錯誤處理
	Error(name string, err error)
}

// Reader 讀取器介面
type Reader interface {
	// FileName 取得檔名物件
	FileName() FileName

	// FromData 讀取資料
	FromData(data []byte, clear bool) error

	// Clear 清除資料
	Clear()
}

// NewFileName 建立檔名資料
func NewFileName(name, ext string) FileName {
	return FileName{
		name: name,
		ext:  ext,
	}
}

// FileName 檔名資料
type FileName struct {
	name string // 名稱
	ext  string // 副檔名
}

// Name 取得名稱
func (this FileName) Name() string {
	return this.name
}

// Ext 取得副檔名
func (this FileName) Ext() string {
	return this.ext
}

// File 取得完整檔名
func (this FileName) File() string {
	return this.name + this.ext
}
