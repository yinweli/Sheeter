// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark05DataCReader struct {
	Benchmark05DataCStorer
}

func (this *Benchmark05DataCReader) FileName() string {
	return "benchmark05DataC.json"
}

func (this *Benchmark05DataCReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark05DataCReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark05DataCReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark05DataCReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark05DataCReader) FromData(data []byte) error {
	this.Benchmark05DataCStorer = Benchmark05DataCStorer{
		Datas: map[int64]Benchmark05DataC{},
	}

	if err := json.Unmarshal(data, &this.Benchmark05DataCStorer); err != nil {
		return err
	}

	return nil
}
