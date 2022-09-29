// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark248DataReader struct {
	Benchmark248DataStorer
}

func (this *Benchmark248DataReader) FileName() string {
	return "benchmark248Data.pbd"
}

func (this *Benchmark248DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark248DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark248DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark248DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark248DataReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark248DataStorer); err != nil {
		return err
	}

	return nil
}