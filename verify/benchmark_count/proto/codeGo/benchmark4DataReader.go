// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark4DataReader struct {
	Benchmark4DataStorer
}

func (this *Benchmark4DataReader) FileName() string {
	return "benchmark4Data.pbd"
}

func (this *Benchmark4DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark4DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark4DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark4DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark4DataReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark4DataStorer); err != nil {
		return err
	}

	return nil
}