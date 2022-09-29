// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark152DataReader struct {
	Benchmark152DataStorer
}

func (this *Benchmark152DataReader) FileName() string {
	return "benchmark152Data.json"
}

func (this *Benchmark152DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark152DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark152DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark152DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark152DataReader) FromData(data []byte) error {
	this.Benchmark152DataStorer = Benchmark152DataStorer{
		Datas: map[int64]Benchmark152Data{},
	}

	if err := json.Unmarshal(data, &this.Benchmark152DataStorer); err != nil {
		return err
	}

	return nil
}
