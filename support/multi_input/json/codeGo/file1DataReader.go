// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
)

type File1DataReader struct {
	*File1DataStorer
}

func (this *File1DataReader) DataName() string {
	return "file1Data"
}

func (this *File1DataReader) DataExt() string {
	return ".json"
}

func (this *File1DataReader) DataFile() string {
	return "file1Data.json"
}

func (this *File1DataReader) FromData(data []byte) error {
	this.File1DataStorer = &File1DataStorer{
		Datas: map[int64]*File1Data{},
	}

	if err := json.Unmarshal(data, this.File1DataStorer); err != nil {
		return fmt.Errorf("from data failed: %w", err)
	}

	return nil
}

func (this *File1DataReader) MergeData(data []byte) error {
	tmpl := &File1DataStorer{
		Datas: map[int64]*File1Data{},
	}

	if err := json.Unmarshal(data, tmpl); err != nil {
		return fmt.Errorf("merge data failed: %w", err)
	}

	if this.File1DataStorer == nil {
		this.File1DataStorer = &File1DataStorer{
			Datas: map[int64]*File1Data{},
		}
	}

	for k, v := range tmpl.Datas {
		if _, ok := this.File1DataStorer.Datas[k]; ok {
			return fmt.Errorf("merge data failed: key repeat")
		}

		this.File1DataStorer.Datas[k] = v
	}

	return nil
}

func (this *File1DataReader) Clear() {
	this.File1DataStorer = nil
}

func (this *File1DataReader) Get(key int64) (result *File1Data, ok bool) {
	result, ok = this.Datas[key]
	return result, ok
}

func (this *File1DataReader) Keys() (result []int64) {
	for itor := range this.Datas {
		result = append(result, itor)
	}

	return result
}

func (this *File1DataReader) Values() (result []*File1Data) {
	for _, itor := range this.Datas {
		result = append(result, itor)
	}

	return result
}

func (this *File1DataReader) Count() int {
	return len(this.Datas)
}