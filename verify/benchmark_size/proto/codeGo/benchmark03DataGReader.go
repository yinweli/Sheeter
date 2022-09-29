// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark03DataGReader struct {
	Benchmark03DataGStorer
}

func (this *Benchmark03DataGReader) FileName() string {
	return "benchmark03DataG.pbd"
}

func (this *Benchmark03DataGReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark03DataGReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark03DataGReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark03DataGReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark03DataGReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark03DataGStorer); err != nil {
		return err
	}

	return nil
}