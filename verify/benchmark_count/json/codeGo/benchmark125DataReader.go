// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark125DataReader struct {
	Benchmark125DataStorer
}

func (this *Benchmark125DataReader) FileName() string {
	return "benchmark125Data.json"
}

func (this *Benchmark125DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark125DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark125DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark125DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark125DataReader) FromData(data []byte) error {
	this.Benchmark125DataStorer = Benchmark125DataStorer{
		Datas: map[int64]Benchmark125Data{},
	}

	if err := json.Unmarshal(data, &this.Benchmark125DataStorer); err != nil {
		return err
	}

	return nil
}
