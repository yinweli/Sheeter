// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark179DataReader struct {
	Benchmark179DataStorer
}

func (this *Benchmark179DataReader) FileName() string {
	return "benchmark179Data.json"
}

func (this *Benchmark179DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark179DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark179DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark179DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark179DataReader) FromData(data []byte) error {
	this.Benchmark179DataStorer = Benchmark179DataStorer{
		Datas: map[int64]Benchmark179Data{},
	}

	if err := json.Unmarshal(data, &this.Benchmark179DataStorer); err != nil {
		return err
	}

	return nil
}