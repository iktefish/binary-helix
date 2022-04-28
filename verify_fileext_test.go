package main

import (
	"testing"

	"github.com/iktefish/binary-helix/utils"
)

func TestVerify_FileExt(t *testing.T) {
	list := []string{"a.out", "b.bo", "c.ca", "f.fa", "c.fastq"}

	if utils.Verify_FileExt(list[0]) == true {
		t.Error("Expected false")
	}
}

func BenchmarkVerify_FileExt(b *testing.B) {
	list := []string{"a.out", "b.bo", "c.ca", "f.fa", "c.fastq"}
	for i := 0; i < b.N; i++ {
		for _, r := range list {
			utils.Verify_FileExt(r)
		}
	}
}
