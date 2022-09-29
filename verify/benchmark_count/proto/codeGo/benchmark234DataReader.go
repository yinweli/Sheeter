// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark234DataReader struct {
	Benchmark234DataStorer
}

func (this *Benchmark234DataReader) FileName() string {
	return "benchmark234Data.pbd"
}

func (this *Benchmark234DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark234DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark234DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark234DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark234DataReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark234DataStorer); err != nil {
		return err
	}

	return nil
}
