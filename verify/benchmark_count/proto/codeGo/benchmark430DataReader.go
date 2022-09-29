// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark430DataReader struct {
	Benchmark430DataStorer
}

func (this *Benchmark430DataReader) FileName() string {
	return "benchmark430Data.pbd"
}

func (this *Benchmark430DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark430DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark430DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark430DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark430DataReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark430DataStorer); err != nil {
		return err
	}

	return nil
}