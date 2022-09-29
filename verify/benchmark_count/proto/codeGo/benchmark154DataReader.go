// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark154DataReader struct {
	Benchmark154DataStorer
}

func (this *Benchmark154DataReader) FileName() string {
	return "benchmark154Data.pbd"
}

func (this *Benchmark154DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark154DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark154DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark154DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark154DataReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark154DataStorer); err != nil {
		return err
	}

	return nil
}
