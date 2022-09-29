// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark272DataReader struct {
	Benchmark272DataStorer
}

func (this *Benchmark272DataReader) FileName() string {
	return "benchmark272Data.pbd"
}

func (this *Benchmark272DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark272DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark272DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark272DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark272DataReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark272DataStorer); err != nil {
		return err
	}

	return nil
}