// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark04DataGReader struct {
	Benchmark04DataGStorer
}

func (this *Benchmark04DataGReader) FileName() string {
	return "benchmark04DataG.pbd"
}

func (this *Benchmark04DataGReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark04DataGReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark04DataGReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark04DataGReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark04DataGReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark04DataGStorer); err != nil {
		return err
	}

	return nil
}