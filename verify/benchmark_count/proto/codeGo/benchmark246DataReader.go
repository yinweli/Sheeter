// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark246DataReader struct {
	Benchmark246DataStorer
}

func (this *Benchmark246DataReader) FileName() string {
	return "benchmark246Data.pbd"
}

func (this *Benchmark246DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark246DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark246DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark246DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark246DataReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark246DataStorer); err != nil {
		return err
	}

	return nil
}