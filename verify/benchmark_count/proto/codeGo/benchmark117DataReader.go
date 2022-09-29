// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark117DataReader struct {
	Benchmark117DataStorer
}

func (this *Benchmark117DataReader) FileName() string {
	return "benchmark117Data.pbd"
}

func (this *Benchmark117DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark117DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark117DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark117DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark117DataReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark117DataStorer); err != nil {
		return err
	}

	return nil
}