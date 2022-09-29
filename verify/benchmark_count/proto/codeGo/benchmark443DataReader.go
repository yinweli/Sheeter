// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark443DataReader struct {
	Benchmark443DataStorer
}

func (this *Benchmark443DataReader) FileName() string {
	return "benchmark443Data.pbd"
}

func (this *Benchmark443DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark443DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark443DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark443DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark443DataReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark443DataStorer); err != nil {
		return err
	}

	return nil
}
