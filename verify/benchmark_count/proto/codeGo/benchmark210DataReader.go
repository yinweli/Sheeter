// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark210DataReader struct {
	Benchmark210DataStorer
}

func (this *Benchmark210DataReader) FileName() string {
	return "benchmark210Data.pbd"
}

func (this *Benchmark210DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark210DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark210DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark210DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark210DataReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark210DataStorer); err != nil {
		return err
	}

	return nil
}
