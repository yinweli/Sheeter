// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"

	"google.golang.org/protobuf/proto"
)

type Path1DataReader struct {
	*Path1DataStorer
}

func (this *Path1DataReader) DataName() string {
	return "path1Data"
}

func (this *Path1DataReader) DataExt() string {
	return ".bytes"
}

func (this *Path1DataReader) DataFile() string {
	return "path1Data.bytes"
}

func (this *Path1DataReader) FromData(data []byte) error {
	this.Path1DataStorer = &Path1DataStorer{
		Datas: map[int64]*Path1Data{},
	}

	if err := proto.Unmarshal(data, this.Path1DataStorer); err != nil {
		return fmt.Errorf("from data failed: %w", err)
	}

	return nil
}

func (this *Path1DataReader) MergeData(data []byte) error {
	tmpl := &Path1DataStorer{
		Datas: map[int64]*Path1Data{},
	}

	if err := proto.Unmarshal(data, tmpl); err != nil {
		return fmt.Errorf("merge data failed: %w", err)
	}

	if this.Path1DataStorer == nil {
		this.Path1DataStorer = &Path1DataStorer{
			Datas: map[int64]*Path1Data{},
		}
	}

	for k, v := range tmpl.Datas {
		if _, ok := this.Path1DataStorer.Datas[k]; ok {
			return fmt.Errorf("merge data failed: key repeat")
		}

		this.Path1DataStorer.Datas[k] = v
	}

	return nil
}

func (this *Path1DataReader) Clear() {
	this.Path1DataStorer = nil
}

func (this *Path1DataReader) Get(key int64) (result *Path1Data, ok bool) {
	result, ok = this.Datas[key]
	return result, ok
}

func (this *Path1DataReader) Keys() (result []int64) {
	for itor := range this.Datas {
		result = append(result, itor)
	}

	return result
}

func (this *Path1DataReader) Values() (result []*Path1Data) {
	for _, itor := range this.Datas {
		result = append(result, itor)
	}

	return result
}

func (this *Path1DataReader) Count() int {
	return len(this.Datas)
}
