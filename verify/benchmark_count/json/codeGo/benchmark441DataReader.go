// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark441DataReader struct {
	Benchmark441DataStorer
}

func (this *Benchmark441DataReader) FileName() string {
	return "benchmark441Data.json"
}

func (this *Benchmark441DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark441DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark441DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark441DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark441DataReader) FromData(data []byte) error {
	this.Benchmark441DataStorer = Benchmark441DataStorer{
		Datas: map[int64]Benchmark441Data{},
	}

	if err := json.Unmarshal(data, &this.Benchmark441DataStorer); err != nil {
		return err
	}

	return nil
}