// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark23DataReader struct {
	Benchmark23DataStorer
}

func (this *Benchmark23DataReader) FileName() string {
	return "benchmark23Data.pbd"
}

func (this *Benchmark23DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark23DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark23DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark23DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark23DataReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark23DataStorer); err != nil {
		return err
	}

	return nil
}
