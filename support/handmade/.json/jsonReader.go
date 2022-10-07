// 以下是模板驗證用程式碼

package _json

import (
	"encoding/json"
	"fmt"
)

type RewardReader struct {
	*RewardStorer
}

func (this *RewardReader) DataName() string {
	return "reward"
}

func (this *RewardReader) DataExt() string {
	return "json"
}

func (this *RewardReader) DataFile() string {
	return "reward.json"
}

func (this *RewardReader) FromData(data []byte) error {
	this.RewardStorer = &RewardStorer{
		Datas: map[int64]Reward{},
	}

	if err := json.Unmarshal(data, this.RewardStorer); err != nil {
		return fmt.Errorf("RewardReader: from data failed: %w", err)
	}

	return nil
}

func (this *RewardReader) MergeData(data []byte) (repeats []int64) {
	tmpl := &RewardStorer{
		Datas: map[int64]Reward{},
	}

	if err := json.Unmarshal(data, tmpl); err != nil {
		return repeats
	}

	if this.RewardStorer == nil {
		this.RewardStorer = &RewardStorer{
			Datas: map[int64]Reward{},
		}
	}

	for k, v := range tmpl.Datas {
		if _, ok := this.RewardStorer.Datas[k]; ok == false {
			this.RewardStorer.Datas[k] = v
		} else {
			repeats = append(repeats, k)
		}
	}

	return repeats
}

// 以下是為了通過編譯的程式碼, 不可使用
