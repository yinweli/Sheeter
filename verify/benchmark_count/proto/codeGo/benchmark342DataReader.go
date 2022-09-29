// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark342DataReader struct {
	Benchmark342DataStorer
}

func (this *Benchmark342DataReader) FileName() string {
	return "benchmark342Data.pbd"
}

func (this *Benchmark342DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark342DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark342DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark342DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark342DataReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark342DataStorer); err != nil {
		return err
	}

	return nil
}