// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark48DataReader struct {
	Benchmark48DataStorer
}

func (this *Benchmark48DataReader) FileName() string {
	return "benchmark48Data.pbd"
}

func (this *Benchmark48DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark48DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark48DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark48DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark48DataReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark48DataStorer); err != nil {
		return err
	}

	return nil
}
