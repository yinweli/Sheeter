// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark04DataDReader struct {
	Benchmark04DataDStorer
}

func (this *Benchmark04DataDReader) FileName() string {
	return "benchmark04DataD.pbd"
}

func (this *Benchmark04DataDReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark04DataDReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark04DataDReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark04DataDReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark04DataDReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark04DataDStorer); err != nil {
		return err
	}

	return nil
}
