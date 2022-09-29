// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark386DataReader struct {
	Benchmark386DataStorer
}

func (this *Benchmark386DataReader) FileName() string {
	return "benchmark386Data.json"
}

func (this *Benchmark386DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark386DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark386DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark386DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark386DataReader) FromData(data []byte) error {
	this.Benchmark386DataStorer = Benchmark386DataStorer{
		Datas: map[int64]Benchmark386Data{},
	}

	if err := json.Unmarshal(data, &this.Benchmark386DataStorer); err != nil {
		return err
	}

	return nil
}
