// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark138DataReader struct {
	Benchmark138DataStorer
}

func (this *Benchmark138DataReader) FileName() string {
	return "benchmark138Data.json"
}

func (this *Benchmark138DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark138DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark138DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark138DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark138DataReader) FromData(data []byte) error {
	this.Benchmark138DataStorer = Benchmark138DataStorer{
		Datas: map[int64]Benchmark138Data{},
	}

	if err := json.Unmarshal(data, &this.Benchmark138DataStorer); err != nil {
		return err
	}

	return nil
}