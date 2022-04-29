package main

import (
	"testing"

	"github.com/iktefish/binary-helix/utils"
	"github.com/iktefish/binary-helix/workers"
)

func TestReader(t *testing.T) {
	fileExt, processed, lineCount := workers.Reader("./sample/phix.fa")
	p := "TCTA"
	splits := workers.Splitter(fileExt, processed, lineCount)

    mergedICount := 0
	itemCount := 0
	utils.Admin_EchoDbContents("slices_db", &itemCount, &mergedICount)

	workers.Carrier(splits, utils.AnalyserList[3], p)

    newMergedICount := 0
	newItemCount := 0
	utils.Admin_EchoDbContents("slices_db", &newItemCount, &newMergedICount)

	if newItemCount == itemCount {
		t.Error("Expected false")
	}

	if fileExt != ".fa" {
		t.Error("Proper identification of file extension.")
	}
}

func BenchmarkReader(b *testing.B) {
	for i := 0; i < b.N; i++ {
		workers.Reader("./sample/phix.fa")
	}
}
