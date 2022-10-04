// 以下是模板驗證用程式碼

package _json

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type RewardReader struct {
	RewardStorer
}

func (this *RewardReader) FileName() string {
	return "reward.json"
}

func (this *RewardReader) FromPath(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("RewardReader: from path failed: %w", err)
	}

	return this.FromData(data)
}

func (this *RewardReader) FromData(data []byte) error {
	this.RewardStorer = RewardStorer{
		Datas: map[int64]Reward{},
	}

	if err := json.Unmarshal(data, &this.RewardStorer); err != nil {
		return fmt.Errorf("RewardReader: from data failed: %w", err)
	}

	return nil
}

func (this *RewardReader) MergePath(path []string) (duplicates []int64) {
	for _, itor := range path {
		if data, err := os.ReadFile(filepath.Join(itor, this.FileName())); err == nil {
			duplicates = append(duplicates, this.MergeData(data)...)
		}
	}

	return duplicates
}

func (this *RewardReader) MergeData(data []byte) (duplicates []int64) {
	storer := &RewardStorer{
		Datas: map[int64]Reward{},
	}

	if err := json.Unmarshal(data, storer); err == nil {
		for k, v := range storer.Datas {
			if _, ok := this.RewardStorer.Datas[k]; ok == false {
				this.RewardStorer.Datas[k] = v
			} else {
				duplicates = append(duplicates, k)
			}
		}
	}

	return duplicates
}

// 以下是為了通過編譯的程式碼, 不可使用