// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark208DataReader struct {
	Benchmark208DataStorer
}

func (this *Benchmark208DataReader) FileName() string {
	return "benchmark208Data.pbd"
}

func (this *Benchmark208DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark208DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark208DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark208DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark208DataReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark208DataStorer); err != nil {
		return err
	}

	return nil
}