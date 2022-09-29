// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark498DataReader struct {
	Benchmark498DataStorer
}

func (this *Benchmark498DataReader) FileName() string {
	return "benchmark498Data.json"
}

func (this *Benchmark498DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark498DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark498DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark498DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark498DataReader) FromData(data []byte) error {
	this.Benchmark498DataStorer = Benchmark498DataStorer{
		Datas: map[int64]Benchmark498Data{},
	}

	if err := json.Unmarshal(data, &this.Benchmark498DataStorer); err != nil {
		return err
	}

	return nil
}
