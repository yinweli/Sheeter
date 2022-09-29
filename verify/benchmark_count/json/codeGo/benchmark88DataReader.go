// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark88DataReader struct {
	Benchmark88DataStorer
}

func (this *Benchmark88DataReader) FileName() string {
	return "benchmark88Data.json"
}

func (this *Benchmark88DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark88DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark88DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark88DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark88DataReader) FromData(data []byte) error {
	this.Benchmark88DataStorer = Benchmark88DataStorer{
		Datas: map[int64]Benchmark88Data{},
	}

	if err := json.Unmarshal(data, &this.Benchmark88DataStorer); err != nil {
		return err
	}

	return nil
}