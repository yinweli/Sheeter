// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark52DataReader struct {
	Benchmark52DataStorer
}

func (this *Benchmark52DataReader) FileName() string {
	return "benchmark52Data.pbd"
}

func (this *Benchmark52DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark52DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark52DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark52DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark52DataReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark52DataStorer); err != nil {
		return err
	}

	return nil
}
