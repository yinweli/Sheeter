// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark32DataReader struct {
	Benchmark32DataStorer
}

func (this *Benchmark32DataReader) FileName() string {
	return "benchmark32Data.json"
}

func (this *Benchmark32DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark32DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark32DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark32DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark32DataReader) FromData(data []byte) error {
	this.Benchmark32DataStorer = Benchmark32DataStorer{
		Datas: map[int64]Benchmark32Data{},
	}

	if err := json.Unmarshal(data, &this.Benchmark32DataStorer); err != nil {
		return err
	}

	return nil
}