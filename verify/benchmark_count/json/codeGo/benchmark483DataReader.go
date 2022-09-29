// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark483DataReader struct {
	Benchmark483DataStorer
}

func (this *Benchmark483DataReader) FileName() string {
	return "benchmark483Data.json"
}

func (this *Benchmark483DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark483DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark483DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark483DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark483DataReader) FromData(data []byte) error {
	this.Benchmark483DataStorer = Benchmark483DataStorer{
		Datas: map[int64]Benchmark483Data{},
	}

	if err := json.Unmarshal(data, &this.Benchmark483DataStorer); err != nil {
		return err
	}

	return nil
}