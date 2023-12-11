// 以下是驗證程式碼, 不可使用

package sheeter

import (
	"encoding/json"
	"fmt"
)

// Handmade $結構說明
type Handmade struct {
	Pkey   int32     `json:"pkey"`   // $欄位說明
	Skey   string    `json:"skey"`   // $欄位說明
	Data1  bool      `json:"data1"`  // $欄位說明
	Data2  []bool    `json:"data2"`  // $欄位說明
	Data3  int32     `json:"data3"`  // $欄位說明
	Data4  []int32   `json:"data4"`  // $欄位說明
	Data5  int64     `json:"data5"`  // $欄位說明
	Data6  []int64   `json:"data6"`  // $欄位說明
	Data7  float32   `json:"data7"`  // $欄位說明
	Data8  []float32 `json:"data8"`  // $欄位說明
	Data9  float64   `json:"data9"`  // $欄位說明
	Data10 []float64 `json:"data10"` // $欄位說明
	Data11 string    `json:"data11"` // $欄位說明
	Data12 []string  `json:"data12"` // $欄位說明
}

// HandmadeReader $結構說明
type HandmadeReader struct {
	Data map[string]*Handmade // $欄位說明
}

// FileName 取得檔名物件
func (this *HandmadeReader) FileName() FileName {
	return NewFileName("handmade", ".json")
}

// FromData 讀取資料
func (this *HandmadeReader) FromData(data []byte, clear bool) error {
	tmpl := map[string]*Handmade{}

	if err := json.Unmarshal(data, &tmpl); err != nil {
		return fmt.Errorf("from data: %w", err)
	} // if

	if clear || this.Data == nil {
		this.Data = map[string]*Handmade{}
	} // if

	for k, v := range tmpl {
		if _, ok := this.Data[k]; ok {
			return fmt.Errorf("from data: key duplicate")
		} // if

		this.Data[k] = v
	} // for

	return nil
}

// Clear 清除資料
func (this *HandmadeReader) Clear() {
	this.Data = nil
}

// Get 取得資料
func (this *HandmadeReader) Get(key string) *Handmade {
	if result, ok := this.Data[key]; ok {
		return result
	} // if

	return nil
}

// Keys 取得索引列表
func (this *HandmadeReader) Keys() (result []string) {
	for itor := range this.Data {
		result = append(result, itor)
	} // for

	return result
}

// Values 取得資料列表
func (this *HandmadeReader) Values() (result []*Handmade) {
	for _, itor := range this.Data {
		result = append(result, itor)
	} // for

	return result
}

// Count 取得資料數量
func (this *HandmadeReader) Count() int {
	return len(this.Data)
}
