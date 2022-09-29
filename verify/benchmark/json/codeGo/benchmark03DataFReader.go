// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark03DataFReader struct {
	Benchmark03DataFStorer
}

func (this *Benchmark03DataFReader) FileName() string {
	return "benchmark03DataF.json"
}

func (this *Benchmark03DataFReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark03DataFReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark03DataFReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark03DataFReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark03DataFReader) FromData(data []byte) error {
	this.Benchmark03DataFStorer = Benchmark03DataFStorer{
		Datas: map[int64]Benchmark03DataF{},
	}

	if err := json.Unmarshal(data, &this.Benchmark03DataFStorer); err != nil {
		return err
	}

	return nil
}
