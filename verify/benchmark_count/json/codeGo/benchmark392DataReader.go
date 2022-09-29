// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark392DataReader struct {
	Benchmark392DataStorer
}

func (this *Benchmark392DataReader) FileName() string {
	return "benchmark392Data.json"
}

func (this *Benchmark392DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark392DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark392DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark392DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark392DataReader) FromData(data []byte) error {
	this.Benchmark392DataStorer = Benchmark392DataStorer{
		Datas: map[int64]Benchmark392Data{},
	}

	if err := json.Unmarshal(data, &this.Benchmark392DataStorer); err != nil {
		return err
	}

	return nil
}