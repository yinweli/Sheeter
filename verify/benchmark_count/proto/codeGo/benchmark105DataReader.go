// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark105DataReader struct {
	Benchmark105DataStorer
}

func (this *Benchmark105DataReader) FileName() string {
	return "benchmark105Data.pbd"
}

func (this *Benchmark105DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark105DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark105DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark105DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark105DataReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark105DataStorer); err != nil {
		return err
	}

	return nil
}
