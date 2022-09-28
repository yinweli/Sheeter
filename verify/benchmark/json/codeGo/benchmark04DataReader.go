// generated by sheeter, DO NOT EDIT.

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark04DataReader struct {
	Benchmark04DataStorer
}

func (this *Benchmark04DataReader) FileName() string {
	return "benchmark04Data.json"
}

func (this *Benchmark04DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark04DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark04DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark04DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark04DataReader) FromData(data []byte) error {
	this.Benchmark04DataStorer = Benchmark04DataStorer{
		Datas: map[int64]Benchmark04Data{},
	}

	if err := json.Unmarshal(data, &this.Benchmark04DataStorer); err != nil {
		return err
	}

	return nil
}
