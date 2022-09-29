// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark85DataReader struct {
	Benchmark85DataStorer
}

func (this *Benchmark85DataReader) FileName() string {
	return "benchmark85Data.pbd"
}

func (this *Benchmark85DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark85DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark85DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark85DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark85DataReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark85DataStorer); err != nil {
		return err
	}

	return nil
}
