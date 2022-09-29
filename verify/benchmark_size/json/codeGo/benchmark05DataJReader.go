// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark05DataJReader struct {
	Benchmark05DataJStorer
}

func (this *Benchmark05DataJReader) FileName() string {
	return "benchmark05DataJ.json"
}

func (this *Benchmark05DataJReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark05DataJReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark05DataJReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark05DataJReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark05DataJReader) FromData(data []byte) error {
	this.Benchmark05DataJStorer = Benchmark05DataJStorer{
		Datas: map[int64]Benchmark05DataJ{},
	}

	if err := json.Unmarshal(data, &this.Benchmark05DataJStorer); err != nil {
		return err
	}

	return nil
}