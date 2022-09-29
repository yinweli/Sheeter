// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Benchmark03DataBReader struct {
	Benchmark03DataBStorer
}

func (this *Benchmark03DataBReader) FileName() string {
	return "benchmark03DataB.json"
}

func (this *Benchmark03DataBReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark03DataBReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark03DataBReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark03DataBReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark03DataBReader) FromData(data []byte) error {
	this.Benchmark03DataBStorer = Benchmark03DataBStorer{
		Datas: map[int64]Benchmark03DataB{},
	}

	if err := json.Unmarshal(data, &this.Benchmark03DataBStorer); err != nil {
		return err
	}

	return nil
}