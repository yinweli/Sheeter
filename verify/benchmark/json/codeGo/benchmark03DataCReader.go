// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark03DataCReader struct {
	Benchmark03DataCStorer
}

func (this *Benchmark03DataCReader) FileName() string {
	return "benchmark03DataC.json"
}

func (this *Benchmark03DataCReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark03DataCReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark03DataCReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark03DataCReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark03DataCReader) FromData(data []byte) error {
	this.Benchmark03DataCStorer = Benchmark03DataCStorer{
		Datas: map[int64]Benchmark03DataC{},
	}

	if err := json.Unmarshal(data, &this.Benchmark03DataCStorer); err != nil {
		return err
	}

	return nil
}