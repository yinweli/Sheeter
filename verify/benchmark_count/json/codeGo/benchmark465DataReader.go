// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark465DataReader struct {
	Benchmark465DataStorer
}

func (this *Benchmark465DataReader) FileName() string {
	return "benchmark465Data.json"
}

func (this *Benchmark465DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark465DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark465DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark465DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark465DataReader) FromData(data []byte) error {
	this.Benchmark465DataStorer = Benchmark465DataStorer{
		Datas: map[int64]Benchmark465Data{},
	}

	if err := json.Unmarshal(data, &this.Benchmark465DataStorer); err != nil {
		return err
	}

	return nil
}
