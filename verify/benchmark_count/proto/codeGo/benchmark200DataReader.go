// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark200DataReader struct {
	Benchmark200DataStorer
}

func (this *Benchmark200DataReader) FileName() string {
	return "benchmark200Data.pbd"
}

func (this *Benchmark200DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark200DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark200DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark200DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark200DataReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark200DataStorer); err != nil {
		return err
	}

	return nil
}