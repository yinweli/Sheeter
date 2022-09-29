// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark184DataReader struct {
	Benchmark184DataStorer
}

func (this *Benchmark184DataReader) FileName() string {
	return "benchmark184Data.json"
}

func (this *Benchmark184DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark184DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark184DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark184DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark184DataReader) FromData(data []byte) error {
	this.Benchmark184DataStorer = Benchmark184DataStorer{
		Datas: map[int64]Benchmark184Data{},
	}

	if err := json.Unmarshal(data, &this.Benchmark184DataStorer); err != nil {
		return err
	}

	return nil
}
