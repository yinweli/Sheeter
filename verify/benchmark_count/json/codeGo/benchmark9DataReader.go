// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark9DataReader struct {
	Benchmark9DataStorer
}

func (this *Benchmark9DataReader) FileName() string {
	return "benchmark9Data.json"
}

func (this *Benchmark9DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark9DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark9DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark9DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark9DataReader) FromData(data []byte) error {
	this.Benchmark9DataStorer = Benchmark9DataStorer{
		Datas: map[int64]Benchmark9Data{},
	}

	if err := json.Unmarshal(data, &this.Benchmark9DataStorer); err != nil {
		return err
	}

	return nil
}
