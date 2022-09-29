// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark37DataReader struct {
	Benchmark37DataStorer
}

func (this *Benchmark37DataReader) FileName() string {
	return "benchmark37Data.pbd"
}

func (this *Benchmark37DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark37DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark37DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark37DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark37DataReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark37DataStorer); err != nil {
		return err
	}

	return nil
}
