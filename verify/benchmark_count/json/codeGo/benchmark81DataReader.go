// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark81DataReader struct {
	Benchmark81DataStorer
}

func (this *Benchmark81DataReader) FileName() string {
	return "benchmark81Data.json"
}

func (this *Benchmark81DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark81DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark81DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark81DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark81DataReader) FromData(data []byte) error {
	this.Benchmark81DataStorer = Benchmark81DataStorer{
		Datas: map[int64]Benchmark81Data{},
	}

	if err := json.Unmarshal(data, &this.Benchmark81DataStorer); err != nil {
		return err
	}

	return nil
}