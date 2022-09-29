// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark101DataReader struct {
	Benchmark101DataStorer
}

func (this *Benchmark101DataReader) FileName() string {
	return "benchmark101Data.pbd"
}

func (this *Benchmark101DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark101DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark101DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark101DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark101DataReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark101DataStorer); err != nil {
		return err
	}

	return nil
}