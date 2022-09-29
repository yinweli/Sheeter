// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark238DataReader struct {
	Benchmark238DataStorer
}

func (this *Benchmark238DataReader) FileName() string {
	return "benchmark238Data.json"
}

func (this *Benchmark238DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark238DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark238DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark238DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark238DataReader) FromData(data []byte) error {
	this.Benchmark238DataStorer = Benchmark238DataStorer{
		Datas: map[int64]Benchmark238Data{},
	}

	if err := json.Unmarshal(data, &this.Benchmark238DataStorer); err != nil {
		return err
	}

	return nil
}