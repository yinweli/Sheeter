// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark04DataJReader struct {
	Benchmark04DataJStorer
}

func (this *Benchmark04DataJReader) FileName() string {
	return "benchmark04DataJ.pbd"
}

func (this *Benchmark04DataJReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark04DataJReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark04DataJReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark04DataJReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark04DataJReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark04DataJStorer); err != nil {
		return err
	}

	return nil
}
