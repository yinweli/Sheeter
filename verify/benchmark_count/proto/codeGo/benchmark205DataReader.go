// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark205DataReader struct {
	Benchmark205DataStorer
}

func (this *Benchmark205DataReader) FileName() string {
	return "benchmark205Data.pbd"
}

func (this *Benchmark205DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark205DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark205DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark205DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark205DataReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark205DataStorer); err != nil {
		return err
	}

	return nil
}
