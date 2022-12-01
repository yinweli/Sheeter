// 以下是模板驗證用程式碼

package _json

import (
	"encoding/json"
	"fmt"
)

type RewardReader struct {
	*RewardStorer
}

func (this *RewardReader) FileName() FileName {
	return NewFileName("reward", ".json")
}

func (this *RewardReader) FromData(data []byte) error {
	this.RewardStorer = &RewardStorer{
		Datas: map[int64]*Reward{},
	}

	if err := json.Unmarshal(data, this.RewardStorer); err != nil {
		return fmt.Errorf("from data failed: %w", err)
	}

	return nil
}

func (this *RewardReader) MergeData(data []byte) error {
	tmpl := &RewardStorer{
		Datas: map[int64]*Reward{},
	}

	if err := json.Unmarshal(data, tmpl); err != nil {
		return fmt.Errorf("merge data failed: %w", err)
	}

	if this.RewardStorer == nil {
		this.RewardStorer = &RewardStorer{
			Datas: map[int64]*Reward{},
		}
	}

	for k, v := range tmpl.Datas {
		if _, ok := this.RewardStorer.Datas[k]; ok {
			return fmt.Errorf("merge data failed: key repeat")
		}

		this.RewardStorer.Datas[k] = v
	}

	return nil
}

func (this *RewardReader) Clear() {
	this.RewardStorer = nil
}

func (this *RewardReader) Get(key int64) (result *Reward, ok bool) {
	result, ok = this.Datas[key]
	return result, ok
}

func (this *RewardReader) Keys() (result []int64) {
	for itor := range this.Datas {
		result = append(result, itor)
	}

	return result
}

func (this *RewardReader) Values() (result []*Reward) {
	for _, itor := range this.Datas {
		result = append(result, itor)
	}

	return result
}

func (this *RewardReader) Count() int {
	return len(this.Datas)
}

// 以下是為了通過編譯的程式碼, 不可使用
