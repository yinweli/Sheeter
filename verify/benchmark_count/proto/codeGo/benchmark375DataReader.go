// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark375DataReader struct {
	Benchmark375DataStorer
}

func (this *Benchmark375DataReader) FileName() string {
	return "benchmark375Data.pbd"
}

func (this *Benchmark375DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark375DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark375DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark375DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark375DataReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark375DataStorer); err != nil {
		return err
	}

	return nil
}