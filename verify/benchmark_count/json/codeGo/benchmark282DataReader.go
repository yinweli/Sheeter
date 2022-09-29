// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark282DataReader struct {
	Benchmark282DataStorer
}

func (this *Benchmark282DataReader) FileName() string {
	return "benchmark282Data.json"
}

func (this *Benchmark282DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark282DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark282DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark282DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark282DataReader) FromData(data []byte) error {
	this.Benchmark282DataStorer = Benchmark282DataStorer{
		Datas: map[int64]Benchmark282Data{},
	}

	if err := json.Unmarshal(data, &this.Benchmark282DataStorer); err != nil {
		return err
	}

	return nil
}
