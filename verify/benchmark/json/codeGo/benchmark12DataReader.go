// generated by sheeter, DO NOT EDIT.

package sheeter

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark12DataReader struct {
	Datas Benchmark12DataStorer
}

type Benchmark12DataStorer = map[int64]Benchmark12Data

func (this *Benchmark12DataReader) FileName() string {
	return "benchmark12Data.json"
}

func (this *Benchmark12DataReader) FromFullPath(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark12DataReader: from full path failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark12DataReader) FromHalfPath(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark12DataReader: from half path failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark12DataReader) FromData(data []byte) error {
	datas := Benchmark12DataStorer{}

	if err := json.Unmarshal(data, &datas); err != nil {
		return err
	}

	this.Datas = datas
	return nil
}
