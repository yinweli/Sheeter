// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark467DataReader struct {
	Benchmark467DataStorer
}

func (this *Benchmark467DataReader) FileName() string {
	return "benchmark467Data.json"
}

func (this *Benchmark467DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark467DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark467DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark467DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark467DataReader) FromData(data []byte) error {
	this.Benchmark467DataStorer = Benchmark467DataStorer{
		Datas: map[int64]Benchmark467Data{},
	}

	if err := json.Unmarshal(data, &this.Benchmark467DataStorer); err != nil {
		return err
	}

	return nil
}
