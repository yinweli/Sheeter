// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark02DataCReader struct {
	Benchmark02DataCStorer
}

func (this *Benchmark02DataCReader) FileName() string {
	return "benchmark02DataC.pbd"
}

func (this *Benchmark02DataCReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark02DataCReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark02DataCReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark02DataCReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark02DataCReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark02DataCStorer); err != nil {
		return err
	}

	return nil
}
