// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark142DataReader struct {
	Benchmark142DataStorer
}

func (this *Benchmark142DataReader) FileName() string {
	return "benchmark142Data.pbd"
}

func (this *Benchmark142DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark142DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark142DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark142DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark142DataReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark142DataStorer); err != nil {
		return err
	}

	return nil
}
