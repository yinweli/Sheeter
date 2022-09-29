// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark463DataReader struct {
	Benchmark463DataStorer
}

func (this *Benchmark463DataReader) FileName() string {
	return "benchmark463Data.json"
}

func (this *Benchmark463DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark463DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark463DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark463DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark463DataReader) FromData(data []byte) error {
	this.Benchmark463DataStorer = Benchmark463DataStorer{
		Datas: map[int64]Benchmark463Data{},
	}

	if err := json.Unmarshal(data, &this.Benchmark463DataStorer); err != nil {
		return err
	}

	return nil
}
