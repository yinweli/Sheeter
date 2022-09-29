// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark360DataReader struct {
	Benchmark360DataStorer
}

func (this *Benchmark360DataReader) FileName() string {
	return "benchmark360Data.pbd"
}

func (this *Benchmark360DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark360DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark360DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark360DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark360DataReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark360DataStorer); err != nil {
		return err
	}

	return nil
}
