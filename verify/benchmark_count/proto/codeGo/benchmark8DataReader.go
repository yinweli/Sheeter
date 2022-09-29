// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark8DataReader struct {
	Benchmark8DataStorer
}

func (this *Benchmark8DataReader) FileName() string {
	return "benchmark8Data.pbd"
}

func (this *Benchmark8DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark8DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark8DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark8DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark8DataReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark8DataStorer); err != nil {
		return err
	}

	return nil
}