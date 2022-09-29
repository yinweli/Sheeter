// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark367DataReader struct {
	Benchmark367DataStorer
}

func (this *Benchmark367DataReader) FileName() string {
	return "benchmark367Data.json"
}

func (this *Benchmark367DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark367DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark367DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark367DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark367DataReader) FromData(data []byte) error {
	this.Benchmark367DataStorer = Benchmark367DataStorer{
		Datas: map[int64]Benchmark367Data{},
	}

	if err := json.Unmarshal(data, &this.Benchmark367DataStorer); err != nil {
		return err
	}

	return nil
}
