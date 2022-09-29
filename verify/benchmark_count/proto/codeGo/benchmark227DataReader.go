// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark227DataReader struct {
	Benchmark227DataStorer
}

func (this *Benchmark227DataReader) FileName() string {
	return "benchmark227Data.pbd"
}

func (this *Benchmark227DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark227DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark227DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark227DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark227DataReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark227DataStorer); err != nil {
		return err
	}

	return nil
}