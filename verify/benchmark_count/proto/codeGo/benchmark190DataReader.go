// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark190DataReader struct {
	Benchmark190DataStorer
}

func (this *Benchmark190DataReader) FileName() string {
	return "benchmark190Data.pbd"
}

func (this *Benchmark190DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark190DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark190DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark190DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark190DataReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark190DataStorer); err != nil {
		return err
	}

	return nil
}