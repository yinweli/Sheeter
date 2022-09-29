// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark124DataReader struct {
	Benchmark124DataStorer
}

func (this *Benchmark124DataReader) FileName() string {
	return "benchmark124Data.json"
}

func (this *Benchmark124DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark124DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark124DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark124DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark124DataReader) FromData(data []byte) error {
	this.Benchmark124DataStorer = Benchmark124DataStorer{
		Datas: map[int64]Benchmark124Data{},
	}

	if err := json.Unmarshal(data, &this.Benchmark124DataStorer); err != nil {
		return err
	}

	return nil
}