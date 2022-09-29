// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark132DataReader struct {
	Benchmark132DataStorer
}

func (this *Benchmark132DataReader) FileName() string {
	return "benchmark132Data.json"
}

func (this *Benchmark132DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark132DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark132DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark132DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark132DataReader) FromData(data []byte) error {
	this.Benchmark132DataStorer = Benchmark132DataStorer{
		Datas: map[int64]Benchmark132Data{},
	}

	if err := json.Unmarshal(data, &this.Benchmark132DataStorer); err != nil {
		return err
	}

	return nil
}
