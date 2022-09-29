// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark05DataAReader struct {
	Benchmark05DataAStorer
}

func (this *Benchmark05DataAReader) FileName() string {
	return "benchmark05DataA.pbd"
}

func (this *Benchmark05DataAReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark05DataAReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark05DataAReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark05DataAReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark05DataAReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark05DataAStorer); err != nil {
		return err
	}

	return nil
}
