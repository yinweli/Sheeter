// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark05DataBReader struct {
	Benchmark05DataBStorer
}

func (this *Benchmark05DataBReader) FileName() string {
	return "benchmark05DataB.pbd"
}

func (this *Benchmark05DataBReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark05DataBReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark05DataBReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark05DataBReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark05DataBReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark05DataBStorer); err != nil {
		return err
	}

	return nil
}
