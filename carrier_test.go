package main

import (
	"testing"

	"github.com/iktefish/binary-helix/utils"
	"github.com/iktefish/binary-helix/workers"
)

func TestCarrier(t *testing.T) {
	fileExt, processed, lineCount := workers.Reader("./sample/phix.fa")
	p := "TCTA"
	splits := workers.Splitter(fileExt, processed, lineCount)

	itemCount := 0
	utils.Admin_EchoDbContents("slices_db", &itemCount)

	workers.Carrier(splits, utils.AnalyserList[3], p)

	newItemCount := 0
	utils.Admin_EchoDbContents("slices_db", &newItemCount)

	if newItemCount == itemCount {
		t.Error("Expected false")
	}
}

func BenchmarkCarrier(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fileExt, processed, lineCount := workers.Reader("./sample/phix.fa")
		p := "TCTA"
		splits := workers.Splitter(fileExt, processed, lineCount)
		workers.Carrier(splits, utils.AnalyserList[3], p)
	}
}
