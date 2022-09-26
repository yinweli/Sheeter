// generated by sheeter, DO NOT EDIT.

package sheeter

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type VerifyData1Reader struct {
	Datas VerifyData1Storer
}

type VerifyData1Storer = map[int64]VerifyData1

func (this *VerifyData1Reader) FileName() string {
	return "verifyData1.json"
}

func (this *VerifyData1Reader) FromFullPath(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("VerifyData1Reader: from full path failed: %w", err)
	}

	return this.FromData(data)
}

func (this *VerifyData1Reader) FromHalfPath(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("VerifyData1Reader: from half path failed: %w", err)
	}

	return this.FromData(data)
}

func (this *VerifyData1Reader) FromData(data []byte) error {
	datas := VerifyData1Storer{}

	if err := json.Unmarshal(data, &datas); err != nil {
		return err
	}

	this.Datas = datas
	return nil
}
