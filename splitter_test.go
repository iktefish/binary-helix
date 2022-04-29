package main

import (
	"testing"

	"github.com/iktefish/binary-helix/utils"
	"github.com/iktefish/binary-helix/workers"
)

func TestSplitter(t *testing.T) {
	fileExt, processed, lineCount := workers.Reader("./sample/phix.fa")
	splits := workers.Splitter(fileExt, processed, lineCount)

	nodeCount := utils.Get_ActiveNodeCount()

	if len(splits) != nodeCount {
		t.Error("Expected an even split between nodes!")
	}
}

func BenchmarkSplitter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fileExt, processed, lineCount := workers.Reader("./sample/phix.fa")
		workers.Splitter(fileExt, processed, lineCount)
	}
}
