// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark215DataReader struct {
	Benchmark215DataStorer
}

func (this *Benchmark215DataReader) FileName() string {
	return "benchmark215Data.json"
}

func (this *Benchmark215DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark215DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark215DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark215DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark215DataReader) FromData(data []byte) error {
	this.Benchmark215DataStorer = Benchmark215DataStorer{
		Datas: map[int64]Benchmark215Data{},
	}

	if err := json.Unmarshal(data, &this.Benchmark215DataStorer); err != nil {
		return err
	}

	return nil
}
