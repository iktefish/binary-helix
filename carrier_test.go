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

    mergedICounter := 0
	itemCount := 0
	utils.Admin_EchoDbContents("slices_db", &itemCount, &mergedICounter)

	workers.Carrier(splits, utils.AnalyserList[3], p)

	newItemCount := 0
	utils.Admin_EchoDbContents("slices_db", &newItemCount, &mergedICounter)

	if newItemCount == itemCount {
		t.Error("Expected false")
	}
}

func BenchmarkCarrier(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fileExt, processed, lineCount := workers.Reader("./sample/phix.fa")
		p := "TCTA"
		t := "trial"
		splits := workers.Splitter(fileExt, processed, lineCount)
		workers.Carrier(splits, t, p)
	}
}
