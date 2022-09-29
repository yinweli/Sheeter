// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark413DataReader struct {
	Benchmark413DataStorer
}

func (this *Benchmark413DataReader) FileName() string {
	return "benchmark413Data.json"
}

func (this *Benchmark413DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark413DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark413DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark413DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark413DataReader) FromData(data []byte) error {
	this.Benchmark413DataStorer = Benchmark413DataStorer{
		Datas: map[int64]Benchmark413Data{},
	}

	if err := json.Unmarshal(data, &this.Benchmark413DataStorer); err != nil {
		return err
	}

	return nil
}