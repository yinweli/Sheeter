// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark328DataReader struct {
	Benchmark328DataStorer
}

func (this *Benchmark328DataReader) FileName() string {
	return "benchmark328Data.json"
}

func (this *Benchmark328DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark328DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark328DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark328DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark328DataReader) FromData(data []byte) error {
	this.Benchmark328DataStorer = Benchmark328DataStorer{
		Datas: map[int64]Benchmark328Data{},
	}

	if err := json.Unmarshal(data, &this.Benchmark328DataStorer); err != nil {
		return err
	}

	return nil
}