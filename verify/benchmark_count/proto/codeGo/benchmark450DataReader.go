// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark450DataReader struct {
	Benchmark450DataStorer
}

func (this *Benchmark450DataReader) FileName() string {
	return "benchmark450Data.pbd"
}

func (this *Benchmark450DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark450DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark450DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark450DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark450DataReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark450DataStorer); err != nil {
		return err
	}

	return nil
}
