// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark121DataReader struct {
	Benchmark121DataStorer
}

func (this *Benchmark121DataReader) FileName() string {
	return "benchmark121Data.pbd"
}

func (this *Benchmark121DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark121DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark121DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark121DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark121DataReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark121DataStorer); err != nil {
		return err
	}

	return nil
}