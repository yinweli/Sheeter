// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark231DataReader struct {
	Benchmark231DataStorer
}

func (this *Benchmark231DataReader) FileName() string {
	return "benchmark231Data.json"
}

func (this *Benchmark231DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark231DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark231DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark231DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark231DataReader) FromData(data []byte) error {
	this.Benchmark231DataStorer = Benchmark231DataStorer{
		Datas: map[int64]Benchmark231Data{},
	}

	if err := json.Unmarshal(data, &this.Benchmark231DataStorer); err != nil {
		return err
	}

	return nil
}
