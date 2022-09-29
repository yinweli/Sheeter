// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark302DataReader struct {
	Benchmark302DataStorer
}

func (this *Benchmark302DataReader) FileName() string {
	return "benchmark302Data.json"
}

func (this *Benchmark302DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark302DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark302DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark302DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark302DataReader) FromData(data []byte) error {
	this.Benchmark302DataStorer = Benchmark302DataStorer{
		Datas: map[int64]Benchmark302Data{},
	}

	if err := json.Unmarshal(data, &this.Benchmark302DataStorer); err != nil {
		return err
	}

	return nil
}