// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark96DataReader struct {
	Benchmark96DataStorer
}

func (this *Benchmark96DataReader) FileName() string {
	return "benchmark96Data.json"
}

func (this *Benchmark96DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark96DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark96DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark96DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark96DataReader) FromData(data []byte) error {
	this.Benchmark96DataStorer = Benchmark96DataStorer{
		Datas: map[int64]Benchmark96Data{},
	}

	if err := json.Unmarshal(data, &this.Benchmark96DataStorer); err != nil {
		return err
	}

	return nil
}