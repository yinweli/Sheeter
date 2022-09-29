// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark221DataReader struct {
	Benchmark221DataStorer
}

func (this *Benchmark221DataReader) FileName() string {
	return "benchmark221Data.pbd"
}

func (this *Benchmark221DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark221DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark221DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark221DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark221DataReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark221DataStorer); err != nil {
		return err
	}

	return nil
}