// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark163DataReader struct {
	Benchmark163DataStorer
}

func (this *Benchmark163DataReader) FileName() string {
	return "benchmark163Data.pbd"
}

func (this *Benchmark163DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark163DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark163DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark163DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark163DataReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark163DataStorer); err != nil {
		return err
	}

	return nil
}
