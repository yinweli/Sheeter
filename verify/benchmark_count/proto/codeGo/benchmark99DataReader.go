// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark99DataReader struct {
	Benchmark99DataStorer
}

func (this *Benchmark99DataReader) FileName() string {
	return "benchmark99Data.pbd"
}

func (this *Benchmark99DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark99DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark99DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark99DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark99DataReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark99DataStorer); err != nil {
		return err
	}

	return nil
}
