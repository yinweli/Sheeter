// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark411DataReader struct {
	Benchmark411DataStorer
}

func (this *Benchmark411DataReader) FileName() string {
	return "benchmark411Data.json"
}

func (this *Benchmark411DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark411DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark411DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark411DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark411DataReader) FromData(data []byte) error {
	this.Benchmark411DataStorer = Benchmark411DataStorer{
		Datas: map[int64]Benchmark411Data{},
	}

	if err := json.Unmarshal(data, &this.Benchmark411DataStorer); err != nil {
		return err
	}

	return nil
}