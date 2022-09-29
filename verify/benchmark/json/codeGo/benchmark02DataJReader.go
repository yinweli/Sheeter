// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark02DataJReader struct {
	Benchmark02DataJStorer
}

func (this *Benchmark02DataJReader) FileName() string {
	return "benchmark02DataJ.json"
}

func (this *Benchmark02DataJReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark02DataJReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark02DataJReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark02DataJReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark02DataJReader) FromData(data []byte) error {
	this.Benchmark02DataJStorer = Benchmark02DataJStorer{
		Datas: map[int64]Benchmark02DataJ{},
	}

	if err := json.Unmarshal(data, &this.Benchmark02DataJStorer); err != nil {
		return err
	}

	return nil
}
