// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark02DataEReader struct {
	Benchmark02DataEStorer
}

func (this *Benchmark02DataEReader) FileName() string {
	return "benchmark02DataE.pbd"
}

func (this *Benchmark02DataEReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark02DataEReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark02DataEReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark02DataEReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark02DataEReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark02DataEStorer); err != nil {
		return err
	}

	return nil
}
