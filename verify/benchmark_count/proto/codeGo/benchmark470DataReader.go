// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark470DataReader struct {
	Benchmark470DataStorer
}

func (this *Benchmark470DataReader) FileName() string {
	return "benchmark470Data.pbd"
}

func (this *Benchmark470DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark470DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark470DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark470DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark470DataReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark470DataStorer); err != nil {
		return err
	}

	return nil
}