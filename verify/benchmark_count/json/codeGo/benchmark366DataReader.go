// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark366DataReader struct {
	Benchmark366DataStorer
}

func (this *Benchmark366DataReader) FileName() string {
	return "benchmark366Data.json"
}

func (this *Benchmark366DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark366DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark366DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark366DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark366DataReader) FromData(data []byte) error {
	this.Benchmark366DataStorer = Benchmark366DataStorer{
		Datas: map[int64]Benchmark366Data{},
	}

	if err := json.Unmarshal(data, &this.Benchmark366DataStorer); err != nil {
		return err
	}

	return nil
}
