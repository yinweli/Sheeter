// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark496DataReader struct {
	Benchmark496DataStorer
}

func (this *Benchmark496DataReader) FileName() string {
	return "benchmark496Data.pbd"
}

func (this *Benchmark496DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark496DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark496DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark496DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark496DataReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark496DataStorer); err != nil {
		return err
	}

	return nil
}
