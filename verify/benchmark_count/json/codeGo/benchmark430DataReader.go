// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark430DataReader struct {
	Benchmark430DataStorer
}

func (this *Benchmark430DataReader) FileName() string {
	return "benchmark430Data.json"
}

func (this *Benchmark430DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark430DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark430DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark430DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark430DataReader) FromData(data []byte) error {
	this.Benchmark430DataStorer = Benchmark430DataStorer{
		Datas: map[int64]Benchmark430Data{},
	}

	if err := json.Unmarshal(data, &this.Benchmark430DataStorer); err != nil {
		return err
	}

	return nil
}