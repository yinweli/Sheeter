// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
)

type VerifyData1Reader struct {
	*VerifyData1Storer
}

func (this *VerifyData1Reader) DataName() string {
	return "verifyData1"
}

func (this *VerifyData1Reader) DataExt() string {
	return "json"
}

func (this *VerifyData1Reader) DataFile() string {
	return "verifyData1.json"
}

func (this *VerifyData1Reader) FromData(data []byte) error {
	this.VerifyData1Storer = &VerifyData1Storer{
		Datas: map[int64]*VerifyData1{},
	}

	if err := json.Unmarshal(data, this.VerifyData1Storer); err != nil {
		return fmt.Errorf("from data failed: %w", err)
	}

	return nil
}

func (this *VerifyData1Reader) MergeData(data []byte) error {
	tmpl := &VerifyData1Storer{
		Datas: map[int64]*VerifyData1{},
	}

	if err := json.Unmarshal(data, tmpl); err != nil {
		return fmt.Errorf("merge data failed: %w", err)
	}

	if this.VerifyData1Storer == nil {
		this.VerifyData1Storer = &VerifyData1Storer{
			Datas: map[int64]*VerifyData1{},
		}
	}

	for k, v := range tmpl.Datas {
		if _, ok := this.VerifyData1Storer.Datas[k]; ok {
			return fmt.Errorf("merge data failed: key repeat")
		}

		this.VerifyData1Storer.Datas[k] = v
	}

	return nil
}

func (this *VerifyData1Reader) Get(key int64) (result *VerifyData1, ok bool) {
	result, ok = this.Datas[key]
	return result, ok
}

func (this *VerifyData1Reader) Keys() (result []int64) {
	for itor := range this.Datas {
		result = append(result, itor)
	}

	return result
}

func (this *VerifyData1Reader) Values() (result []*VerifyData1) {
	for _, itor := range this.Datas {
		result = append(result, itor)
	}

	return result
}

func (this *VerifyData1Reader) Count() int {
	return len(this.Datas)
}
