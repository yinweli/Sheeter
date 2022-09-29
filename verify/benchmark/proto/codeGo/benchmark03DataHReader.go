// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark03DataHReader struct {
	Benchmark03DataHStorer
}

func (this *Benchmark03DataHReader) FileName() string {
	return "benchmark03DataH.pbd"
}

func (this *Benchmark03DataHReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark03DataHReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark03DataHReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark03DataHReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark03DataHReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark03DataHStorer); err != nil {
		return err
	}

	return nil
}
