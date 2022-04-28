package main

import (
	"testing"
)

func TestCarrier(t *testing.T) {
		fileExt, processed, lineCount := workers.Reader("./test/input/phix.fa")
		splits := workers.Splitter(fileExt, processed, lineCount)
		workers.Carrier(splits, utils.AnalyserList[3], p)

    if utils.Get_ActiveNodeCount() < 0 {
		t.Error("Expected >= 0")
    }
}

func BenchmarkCarrier(b *testing.B) {
	for i := 0; i < b.N; i++ {
        utils.Get_ActiveNodeCount()
	}
}
