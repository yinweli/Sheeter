// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
)

type ExampleDataReader struct {
	*ExampleDataStorer
}

func (this *ExampleDataReader) DataName() string {
	return "exampleData"
}

func (this *ExampleDataReader) DataExt() string {
	return "json"
}

func (this *ExampleDataReader) DataFile() string {
	return "exampleData.json"
}

func (this *ExampleDataReader) FromData(data []byte) error {
	this.ExampleDataStorer = &ExampleDataStorer{
		Datas: map[int64]*ExampleData{},
	}

	if err := json.Unmarshal(data, this.ExampleDataStorer); err != nil {
		return fmt.Errorf("from data failed: %w", err)
	}

	return nil
}

func (this *ExampleDataReader) MergeData(data []byte) error {
	tmpl := &ExampleDataStorer{
		Datas: map[int64]*ExampleData{},
	}

	if err := json.Unmarshal(data, tmpl); err != nil {
		return fmt.Errorf("merge data failed: %w", err)
	}

	if this.ExampleDataStorer == nil {
		this.ExampleDataStorer = &ExampleDataStorer{
			Datas: map[int64]*ExampleData{},
		}
	}

	for k, v := range tmpl.Datas {
		if _, ok := this.ExampleDataStorer.Datas[k]; ok {
			return fmt.Errorf("merge data failed: key repeat")
		}

		this.ExampleDataStorer.Datas[k] = v
	}

	return nil
}

func (this *ExampleDataReader) Clear() {
	this.ExampleDataStorer = nil
}

func (this *ExampleDataReader) Get(key int64) (result *ExampleData, ok bool) {
	result, ok = this.Datas[key]
	return result, ok
}

func (this *ExampleDataReader) Keys() (result []int64) {
	for itor := range this.Datas {
		result = append(result, itor)
	}

	return result
}

func (this *ExampleDataReader) Values() (result []*ExampleData) {
	for _, itor := range this.Datas {
		result = append(result, itor)
	}

	return result
}

func (this *ExampleDataReader) Count() int {
	return len(this.Datas)
}