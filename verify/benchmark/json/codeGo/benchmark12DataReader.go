// generated by sheeter, DO NOT EDIT.

package sheeter

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark12DataReader struct {
	Benchmark12DataStorer
}

func (this *Benchmark12DataReader) FileName() string {
	return "benchmark12Data.json"
}

func (this *Benchmark12DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark12DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark12DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark12DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark12DataReader) FromData(data []byte) error {
	this.Benchmark12DataStorer = Benchmark12DataStorer{
		Datas: map[int64]Benchmark12Data{},
	}

	if err := json.Unmarshal(data, &this.Benchmark12DataStorer); err != nil {
		return err
	}

	return nil
}
