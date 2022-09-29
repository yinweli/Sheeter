// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark428DataReader struct {
	Benchmark428DataStorer
}

func (this *Benchmark428DataReader) FileName() string {
	return "benchmark428Data.pbd"
}

func (this *Benchmark428DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark428DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark428DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark428DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark428DataReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark428DataStorer); err != nil {
		return err
	}

	return nil
}
