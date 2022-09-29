// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark439DataReader struct {
	Benchmark439DataStorer
}

func (this *Benchmark439DataReader) FileName() string {
	return "benchmark439Data.json"
}

func (this *Benchmark439DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark439DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark439DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark439DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark439DataReader) FromData(data []byte) error {
	this.Benchmark439DataStorer = Benchmark439DataStorer{
		Datas: map[int64]Benchmark439Data{},
	}

	if err := json.Unmarshal(data, &this.Benchmark439DataStorer); err != nil {
		return err
	}

	return nil
}
