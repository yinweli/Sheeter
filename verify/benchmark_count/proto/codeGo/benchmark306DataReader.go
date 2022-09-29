// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark306DataReader struct {
	Benchmark306DataStorer
}

func (this *Benchmark306DataReader) FileName() string {
	return "benchmark306Data.pbd"
}

func (this *Benchmark306DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark306DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark306DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark306DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark306DataReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark306DataStorer); err != nil {
		return err
	}

	return nil
}
