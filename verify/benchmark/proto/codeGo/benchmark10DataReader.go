// generated by sheeter, DO NOT EDIT.

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark10DataReader struct {
	Benchmark10DataStorer
}

func (this *Benchmark10DataReader) FileName() string {
	return "benchmark10Data.pbd"
}

func (this *Benchmark10DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark10DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark10DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark10DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark10DataReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark10DataStorer); err != nil {
		return err
	}

	return nil
}
