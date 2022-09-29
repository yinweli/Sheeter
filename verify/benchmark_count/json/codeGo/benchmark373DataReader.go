// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark373DataReader struct {
	Benchmark373DataStorer
}

func (this *Benchmark373DataReader) FileName() string {
	return "benchmark373Data.json"
}

func (this *Benchmark373DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark373DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark373DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark373DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark373DataReader) FromData(data []byte) error {
	this.Benchmark373DataStorer = Benchmark373DataStorer{
		Datas: map[int64]Benchmark373Data{},
	}

	if err := json.Unmarshal(data, &this.Benchmark373DataStorer); err != nil {
		return err
	}

	return nil
}
