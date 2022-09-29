// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark260DataReader struct {
	Benchmark260DataStorer
}

func (this *Benchmark260DataReader) FileName() string {
	return "benchmark260Data.pbd"
}

func (this *Benchmark260DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark260DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark260DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark260DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark260DataReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark260DataStorer); err != nil {
		return err
	}

	return nil
}
