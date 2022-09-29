// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark487DataReader struct {
	Benchmark487DataStorer
}

func (this *Benchmark487DataReader) FileName() string {
	return "benchmark487Data.json"
}

func (this *Benchmark487DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark487DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark487DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark487DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark487DataReader) FromData(data []byte) error {
	this.Benchmark487DataStorer = Benchmark487DataStorer{
		Datas: map[int64]Benchmark487Data{},
	}

	if err := json.Unmarshal(data, &this.Benchmark487DataStorer); err != nil {
		return err
	}

	return nil
}