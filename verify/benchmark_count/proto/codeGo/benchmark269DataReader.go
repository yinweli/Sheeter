// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark269DataReader struct {
	Benchmark269DataStorer
}

func (this *Benchmark269DataReader) FileName() string {
	return "benchmark269Data.pbd"
}

func (this *Benchmark269DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark269DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark269DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark269DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark269DataReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark269DataStorer); err != nil {
		return err
	}

	return nil
}