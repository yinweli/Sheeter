// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark43DataReader struct {
	Benchmark43DataStorer
}

func (this *Benchmark43DataReader) FileName() string {
	return "benchmark43Data.pbd"
}

func (this *Benchmark43DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark43DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark43DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark43DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark43DataReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark43DataStorer); err != nil {
		return err
	}

	return nil
}