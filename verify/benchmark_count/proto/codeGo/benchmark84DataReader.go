// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark84DataReader struct {
	Benchmark84DataStorer
}

func (this *Benchmark84DataReader) FileName() string {
	return "benchmark84Data.pbd"
}

func (this *Benchmark84DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark84DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark84DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark84DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark84DataReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark84DataStorer); err != nil {
		return err
	}

	return nil
}
