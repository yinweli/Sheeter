// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark110DataReader struct {
	Benchmark110DataStorer
}

func (this *Benchmark110DataReader) FileName() string {
	return "benchmark110Data.pbd"
}

func (this *Benchmark110DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark110DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark110DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark110DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark110DataReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark110DataStorer); err != nil {
		return err
	}

	return nil
}
