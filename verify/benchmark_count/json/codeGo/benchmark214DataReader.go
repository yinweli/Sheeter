// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark214DataReader struct {
	Benchmark214DataStorer
}

func (this *Benchmark214DataReader) FileName() string {
	return "benchmark214Data.json"
}

func (this *Benchmark214DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark214DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark214DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark214DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark214DataReader) FromData(data []byte) error {
	this.Benchmark214DataStorer = Benchmark214DataStorer{
		Datas: map[int64]Benchmark214Data{},
	}

	if err := json.Unmarshal(data, &this.Benchmark214DataStorer); err != nil {
		return err
	}

	return nil
}