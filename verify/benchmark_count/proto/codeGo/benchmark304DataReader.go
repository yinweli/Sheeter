// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark304DataReader struct {
	Benchmark304DataStorer
}

func (this *Benchmark304DataReader) FileName() string {
	return "benchmark304Data.pbd"
}

func (this *Benchmark304DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark304DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark304DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark304DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark304DataReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark304DataStorer); err != nil {
		return err
	}

	return nil
}
