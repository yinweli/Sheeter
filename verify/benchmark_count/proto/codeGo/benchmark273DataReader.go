// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark273DataReader struct {
	Benchmark273DataStorer
}

func (this *Benchmark273DataReader) FileName() string {
	return "benchmark273Data.pbd"
}

func (this *Benchmark273DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark273DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark273DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark273DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark273DataReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark273DataStorer); err != nil {
		return err
	}

	return nil
}
