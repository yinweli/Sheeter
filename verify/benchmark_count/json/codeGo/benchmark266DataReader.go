// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark266DataReader struct {
	Benchmark266DataStorer
}

func (this *Benchmark266DataReader) FileName() string {
	return "benchmark266Data.json"
}

func (this *Benchmark266DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark266DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark266DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark266DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark266DataReader) FromData(data []byte) error {
	this.Benchmark266DataStorer = Benchmark266DataStorer{
		Datas: map[int64]Benchmark266Data{},
	}

	if err := json.Unmarshal(data, &this.Benchmark266DataStorer); err != nil {
		return err
	}

	return nil
}