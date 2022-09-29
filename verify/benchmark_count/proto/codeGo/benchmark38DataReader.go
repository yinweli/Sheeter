// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark38DataReader struct {
	Benchmark38DataStorer
}

func (this *Benchmark38DataReader) FileName() string {
	return "benchmark38Data.pbd"
}

func (this *Benchmark38DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark38DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark38DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark38DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark38DataReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark38DataStorer); err != nil {
		return err
	}

	return nil
}
