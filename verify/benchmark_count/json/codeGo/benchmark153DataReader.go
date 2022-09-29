// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark153DataReader struct {
	Benchmark153DataStorer
}

func (this *Benchmark153DataReader) FileName() string {
	return "benchmark153Data.json"
}

func (this *Benchmark153DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark153DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark153DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark153DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark153DataReader) FromData(data []byte) error {
	this.Benchmark153DataStorer = Benchmark153DataStorer{
		Datas: map[int64]Benchmark153Data{},
	}

	if err := json.Unmarshal(data, &this.Benchmark153DataStorer); err != nil {
		return err
	}

	return nil
}
