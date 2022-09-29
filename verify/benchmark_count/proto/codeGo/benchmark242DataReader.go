// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark242DataReader struct {
	Benchmark242DataStorer
}

func (this *Benchmark242DataReader) FileName() string {
	return "benchmark242Data.pbd"
}

func (this *Benchmark242DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark242DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark242DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark242DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark242DataReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark242DataStorer); err != nil {
		return err
	}

	return nil
}