// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark497DataReader struct {
	Benchmark497DataStorer
}

func (this *Benchmark497DataReader) FileName() string {
	return "benchmark497Data.json"
}

func (this *Benchmark497DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark497DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark497DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark497DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark497DataReader) FromData(data []byte) error {
	this.Benchmark497DataStorer = Benchmark497DataStorer{
		Datas: map[int64]Benchmark497Data{},
	}

	if err := json.Unmarshal(data, &this.Benchmark497DataStorer); err != nil {
		return err
	}

	return nil
}