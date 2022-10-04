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

// 以下是為了通過編譯的程式碼, 不可使用
