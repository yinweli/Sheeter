// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark348DataReader struct {
	Benchmark348DataStorer
}

func (this *Benchmark348DataReader) FileName() string {
	return "benchmark348Data.pbd"
}

func (this *Benchmark348DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark348DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark348DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark348DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark348DataReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark348DataStorer); err != nil {
		return err
	}

	return nil
}
