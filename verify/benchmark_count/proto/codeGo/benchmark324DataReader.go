// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark324DataReader struct {
	Benchmark324DataStorer
}

func (this *Benchmark324DataReader) FileName() string {
	return "benchmark324Data.pbd"
}

func (this *Benchmark324DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark324DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark324DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark324DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark324DataReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark324DataStorer); err != nil {
		return err
	}

	return nil
}