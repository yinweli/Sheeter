// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark151DataReader struct {
	Benchmark151DataStorer
}

func (this *Benchmark151DataReader) FileName() string {
	return "benchmark151Data.json"
}

func (this *Benchmark151DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark151DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark151DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark151DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark151DataReader) FromData(data []byte) error {
	this.Benchmark151DataStorer = Benchmark151DataStorer{
		Datas: map[int64]Benchmark151Data{},
	}

	if err := json.Unmarshal(data, &this.Benchmark151DataStorer); err != nil {
		return err
	}

	return nil
}
