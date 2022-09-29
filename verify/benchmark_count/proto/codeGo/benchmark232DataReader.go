// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark232DataReader struct {
	Benchmark232DataStorer
}

func (this *Benchmark232DataReader) FileName() string {
	return "benchmark232Data.pbd"
}

func (this *Benchmark232DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark232DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark232DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark232DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark232DataReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark232DataStorer); err != nil {
		return err
	}

	return nil
}