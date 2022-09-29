// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark499DataReader struct {
	Benchmark499DataStorer
}

func (this *Benchmark499DataReader) FileName() string {
	return "benchmark499Data.json"
}

func (this *Benchmark499DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark499DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark499DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark499DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark499DataReader) FromData(data []byte) error {
	this.Benchmark499DataStorer = Benchmark499DataStorer{
		Datas: map[int64]Benchmark499Data{},
	}

	if err := json.Unmarshal(data, &this.Benchmark499DataStorer); err != nil {
		return err
	}

	return nil
}
