// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark89DataReader struct {
	Benchmark89DataStorer
}

func (this *Benchmark89DataReader) FileName() string {
	return "benchmark89Data.pbd"
}

func (this *Benchmark89DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark89DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark89DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark89DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark89DataReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark89DataStorer); err != nil {
		return err
	}

	return nil
}
