// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark353DataReader struct {
	Benchmark353DataStorer
}

func (this *Benchmark353DataReader) FileName() string {
	return "benchmark353Data.json"
}

func (this *Benchmark353DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark353DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark353DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark353DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark353DataReader) FromData(data []byte) error {
	this.Benchmark353DataStorer = Benchmark353DataStorer{
		Datas: map[int64]Benchmark353Data{},
	}

	if err := json.Unmarshal(data, &this.Benchmark353DataStorer); err != nil {
		return err
	}

	return nil
}
