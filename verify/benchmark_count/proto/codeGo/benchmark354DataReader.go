// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark354DataReader struct {
	Benchmark354DataStorer
}

func (this *Benchmark354DataReader) FileName() string {
	return "benchmark354Data.pbd"
}

func (this *Benchmark354DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark354DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark354DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark354DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark354DataReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark354DataStorer); err != nil {
		return err
	}

	return nil
}
