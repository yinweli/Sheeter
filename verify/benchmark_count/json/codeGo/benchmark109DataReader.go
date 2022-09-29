// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark109DataReader struct {
	Benchmark109DataStorer
}

func (this *Benchmark109DataReader) FileName() string {
	return "benchmark109Data.json"
}

func (this *Benchmark109DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark109DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark109DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark109DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark109DataReader) FromData(data []byte) error {
	this.Benchmark109DataStorer = Benchmark109DataStorer{
		Datas: map[int64]Benchmark109Data{},
	}

	if err := json.Unmarshal(data, &this.Benchmark109DataStorer); err != nil {
		return err
	}

	return nil
}
