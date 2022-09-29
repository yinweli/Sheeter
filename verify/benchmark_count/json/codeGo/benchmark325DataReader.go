// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark325DataReader struct {
	Benchmark325DataStorer
}

func (this *Benchmark325DataReader) FileName() string {
	return "benchmark325Data.json"
}

func (this *Benchmark325DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark325DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark325DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark325DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark325DataReader) FromData(data []byte) error {
	this.Benchmark325DataStorer = Benchmark325DataStorer{
		Datas: map[int64]Benchmark325Data{},
	}

	if err := json.Unmarshal(data, &this.Benchmark325DataStorer); err != nil {
		return err
	}

	return nil
}
