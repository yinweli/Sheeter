// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark267DataReader struct {
	Benchmark267DataStorer
}

func (this *Benchmark267DataReader) FileName() string {
	return "benchmark267Data.pbd"
}

func (this *Benchmark267DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark267DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark267DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark267DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark267DataReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark267DataStorer); err != nil {
		return err
	}

	return nil
}
