// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark308DataReader struct {
	Benchmark308DataStorer
}

func (this *Benchmark308DataReader) FileName() string {
	return "benchmark308Data.json"
}

func (this *Benchmark308DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark308DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark308DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark308DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark308DataReader) FromData(data []byte) error {
	this.Benchmark308DataStorer = Benchmark308DataStorer{
		Datas: map[int64]Benchmark308Data{},
	}

	if err := json.Unmarshal(data, &this.Benchmark308DataStorer); err != nil {
		return err
	}

	return nil
}
