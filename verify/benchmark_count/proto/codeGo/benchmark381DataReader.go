// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark381DataReader struct {
	Benchmark381DataStorer
}

func (this *Benchmark381DataReader) FileName() string {
	return "benchmark381Data.pbd"
}

func (this *Benchmark381DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark381DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark381DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark381DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark381DataReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark381DataStorer); err != nil {
		return err
	}

	return nil
}
