// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark47DataReader struct {
	Benchmark47DataStorer
}

func (this *Benchmark47DataReader) FileName() string {
	return "benchmark47Data.pbd"
}

func (this *Benchmark47DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark47DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark47DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark47DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark47DataReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark47DataStorer); err != nil {
		return err
	}

	return nil
}
