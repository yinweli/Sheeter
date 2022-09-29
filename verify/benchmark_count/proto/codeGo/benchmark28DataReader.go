// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark28DataReader struct {
	Benchmark28DataStorer
}

func (this *Benchmark28DataReader) FileName() string {
	return "benchmark28Data.pbd"
}

func (this *Benchmark28DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark28DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark28DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark28DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark28DataReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark28DataStorer); err != nil {
		return err
	}

	return nil
}
