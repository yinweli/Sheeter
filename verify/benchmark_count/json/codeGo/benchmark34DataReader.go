// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark34DataReader struct {
	Benchmark34DataStorer
}

func (this *Benchmark34DataReader) FileName() string {
	return "benchmark34Data.json"
}

func (this *Benchmark34DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark34DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark34DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark34DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark34DataReader) FromData(data []byte) error {
	this.Benchmark34DataStorer = Benchmark34DataStorer{
		Datas: map[int64]Benchmark34Data{},
	}

	if err := json.Unmarshal(data, &this.Benchmark34DataStorer); err != nil {
		return err
	}

	return nil
}