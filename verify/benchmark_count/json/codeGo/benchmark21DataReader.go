// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark21DataReader struct {
	Benchmark21DataStorer
}

func (this *Benchmark21DataReader) FileName() string {
	return "benchmark21Data.json"
}

func (this *Benchmark21DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark21DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark21DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark21DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark21DataReader) FromData(data []byte) error {
	this.Benchmark21DataStorer = Benchmark21DataStorer{
		Datas: map[int64]Benchmark21Data{},
	}

	if err := json.Unmarshal(data, &this.Benchmark21DataStorer); err != nil {
		return err
	}

	return nil
}
