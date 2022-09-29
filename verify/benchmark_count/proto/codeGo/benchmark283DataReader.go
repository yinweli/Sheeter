// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

package sheeterProto

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Benchmark283DataReader struct {
	Benchmark283DataStorer
}

func (this *Benchmark283DataReader) FileName() string {
	return "benchmark283Data.pbd"
}

func (this *Benchmark283DataReader) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Benchmark283DataReader: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark283DataReader) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("Benchmark283DataReader: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *Benchmark283DataReader) FromData(data []byte) error {
	if err := proto.Unmarshal(data, &this.Benchmark283DataStorer); err != nil {
		return err
	}

	return nil
}
