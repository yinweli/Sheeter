// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark329DataReader struct {
	Benchmark329DataStorer
}

func (this *Benchmark329DataReader) FileName() string {
	return "benchmark329Data.json"
}

func (this *Benchmark329DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark329DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark329DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark329DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark329DataReader) FromData(data []byte) error {
	this.Benchmark329DataStorer = Benchmark329DataStorer{
		Datas: map[int64]Benchmark329Data{},
	}

	if err := json.Unmarshal(data, &this.Benchmark329DataStorer); err != nil {
		return err
	}

	return nil
}