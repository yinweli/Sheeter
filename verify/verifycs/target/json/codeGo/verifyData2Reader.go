// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type VerifyData2Reader struct {
	VerifyData2Storer
}

func (this *VerifyData2Reader) FileName() string {
	return "verifyData2.json"
}

func (this *VerifyData2Reader) FromPath(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("VerifyData2Reader: from path failed: %w", err)
	}

	return this.FromData(data)
}

func (this *VerifyData2Reader) FromData(data []byte) error {
	this.VerifyData2Storer = VerifyData2Storer{
		Datas: map[int64]VerifyData2{},
	}

	if err := json.Unmarshal(data, &this.VerifyData2Storer); err != nil {
		return fmt.Errorf("VerifyData2Reader: from data failed: %w", err)
	}

	return nil
}
