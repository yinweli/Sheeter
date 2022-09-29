// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark309DataReader struct {
	Benchmark309DataStorer
}

func (this *Benchmark309DataReader) FileName() string {
	return "benchmark309Data.json"
}

func (this *Benchmark309DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark309DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark309DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark309DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark309DataReader) FromData(data []byte) error {
	this.Benchmark309DataStorer = Benchmark309DataStorer{
		Datas: map[int64]Benchmark309Data{},
	}

	if err := json.Unmarshal(data, &this.Benchmark309DataStorer); err != nil {
		return err
	}

	return nil
}