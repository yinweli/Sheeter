// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark71DataReader struct {
	Benchmark71DataStorer
}

func (this *Benchmark71DataReader) FileName() string {
	return "benchmark71Data.pbd"
}

func (this *Benchmark71DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark71DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark71DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark71DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark71DataReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark71DataStorer); err != nil {
		return err
	}

	return nil
}
