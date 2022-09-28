// generated by sheeter, DO NOT EDIT.

package sheeter

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark09DataReader struct {
	Benchmark09DataStorer
}

func (this *Benchmark09DataReader) FileName() string {
	return "benchmark09Data.json"
}

func (this *Benchmark09DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark09DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark09DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark09DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark09DataReader) FromData(data []byte) error {
	this.Benchmark09DataStorer = Benchmark09DataStorer{
		Datas: map[int64]Benchmark09Data{},
	}

	if err := json.Unmarshal(data, &this.Benchmark09DataStorer); err != nil {
		return err
	}

	return nil
}
