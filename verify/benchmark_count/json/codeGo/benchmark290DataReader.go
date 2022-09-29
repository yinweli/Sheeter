// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark290DataReader struct {
	Benchmark290DataStorer
}

func (this *Benchmark290DataReader) FileName() string {
	return "benchmark290Data.json"
}

func (this *Benchmark290DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark290DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark290DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark290DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark290DataReader) FromData(data []byte) error {
	this.Benchmark290DataStorer = Benchmark290DataStorer{
		Datas: map[int64]Benchmark290Data{},
	}

	if err := json.Unmarshal(data, &this.Benchmark290DataStorer); err != nil {
		return err
	}

	return nil
}