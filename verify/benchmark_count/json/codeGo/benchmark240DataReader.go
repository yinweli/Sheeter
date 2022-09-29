// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark240DataReader struct {
	Benchmark240DataStorer
}

func (this *Benchmark240DataReader) FileName() string {
	return "benchmark240Data.json"
}

func (this *Benchmark240DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark240DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark240DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark240DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark240DataReader) FromData(data []byte) error {
	this.Benchmark240DataStorer = Benchmark240DataStorer{
		Datas: map[int64]Benchmark240Data{},
	}

	if err := json.Unmarshal(data, &this.Benchmark240DataStorer); err != nil {
		return err
	}

	return nil
}