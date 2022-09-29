// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark189DataReader struct {
	Benchmark189DataStorer
}

func (this *Benchmark189DataReader) FileName() string {
	return "benchmark189Data.json"
}

func (this *Benchmark189DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark189DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark189DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark189DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark189DataReader) FromData(data []byte) error {
	this.Benchmark189DataStorer = Benchmark189DataStorer{
		Datas: map[int64]Benchmark189Data{},
	}

	if err := json.Unmarshal(data, &this.Benchmark189DataStorer); err != nil {
		return err
	}

	return nil
}