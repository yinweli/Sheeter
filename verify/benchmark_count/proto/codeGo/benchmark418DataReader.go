// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark418DataReader struct {
	Benchmark418DataStorer
}

func (this *Benchmark418DataReader) FileName() string {
	return "benchmark418Data.pbd"
}

func (this *Benchmark418DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark418DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark418DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark418DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark418DataReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark418DataStorer); err != nil {
		return err
	}

	return nil
}