// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark05DataCReader struct {
	Benchmark05DataCStorer
}

func (this *Benchmark05DataCReader) FileName() string {
	return "benchmark05DataC.pbd"
}

func (this *Benchmark05DataCReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark05DataCReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark05DataCReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark05DataCReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark05DataCReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark05DataCStorer); err != nil {
		return err
	}

	return nil
}
