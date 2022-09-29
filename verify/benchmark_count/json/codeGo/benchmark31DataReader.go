// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark31DataReader struct {
	Benchmark31DataStorer
}

func (this *Benchmark31DataReader) FileName() string {
	return "benchmark31Data.json"
}

func (this *Benchmark31DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark31DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark31DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark31DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark31DataReader) FromData(data []byte) error {
	this.Benchmark31DataStorer = Benchmark31DataStorer{
		Datas: map[int64]Benchmark31Data{},
	}

	if err := json.Unmarshal(data, &this.Benchmark31DataStorer); err != nil {
		return err
	}

	return nil
}