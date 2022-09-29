// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark140DataReader struct {
	Benchmark140DataStorer
}

func (this *Benchmark140DataReader) FileName() string {
	return "benchmark140Data.pbd"
}

func (this *Benchmark140DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark140DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark140DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark140DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark140DataReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark140DataStorer); err != nil {
		return err
	}

	return nil
}
