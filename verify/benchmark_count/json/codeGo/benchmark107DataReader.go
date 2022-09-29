// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark107DataReader struct {
	Benchmark107DataStorer
}

func (this *Benchmark107DataReader) FileName() string {
	return "benchmark107Data.json"
}

func (this *Benchmark107DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark107DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark107DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark107DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark107DataReader) FromData(data []byte) error {
	this.Benchmark107DataStorer = Benchmark107DataStorer{
		Datas: map[int64]Benchmark107Data{},
	}

	if err := json.Unmarshal(data, &this.Benchmark107DataStorer); err != nil {
		return err
	}

	return nil
}
