// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark182DataReader struct {
	Benchmark182DataStorer
}

func (this *Benchmark182DataReader) FileName() string {
	return "benchmark182Data.pbd"
}

func (this *Benchmark182DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark182DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark182DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark182DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark182DataReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark182DataStorer); err != nil {
		return err
	}

	return nil
}
