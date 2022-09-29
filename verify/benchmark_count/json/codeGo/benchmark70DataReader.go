// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark70DataReader struct {
	Benchmark70DataStorer
}

func (this *Benchmark70DataReader) FileName() string {
	return "benchmark70Data.json"
}

func (this *Benchmark70DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark70DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark70DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark70DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark70DataReader) FromData(data []byte) error {
	this.Benchmark70DataStorer = Benchmark70DataStorer{
		Datas: map[int64]Benchmark70Data{},
	}

	if err := json.Unmarshal(data, &this.Benchmark70DataStorer); err != nil {
		return err
	}

	return nil
}