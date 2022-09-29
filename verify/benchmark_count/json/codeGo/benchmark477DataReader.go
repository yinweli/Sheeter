// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark477DataReader struct {
	Benchmark477DataStorer
}

func (this *Benchmark477DataReader) FileName() string {
	return "benchmark477Data.json"
}

func (this *Benchmark477DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark477DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark477DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark477DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark477DataReader) FromData(data []byte) error {
	this.Benchmark477DataStorer = Benchmark477DataStorer{
		Datas: map[int64]Benchmark477Data{},
	}

	if err := json.Unmarshal(data, &this.Benchmark477DataStorer); err != nil {
		return err
	}

	return nil
}
