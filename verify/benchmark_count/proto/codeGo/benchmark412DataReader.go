// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark412DataReader struct {
	Benchmark412DataStorer
}

func (this *Benchmark412DataReader) FileName() string {
	return "benchmark412Data.pbd"
}

func (this *Benchmark412DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark412DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark412DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark412DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark412DataReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark412DataStorer); err != nil {
		return err
	}

	return nil
}
