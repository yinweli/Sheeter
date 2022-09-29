// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark02DataAReader struct {
	Benchmark02DataAStorer
}

func (this *Benchmark02DataAReader) FileName() string {
	return "benchmark02DataA.json"
}

func (this *Benchmark02DataAReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark02DataAReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark02DataAReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark02DataAReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark02DataAReader) FromData(data []byte) error {
	this.Benchmark02DataAStorer = Benchmark02DataAStorer{
		Datas: map[int64]Benchmark02DataA{},
	}

	if err := json.Unmarshal(data, &this.Benchmark02DataAStorer); err != nil {
		return err
	}

	return nil
}