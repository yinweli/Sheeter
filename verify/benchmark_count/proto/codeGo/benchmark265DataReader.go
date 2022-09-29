// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark265DataReader struct {
	Benchmark265DataStorer
}

func (this *Benchmark265DataReader) FileName() string {
	return "benchmark265Data.pbd"
}

func (this *Benchmark265DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark265DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark265DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark265DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark265DataReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark265DataStorer); err != nil {
		return err
	}

	return nil
}
