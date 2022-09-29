// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark349DataReader struct {
	Benchmark349DataStorer
}

func (this *Benchmark349DataReader) FileName() string {
	return "benchmark349Data.json"
}

func (this *Benchmark349DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark349DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark349DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark349DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark349DataReader) FromData(data []byte) error {
	this.Benchmark349DataStorer = Benchmark349DataStorer{
		Datas: map[int64]Benchmark349Data{},
	}

	if err := json.Unmarshal(data, &this.Benchmark349DataStorer); err != nil {
		return err
	}

	return nil
}
