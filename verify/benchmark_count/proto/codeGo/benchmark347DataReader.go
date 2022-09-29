// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark347DataReader struct {
	Benchmark347DataStorer
}

func (this *Benchmark347DataReader) FileName() string {
	return "benchmark347Data.pbd"
}

func (this *Benchmark347DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark347DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark347DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark347DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark347DataReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark347DataStorer); err != nil {
		return err
	}

	return nil
}