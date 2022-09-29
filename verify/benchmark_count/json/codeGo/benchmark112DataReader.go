// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark112DataReader struct {
	Benchmark112DataStorer
}

func (this *Benchmark112DataReader) FileName() string {
	return "benchmark112Data.json"
}

func (this *Benchmark112DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark112DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark112DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark112DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark112DataReader) FromData(data []byte) error {
	this.Benchmark112DataStorer = Benchmark112DataStorer{
		Datas: map[int64]Benchmark112Data{},
	}

	if err := json.Unmarshal(data, &this.Benchmark112DataStorer); err != nil {
		return err
	}

	return nil
}