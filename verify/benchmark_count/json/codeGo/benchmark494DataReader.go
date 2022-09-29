// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark494DataReader struct {
	Benchmark494DataStorer
}

func (this *Benchmark494DataReader) FileName() string {
	return "benchmark494Data.json"
}

func (this *Benchmark494DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark494DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark494DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark494DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark494DataReader) FromData(data []byte) error {
	this.Benchmark494DataStorer = Benchmark494DataStorer{
		Datas: map[int64]Benchmark494Data{},
	}

	if err := json.Unmarshal(data, &this.Benchmark494DataStorer); err != nil {
		return err
	}

	return nil
}