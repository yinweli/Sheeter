// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark80DataReader struct {
	Benchmark80DataStorer
}

func (this *Benchmark80DataReader) FileName() string {
	return "benchmark80Data.pbd"
}

func (this *Benchmark80DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark80DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark80DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark80DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark80DataReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark80DataStorer); err != nil {
		return err
	}

	return nil
}
