// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"

	"google.golang.org/protobuf/proto"
)

type File2DataReader struct {
	*File2DataStorer
}

func (this *File2DataReader) DataName() string {
	return "file2Data"
}

func (this *File2DataReader) DataExt() string {
	return ".bytes"
}

func (this *File2DataReader) DataFile() string {
	return "file2Data.bytes"
}

func (this *File2DataReader) FromData(data []byte) error {
	this.File2DataStorer = &File2DataStorer{
		Datas: map[int64]*File2Data{},
	}

	if err := proto.Unmarshal(data, this.File2DataStorer); err != nil {
		return fmt.Errorf("from data failed: %w", err)
	}

	return nil
}

func (this *File2DataReader) MergeData(data []byte) error {
	tmpl := &File2DataStorer{
		Datas: map[int64]*File2Data{},
	}

	if err := proto.Unmarshal(data, tmpl); err != nil {
		return fmt.Errorf("merge data failed: %w", err)
	}

	if this.File2DataStorer == nil {
		this.File2DataStorer = &File2DataStorer{
			Datas: map[int64]*File2Data{},
		}
	}

	for k, v := range tmpl.Datas {
		if _, ok := this.File2DataStorer.Datas[k]; ok {
			return fmt.Errorf("merge data failed: key repeat")
		}

		this.File2DataStorer.Datas[k] = v
	}

	return nil
}

func (this *File2DataReader) Clear() {
	this.File2DataStorer = nil
}

func (this *File2DataReader) Get(key int64) (result *File2Data, ok bool) {
	result, ok = this.Datas[key]
	return result, ok
}

func (this *File2DataReader) Keys() (result []int64) {
	for itor := range this.Datas {
		result = append(result, itor)
	}

	return result
}

func (this *File2DataReader) Values() (result []*File2Data) {
	for _, itor := range this.Datas {
		result = append(result, itor)
	}

	return result
}

func (this *File2DataReader) Count() int {
	return len(this.Datas)
}
