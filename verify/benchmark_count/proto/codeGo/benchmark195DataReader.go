// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark195DataReader struct {
	Benchmark195DataStorer
}

func (this *Benchmark195DataReader) FileName() string {
	return "benchmark195Data.pbd"
}

func (this *Benchmark195DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark195DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark195DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark195DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark195DataReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark195DataStorer); err != nil {
		return err
	}

	return nil
}
