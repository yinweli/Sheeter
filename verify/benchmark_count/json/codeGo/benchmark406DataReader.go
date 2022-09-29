// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark406DataReader struct {
	Benchmark406DataStorer
}

func (this *Benchmark406DataReader) FileName() string {
	return "benchmark406Data.json"
}

func (this *Benchmark406DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark406DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark406DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark406DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark406DataReader) FromData(data []byte) error {
	this.Benchmark406DataStorer = Benchmark406DataStorer{
		Datas: map[int64]Benchmark406Data{},
	}

	if err := json.Unmarshal(data, &this.Benchmark406DataStorer); err != nil {
		return err
	}

	return nil
}
