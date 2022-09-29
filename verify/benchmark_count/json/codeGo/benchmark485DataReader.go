// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark485DataReader struct {
	Benchmark485DataStorer
}

func (this *Benchmark485DataReader) FileName() string {
	return "benchmark485Data.json"
}

func (this *Benchmark485DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark485DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark485DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark485DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark485DataReader) FromData(data []byte) error {
	this.Benchmark485DataStorer = Benchmark485DataStorer{
		Datas: map[int64]Benchmark485Data{},
	}

	if err := json.Unmarshal(data, &this.Benchmark485DataStorer); err != nil {
		return err
	}

	return nil
}
