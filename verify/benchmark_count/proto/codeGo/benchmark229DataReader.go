// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark229DataReader struct {
	Benchmark229DataStorer
}

func (this *Benchmark229DataReader) FileName() string {
	return "benchmark229Data.pbd"
}

func (this *Benchmark229DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark229DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark229DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark229DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark229DataReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark229DataStorer); err != nil {
		return err
	}

	return nil
}
