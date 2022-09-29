// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark460DataReader struct {
	Benchmark460DataStorer
}

func (this *Benchmark460DataReader) FileName() string {
	return "benchmark460Data.json"
}

func (this *Benchmark460DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark460DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark460DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark460DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark460DataReader) FromData(data []byte) error {
	this.Benchmark460DataStorer = Benchmark460DataStorer{
		Datas: map[int64]Benchmark460Data{},
	}

	if err := json.Unmarshal(data, &this.Benchmark460DataStorer); err != nil {
		return err
	}

	return nil
}
