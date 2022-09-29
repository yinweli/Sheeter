// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark312DataReader struct {
	Benchmark312DataStorer
}

func (this *Benchmark312DataReader) FileName() string {
	return "benchmark312Data.pbd"
}

func (this *Benchmark312DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark312DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark312DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark312DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark312DataReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark312DataStorer); err != nil {
		return err
	}

	return nil
}