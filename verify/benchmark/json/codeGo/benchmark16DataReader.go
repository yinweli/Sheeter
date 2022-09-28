// generated by sheeter, DO NOT EDIT.

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark16DataReader struct {
	Benchmark16DataStorer
}

func (this *Benchmark16DataReader) FileName() string {
	return "benchmark16Data.json"
}

func (this *Benchmark16DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark16DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark16DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark16DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark16DataReader) FromData(data []byte) error {
	this.Benchmark16DataStorer = Benchmark16DataStorer{
		Datas: map[int64]Benchmark16Data{},
	}

	if err := json.Unmarshal(data, &this.Benchmark16DataStorer); err != nil {
		return err
	}

	return nil
}
