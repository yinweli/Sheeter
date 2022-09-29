// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark148DataReader struct {
	Benchmark148DataStorer
}

func (this *Benchmark148DataReader) FileName() string {
	return "benchmark148Data.pbd"
}

func (this *Benchmark148DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark148DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark148DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark148DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark148DataReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark148DataStorer); err != nil {
		return err
	}

	return nil
}