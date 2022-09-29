// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark284DataReader struct {
	Benchmark284DataStorer
}

func (this *Benchmark284DataReader) FileName() string {
	return "benchmark284Data.json"
}

func (this *Benchmark284DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark284DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark284DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark284DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark284DataReader) FromData(data []byte) error {
	this.Benchmark284DataStorer = Benchmark284DataStorer{
		Datas: map[int64]Benchmark284Data{},
	}

	if err := json.Unmarshal(data, &this.Benchmark284DataStorer); err != nil {
		return err
	}

	return nil
}
