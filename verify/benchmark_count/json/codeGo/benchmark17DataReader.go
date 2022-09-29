// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark17DataReader struct {
	Benchmark17DataStorer
}

func (this *Benchmark17DataReader) FileName() string {
	return "benchmark17Data.json"
}

func (this *Benchmark17DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark17DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark17DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark17DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark17DataReader) FromData(data []byte) error {
	this.Benchmark17DataStorer = Benchmark17DataStorer{
		Datas: map[int64]Benchmark17Data{},
	}

	if err := json.Unmarshal(data, &this.Benchmark17DataStorer); err != nil {
		return err
	}

	return nil
}
