// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark144DataReader struct {
	Benchmark144DataStorer
}

func (this *Benchmark144DataReader) FileName() string {
	return "benchmark144Data.json"
}

func (this *Benchmark144DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark144DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark144DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark144DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark144DataReader) FromData(data []byte) error {
	this.Benchmark144DataStorer = Benchmark144DataStorer{
		Datas: map[int64]Benchmark144Data{},
	}

	if err := json.Unmarshal(data, &this.Benchmark144DataStorer); err != nil {
		return err
	}

	return nil
}
