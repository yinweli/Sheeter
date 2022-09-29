// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark86DataReader struct {
	Benchmark86DataStorer
}

func (this *Benchmark86DataReader) FileName() string {
	return "benchmark86Data.json"
}

func (this *Benchmark86DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark86DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark86DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark86DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark86DataReader) FromData(data []byte) error {
	this.Benchmark86DataStorer = Benchmark86DataStorer{
		Datas: map[int64]Benchmark86Data{},
	}

	if err := json.Unmarshal(data, &this.Benchmark86DataStorer); err != nil {
		return err
	}

	return nil
}