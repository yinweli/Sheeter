// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark363DataReader struct {
	Benchmark363DataStorer
}

func (this *Benchmark363DataReader) FileName() string {
	return "benchmark363Data.json"
}

func (this *Benchmark363DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark363DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark363DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark363DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark363DataReader) FromData(data []byte) error {
	this.Benchmark363DataStorer = Benchmark363DataStorer{
		Datas: map[int64]Benchmark363Data{},
	}

	if err := json.Unmarshal(data, &this.Benchmark363DataStorer); err != nil {
		return err
	}

	return nil
}