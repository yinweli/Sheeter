// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark472DataReader struct {
	Benchmark472DataStorer
}

func (this *Benchmark472DataReader) FileName() string {
	return "benchmark472Data.json"
}

func (this *Benchmark472DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark472DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark472DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark472DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark472DataReader) FromData(data []byte) error {
	this.Benchmark472DataStorer = Benchmark472DataStorer{
		Datas: map[int64]Benchmark472Data{},
	}

	if err := json.Unmarshal(data, &this.Benchmark472DataStorer); err != nil {
		return err
	}

	return nil
}
