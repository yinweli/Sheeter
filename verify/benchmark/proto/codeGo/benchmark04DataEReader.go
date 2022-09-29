// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark04DataEReader struct {
	Benchmark04DataEStorer
}

func (this *Benchmark04DataEReader) FileName() string {
	return "benchmark04DataE.pbd"
}

func (this *Benchmark04DataEReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark04DataEReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark04DataEReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark04DataEReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark04DataEReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark04DataEStorer); err != nil {
		return err
	}

	return nil
}
