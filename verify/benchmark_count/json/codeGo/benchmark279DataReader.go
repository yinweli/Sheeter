// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark279DataReader struct {
	Benchmark279DataStorer
}

func (this *Benchmark279DataReader) FileName() string {
	return "benchmark279Data.json"
}

func (this *Benchmark279DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark279DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark279DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark279DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark279DataReader) FromData(data []byte) error {
	this.Benchmark279DataStorer = Benchmark279DataStorer{
		Datas: map[int64]Benchmark279Data{},
	}

	if err := json.Unmarshal(data, &this.Benchmark279DataStorer); err != nil {
		return err
	}

	return nil
}