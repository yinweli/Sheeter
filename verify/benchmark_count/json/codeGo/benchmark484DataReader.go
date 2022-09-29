// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark484DataReader struct {
	Benchmark484DataStorer
}

func (this *Benchmark484DataReader) FileName() string {
	return "benchmark484Data.json"
}

func (this *Benchmark484DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark484DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark484DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark484DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark484DataReader) FromData(data []byte) error {
	this.Benchmark484DataStorer = Benchmark484DataStorer{
		Datas: map[int64]Benchmark484Data{},
	}

	if err := json.Unmarshal(data, &this.Benchmark484DataStorer); err != nil {
		return err
	}

	return nil
}