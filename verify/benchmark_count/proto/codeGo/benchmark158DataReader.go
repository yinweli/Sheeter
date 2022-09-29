// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark158DataReader struct {
	Benchmark158DataStorer
}

func (this *Benchmark158DataReader) FileName() string {
	return "benchmark158Data.pbd"
}

func (this *Benchmark158DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark158DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark158DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark158DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark158DataReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark158DataStorer); err != nil {
		return err
	}

	return nil
}