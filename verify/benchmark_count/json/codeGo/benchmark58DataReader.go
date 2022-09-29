// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark58DataReader struct {
	Benchmark58DataStorer
}

func (this *Benchmark58DataReader) FileName() string {
	return "benchmark58Data.json"
}

func (this *Benchmark58DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark58DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark58DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark58DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark58DataReader) FromData(data []byte) error {
	this.Benchmark58DataStorer = Benchmark58DataStorer{
		Datas: map[int64]Benchmark58Data{},
	}

	if err := json.Unmarshal(data, &this.Benchmark58DataStorer); err != nil {
		return err
	}

	return nil
}
