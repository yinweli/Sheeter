// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark314DataReader struct {
	Benchmark314DataStorer
}

func (this *Benchmark314DataReader) FileName() string {
	return "benchmark314Data.pbd"
}

func (this *Benchmark314DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark314DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark314DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark314DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark314DataReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark314DataStorer); err != nil {
		return err
	}

	return nil
}