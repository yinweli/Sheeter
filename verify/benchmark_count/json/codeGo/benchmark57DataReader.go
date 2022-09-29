// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark57DataReader struct {
	Benchmark57DataStorer
}

func (this *Benchmark57DataReader) FileName() string {
	return "benchmark57Data.json"
}

func (this *Benchmark57DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark57DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark57DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark57DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark57DataReader) FromData(data []byte) error {
	this.Benchmark57DataStorer = Benchmark57DataStorer{
		Datas: map[int64]Benchmark57Data{},
	}

	if err := json.Unmarshal(data, &this.Benchmark57DataStorer); err != nil {
		return err
	}

	return nil
}
